/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	asimconfig "bennypowers.dev/asimonim/config"
	"bennypowers.dev/asimonim/specifier"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/set"
)

func normalizeGitURL(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return ""
	}

	var host, path string

	if after, ok := strings.CutPrefix(rawURL, "ssh://"); ok {
		after = strings.TrimPrefix(after, "git@")
		host, path, ok = strings.Cut(after, "/")
		if !ok {
			return rawURL
		}
		// strip optional :port from host
		if h, _, hasPort := strings.Cut(host, ":"); hasPort {
			host = h
		}
	} else if after, ok := strings.CutPrefix(rawURL, "git@"); ok {
		host, path, ok = strings.Cut(after, ":")
		if !ok {
			return rawURL
		}
	} else if after, ok := strings.CutPrefix(rawURL, "https://"); ok {
		host, path, ok = strings.Cut(after, "/")
		if !ok {
			return rawURL
		}
	} else if after, ok := strings.CutPrefix(rawURL, "http://"); ok {
		host, path, ok = strings.Cut(after, "/")
		if !ok {
			return rawURL
		}
	} else {
		return rawURL
	}

	path = strings.TrimSuffix(path, ".git")
	path = strings.TrimSuffix(path, "/")

	return fmt.Sprintf("https://%s/%s/", host, path)
}

var skipDirs = set.NewSet("node_modules", "dist", "test", "demo", ".git", "build", "_site")

type depthRange struct {
	min, max int
	seen     bool
}

func (d *depthRange) observe(depth int) {
	if !d.seen {
		d.min = depth
		d.max = depth
		d.seen = true
		return
	}
	d.min = min(d.min, depth)
	d.max = max(d.max, depth)
}

func detectSourceFiles(fsys fs.FS) ([]string, error) {
	dirDepths := map[string]*depthRange{}

	err := platform.WalkDir(fsys, ".", skipDirs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".ts") {
			return nil
		}
		if strings.HasSuffix(path, ".d.ts") {
			return nil
		}

		parts := strings.Split(filepath.ToSlash(path), "/")
		if len(parts) >= 2 {
			topDir := parts[0]
			depth := len(parts) - 2
			if dirDepths[topDir] == nil {
				dirDepths[topDir] = &depthRange{}
			}
			dirDepths[topDir].observe(depth)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(dirDepths) == 0 {
		return nil, nil
	}

	dirs := make([]string, 0, len(dirDepths))
	for d := range dirDepths {
		dirs = append(dirs, d)
	}
	slices.Sort(dirs)

	patterns := make([]string, len(dirs))
	for i, dir := range dirs {
		dr := dirDepths[dir]
		if dr.min != dr.max {
			patterns[i] = dir + "/**/*.ts"
		} else {
			patterns[i] = dir + "/" + globStars(dr.max) + "*.ts"
		}
	}
	return patterns, nil
}

// globStars returns the glob wildcard prefix for a given nesting depth.
// depth 0 -> "" (files directly in dir), depth 1 -> "*/" , depth 2+ -> "**/"
func globStars(depth int) string {
	switch {
	case depth <= 0:
		return ""
	case depth == 1:
		return "*/"
	default:
		return "**/"
	}
}

func detectDemoFiles(fsys fs.FS) (fileGlob string, urlPattern string, err error) {
	dirMaxDepth := map[string]int{}

	walkErr := platform.WalkDir(fsys, ".", set.NewSet("node_modules", ".git"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".html") {
			return nil
		}

		parts := strings.Split(filepath.ToSlash(path), "/")
		for i, part := range parts {
			if part == "demo" && i > 0 && i < len(parts)-1 {
				topDir := parts[0]
				depth := i - 1
				if d, ok := dirMaxDepth[topDir]; !ok || depth > d {
					dirMaxDepth[topDir] = depth
				}
				break
			}
		}
		return nil
	})
	if walkErr != nil {
		return "", "", walkErr
	}

	if len(dirMaxDepth) == 0 {
		return "", "", nil
	}

	dirs := make([]string, 0, len(dirMaxDepth))
	for d := range dirMaxDepth {
		dirs = append(dirs, d)
	}
	slices.Sort(dirs)
	topDir := dirs[0]

	depth := dirMaxDepth[topDir]
	star := globStars(depth)
	fileGlob = topDir + "/" + star + "demo/*.html"
	if depth > 0 {
		urlPattern = topDir + "/:tag/demo/:demo.html"
	} else {
		urlPattern = topDir + "/demo/:demo.html"
	}
	return fileGlob, urlPattern, nil
}

