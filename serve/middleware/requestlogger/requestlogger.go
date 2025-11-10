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

package requestlogger

import (
	"net/http"

	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
)

// New creates a logging middleware that logs all HTTP requests.
// Internal polling endpoints (/__cem-logs, /__cem/reload) are not logged.
func New(log logger.Logger) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip logging for internal polling endpoints
			if r.URL.Path != "/__cem-logs" && r.URL.Path != "/__cem/reload" {
				log.Info("%s %s", r.Method, r.URL.Path)
			}
			next.ServeHTTP(w, r)
		})
	}
}
