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
	"embed"
	"html/template"
)

//go:embed templates/default-index.html
var defaultIndexTemplate string

//go:embed templates/websocket-client.js
var WebSocketClientScript string

//go:embed templates/js/*.js
var internalModules embed.FS

// DefaultIndexTemplate is the parsed template for the default index page
var DefaultIndexTemplate = template.Must(template.New("default-index").Parse(defaultIndexTemplate))
