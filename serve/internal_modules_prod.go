//go:build !cemdev

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

package serve

import (
	"bennypowers.dev/cem/serve/middleware/routes"
)

// readInternalModule reads an internal module from the embedded filesystem
// This is the production version that uses embed.FS
func readInternalModule(path string) ([]byte, error) {
	return routes.InternalModules.ReadFile(path)
}

// setupDevWatcher is a no-op in production builds
func setupDevWatcher(server *Server) error {
	return nil
}
