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
package queries

import (
	"errors"
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/pterm/pterm"
	ts "github.com/tree-sitter/go-tree-sitter"
)

var ErrNoQueryManager = errors.New("QueryManager is nil")

type QueryManagerI interface {
	Close()
	getQuery(name string) (*ts.Query, error)
}

type QueryManager struct {
	typescript map[string]*ts.Query
	jsdoc      map[string]*ts.Query
	css        map[string]*ts.Query
	html       map[string]*ts.Query
	tsx        map[string]*ts.Query
	blade      map[string]*ts.Query
}

func NewQueryManager(selector QuerySelector) (*QueryManager, error) {
	start := time.Now()
	qm := &QueryManager{
		typescript: make(map[string]*ts.Query),
		jsdoc:      make(map[string]*ts.Query),
		css:        make(map[string]*ts.Query),
		html:       make(map[string]*ts.Query),
		tsx:        make(map[string]*ts.Query),
		blade:      make(map[string]*ts.Query),
	}

	// Load only the requested queries
	for _, queryName := range selector.HTML {
		if err := qm.loadQuery("html", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load HTML query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.TypeScript {
		if err := qm.loadQuery("typescript", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load TypeScript query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.CSS {
		if err := qm.loadQuery("css", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load CSS query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.JSDoc {
		if err := qm.loadQuery("jsdoc", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load JSDoc query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.TSX {
		if err := qm.loadQuery("tsx", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load TSX query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.Blade {
		if err := qm.loadQuery("blade", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load Blade query %s: %w", queryName, err)
		}
	}

	pterm.Debug.Println("Constructing selected queries took", time.Since(start))
	return qm, nil
}

func (qm *QueryManager) loadQuery(language, queryName string) error {
	// Blade reuses html/*.scm files compiled against the Blade grammar
	sourceDir := language
	if language == "blade" {
		sourceDir = "html"
	}
	// Use path.Join (not filepath.Join) - embed.FS requires POSIX / separators
	queryPath := path.Join(sourceDir, queryName+".scm")
	data, err := queries.ReadFile(queryPath)
	if err != nil {
		return fmt.Errorf("failed to read query file %s: %w", queryPath, err)
	}

	queryText := string(data)
	var tsLang *ts.Language
	switch language {
	case "typescript":
		tsLang = languages.typescript
	case "jsdoc":
		tsLang = languages.jsdoc
	case "css":
		tsLang = languages.css
	case "html":
		tsLang = languages.html
	case "tsx":
		tsLang = languages.tsx
	case "blade":
		tsLang = languages.blade
	default:
		return fmt.Errorf("unknown language %s", language)
	}

	query, qerr := ts.NewQuery(tsLang, queryText)
	if qerr != nil {
		return fmt.Errorf("failed to parse query %s: %w", queryName, qerr)
	}

	switch language {
	case "typescript":
		qm.typescript[queryName] = query
	case "jsdoc":
		qm.jsdoc[queryName] = query
	case "css":
		qm.css[queryName] = query
	case "html":
		qm.html[queryName] = query
	case "tsx":
		qm.tsx[queryName] = query
	case "blade":
		qm.blade[queryName] = query
	}

	return nil
}

func (qm *QueryManager) Close() {
	for _, query := range qm.typescript {
		query.Close()
	}
	for _, query := range qm.jsdoc {
		query.Close()
	}
	for _, query := range qm.css {
		query.Close()
	}
	for _, query := range qm.html {
		query.Close()
	}
	for _, query := range qm.tsx {
		query.Close()
	}
	for _, query := range qm.blade {
		query.Close()
	}
}

func (qm *QueryManager) getQuery(queryName string, language string) (*ts.Query, error) {
	var q *ts.Query
	var ok bool
	switch language {
	case "typescript":
		q, ok = qm.typescript[queryName]
	case "jsdoc":
		q, ok = qm.jsdoc[queryName]
	case "css":
		q, ok = qm.css[queryName]
	case "html":
		q, ok = qm.html[queryName]
	case "tsx":
		q, ok = qm.tsx[queryName]
	case "blade":
		q, ok = qm.blade[queryName]
	}
	if !ok {
		return nil, fmt.Errorf("unknown query %s", queryName)
	}
	return q, nil
}

// Thread-safe singleton QueryManager (there can be only one!)
var (
	globalQueryManager *QueryManager
	globalQueryOnce    sync.Once
	globalQueryError   error
)

// GetGlobalQueryManager returns the singleton QueryManager instance
func GetGlobalQueryManager() (*QueryManager, error) {
	globalQueryOnce.Do(func() {
		manager, err := NewQueryManager(LSPQueries())
		if err != nil {
			globalQueryError = err
			return
		}
		globalQueryManager = manager
	})

	if globalQueryError != nil {
		return nil, globalQueryError
	}

	if globalQueryManager == nil {
		return nil, fmt.Errorf("failed to initialize global query manager")
	}

	return globalQueryManager, nil
}
