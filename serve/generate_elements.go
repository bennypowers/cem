//go:build ignore

// generate_elements.go compiles LitElement TypeScript sources into JavaScript
// for embedding in the CEM binary.
//
// Usage: go run generate_elements.go
//
// This script:
//  1. Bundles lit and its sub-paths into templates/vendor/ with sourcemaps
//  2. Transpiles each element .ts file into templates/elements/ with sourcemaps,
//     rewriting bare lit imports to /__cem/vendor/ paths
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "generate_elements: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := vendorLit(); err != nil {
		return fmt.Errorf("vendor lit: %w", err)
	}
	if err := transpileElements(); err != nil {
		return fmt.Errorf("transpile elements: %w", err)
	}
	if err := bundleSSR(); err != nil {
		return fmt.Errorf("bundle SSR: %w", err)
	}
	return nil
}

// vendorLit bundles lit and its commonly used sub-paths into templates/vendor/.
// Each entry point gets its own file so components can import granularly.
func vendorLit() error {
	outdir := "middleware/routes/templates/vendor"
	if err := os.MkdirAll(outdir, 0o755); err != nil {
		return err
	}

	entries := []struct {
		entry  string
		outfile string
	}{
		{"lit", "lit.js"},
		{"lit/decorators/custom-element.js", "lit/decorators/custom-element.js"},
		{"lit/decorators/property.js", "lit/decorators/property.js"},
		{"lit/decorators/state.js", "lit/decorators/state.js"},
		{"lit/decorators/query.js", "lit/decorators/query.js"},
		{"lit/decorators/query-all.js", "lit/decorators/query-all.js"},
		{"lit/decorators/event-options.js", "lit/decorators/event-options.js"},
		{"lit/directives/class-map.js", "lit/directives/class-map.js"},
		{"lit/directives/style-map.js", "lit/directives/style-map.js"},
		{"lit/directives/if-defined.js", "lit/directives/if-defined.js"},
		{"lit/directives/repeat.js", "lit/directives/repeat.js"},
		{"lit/directives/live.js", "lit/directives/live.js"},
		{"@lit-labs/ssr-client/lit-element-hydrate-support.js", "lit-element-hydrate-support.js"},
		{"lit/directives/unsafe-html.js", "lit/directives/unsafe-html.js"},
	}

	for _, e := range entries {
		outpath := filepath.Join(outdir, e.outfile)
		if err := os.MkdirAll(filepath.Dir(outpath), 0o755); err != nil {
			return err
		}

		result := api.Build(api.BuildOptions{
			EntryPoints:       []string{e.entry},
			Bundle:            true,
			Format:            api.FormatESModule,
			Target:            api.ES2022,
			Outfile:           outpath,
			Sourcemap:         api.SourceMapLinked,
			MinifySyntax:      true,
			MinifyWhitespace:  false,
			Write:             true,
			NodePaths:         []string{"elements/node_modules"},
			LogLevel:          api.LogLevelWarning,
		})

		if len(result.Errors) > 0 {
			msgs := api.FormatMessages(result.Errors, api.FormatMessagesOptions{})
			return fmt.Errorf("bundling %s:\n%s", e.entry, strings.Join(msgs, "\n"))
		}
	}

	return nil
}

// litCSSPlugin returns an esbuild plugin that transforms CSS imports with
// `{ type: 'css' }` into CSSStyleSheet modules. It also strips the import
// attribute syntax that esbuild doesn't understand.
func litCSSPlugin() api.Plugin {
	return api.Plugin{
		Name: "lit-css",
		Setup: func(build api.PluginBuild) {
			// Resolve .css imports from our elements to our custom namespace
			build.OnResolve(api.OnResolveOptions{Filter: `\.css$`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				// Only handle CSS from our element sources, not from node_modules
				if strings.Contains(args.Importer, "node_modules") {
					return api.OnResolveResult{}, nil
				}
				absPath := filepath.Join(filepath.Dir(args.Importer), args.Path)
				return api.OnResolveResult{
					Path:      absPath,
					Namespace: "lit-css",
				}, nil
			})

			// Load .css files as JS modules that create CSSStyleSheet
			build.OnLoad(api.OnLoadOptions{Filter: `\.css$`, Namespace: "lit-css"}, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
				css, err := os.ReadFile(args.Path)
				if err != nil {
					return api.OnLoadResult{}, err
				}

				// Also copy the raw CSS to the output directory.
				// The CSS path may be under elements/ directly or via a .tmp.ts redirect.
				cssPath := args.Path
				rel, err := filepath.Rel("elements", cssPath)
				if err != nil {
					// Fallback: use basename in same structure
					rel = filepath.Join(filepath.Base(filepath.Dir(cssPath)), filepath.Base(cssPath))
				}
				outCSS := filepath.Join("middleware/routes/templates/elements", rel)
				if err := os.MkdirAll(filepath.Dir(outCSS), 0o755); err != nil {
					return api.OnLoadResult{}, err
				}
				if err := os.WriteFile(outCSS, css, 0o644); err != nil {
					return api.OnLoadResult{}, err
				}

				// Double-encode: JSON.parse needs a JSON string, which esbuild
				// will embed as a JS string literal. So we JSON-encode twice:
				// once for the CSS content, once for the JS string literal.
				jsonCSS, _ := json.Marshal(string(css))
				jsonWrapped, _ := json.Marshal(string(jsonCSS))
				js := fmt.Sprintf(
					"const s=new CSSStyleSheet();s.replaceSync(JSON.parse(%s));export default s;",
					string(jsonWrapped),
				)
				contents := js
				return api.OnLoadResult{
					Contents: &contents,
					Loader:   api.LoaderJS,
				}, nil
			})
		},
	}
}

