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
package platform

import (
	"io/fs"

	"bennypowers.dev/cem/internal/set"
)

var defaultSkipDirs = set.NewSet(".git")

// WalkDir walks the file tree rooted at root on fsys, silently pruning
// .git directories and any additional directories in skip.
// All other behavior matches fs.WalkDir.
func WalkDir(fsys fs.FS, root string, skip set.Set[string], fn fs.WalkDirFunc) error {
	return fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if d != nil && d.IsDir() {
			name := d.Name()
			if defaultSkipDirs.Has(name) || skip.Has(name) {
				return fs.SkipDir
			}
		}
		return fn(path, d, err)
	})
}
