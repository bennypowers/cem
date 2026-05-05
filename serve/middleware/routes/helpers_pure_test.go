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
package routes

import (
	"testing"

	"bennypowers.dev/cem/health"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/serve/middleware"
	"github.com/stretchr/testify/assert"
)

func TestFilterHealthByComponent(t *testing.T) {
	result := &health.HealthResult{
		Modules: []health.ModuleReport{
			{
				Path: "src/button.js",
				Declarations: []health.ComponentReport{
					{TagName: "my-button", Name: "MyButton", Score: 80, MaxScore: 100},
					{TagName: "my-icon", Name: "MyIcon", Score: 60, MaxScore: 100},
				},
			},
		},
		Recommendations: []string{
			"my-button: add description (+5 pts)",
			"my-icon: add summary (+3 pts)",
		},
	}

	t.Run("filters by tag name", func(t *testing.T) {
		filtered := filterHealthByComponent(result, "my-button")
		assert.Len(t, filtered.Modules, 1)
		assert.Len(t, filtered.Modules[0].Declarations, 1)
		assert.Equal(t, "my-button", filtered.Modules[0].Declarations[0].TagName)
		assert.Equal(t, 80, filtered.OverallScore)
		assert.Equal(t, 100, filtered.OverallMax)
	})

	t.Run("filters recommendations", func(t *testing.T) {
		filtered := filterHealthByComponent(result, "my-button")
		assert.Len(t, filtered.Recommendations, 1)
		assert.Contains(t, filtered.Recommendations[0], "my-button")
	})

	t.Run("no match returns empty", func(t *testing.T) {
		filtered := filterHealthByComponent(result, "nonexistent")
		assert.Empty(t, filtered.Modules)
	})
}

func TestMarkdown(t *testing.T) {
	t.Run("converts markdown to HTML", func(t *testing.T) {
		result := markdown("**bold** text")
		assert.Contains(t, string(result), "<strong>bold</strong>")
	})

	t.Run("empty string", func(t *testing.T) {
		result := markdown("")
		assert.Empty(t, string(result))
	})
}

func TestPackageHasDemo(t *testing.T) {
	pkg := PackageNavigation{
		Name: "my-pkg",
		Elements: []ElementListing{
			{TagName: "my-button", Demos: []DemoListing{{URL: "/demo/button/"}}},
			{TagName: "my-card", Demos: []DemoListing{{URL: "/demo/card/"}}},
		},
	}

	assert.True(t, packageHasDemo(pkg, "/demo/button/"))
	assert.True(t, packageHasDemo(pkg, "/demo/card/"))
	assert.False(t, packageHasDemo(pkg, "/demo/icon/"))
}

func TestElementHasDemo(t *testing.T) {
	el := ElementListing{
		TagName: "my-button",
		Demos:   []DemoListing{{URL: "/demo/button/"}},
	}

	assert.True(t, elementHasDemo(el, "/demo/button/"))
	assert.False(t, elementHasDemo(el, "/demo/card/"))
}

func TestBuildPackageListingsFromRoutes(t *testing.T) {
	routes := map[string]*DemoRouteEntry{
		"/demo/button/": {
			TagName:     "my-button",
			PackageName: "@my/elements",
			Demo:        &M.Demo{URL: "/demo/button/", Description: "Button demo"},
			Declaration: &M.CustomElementDeclaration{},
		},
		"/demo/card/": {
			TagName:     "my-card",
			PackageName: "@my/elements",
			Demo:        &M.Demo{URL: "/demo/card/", Description: "Card demo"},
			Declaration: &M.CustomElementDeclaration{},
		},
	}

	pkgs, err := buildPackageListingsFromRoutes(routes)
	assert.NoError(t, err)
	assert.Len(t, pkgs, 1)
	assert.Equal(t, "@my/elements", pkgs[0].Name)
	assert.Len(t, pkgs[0].Elements, 2)
}

func TestResolveViaRoutingTable(t *testing.T) {
	demoRoutes := map[string]*middleware.DemoRouteEntry{
		"/elements/button/demo/": {
			FilePath:    "elements/button/demo/index.html",
			PackageName: "my-pkg",
		},
	}
	ctx := &mockContext{watchDir: "/project", demoRoutes: demoRoutes}
	cfg := Config{Context: ctx}

	t.Run("matches prefix", func(t *testing.T) {
		result := resolveViaRoutingTable("/elements/button/button.css", cfg)
		assert.Contains(t, result, "button.css")
	})

	t.Run("no match", func(t *testing.T) {
		result := resolveViaRoutingTable("/elements/card/card.css", cfg)
		assert.Empty(t, result)
	})

	t.Run("nil routes", func(t *testing.T) {
		nilCtx := &mockContext{watchDir: "/project"}
		result := resolveViaRoutingTable("/anything", Config{Context: nilCtx})
		assert.Empty(t, result)
	})
}

func TestResolveViaDemoSubresource_EdgeCases(t *testing.T) {
	demoRoutes := map[string]*middleware.DemoRouteEntry{
		"/elements/avatar/demo/sizes/": {
			FilePath:    "elements/avatar/demo/sizes.html",
			PackageName: "my-pkg",
		},
	}
	ctx := &mockContext{watchDir: "/project", demoRoutes: demoRoutes}
	cfg := Config{Context: ctx}

	t.Run("resolves subresource", func(t *testing.T) {
		result := resolveViaDemoSubresource("/elements/avatar/demo/sizes/image.jpg", cfg)
		assert.Contains(t, result, "image.jpg")
	})

	t.Run("no match", func(t *testing.T) {
		result := resolveViaDemoSubresource("/unrelated/path", cfg)
		assert.Empty(t, result)
	})

	t.Run("path traversal rejected", func(t *testing.T) {
		result := resolveViaDemoSubresource("/elements/avatar/demo/sizes/../../../etc/passwd", cfg)
		assert.Empty(t, result)
	})

	t.Run("nil routes", func(t *testing.T) {
		nilCtx := &mockContext{watchDir: "/project"}
		result := resolveViaDemoSubresource("/anything", Config{Context: nilCtx})
		assert.Empty(t, result)
	})
}