// stripImportAttributes removes `with { type: 'css' }` from import statements
// since esbuild doesn't support import attributes natively.
func stripImportAttributes(source string) string {
	// Match: with { type: 'css' } or with { type: "css" }
	// This is intentionally simple since we control all source files
	result := source
	for _, pattern := range []string{
		` with { type: 'css' }`,
		` with { type: "css" }`,
	} {
		result = strings.ReplaceAll(result, pattern, "")
	}
	return result
}

// transpileElements finds all .ts files under elements/ and transpiles them
// to .js in templates/elements/. CSS imports with { type: 'css' } are handled
// by the lit-css esbuild plugin.
func transpileElements() error {
	entries, err := filepath.Glob("elements/*/*.ts")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		rel, _ := filepath.Rel("elements", entry)
		outfile := filepath.Join("middleware/routes/templates/elements", strings.TrimSuffix(rel, ".ts")+".js")

		if err := os.MkdirAll(filepath.Dir(outfile), 0o755); err != nil {
			return err
		}

		// Read source, strip import attributes esbuild can't handle
		source, err := os.ReadFile(entry)
		if err != nil {
			return fmt.Errorf("reading %s: %w", entry, err)
		}
		processed := stripImportAttributes(string(source))

		result := api.Build(api.BuildOptions{
			Stdin: &api.StdinOptions{
				Contents:   processed,
				Sourcefile: filepath.Base(entry), // e.g. "pf-v6-badge.ts" for sourcemaps
				Loader:     api.LoaderTS,
				ResolveDir: filepath.Dir(entry), // resolve imports relative to source dir
			},
			Outfile: outfile,
			Bundle:      true,
			Format:      api.FormatESModule,
			Target:      api.ES2022,
			Sourcemap:   api.SourceMapInline,
			Write:       false,
			TsconfigRaw: `{"compilerOptions":{}}`,
			LogLevel:    api.LogLevelWarning,
			Plugins:     []api.Plugin{litCSSPlugin()},
			// Mark everything except .css as external so we don't bundle JS deps
			External:    []string{"lit", "lit/*", "../*", "/__cem/*"},
		})

		if len(result.Errors) > 0 {
			msgs := api.FormatMessages(result.Errors, api.FormatMessagesOptions{})
			return fmt.Errorf("transpiling %s:\n%s", entry, strings.Join(msgs, "\n"))
		}

		for _, f := range result.OutputFiles {
			content := string(f.Contents)
			if strings.HasSuffix(f.Path, ".js") {
				content = rewriteLitImports(content)
			}
			if err := os.WriteFile(f.Path, []byte(content), 0o644); err != nil {
				return err
			}
		}
	}

	return nil
}

// stubExternals returns an esbuild plugin that replaces /__cem/* imports
// with empty modules. These are client-only runtime modules (WebSocket,
// state persistence) that aren't needed for SSR rendering.
func stubExternals() api.Plugin {
	return api.Plugin{
		Name: "stub-externals",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: `^/__cem/`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{
					Path:      args.Path,
					Namespace: "stub",
				}, nil
			})
			build.OnLoad(api.OnLoadOptions{Filter: `.*`, Namespace: "stub"}, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
				var contents string
				if strings.Contains(args.Path, "/__cem/") {
					contents = "export default {}; export class CEMReloadClient {}; export class StatePersistence {}"
				} else {
					// Node built-in stubs with working Buffer for base64
					contents = `export default {};
export const Buffer = {
  from(s, encoding) {
    if (encoding === 'binary') {
      const arr = new Uint8Array(s.length);
      for (let i = 0; i < s.length; i++) arr[i] = s.charCodeAt(i);
      return { toString(enc) {
        if (enc === 'base64') return btoa(String.fromCharCode(...arr));
        return s;
      }};
    }
    if (typeof s === 'string') return new TextEncoder().encode(s);
    return new Uint8Array(s);
  },
  isBuffer() { return false; },
  alloc(n) { return new Uint8Array(n); },
};
export const readFileSync = () => "";`
				}
				return api.OnLoadResult{Contents: &contents, Loader: api.LoaderJS}, nil
			})
		},
	}
}

