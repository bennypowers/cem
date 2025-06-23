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

	M "bennypowers.dev/cem/manifest"
	DT "bennypowers.dev/cem/designtokens"
	DS "github.com/bmatcuk/doublestar"
)

var defaultExcludePatterns = []string{
	"**/*.d.ts",
}

// CLI or config arguments passed to the generate command
type GenerateArgs struct {
	// List of files or file globs to include in the manifest
	Files              []string
	// List of files or file globs to exclude from the manifest
	Exclude            []string
	// Path or `npm:@scope/package/path/to/file.json` spec to DTCG format design
	// tokens json module
	DesignTokensSpec   string
	// Prefix those design tokens use in CSS. If the design tokens are generated
	// by style dictionary and have a `name` field, that will be used instead.
	DesignTokensPrefix string
	// Do not exclude files that are excluded by default e.g. *.d.ts files.
	NoDefaultExcludes  bool
	// File path to write output to. If omitted, output will be written to stdout.
	Output             string
}

func mergeCssPropertyInfoFromDesignTokens(module *M.Module, designTokens DT.DesignTokens) {
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

// Checks whether a file matches any of the given patterns (with doublestar)
func matchesAnyPattern(file string, patterns []string) bool {
	for _, pattern := range patterns {
		match, err := DS.PathMatch(pattern, file)
		if err == nil && match {
			return true
		}
	}
	return false
}

// Generates a custom-elements manifest from a list of typescript files
func Generate(args GenerateArgs) (*string, error) {
	excludePatterns := make([]string, 0, len(args.Exclude)+len(defaultExcludePatterns))
	excludePatterns = append(excludePatterns, args.Exclude...)
	if !args.NoDefaultExcludes {
		excludePatterns = append(excludePatterns, defaultExcludePatterns...)
	}

	var wg sync.WaitGroup
	var errs error
	var designTokens *DT.DesignTokens

	if args.DesignTokensSpec != "" {
		tokens, err := DT.LoadDesignTokens(args.DesignTokensSpec, args.DesignTokensPrefix)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		designTokens = tokens
	}

	modulesChan := make(chan *M.Module, len(args.Files))
	jobsChan := make (chan string, len(args.Files))
	// Fill jobs channel with files to process
	for _, file := range args.Files {
		if !matchesAnyPattern(file, excludePatterns) {
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

				mp := NewModuleProcessor(file, code, queryManager)
				module, err := mp.Collect()
				mp.Close()
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
