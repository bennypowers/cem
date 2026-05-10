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
package platform_test

import (
	"errors"
	"io/fs"
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/set"
)

func TestWalkDir(t *testing.T) {
	fsys := fstest.MapFS{
		"src/main.go":                    &fstest.MapFile{},
		"src/util.go":                    &fstest.MapFile{},
		".git/HEAD":                      &fstest.MapFile{},
		".git/refs/heads/main":           &fstest.MapFile{},
		"sub/.git/HEAD":                  &fstest.MapFile{},
		"node_modules/foo/index.js":      &fstest.MapFile{},
		"dist/bundle.js":                 &fstest.MapFile{},
		"lib/helper.go":                  &fstest.MapFile{},
	}

	t.Run("prunes .git at any depth", func(t *testing.T) {
		var visited []string
		err := platform.WalkDir(fsys, ".", nil, func(path string, d fs.DirEntry, err error) error {
			visited = append(visited, path)
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range visited {
			if v == ".git" || v == "sub/.git" || v == ".git/HEAD" || v == ".git/refs/heads/main" || v == "sub/.git/HEAD" {
				t.Errorf("visited pruned path: %s", v)
			}
		}
		if !slices.Contains(visited, "src/main.go") {
			t.Error("did not visit src/main.go")
		}
	})

	t.Run("prunes multiple dirs", func(t *testing.T) {
		var visited []string
		err := platform.WalkDir(fsys, ".", set.NewSet(".git", "node_modules", "dist"), func(path string, d fs.DirEntry, err error) error {
			visited = append(visited, path)
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		blocked := map[string]bool{".git": true, "node_modules": true, "dist": true}
		for _, v := range visited {
			if blocked[v] {
				t.Errorf("visited pruned dir: %s", v)
			}
		}
	})

	t.Run("passes through normal dirs and files", func(t *testing.T) {
		want := map[string]bool{"src": false, "src/main.go": false, "lib": false, "lib/helper.go": false}
		err := platform.WalkDir(fsys, ".", set.NewSet(".git"), func(path string, d fs.DirEntry, err error) error {
			if _, ok := want[path]; ok {
				want[path] = true
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		for path, seen := range want {
			if !seen {
				t.Errorf("did not visit expected path: %s", path)
			}
		}
	})

	t.Run("caller SkipDir composes", func(t *testing.T) {
		var visited []string
		err := platform.WalkDir(fsys, ".", set.NewSet(".git"), func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() && d.Name() == "node_modules" {
				return fs.SkipDir
			}
			visited = append(visited, path)
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range visited {
			if v == "node_modules" || v == "node_modules/foo/index.js" {
				t.Errorf("visited caller-skipped path: %s", v)
			}
		}
	})

	t.Run("propagates errors from fn", func(t *testing.T) {
		sentinel := errors.New("stop")
		err := platform.WalkDir(fsys, ".", set.NewSet(".git"), func(path string, d fs.DirEntry, err error) error {
			if path == "src/main.go" {
				return sentinel
			}
			return nil
		})
		if !errors.Is(err, sentinel) {
			t.Errorf("expected sentinel error, got: %v", err)
		}
	})

	t.Run("root dot entry not pruned by caller dot-prefix check", func(t *testing.T) {
		var visited []string
		err := platform.WalkDir(fsys, ".", nil, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() && path != "." && strings.HasPrefix(d.Name(), ".") {
				return fs.SkipDir
			}
			visited = append(visited, path)
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Contains(visited, ".") {
			t.Error("root entry '.' must not be pruned")
		}
		if !slices.Contains(visited, "src/main.go") {
			t.Error("files under root must be visited")
		}
	})

	t.Run("nil skip set still prunes .git", func(t *testing.T) {
		var visited []string
		err := platform.WalkDir(fsys, ".", nil, func(path string, d fs.DirEntry, err error) error {
			visited = append(visited, path)
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		if slices.Contains(visited, ".git") {
			t.Error(".git should be pruned even with nil skip set")
		}
		if !slices.Contains(visited, "src/main.go") {
			t.Error("normal files should still be visited")
		}
	})
}
