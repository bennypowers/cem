// Package registry triggers registration of all tree-sitter language
// packages. Import it with a blank identifier to ensure every language
// is available to the query manager before queries are compiled.
package registry

import (
	_ "bennypowers.dev/cem/internal/languages/blade"
	_ "bennypowers.dev/cem/internal/languages/css"
	_ "bennypowers.dev/cem/internal/languages/et"
	_ "bennypowers.dev/cem/internal/languages/handlebars"
	_ "bennypowers.dev/cem/internal/languages/html"
	_ "bennypowers.dev/cem/internal/languages/jinja"
	_ "bennypowers.dev/cem/internal/languages/jsdoc"
	_ "bennypowers.dev/cem/internal/languages/php"
	_ "bennypowers.dev/cem/internal/languages/tsx"
	_ "bennypowers.dev/cem/internal/languages/typescript"
)
