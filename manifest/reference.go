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
package manifest

// Reference to an export of a module.
type Reference struct {
	Name    string `json:"name"`
	Package string `json:"package,omitempty"`
	Module  string `json:"module,omitempty"`
}

func NewReference(name string, pkg string, module string) *Reference {
	return &Reference{
		Name:    name,
		Package: pkg,
		Module:  NormalizeSourcePath(module),
	}
}

// Clone creates a deep copy of the Reference.
func (r Reference) Clone() Reference {
	return Reference{
		Name:    r.Name,
		Package: r.Package,
		Module:  r.Module,
	}
}
