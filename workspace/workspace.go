/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

SPDX-License-Identifier: GPL-3.0-or-later

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

// Package workspace provides an exported constructor for filesystem-backed
// workspace contexts that are usable outside the cem module.
package workspace

import (
	"fmt"
	"io/fs"

	"bennypowers.dev/cem/internal/platform"
	W "bennypowers.dev/cem/internal/workspace"
	"bennypowers.dev/cem/types"
)

// NewFileSystemWorkspaceContext returns a WorkspaceContext rooted at dir.
//
// The context is backed by the host filesystem and uses the same local
// workspace context and Init sequence that cem applies after it resolves a
// package path. Construction runs Init, so ConfigFile, Config, PackageJSON,
// and CustomElementsManifestPath are ready when the constructor returns.
// Manifest contents are still decoded lazily on the first Manifest call, as
// in the internal implementation.
//
// dir is used as the workspace root verbatim, so pass an absolute path to keep
// later operations independent of the process working directory. Config
// defaults may be inherited from an enclosing workspace root.
//
// dir must name an existing directory. A missing directory returns an error
// wrapping fs.ErrNotExist, and a non-directory path returns an error wrapping
// fs.ErrInvalid. Errors discovered by Init, such as an invalid config file,
// are returned wrapped with the workspace directory. Init also applies a
// config file's logLevel or verbose setting to cem's process-wide logger,
// matching internal behavior. An unreadable or invalid manifest is not
// reported here; it surfaces from Manifest.
//
// Cleanup is a no-op for filesystem-backed workspaces and need not be called,
// but calling it is safe.
func NewFileSystemWorkspaceContext(dir string) (types.WorkspaceContext, error) {
	fsys := platform.NewOSFileSystem()

	info, err := fsys.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("workspace: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("workspace: %s: not a directory: %w", dir, fs.ErrInvalid)
	}

	ctx := W.NewFileSystemWorkspaceContext(dir, W.WithFileSystem(fsys))
	if err := ctx.Init(); err != nil {
		return nil, fmt.Errorf("workspace: initializing %s: %w", dir, err)
	}

	return ctx, nil
}
