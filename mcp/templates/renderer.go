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
package templates

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
	"sync"

	"bennypowers.dev/cem/mcp/helpers"
	"bennypowers.dev/cem/mcp/security"
)

// TemplatePool provides thread-safe template rendering with instance pooling
type TemplatePool struct {
	funcMap        template.FuncMap
	securityPolicy security.SecurityPolicy
	pools          map[string]*sync.Pool
	poolsMu        sync.RWMutex
	templateSources map[string]*embed.FS
	templateSourcesMu sync.RWMutex
}

// NewTemplatePool creates a new thread-safe template pool
func NewTemplatePool() *TemplatePool {
	return NewSecureTemplatePool(security.DefaultSecurityPolicy())
}

// NewSecureTemplatePool creates a new template pool with security controls
func NewSecureTemplatePool(policy security.SecurityPolicy) *TemplatePool {
	return &TemplatePool{
		securityPolicy: policy,
		funcMap:        createSecureFuncMap(),
		pools:          make(map[string]*sync.Pool),
		templateSources: make(map[string]*embed.FS),
	}
}

// RegisterTemplateSource registers an embedded filesystem for a specific package
func (tp *TemplatePool) RegisterTemplateSource(packageName string, fs *embed.FS) {
	tp.templateSourcesMu.Lock()
	defer tp.templateSourcesMu.Unlock()
	tp.templateSources[packageName] = fs
}

// getTemplatePool returns or creates a sync.Pool for the given template name
func (tp *TemplatePool) getTemplatePool(templateName string) *sync.Pool {
	tp.poolsMu.RLock()
	pool, exists := tp.pools[templateName]
	tp.poolsMu.RUnlock()
	
	if exists {
		return pool
	}
	
	tp.poolsMu.Lock()
	defer tp.poolsMu.Unlock()
	
	// Double-check after acquiring write lock
	if pool, exists := tp.pools[templateName]; exists {
		return pool
	}
	
	// Create new pool for this template
	pool = &sync.Pool{
		New: func() interface{} {
			return tp.createTemplate(templateName)
		},
	}
	tp.pools[templateName] = pool
	return pool
}

// createTemplate creates a new template instance for the pool
func (tp *TemplatePool) createTemplate(templateName string) *template.Template {
	// Try to load template from registered sources
	tp.templateSourcesMu.RLock()
	defer tp.templateSourcesMu.RUnlock()
	
	var content []byte
	var err error
	
	// Try each registered template source
	for _, fs := range tp.templateSources {
		templatePath := filepath.Join("templates", templateName+".md")
		content, err = fs.ReadFile(templatePath)
		if err == nil {
			break
		}
	}
	
	if err != nil {
		// Return nil to indicate template not found - this will cause Execute to fail
		return nil
	}
	
	// Create and parse template with function map
	tmpl := template.New(templateName).Funcs(tp.funcMap)
	tmpl, err = tmpl.Parse(string(content))
	if err != nil {
		// Return nil to indicate parse error - this will cause Execute to fail
		return nil
	}
	
	return tmpl
}

// Render renders a template using the thread-safe pool
func (tp *TemplatePool) Render(templateName string, data interface{}) (string, error) {
	// Security: Validate template name to prevent path traversal
	if strings.Contains(templateName, "..") || strings.Contains(templateName, "/") {
		return "", fmt.Errorf("invalid template name: %s", templateName)
	}
	
	// Get template instance from pool
	pool := tp.getTemplatePool(templateName)
	tmplInterface := pool.Get()
	defer pool.Put(tmplInterface)
	
	// Check if template creation failed
	tmpl, ok := tmplInterface.(*template.Template)
	if !ok || tmpl == nil {
		return "", fmt.Errorf("template not found: %s", templateName)
	}
	
	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}
	
	return buf.String(), nil
}

// createSecureFuncMap creates a restricted function map for template security
func createSecureFuncMap() template.FuncMap {
	return template.FuncMap{
		"title": helpers.TitleCaser.String,
		"len": func(slice interface{}) int {
			switch s := slice.(type) {
			case []string:
				return len(s)
			// Map types for guidelines templates
			case map[string]interface{}:
				return len(s)
			case map[string]string:
				return len(s)
			case map[string][]string:
				return len(s)
			default:
				// Use reflection for unknown slice/map types
				if s != nil {
					if v, ok := s.(interface{ Len() int }); ok {
						return v.Len()
					}
				}
				return 0
			}
		},
		"index": func(slice interface{}, i int) interface{} {
			switch s := slice.(type) {
			case []string:
				if i >= 0 && i < len(s) {
					return s[i]
				}
			}
			return ""
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"join": func(slice []string, sep string) string {
			return strings.Join(slice, sep)
		},
		"add": func(a, b int) int {
			return a + b
		},
	}
}

// Global thread-safe template pool instance
var globalTemplatePool = NewTemplatePool()

// RegisterTemplateSource registers templates for a package with the global pool
func RegisterTemplateSource(packageName string, fs *embed.FS) {
	globalTemplatePool.RegisterTemplateSource(packageName, fs)
}

// RenderTemplate renders a template using the global thread-safe pool
func RenderTemplate(templateName string, data interface{}) (string, error) {
	return globalTemplatePool.Render(templateName, data)
}