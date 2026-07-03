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
package breaking

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	M "bennypowers.dev/cem/manifest"
)

func ResolveManifestFromGit(ref, manifestPath, workDir string) (*M.Package, error) {
	gitPath := manifestPath
	if workDir != "" {
		rel, err := relativeToGitRoot(workDir, manifestPath)
		if err == nil && rel != "" {
			gitPath = rel
		}
	}

	gitCmd := exec.Command("git", "show", ref+":"+gitPath)
	if workDir != "" {
		gitCmd.Dir = workDir
	}
	out, err := gitCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest at %s from git ref %q: %w", gitPath, ref, err)
	}

	var pkg M.Package
	if err := json.Unmarshal(out, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse manifest at %s from git ref %q: %w", gitPath, ref, err)
	}

	return &pkg, nil
}

func FindLatestSemverTag(workDir string) (string, error) {
	gitCmd := exec.Command("git", "tag", "--list", "--sort=-v:refname")
	if workDir != "" {
		gitCmd.Dir = workDir
	}
	out, err := gitCmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to list git tags: %w", err)
	}

	for line := range strings.SplitSeq(strings.TrimSpace(string(out)), "\n") {
		tag := strings.TrimSpace(line)
		if tag == "" {
			continue
		}
		if isSemverTag(tag) {
			return tag, nil
		}
	}

	return "", fmt.Errorf("no semver tags found")
}

func isSemverTag(tag string) bool {
	tag = strings.TrimPrefix(tag, "v")
	parts := strings.SplitN(tag, ".", 3)
	if len(parts) < 3 {
		return false
	}
	for _, part := range parts {
		p := strings.SplitN(part, "-", 2)[0]
		if p == "" {
			return false
		}
		for _, c := range p {
			if c < '0' || c > '9' {
				return false
			}
		}
	}
	return true
}

func relativeToGitRoot(workDir, path string) (string, error) {
	gitCmd := exec.Command("git", "rev-parse", "--show-prefix")
	gitCmd.Dir = workDir
	out, err := gitCmd.Output()
	if err != nil {
		return "", err
	}
	prefix := strings.TrimSpace(string(out))
	if prefix == "" {
		return path, nil
	}
	return prefix + path, nil
}
