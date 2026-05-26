/*
Copyright © 2026 Benny Powers

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
	"testing"

	"github.com/stretchr/testify/assert"
)

// Inline: pure function, scalar assertions
func TestExportTracker_AddElementSource(t *testing.T) {
	et := &ExportTracker{
		ElementSources: make(map[string][]string),
	}

	et.AddElementSource("my-el", "src/my-el.js")
	assert.Equal(t, []string{"src/my-el.js"}, et.ElementSources["my-el"])

	et.AddElementSource("my-el", "src/my-el.js")
	assert.Len(t, et.ElementSources["my-el"], 1, "no duplicates")

	et.AddElementSource("my-el", "src/index.js")
	assert.Len(t, et.ElementSources["my-el"], 2)

	et.AddElementSource("", "src/x.js")
	et.AddElementSource("my-el", "")
	assert.Len(t, et.ElementSources["my-el"], 2, "empty args ignored")
}

func TestDependencyTracker_AddModuleDependency(t *testing.T) {
	dt := NewDependencyTracker()

	dt.AddModuleDependency("a.js", "b.js")
	assert.Equal(t, []string{"b.js"}, dt.GetModuleDependencies("a.js"))

	dt.AddModuleDependency("a.js", "b.js")
	assert.Len(t, dt.GetModuleDependencies("a.js"), 1, "no duplicates")

	dt.AddModuleDependency("a.js", "c.js")
	assert.Len(t, dt.GetModuleDependencies("a.js"), 2)

	dt.AddModuleDependency("", "b.js")
	dt.AddModuleDependency("a.js", "")
}

func TestDependencyTracker_GetModuleDependencies_Copy(t *testing.T) {
	dt := NewDependencyTracker()
	dt.AddModuleDependency("a.js", "b.js")

	deps := dt.GetModuleDependencies("a.js")
	deps[0] = "mutated"
	assert.Equal(t, "b.js", dt.GetModuleDependencies("a.js")[0], "returns copy")
}

func TestDependencyTracker_GetModuleDependencies_Missing(t *testing.T) {
	dt := NewDependencyTracker()
	assert.Nil(t, dt.GetModuleDependencies("nonexistent.js"))
}

func TestDependencyTracker_AddReExportChain(t *testing.T) {
	dt := NewDependencyTracker()

	dt.AddReExportChain("index.js", "button.js")
	chains := dt.GetReExportChains()
	assert.Equal(t, []string{"button.js"}, chains["index.js"])

	dt.AddReExportChain("index.js", "button.js")
	chains = dt.GetReExportChains()
	assert.Len(t, chains["index.js"], 1, "no duplicates")

	dt.AddReExportChain("", "x.js")
	dt.AddReExportChain("x.js", "")
}

func TestDependencyTracker_GetReExportChains_DeepCopy(t *testing.T) {
	dt := NewDependencyTracker()
	dt.AddReExportChain("a.js", "b.js")

	chains := dt.GetReExportChains()
	chains["a.js"][0] = "mutated"
	assert.Equal(t, "b.js", dt.GetReExportChains()["a.js"][0], "returns deep copy")
}
