package generate

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
	"sync"

	"bennypowers.dev/cem/designtokens"
	M "bennypowers.dev/cem/manifest"
	S "bennypowers.dev/cem/set"
)

func mergeCssPropertyInfoFromDesignTokens(module *M.Module, designTokens designtokens.DesignTokens) {
	for i, d := range module.Declarations {
		if d, ok := d.(*M.CustomElementDeclaration); ok {
			for i, p := range d.CssProperties {
				if token, ok := designTokens.Get(p.Name); ok {
					p.Description = token.Description
					p.Syntax = token.Syntax
					d.CssProperties[i] = p
				}
			}
		}
		module.Declarations[i] = d
	}
}

// Generates a custom-elements manifest from a list of typescript files
func Generate(
	files []string,
	exclude []string,
	designTokensSpec string,
	designTokensPrefix string,
) (*string, error) {
	excludeSet := S.NewSet(exclude...)

	var wg sync.WaitGroup
	var errs error
	var designTokens *designtokens.DesignTokens

	if designTokensSpec != "" {
		tokens, err := designtokens.LoadDesignTokens(designTokensSpec, designTokensPrefix)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		designTokens = tokens
	}

	modulesChan := make(chan *M.Module, len(files))
	jobsChan := make (chan string, len(files))
	// Fill jobs channel with files to process
	for _, file := range files {
		if !excludeSet.Has(file) {
			jobsChan <- file
		}
	}
	close(jobsChan)

	numWorkers := runtime.NumCPU()
	wg.Add(numWorkers)

	queryManager, err := NewQueryManager()
	if err != nil {
		log.Fatal(err)
	}

	for range numWorkers {
		go func() {
			defer wg.Done()
			for file := range jobsChan {
				code, err := os.ReadFile(file)
				if err != nil {
					errs = errors.Join(errs, errors.New(fmt.Sprintf("Could not Read file %s", file)), err)
					continue
				}

				err, module := generateModule(file, code, queryManager)
				if err != nil {
					errs = errors.Join(errs, errors.New(fmt.Sprintf("Problem generating module %s", file)), err)
					continue
				}
				if module != nil {
					modulesChan <- module
				}
			}
		}()
	}

	wg.Wait()
	close(modulesChan)

	modules := make([]M.Module, 0)
	for module := range modulesChan {
		if designTokens != nil {
			mergeCssPropertyInfoFromDesignTokens(module, *designTokens)
		}
		modules = append(modules, *module)
	}

	pkg := M.NewPackage(modules)

	slices.SortStableFunc(pkg.Modules, func(a, b M.Module) int {
		return cmp.Compare(a.Path, b.Path)
	})

	manifest, err := M.SerializeToString(&pkg)

	queryManager.Close()

	if err != nil {
		return nil, errors.Join(errs, err)
	}

	return &manifest, errs
}
