package generate

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"maps"
	"os"
	"runtime"
	"slices"
	"sync"

	C "bennypowers.dev/cem/cmd/config"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	DT "bennypowers.dev/cem/designtokens"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/generate/queries"

	DS "github.com/bmatcuk/doublestar"
)

var defaultExcludePatterns = []string{
	"**/*.d.ts",
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

type preprocessResult struct {
	demoFiles []string
	designTokens *DT.DesignTokens
	excludePatterns []string
}

func preprocess(cfg *C.CemConfig) (r preprocessResult, errs error) {
	r.excludePatterns = make([]string, 0, len(cfg.Generate.Exclude)+len(defaultExcludePatterns))
	r.excludePatterns = append(r.excludePatterns, cfg.Generate.Exclude...)
	if !cfg.Generate.NoDefaultExcludes {
		r.excludePatterns = append(r.excludePatterns, defaultExcludePatterns...)
	}
	if cfg.Generate.DesignTokensSpec != "" {
		tokens, err := DT.LoadDesignTokens(cfg.Generate.DesignTokensSpec, cfg.Generate.DesignTokensPrefix)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		r.designTokens = tokens
	}
	if cfg.Generate.DemoDiscovery.FilePattern != "" {
		r.demoFiles, _ = DS.Glob(cfg.Generate.DemoDiscovery.FilePattern)
	}
	return r, errs
}

func process(
	cfg *C.CemConfig,
	result preprocessResult,
	qm *Q.QueryManager,
) (modules []M.Module, aliases map[string]string, errs error) {
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	modulesChan := make(chan *M.Module, len(cfg.Generate.Files))
	jobsChan := make (chan string, len(cfg.Generate.Files))
	// Fill jobs channel with files to process
	for _, file := range cfg.Generate.Files {
		if !matchesAnyPattern(file, result.excludePatterns) {
			jobsChan <- file
		}
	}
	close(jobsChan)

	for range numWorkers {
		go func() {
			defer wg.Done()
			for file := range jobsChan {
				code, err := os.ReadFile(file)
				if err != nil {
					errs = errors.Join(errs, errors.New(fmt.Sprintf("Could not Read file %s", file)), err)
					continue
				}

				mp := NewModuleProcessor(file, code, qm)
				module, tagAliases, err := mp.Collect()
				maps.Copy(aliases, tagAliases)
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

	for module := range modulesChan {
		modules = append(modules, *module)
	}

	return modules, aliases, errs
}

func postprocess(
	cfg *C.CemConfig,
	result preprocessResult,
	allTagAliases map[string]string,
	qm *Q.QueryManager,
	modules []M.Module,
) (pkg M.Package, errs error) {
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	modulesChan := make(chan *M.Module, len(modules))
	for _, module := range modules {
		go func() {
			if result.designTokens != nil {
				DT.MergeDesignTokensToModule(&module, *result.designTokens)
			}
			// Discover demos and attach to manifest
			if len(result.demoFiles) > 0 {
				err := DD.DiscoverDemos(cfg, allTagAliases, &module, qm, result.demoFiles)
				if err != nil {
					errs = errors.Join(errs, err)
				}
			}
		}()
	}
	wg.Wait()
	close(modulesChan)
	pkg = M.NewPackage(modules)
	slices.SortStableFunc(pkg.Modules, func(a, b M.Module) int {
		return cmp.Compare(a.Path, b.Path)
	})
	return pkg, errs
}

// Generates a custom-elements manifest from a list of typescript files
func Generate(cfg *C.CemConfig) (manifest *string, errs error) {
	qm, err := Q.NewQueryManager()
	defer qm.Close()
	if err != nil {
		log.Fatal(err)
	}
	result, err := preprocess(cfg)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	modules, aliases, err := process(cfg, result, qm)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	pkg, err := postprocess(cfg, result, aliases, qm, modules)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	*manifest, err = M.SerializeToString(&pkg)
	if err != nil {
		return nil, errors.Join(errs, err)
	}
	return manifest, errs
}
