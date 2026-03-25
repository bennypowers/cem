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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	G "bennypowers.dev/cem/generate"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	"github.com/evanw/esbuild/pkg/api"

	litssr "bennypowers.dev/lit-ssr-wasm/go"
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
	if err := generateManifests(); err != nil {
		return fmt.Errorf("generate manifests: %w", err)
	}
	if err := bundleSSR(); err != nil {
		return fmt.Errorf("bundle SSR: %w", err)
	}
	if err := compileSSRBytecode(); err != nil {
		return fmt.Errorf("compile SSR bytecode: %w", err)
	}
	if err := bundleChromeBrowser(); err != nil {
		return fmt.Errorf("bundle chrome browser: %w", err)
	}
	return nil
}

// compileSSRBytecode compiles the SSR bundle to QuickJS bytecode for
// faster worker init (~100ms vs ~1s per worker).
func compileSSRBytecode() error {
	source, err := os.ReadFile("middleware/routes/templates/ssr-bundle.js")
	if err != nil {
		return err
	}
	bytecode, err := litssr.CompileSource(context.Background(), string(source))
	if err != nil {
		return fmt.Errorf("compile bytecode: %w", err)
	}
	return os.WriteFile("middleware/routes/templates/ssr-bundle.qbc", bytecode, 0o644)
}

// generateManifests runs cem generate on the element sources to produce
// custom-elements.json manifests for the two internal packages.
func generateManifests() error {
	pkgs := []struct {
		glob    string
		outPath string
	}{
		{"pf-v6-*/*.ts", "elements/patternfly/custom-elements.json"},
		{"cem-*/*.ts", "elements/chrome/custom-elements.json"},
	}

	for _, pkg := range pkgs {
		wsCtx := W.NewFileSystemWorkspaceContext("elements")
		if err := wsCtx.Init(); err != nil {
			return fmt.Errorf("init workspace for %s: %w", pkg.glob, err)
		}
		cfg, err := wsCtx.Config()
		if err != nil {
			return fmt.Errorf("config for %s: %w", pkg.glob, err)
		}
		cfg.Generate.Files = []string{pkg.glob}

		session, err := G.NewGenerateSession(wsCtx)
		if err != nil {
			return fmt.Errorf("session for %s: %w", pkg.glob, err)
		}

		manifest, err := session.GenerateFullManifest(context.Background())
		session.Close()
		if err != nil {
			return fmt.Errorf("generate %s: %w", pkg.glob, err)
		}

		jsonStr, err := M.SerializeToString(manifest)
		if err != nil {
			return fmt.Errorf("serialize %s: %w", pkg.glob, err)
		}

		if err := os.MkdirAll(filepath.Dir(pkg.outPath), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(pkg.outPath, []byte(jsonStr), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", pkg.outPath, err)
		}
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
				// Use path relative to cwd so source maps are stable
				// across machines and worktrees
				cwd, _ := os.Getwd()
				relPath, err := filepath.Rel(cwd, absPath)
				if err != nil {
					relPath = absPath
				}
				return api.OnResolveResult{
					Path:      relPath,
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

// stubCemExternals stubs client-only /__cem/* imports for the SSR bundle.
func stubCemExternals() api.Plugin {
	return api.Plugin{
		Name: "stub-cem-externals",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: `^/__cem/`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{Path: args.Path, Namespace: "cem-stub"}, nil
			})
			build.OnLoad(api.OnLoadOptions{Filter: `.*`, Namespace: "cem-stub"}, func(_ api.OnLoadArgs) (api.OnLoadResult, error) {
				contents := "export default {}; export class CEMReloadClient {}; export class StatePersistence {}"
				return api.OnLoadResult{Contents: &contents, Loader: api.LoaderJS}, nil
			})
		},
	}
}

// litSsrWasmPlugin resolves @lit-labs/ssr-dom-shim to globalThis re-exports,
// so the consumer's Lit copy shares the same customElements registry and DOM
// shims that the lit-ssr-wasm runtime provides.
// Port of lit-ssr-wasm's src/esbuild-plugin.ts.
func litSsrWasmPlugin() api.Plugin {
	shimBridge := `
export const customElements = globalThis.customElements;
export const HTMLElement = globalThis.HTMLElement;
export const Element = globalThis.Element;
export const CSSStyleSheet = globalThis.CSSStyleSheet;
export const CustomElementRegistry = globalThis.CustomElementRegistry;
export const Event = globalThis.Event;
export const CustomEvent = globalThis.CustomEvent;
export const EventTarget = globalThis.EventTarget;
export const ariaMixinAttributes = globalThis.ariaMixinAttributes ?? {};
export const HYDRATE_INTERNALS_ATTR_PREFIX = globalThis.HYDRATE_INTERNALS_ATTR_PREFIX ?? 'internals-';
export const ElementInternals = globalThis.ElementInternals ?? class ElementInternals {};
`
	return api.Plugin{
		Name: "lit-ssr-wasm",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: `^@lit-labs/ssr-dom-shim`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{Path: args.Path, Namespace: "lit-ssr-wasm-shim"}, nil
			})
			build.OnLoad(api.OnLoadOptions{Filter: `.*`, Namespace: "lit-ssr-wasm-shim"}, func(_ api.OnLoadArgs) (api.OnLoadResult, error) {
				return api.OnLoadResult{Contents: &shimBridge, Loader: api.LoaderJS}, nil
			})
		},
	}
}

