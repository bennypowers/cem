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

package routes_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/health"
	"bennypowers.dev/cem/internal/platform"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/routes"
)

type frontmatterTestContext struct {
	manifest   []byte
	watchDir   string
	demoRoutes map[string]*middleware.DemoRouteEntry
	fs         platform.FileSystem
}

func (c *frontmatterTestContext) WatchDir() string                                  { return c.watchDir }
func (c *frontmatterTestContext) IsWorkspace() bool                                 { return false }
func (c *frontmatterTestContext) WorkspacePackages() []middleware.WorkspacePackage   { return nil }
func (c *frontmatterTestContext) Manifest() ([]byte, error)                         { return c.manifest, nil }
func (c *frontmatterTestContext) ImportMap() middleware.ImportMap                    { return nil }
func (c *frontmatterTestContext) DemoRoutes() map[string]*middleware.DemoRouteEntry  { return c.demoRoutes }
func (c *frontmatterTestContext) SourceControlRootURL() string                      { return "" }
func (c *frontmatterTestContext) Logger() logger.Logger                             { return logger.NewDefaultLogger() }
func (c *frontmatterTestContext) FileSystem() platform.FileSystem                   { return c.fs }
func (c *frontmatterTestContext) PackageJSON() (*middleware.PackageJSON, error)      { return nil, nil }
func (c *frontmatterTestContext) BroadcastError(title, message, file string) error   { return nil }
func (c *frontmatterTestContext) DemoRenderingMode() string                         { return "light" }
func (c *frontmatterTestContext) URLRewrites() []config.URLRewrite                  { return nil }
func (c *frontmatterTestContext) PathResolver() middleware.PathResolver              { return nil }
func (c *frontmatterTestContext) HealthResult() (*health.HealthResult, error)        { return nil, nil }
func (c *frontmatterTestContext) IsStaticBuild() bool                               { return false }

func TestFrontmatterStrippedFromDemoHTTPResponse(t *testing.T) {
	demoContent := []byte("---\ndescription: Visible frontmatter bug\n---\n<test-el>content</test-el>\n")

	manifest := map[string]any{
		"schemaVersion": "1.0.0",
		"modules": []map[string]any{{
			"kind": "javascript-module",
			"path": "src/test-el.ts",
			"declarations": []map[string]any{{
				"kind":          "class",
				"customElement": true,
				"name":          "TestEl",
				"tagName":       "test-el",
				"demos": []map[string]any{{
					"description": "Frontmatter demo",
					"url":         "./demo/frontmatter.html",
				}},
			}},
		}},
	}
	manifestBytes, err := json.Marshal(manifest)
	if err != nil {
		t.Fatalf("marshal manifest: %v", err)
	}

	watchDir := t.TempDir()
	fs := platform.NewOSFileSystem()
	demoDir := watchDir + "/demo"
	if err := fs.MkdirAll(demoDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := fs.WriteFile(demoDir+"/frontmatter.html", demoContent, 0644); err != nil {
		t.Fatalf("write demo: %v", err)
	}

	ctx := &frontmatterTestContext{
		manifest: manifestBytes,
		watchDir: watchDir,
		fs:       fs,
		demoRoutes: map[string]*middleware.DemoRouteEntry{
			"/demo/frontmatter.html": {
				TagName:  "test-el",
				FilePath: "demo/frontmatter.html",
				Demo: &M.Demo{
					Description: "Frontmatter demo",
					URL:         "./demo/frontmatter.html",
				},
				LocalRoute: "/demo/frontmatter.html",
			},
		},
	}

	mw := routes.New(routes.Config{Context: ctx})
	handler := mw(http.NotFoundHandler())

	req := httptest.NewRequest("GET", "/demo/frontmatter.html", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// Frontmatter must not appear in rendered output
	if strings.Contains(body, "Visible frontmatter bug") {
		t.Error("frontmatter description leaked into HTTP response body")
	}

	// Demo content must be present
	if !strings.Contains(body, "<test-el>content</test-el>") {
		t.Error("demo element content missing from HTTP response")
	}
}
