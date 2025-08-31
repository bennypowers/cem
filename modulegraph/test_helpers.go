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
package modulegraph

import (
	M "bennypowers.dev/cem/manifest"
)

// MockElementDefinition implements types.ElementDefinition for testing
type MockElementDefinition struct {
	TagName     string
	ClassName   string
	ModulePath  string
	PackageName string
	SourceHref  string
	ElementPtr  *M.CustomElement
}

func (m *MockElementDefinition) GetTagName() string {
	return m.TagName
}

func (m *MockElementDefinition) GetClassName() string {
	return m.ClassName
}

func (m *MockElementDefinition) GetModulePath() string {
	return m.ModulePath
}

func (m *MockElementDefinition) GetPackageName() string {
	return m.PackageName
}

func (m *MockElementDefinition) GetSourceHref() string {
	return m.SourceHref
}

func (m *MockElementDefinition) Element() *M.CustomElement {
	return m.ElementPtr
}