// stubNodeBuiltins replaces Node.js built-in imports with stubs for QuickJS.
// Matches the approach from lit-ssr-wasm's build script.
func stubNodeBuiltins() api.Plugin {
	return api.Plugin{
		Name: "stub-node-builtins",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: `^(node:|buffer$|fs$|path$|stream|util$|events$|crypto$|os$|url$|http$|https$|net$|tls$|child_process$|module$|vm$|zlib$|node-fetch)|register-css-hook`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{Path: args.Path, Namespace: "stub"}, nil
			})
		},
	}
}

// rewriteLitImports replaces bare lit specifiers with /__cem/vendor/ paths.
func rewriteLitImports(code string) string {
	replacements := []struct {
		from, to string
	}{
		// Order matters: longer paths first to avoid partial matches
		{`from "lit/decorators/custom-element.js"`, `from "/__cem/vendor/lit/decorators/custom-element.js"`},
		{`from "lit/decorators/property.js"`, `from "/__cem/vendor/lit/decorators/property.js"`},
		{`from "lit/decorators/state.js"`, `from "/__cem/vendor/lit/decorators/state.js"`},
		{`from "lit/decorators/query.js"`, `from "/__cem/vendor/lit/decorators/query.js"`},
		{`from "lit/decorators/query-all.js"`, `from "/__cem/vendor/lit/decorators/query-all.js"`},
		{`from "lit/decorators/event-options.js"`, `from "/__cem/vendor/lit/decorators/event-options.js"`},
		{`from "lit/directives/class-map.js"`, `from "/__cem/vendor/lit/directives/class-map.js"`},
		{`from "lit/directives/style-map.js"`, `from "/__cem/vendor/lit/directives/style-map.js"`},
		{`from "lit/directives/if-defined.js"`, `from "/__cem/vendor/lit/directives/if-defined.js"`},
		{`from "lit/directives/repeat.js"`, `from "/__cem/vendor/lit/directives/repeat.js"`},
		{`from "lit/directives/live.js"`, `from "/__cem/vendor/lit/directives/live.js"`},
		{`from "lit/directives/unsafe-html.js"`, `from "/__cem/vendor/lit/directives/unsafe-html.js"`},
		{`from "lit"`, `from "/__cem/vendor/lit.js"`},
	}
	for _, r := range replacements {
		code = strings.ReplaceAll(code, r.from, r.to)
	}
	return code
}

// bundleSSR creates a single fully-bundled JS file containing all element
// components with lit included. This bundle is used by the lit-ssr-wasm
// renderer for server-side rendering. Unlike the browser-served modules
// (which use external imports), this bundle is self-contained.
func bundleSSR() error {
	// Find all element .ts files
	entries, err := filepath.Glob("elements/*/*.ts")
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return nil
	}

	// Self-contained SSR bundle:
	//   1. ssr-polyfills.js (URL, Buffer for QuickJS)
	//   2. @lit-labs/ssr DOM shims (HTMLElement, document, CSSStyleSheet, etc.)
	//   3. lit (loads against the shims)
	//   4. All chrome components
	var imports strings.Builder
	fmt.Fprintln(&imports, "import './elements/ssr-polyfills.js';")
	fmt.Fprintln(&imports, "import '@lit-labs/ssr/lib/install-global-dom-shim.js';")
	for _, entry := range entries {
		fmt.Fprintf(&imports, "import './%s';\n", entry)
	}

	outfile := "middleware/routes/templates/ssr-bundle.js"
	source := stripImportAttributes(imports.String())

	result := api.Build(api.BuildOptions{
		Stdin: &api.StdinOptions{
			Contents:   source,
			Sourcefile: "ssr-entry.ts",
			Loader:     api.LoaderTS,
			ResolveDir: ".",
		},
		Outfile:    outfile,
		Bundle:     true,
		Format:     api.FormatESModule,
		Target:     api.ES2022,
		Platform:   api.PlatformNode,
		Conditions: []string{"node"},
		Write:      true,
		TsconfigRaw: `{"compilerOptions":{}}`,
		LogLevel:   api.LogLevelWarning,
		NodePaths:  []string{"elements/node_modules"},
		Define:     map[string]string{"process.env.NODE_ENV": `"production"`},
		Plugins:    []api.Plugin{litCSSPlugin(), stubExternals(), stubNodeBuiltins()},
	})

	if len(result.Errors) > 0 {
		msgs := api.FormatMessages(result.Errors, api.FormatMessagesOptions{})
		return fmt.Errorf("bundling SSR:\n%s", strings.Join(msgs, "\n"))
	}

	return nil
}
