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
package modulegraph_test

import (
	"testing"
	"time"

	"bennypowers.dev/cem/internal/modulegraph"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, scalar assertions
func TestDefaultMetricsCollector_Counters(t *testing.T) {
	m := modulegraph.NewDefaultMetricsCollector()

	assert.Equal(t, int64(0), m.GetCounterValue("files"))

	m.IncrementCounter("files")
	assert.Equal(t, int64(1), m.GetCounterValue("files"))

	m.IncrementCounter("files")
	m.IncrementCounter("files")
	assert.Equal(t, int64(3), m.GetCounterValue("files"))

	assert.Equal(t, int64(0), m.GetCounterValue("nonexistent"))
}

func TestDefaultMetricsCollector_Gauges(t *testing.T) {
	m := modulegraph.NewDefaultMetricsCollector()

	assert.Equal(t, int64(0), m.GetGaugeValue("modules"))

	m.SetGauge("modules", 42)
	assert.Equal(t, int64(42), m.GetGaugeValue("modules"))

	m.SetGauge("modules", 10)
	assert.Equal(t, int64(10), m.GetGaugeValue("modules"))

	assert.Equal(t, int64(0), m.GetGaugeValue("nonexistent"))
}

func TestDefaultMetricsCollector_Durations(t *testing.T) {
	m := modulegraph.NewDefaultMetricsCollector()
	m.RecordDuration("parse", 100*time.Millisecond)
	m.RecordDuration("parse", 200*time.Millisecond)
}

func TestDefaultMetricsCollector_Histograms(t *testing.T) {
	m := modulegraph.NewDefaultMetricsCollector()
	m.AddHistogramValue("depth", 1.0)
	m.AddHistogramValue("depth", 2.5)
	m.AddHistogramValue("depth", 3.0)
}

func TestNoOpMetricsCollector(t *testing.T) {
	m := &modulegraph.NoOpMetricsCollector{}
	m.IncrementCounter("x")
	m.SetGauge("x", 1)
	m.RecordDuration("x", time.Second)
	m.AddHistogramValue("x", 1.0)
}

func TestModuleGraph_SettersAndGetters(t *testing.T) {
	mg := modulegraph.NewModuleGraph(nil, nil)

	t.Run("GetMetrics returns collector", func(t *testing.T) {
		assert.NotNil(t, mg.GetMetrics())
	})

	t.Run("SetMaxTransitiveDepth", func(t *testing.T) {
		mg.SetMaxTransitiveDepth(5)
		assert.Equal(t, 5, mg.MaxTransitiveDepth)
	})

	t.Run("SetMaxTransitiveDepth zero uses default", func(t *testing.T) {
		mg.SetMaxTransitiveDepth(0)
		assert.Equal(t, modulegraph.DefaultMaxTransitiveDepth, mg.MaxTransitiveDepth)
	})

	t.Run("SetMaxTransitiveDepth negative uses default", func(t *testing.T) {
		mg.SetMaxTransitiveDepth(-1)
		assert.Equal(t, modulegraph.DefaultMaxTransitiveDepth, mg.MaxTransitiveDepth)
	})

	t.Run("ClearTransitiveElementsCache", func(t *testing.T) {
		mg.TransitiveElementsCache.Store("key", []string{"a"})
		mg.ClearTransitiveElementsCache()
		_, found := mg.TransitiveElementsCache.Load("key")
		assert.False(t, found)
	})

	t.Run("SetManifestResolver", func(t *testing.T) {
		mg.SetManifestResolver(nil)
	})
}

func TestModuleGraphWithMetrics(t *testing.T) {
	mg := modulegraph.NewModuleGraphWithMetrics(nil, modulegraph.NewDefaultMetricsCollector(), nil)
	assert.NotNil(t, mg)
	metrics := mg.GetMetrics()
	_, ok := metrics.(*modulegraph.DefaultMetricsCollector)
	assert.True(t, ok, "should use DefaultMetricsCollector")
}

func TestExportTracker_GetAllModulePaths(t *testing.T) {
	et := modulegraph.NewExportTracker()
	et.AddDirectExport("a.js", "Foo", "x-foo")
	et.AddDirectExport("b.js", "Bar", "x-bar")

	paths := et.GetAllModulePaths()
	assert.Len(t, paths, 2)
	assert.Contains(t, paths, "a.js")
	assert.Contains(t, paths, "b.js")
}

func TestExportTracker_GetModuleExports(t *testing.T) {
	et := modulegraph.NewExportTracker()
	et.AddDirectExport("a.js", "Foo", "x-foo")

	exports := et.GetModuleExports("a.js")
	assert.Len(t, exports, 1)
	assert.Equal(t, "Foo", exports[0].ElementName)

	exports[0].ElementName = "mutated"
	assert.Equal(t, "Foo", et.GetModuleExports("a.js")[0].ElementName, "returns copy")

	assert.Nil(t, et.GetModuleExports("nonexistent.js"))
}

func TestExportTracker_GetAllTagNames(t *testing.T) {
	et := modulegraph.NewExportTracker()
	et.AddDirectExport("a.js", "Foo", "x-foo")
	et.AddDirectExport("b.js", "Bar", "x-bar")

	tags := et.GetAllTagNames()
	assert.Len(t, tags, 2)
	assert.Contains(t, tags, "x-foo")
	assert.Contains(t, tags, "x-bar")
}

func TestExportTracker_GetElementSources(t *testing.T) {
	et := modulegraph.NewExportTracker()
	et.AddElementSource("x-foo", "a.js")
	et.AddElementSource("x-foo", "b.js")

	sources := et.GetElementSources("x-foo")
	assert.Len(t, sources, 2)
	assert.Nil(t, et.GetElementSources("nonexistent"))
}

func TestExportTracker_AddReExport(t *testing.T) {
	et := modulegraph.NewExportTracker()
	et.AddReExport("index.js", "button.js", "Button", "x-button")

	exports := et.GetModuleExports("index.js")
	assert.Len(t, exports, 1)
	assert.Equal(t, "button.js", exports[0].SourceModule)
}

type mockManifestResolver struct {
	elements map[string][]string
}

func (m *mockManifestResolver) FindManifestModulesForImportPath(importPath string) []string {
	return nil
}

func (m *mockManifestResolver) GetManifestModulePath(filePath string) string {
	return ""
}

func (m *mockManifestResolver) GetElementsFromManifestModule(manifestModulePath string) []string {
	return m.elements[manifestModulePath]
}

func TestGetTransitiveElements(t *testing.T) {
	t.Run("empty path returns nil", func(t *testing.T) {
		mg := modulegraph.NewModuleGraph(nil, nil)
		assert.Nil(t, mg.GetTransitiveElements(""))
	})

	t.Run("single module with resolver", func(t *testing.T) {
		mg := modulegraph.NewModuleGraphWithMetrics(nil, modulegraph.NewDefaultMetricsCollector(), nil)
		mg.SetManifestResolver(&mockManifestResolver{
			elements: map[string][]string{
				"button.js": {"x-button"},
			},
		})
		mg.AddModuleDependency("index.js", "button.js")

		result := mg.GetTransitiveElements("index.js")
		assert.Contains(t, result, "x-button")
	})

	t.Run("transitive chain", func(t *testing.T) {
		mg := modulegraph.NewModuleGraphWithMetrics(nil, modulegraph.NewDefaultMetricsCollector(), nil)
		mg.SetManifestResolver(&mockManifestResolver{
			elements: map[string][]string{
				"tab.js":  {"x-tab"},
				"icon.js": {"x-icon"},
			},
		})
		mg.AddModuleDependency("index.js", "tabs.js")
		mg.AddModuleDependency("tabs.js", "tab.js")
		mg.AddModuleDependency("tabs.js", "icon.js")

		result := mg.GetTransitiveElements("index.js")
		assert.Contains(t, result, "x-tab")
		assert.Contains(t, result, "x-icon")
	})

	t.Run("caches results", func(t *testing.T) {
		mg := modulegraph.NewModuleGraphWithMetrics(nil, modulegraph.NewDefaultMetricsCollector(), nil)
		mg.SetManifestResolver(&mockManifestResolver{
			elements: map[string][]string{
				"a.js": {"x-a"},
			},
		})
		mg.AddModuleDependency("index.js", "a.js")

		r1 := mg.GetTransitiveElements("index.js")
		r2 := mg.GetTransitiveElements("index.js")
		assert.Equal(t, r1, r2)

		metrics := mg.GetMetrics().(*modulegraph.DefaultMetricsCollector)
		assert.Equal(t, int64(1), metrics.GetCounterValue("cache_misses"))
		assert.True(t, metrics.GetCounterValue("cache_hits") >= 1)
	})

	t.Run("no manifest resolver returns empty", func(t *testing.T) {
		mg := modulegraph.NewModuleGraph(nil, nil)
		result := mg.GetTransitiveElements("index.js")
		assert.Empty(t, result)
	})

	t.Run("circular dependency handled", func(t *testing.T) {
		mg := modulegraph.NewModuleGraphWithMetrics(nil, modulegraph.NewDefaultMetricsCollector(), nil)
		mg.SetManifestResolver(&mockManifestResolver{
			elements: map[string][]string{
				"a.js": {"x-a"},
				"b.js": {"x-b"},
			},
		})
		mg.AddModuleDependency("a.js", "b.js")
		mg.AddModuleDependency("b.js", "a.js")

		result := mg.GetTransitiveElements("a.js")
		assert.Contains(t, result, "x-a")
		assert.Contains(t, result, "x-b")
	})
}
