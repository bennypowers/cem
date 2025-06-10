package generate

import (
	"cmp"
	"fmt"
	"os"
	"runtime"
	"slices"
	"sync"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/set"
)

func processFile(file string, channel chan<-manifest.Module, wg *sync.WaitGroup) {
	defer wg.Done()

	code, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return
	}

	err, module := generateModule(file, code)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	if module != nil {
		channel <- *module
	}
}

// Generates a custom-elements manifest from a list of typescript files
func Generate(files []string, exclude []string) (error, *string) {
	excludeSet := set.NewSet(exclude...)

	var wg sync.WaitGroup

	modulesChan := make(chan manifest.Module, len(files))
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

	for range numWorkers {
		go func() {
			defer wg.Done()
			for file := range jobsChan {
				code, err := os.ReadFile(file)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					continue
				}

				err, module := generateModule(file, code)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				}
				if module != nil {
					modulesChan <- *module
				}
			}
		}()
	}

	wg.Wait()
	close(modulesChan)

	modules := make([]manifest.Module, 0)
	for module := range modulesChan {
		modules = append(modules, module)
	}

	pkg := manifest.NewPackage(modules)

	slices.SortFunc(pkg.Modules, func(a, b manifest.Module) int {
		return cmp.Compare(a.Path, b.Path)
	})

	manifest, err := manifest.SerializeToString(&pkg)

	if err != nil {
		return err, nil
	}

	return nil, &manifest
}