// stubNodeBuiltins stubs Node.js built-in modules for QuickJS.
// Port of lit-ssr-wasm's src/esbuild-stubs.ts.
func stubNodeBuiltins() api.Plugin {
	bufferStub := `export default {};
export const Buffer = {
  from(x, encoding) {
    if (typeof x === "string") {
      if (encoding === "binary") {
        return { toString(enc) { if (enc === "base64") return globalThis.btoa(x); return x; } };
      }
      return new TextEncoder().encode(x);
    }
    return new Uint8Array(x);
  },
  isBuffer: () => false,
  alloc: n => new Uint8Array(n),
};
export const readFileSync = () => "";`
	return api.Plugin{
		Name: "stub-node-builtins",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: `^(?:node:)?(?:buffer|fs|path|stream|util|events|crypto|os|url|http|https|net|tls|child_process|module|vm|zlib|node-fetch)(?:/|$)`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{Path: args.Path, Namespace: "node-stub"}, nil
			})
			build.OnLoad(api.OnLoadOptions{Filter: `.*`, Namespace: "node-stub"}, func(_ api.OnLoadArgs) (api.OnLoadResult, error) {
				return api.OnLoadResult{Contents: &bufferStub, Loader: api.LoaderJS}, nil
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

	// SSR bundle: lit + all chrome components.
	// The lit-ssr-wasm runtime provides DOM shims via globalThis.
	// litSsrWasmPlugin bridges @lit-labs/ssr-dom-shim to those globals.
	var imports strings.Builder
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
		Plugins:    []api.Plugin{litCSSPlugin(), stubCemExternals(), litSsrWasmPlugin(), stubNodeBuiltins()},
	})

	if len(result.Errors) > 0 {
		msgs := api.FormatMessages(result.Errors, api.FormatMessagesOptions{})
		return fmt.Errorf("bundling SSR:\n%s", strings.Join(msgs, "\n"))
	}

	return nil
}

// bundleChromeBrowser creates a single browser-targeted JS bundle containing
// all chrome components with lit inlined. Unlike the SSR bundle (which includes
// @lit-labs/ssr shims for QuickJS), this bundle is for the browser and uses
// standard lit without any SSR bridging.
//
// Used by `cem serve --build` for static site output: one <script> tag for
// the entire chrome UI instead of 53+ individual module requests.
func bundleChromeBrowser() error {
	entries, err := filepath.Glob("elements/*/*.ts")
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return nil
	}

	var imports strings.Builder
	for _, entry := range entries {
		fmt.Fprintf(&imports, "import './%s';\n", entry)
	}

	// Hydration support must load before any LitElement definitions
	// so Lit can pick up SSR-rendered DSD content instead of re-rendering
	source := "import '@lit-labs/ssr-client/lit-element-hydrate-support.js';\n" +
		stripImportAttributes(imports.String())
	outfile := "middleware/routes/templates/chrome-bundle.js"

	result := api.Build(api.BuildOptions{
		Stdin: &api.StdinOptions{
			Contents:   source,
			Sourcefile: "chrome-entry.ts",
			Loader:     api.LoaderTS,
			ResolveDir: ".",
		},
		Outfile:          outfile,
		Bundle:           true,
		Format:           api.FormatESModule,
		Target:           api.ES2022,
		Write:            true,
		MinifySyntax:     true,
		MinifyWhitespace: true,
		MinifyIdentifiers: true,
		Sourcemap:        api.SourceMapLinked,
		TsconfigRaw:      `{"compilerOptions":{}}`,
		LogLevel:         api.LogLevelWarning,
		NodePaths:        []string{"elements/node_modules"},
		Plugins:          []api.Plugin{litCSSPlugin()},
		External:         []string{"/__cem/*"},
	})

	if len(result.Errors) > 0 {
		msgs := api.FormatMessages(result.Errors, api.FormatMessagesOptions{})
		return fmt.Errorf("bundling chrome browser:\n%s", strings.Join(msgs, "\n"))
	}

	return nil
}
