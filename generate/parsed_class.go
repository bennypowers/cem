package generate

import (
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
)

// ParsedClass is an intermediate structure for class processing in the generate package.
type ParsedClass struct {
	Name            string
	Alias           string
	Captures        Q.CaptureMap
	IsCustomElement bool
	IsHTMLElement   bool
	CEMDeclaration  M.Declaration
	CssProperties   []M.CssCustomProperty
}
