package generate

import (
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/generate/queries"
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
