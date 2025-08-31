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

// SourceReference is a reference to the source of a declaration or member.
type SourceReference struct {
	Href string `json:"href"`
}

// Clone creates a deep copy of the SourceReference.
func (s SourceReference) Clone() SourceReference {
	return SourceReference{
		Href: s.Href,
	}
}
