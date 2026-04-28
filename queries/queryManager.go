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
	"sync"
	"time"

	"bennypowers.dev/cem/internal/languages"
	"github.com/pterm/pterm"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// ErrNoQueryManager is returned when a nil QueryManager is passed.
var ErrNoQueryManager = errors.New("QueryManager is nil")

// QuerySelector maps language names to the query names to load.
type QuerySelector map[string][]string

// LSPQueries returns a selector loading queries needed for LSP operations.
// Each registered language provides its own LSP query list.
func LSPQueries() QuerySelector {
	return queriesForScope(languages.ScopeLSP)
}

// GenerateQueries returns a selector loading queries needed for generation.
func GenerateQueries() QuerySelector {
	return queriesForScope(languages.ScopeGenerate)
}

// AllQueries returns a selector loading all available queries.
func AllQueries() QuerySelector {
	return queriesForScope(languages.ScopeAll)
}

func queriesForScope(scope languages.Scope) QuerySelector {
	sel := QuerySelector{}
	for _, lang := range languages.All() {
		names := lang.QueryNames(scope)
		if len(names) > 0 {
			sel[lang.Name()] = names
		}
	}
	return sel
}

// QueryManager holds compiled tree-sitter queries organized by language.
type QueryManager struct {
	queries map[string]map[string]*ts.Query // language -> queryName -> query
}

// NewQueryManager creates a QueryManager and loads queries per the selector.
func NewQueryManager(selector QuerySelector) (*QueryManager, error) {
	start := time.Now()
	qm := &QueryManager{
		queries: make(map[string]map[string]*ts.Query),
	}

	for langName, queryNames := range selector {
		qm.queries[langName] = make(map[string]*ts.Query)
		for _, queryName := range queryNames {
			if _, loaded := qm.queries[langName][queryName]; loaded {
				continue
			}
			if err := qm.loadQuery(langName, queryName); err != nil {
				qm.Close()
				return nil, fmt.Errorf("failed to load %s query %s: %w", langName, queryName, err)
			}
		}
	}

	pterm.Debug.Println("Constructing selected queries took", time.Since(start))
	return qm, nil
}

func (qm *QueryManager) loadQuery(langName, queryName string) error {
	lang := languages.Get(langName)
	if lang == nil {
		return fmt.Errorf("unknown language %s", langName)
	}

	queryFile := "queries/" + queryName + ".scm"
	data, err := lang.QueryFS().ReadFile(queryFile)
	if err != nil {
		return fmt.Errorf("failed to read query file %s from %s: %w", queryFile, langName, err)
	}

	query, qerr := ts.NewQuery(lang.TSLanguage(), string(data))
	if qerr != nil {
		return fmt.Errorf("failed to parse query %s: %w", queryName, qerr)
	}

	qm.queries[langName][queryName] = query
	return nil
}

func (qm *QueryManager) getQuery(queryName string, language string) (*ts.Query, error) {
	if langQueries, ok := qm.queries[language]; ok {
		if q, ok := langQueries[queryName]; ok {
			return q, nil
		}
	}
	return nil, fmt.Errorf("unknown query %s for language %s", queryName, language)
}

// Close releases all compiled queries.
func (qm *QueryManager) Close() {
	for _, langQueries := range qm.queries {
		for _, query := range langQueries {
			query.Close()
		}
	}
}

// Thread-safe singleton QueryManager (there can be only one!)
var (
	globalQueryManager *QueryManager
	globalQueryOnce    sync.Once
	globalQueryError   error
)

// GetGlobalQueryManager returns the singleton QueryManager instance.
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
