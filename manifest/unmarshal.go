/*
Copyright © 2025 Benny Powers

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

import (
	"bytes"
	"encoding/json"
)

func decodeDeprecatedField(dst *Deprecated, data json.RawMessage) bool {
	if dst == nil {
		return true
	}
	if len(data) == 0 || string(bytes.TrimSpace(data)) == "null" {
		*dst = nil
		return true
	}
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*dst = NewDeprecated(b)
		return true
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*dst = NewDeprecated(s)
		return true
	}
	return false // unknown type
}
