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
package list

import (
	"context"
	"reflect"
	"strings"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/internal/tui"
	treeview "github.com/Digital-Shane/treeview/v2"
)

func RenderTree(title string, renderable M.Renderable, pred M.PredicateFunc) (string, error) {
	if renderable == nil || isTypedNil(renderable) {
		return "No deprecations found", nil
	}
	root := renderable.ToTreeNode(pred)

	if !root.HasChildren() {
		return "", nil
	}

	tree := treeview.NewTree([]*treeview.Node[M.DisplayNode]{root},
		treeview.WithExpandFunc[M.DisplayNode](func(_ *treeview.Node[M.DisplayNode]) bool { return true }),
	)

	s, err := tree.Render(context.Background())
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.WriteString(tui.HeaderStyle.Render(title))
	builder.WriteString("\n")
	builder.WriteString(s)
	return builder.String(), nil
}

func isTypedNil(i any) bool {
	// Use reflection to detect typed nil
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}
	return false
}
