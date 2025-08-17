/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package server

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// ServerContext provides the dependencies needed for server lifecycle
type ServerContext interface {
	DocumentManager() DocumentManager
	Workspace() Workspace
	DebugLog(format string, args ...any)
	InitializeManifests() error
	UpdateWorkspaceFromLSP(rootURI *string, workspaceFolders []protocol.WorkspaceFolder) error
}

// DocumentManager interface for cleanup
type DocumentManager interface {
	Close()
}

// Workspace interface for cleanup
type Workspace interface {
	Cleanup() error
}