func detectCSSPatterns(sourcePatterns []string, fsys fs.FS) (include []string, exclude []string) {
	cssDirs := set.NewSet[string]()

	_ = platform.WalkDir(fsys, ".", skipDirs, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".css") {
			return nil
		}

		parts := strings.SplitN(filepath.ToSlash(path), "/", 2)
		if len(parts) >= 2 {
			topDir := parts[0]
			for _, sp := range sourcePatterns {
				if strings.HasPrefix(sp, topDir+"/") {
					cssDirs.Add(topDir)
					break
				}
			}
		}
		return nil
	})

	if len(cssDirs) == 0 {
		return nil, nil
	}

	dirs := cssDirs.Members()
	slices.Sort(dirs)

	include = make([]string, len(dirs))
	for i, dir := range dirs {
		include[i] = dir + "/**/*.css"
	}

	exclude = []string{"demo/**/*.css", "**/*.min.css"}
	return include, exclude
}

type gitRemoteResult struct {
	url    string
	branch string
}

func detectGitRemote(root string) (string, error) {
	gitCmd := exec.Command("git", "config", "--get", "remote.origin.url")
	gitCmd.Dir = root
	out, err := gitCmd.Output()
	if err != nil {
		return "", nil
	}
	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return "", nil
	}

	result := gitRemoteResult{
		url:    normalizeGitURL(raw),
		branch: detectDefaultBranch(root),
	}

	if ghPath, ghErr := exec.LookPath("gh"); ghErr == nil && ghPath != "" {
		result = detectForkUpstream(root, result)
	}

	return result.url + "tree/" + result.branch + "/", nil
}

func detectForkUpstream(root string, fallback gitRemoteResult) gitRemoteResult {
	ghCmd := exec.Command("gh", "repo", "view", "--json", "parent")
	ghCmd.Dir = root
	out, err := ghCmd.Output()
	if err != nil {
		return fallback
	}

	var result struct {
		Parent struct {
			URL   string `json:"url"`
			Owner struct {
				Login string `json:"login"`
			} `json:"owner"`
			Name string `json:"name"`
		} `json:"parent"`
	}
	if json.Unmarshal(out, &result) != nil || result.Parent.URL == "" {
		return fallback
	}

	upstream := normalizeGitURL(result.Parent.URL)
	if upstream == "" {
		return fallback
	}

	branch := detectRepoDefaultBranch(result.Parent.Owner.Login, result.Parent.Name)
	if branch == "" {
		branch = fallback.branch
	}

	return gitRemoteResult{url: upstream, branch: branch}
}

func detectRepoDefaultBranch(owner, name string) string {
	if owner == "" || name == "" {
		return ""
	}
	ghCmd := exec.Command("gh", "repo", "view", owner+"/"+name, "--json", "defaultBranchRef")
	out, err := ghCmd.Output()
	if err != nil {
		return ""
	}
	var result struct {
		DefaultBranchRef struct {
			Name string `json:"name"`
		} `json:"defaultBranchRef"`
	}
	if json.Unmarshal(out, &result) != nil {
		return ""
	}
	return result.DefaultBranchRef.Name
}

func detectDefaultBranch(root string) string {
	refCmd := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD")
	refCmd.Dir = root
	out, err := refCmd.Output()
	if err == nil {
		ref := strings.TrimSpace(string(out))
		if _, branch, ok := strings.Cut(ref, "refs/remotes/origin/"); ok && branch != "" {
			return branch
		}
	}

	return "main"
}

func detectDesignTokens(root string, fsys platform.FileSystem) ([]*specifier.ResolvedFile, error) {
	return asimconfig.DiscoverDesignTokens(fsys, root)
}

func detectTypeScript(root string, fsys platform.FileSystem) (hasTsConfig bool, target string) {
	_, err := fsys.Stat(filepath.Join(root, "tsconfig.json"))
	if err != nil {
		return false, ""
	}
	return true, ""
}
