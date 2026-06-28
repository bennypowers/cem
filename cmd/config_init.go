/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"bennypowers.dev/asimonim/specifier"
	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/tui"
	W "bennypowers.dev/cem/internal/workspace"
	"charm.land/bubbles/v2/key"
	"charm.land/huh/v2"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var validFormats = []string{"yaml", "json", "jsonc"}

func init() {
	configInitCmd.Flags().BoolP("yes", "y", false, "accept defaults without prompting")
	configInitCmd.Flags().String("format", "yaml", "output format (yaml, json, jsonc)")
	configInitCmd.SilenceUsage = true
	configInitCmd.RunE = runConfigInit
	configInitCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		verboseCount, _ := cmd.Flags().GetCount("verbose")
		quiet, _ := cmd.Flags().GetBool("quiet")
		if verboseCount > 0 && quiet {
			return fmt.Errorf("cannot use both --verbose and --quiet flags together")
		}
		if quiet {
			logging.SetVerbosity(logging.VerbosityQuiet)
		} else if verboseCount > 0 {
			flagLevel := logging.VerbosityVerbose
			switch {
			case verboseCount >= 3:
				flagLevel = logging.VerbosityTrace
			case verboseCount == 2:
				flagLevel = logging.VerbosityDebug
			}
			if flagLevel > logging.CurrentVerbosity() {
				logging.SetVerbosity(flagLevel)
			}
		}
		return nil
	}
}

func runConfigInit(cmd *cobra.Command, args []string) error {
	yes, _ := cmd.Flags().GetBool("yes")
	format, _ := cmd.Flags().GetString("format")
	interactive := term.IsTerminal(int(os.Stdin.Fd()))

	if !interactive && !yes {
		return fmt.Errorf("no TTY detected; use --yes to accept defaults")
	}

	if !slices.Contains(validFormats, format) {
		return fmt.Errorf("unsupported format %q; valid: %s", format, strings.Join(validFormats, ", "))
	}

	if interactive && !yes {
		if err := showWelcomePage(); err != nil {
			if errors.Is(err, tui.ErrCancelled) {
				return nil
			}
			return err
		}
	}

	root, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get working directory: %w", err)
	}
	if p, _ := cmd.Flags().GetString("package"); p != "" {
		root, err = filepath.Abs(p)
		if err != nil {
			return fmt.Errorf("unable to resolve package path: %w", err)
		}
	}

	var cfgPath string
	var existing *IC.CemConfig
	var existingData []byte
	var existingFormat string
	if found := IC.FindConfigFile(root); found != "" {
		cfgPath = found
		existing, err = IC.LoadConfig(cfgPath)
		if err != nil {
			logging.Warning("Failed to load existing config: %v", err)
		} else {
			existingFormat = IC.FormatFromPath(cfgPath)
			existingData, _ = os.ReadFile(cfgPath)
			logging.Info("Found existing config: %s", cfgPath)
		}
	}

	osFS := &platform.OSFileSystem{}
	hasPkgJSON := osFS.Exists(filepath.Join(root, "package.json"))
	dirFS := os.DirFS(root)

	cfg := &IC.CemConfig{}
	if existing != nil {
		*cfg = *existing
	}

	// Detection phase
	detectedFiles, detectErr := detectSourceFiles(dirFS)
	if detectErr != nil {
		logging.Debug("source file detection: %v", detectErr)
	}
	detectedGitRemote, detectErr := detectGitRemote(root)
	if detectErr != nil {
		logging.Debug("git remote detection: %v", detectErr)
	}
	detectedGlob, detectedPattern, detectErr := detectDemoFiles(dirFS)
	if detectErr != nil {
		logging.Debug("demo file detection: %v", detectErr)
	}
	var detectedTokens []*specifier.ResolvedFile
	if hasPkgJSON {
		detectedTokens, detectErr = detectDesignTokens(root, osFS)
		if detectErr != nil {
			logging.Debug("design token detection: %v", detectErr)
		}
	}
	var detectedPkgCE string
	if hasPkgJSON {
		if data, readErr := os.ReadFile(filepath.Join(root, "package.json")); readErr == nil {
			var pkg struct {
				CustomElements string `json:"customElements"`
			}
			if json.Unmarshal(data, &pkg) == nil {
				detectedPkgCE = pkg.CustomElements
			}
		}
	}
	detectedCSSInclude, detectedCSSExclude := detectCSSPatterns(
		append(detectedFiles, cfg.Generate.Files...), dirFS)
	hasTsConfig, _ := detectTypeScript(root, osFS)

	// Resolve field values
	filesFV := fieldValue{
		Title: "Source file patterns",
		Description: "Glob patterns for TypeScript files containing custom element definitions.\n" +
			"Detected by scanning for .ts files in top-level directories.\n" +
			"Learn more: https://bennypowers.dev/cem/docs/reference/commands/generate/",
		Existing: strings.Join(cfg.Generate.Files, ", "),
		Detected: strings.Join(detectedFiles, ", "),
		Fallback: "**/*.ts",
	}
	outputFV := fieldValue{
		Title: "Manifest output path",
		Description: "Path for the generated custom elements manifest.\n" +
			"Tools and IDEs use this file for component metadata.\n" +
			"Learn more: https://bennypowers.dev/cem/docs/reference/configuration/",
		Existing: cfg.Generate.Output,
		Detected: detectedPkgCE,
		Fallback: "custom-elements.json",
	}
	scurlFV := fieldValue{
		Title: "Source control root URL",
		Description: "Base URL for linking manifest entries back to source files in your repository.\n" +
			"Learn more: https://bennypowers.dev/cem/docs/reference/configuration/",
		Placeholder: "https://github.com/user/repo/tree/main/",
		Existing:    cfg.SourceControlRootUrl,
		Detected:    detectedGitRemote,
	}
	globFV := fieldValue{
		Title: "Demo file glob",
		Description: "Glob pattern to discover demo HTML files in your project.\n" +
			"Demos are HTML partials that showcase your custom elements.\n" +
			"Learn more: https://bennypowers.dev/cem/docs/usage/demos/",
		Existing: cfg.Generate.DemoDiscovery.FileGlob,
		Detected: detectedGlob,
	}
	patternFV := fieldValue{
		Title: "Demo URL pattern",
		Description: "URL pattern with named params for matching demo file paths to elements.\n" +
			"Used by the dev server to route demo requests.\n" +
			"Learn more: https://bennypowers.dev/cem/docs/usage/demos/",
		Placeholder: "elements/:tag/demo/:demo.html",
		Existing:    cfg.Generate.DemoDiscovery.URLPattern,
		Detected:    detectedPattern,
	}

	templateFV := fieldValue{
		Title: "Demo URL template",
		Description: "Go template for generating public demo URLs from matched parameters.\n" +
			"Leave empty if you don't need public demo links in the manifest.",
		Placeholder: "https://example.com/{{.tag}}/{{.demo}}/",
		Existing:    cfg.Generate.DemoDiscovery.URLTemplate,
	}

	var detectedTokenSpec string
	if len(detectedTokens) == 1 {
		detectedTokenSpec = detectedTokens[0].Specifier
		if detectedTokenSpec == "" {
			detectedTokenSpec = detectedTokens[0].Path
		}
	}
	tokenSpecFV := fieldValue{
		Title: "Design token spec",
		Description: "Path or npm specifier for a DTCG design token file.\n" +
			"CEM documents CSS custom properties from design tokens in the manifest.\n" +
			"Learn more: https://bennypowers.dev/cem/docs/reference/configuration/",
		Existing: cfg.Generate.DesignTokens.Spec,
		Detected: detectedTokenSpec,
	}
	tokenPrefixFV := fieldValue{
		Title: "Token prefix",
		Description: "CSS custom property prefix applied to token names.\n" +
			"For example, prefix 'my-ds' produces --my-ds-color-primary.",
		Existing: cfg.Generate.DesignTokens.Prefix,
	}

	existingPort := ""
	if cfg.Serve.Port != 0 {
		existingPort = strconv.Itoa(cfg.Serve.Port)
	}
	portFV := fieldValue{
		Title: "Dev server port",
		Description: "Port for the cem dev server.\n" +
			"Learn more: https://bennypowers.dev/cem/docs/reference/commands/serve/",
		Existing: existingPort,
		Fallback: "8000",
	}

	if interactive && !yes {
		// Build the main form
		var groups []*huh.Group

		// === Generate ===
		configureGenerate := true
		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure manifest generation?").
				Value(&configureGenerate),
		).Title("Generate").
			Description("CEM analyzes TypeScript source files and produces a custom-elements.json manifest\n"+
				"describing your web components, their attributes, properties, events, slots, and CSS.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/reference/commands/generate/"))

		genGate := func() bool { return configureGenerate }
		filesFV.gate = genGate
		outputFV.gate = genGate
		scurlFV.gate = genGate
		groups = append(groups, filesFV.Groups()...)
		groups = append(groups, outputFV.Groups()...)
		groups = append(groups, scurlFV.Groups()...)

		// === Demo Discovery ===
		hasDemos := globFV.Value() != "" || patternFV.Value() != ""
		configureDemos := hasDemos || len(detectedFiles) > 0
		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure demo discovery?").
				Value(&configureDemos),
		).Title("Demo Discovery").
			Description("Demos are HTML partials that showcase your custom elements.\n"+
				"The dev server discovers and serves them with live reload.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/usage/demos/"))

		demoGate := func() bool { return configureDemos }
		globFV.gate = demoGate
		patternFV.gate = demoGate
		templateFV.gate = demoGate
		groups = append(groups, globFV.Groups()...)
		groups = append(groups, patternFV.Groups()...)
		groups = append(groups, templateFV.Groups()...)

		// === Design Tokens ===
		configureTokens := len(detectedTokens) > 0 || cfg.Generate.DesignTokens.Spec != ""
		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure design tokens?").
				Value(&configureTokens),
		).Title("Design Tokens").
			Description("Document CSS custom properties from DTCG design token files in the manifest.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/reference/configuration/"))

		tokenGate := func() bool { return configureTokens }
		tokenSpecFV.gate = tokenGate
		tokenPrefixFV.gate = tokenGate
		groups = append(groups, tokenSpecFV.Groups()...)
		groups = append(groups, tokenPrefixFV.Groups()...)

		// === Dev Server ===
		configureServe := true
		rendering := cfg.Serve.Demos.Rendering
		if rendering == "" {
			rendering = "shadow"
		}
		importMapGen := cfg.Serve.ImportMap.Generate
		if existing == nil {
			importMapGen = true
		}
		enableCSS := cfg.Serve.Transforms.CSS.Enabled || len(detectedCSSInclude) > 0

		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure the dev server?").
				Value(&configureServe),
		).Title("Dev Server").
			Description("The cem dev server serves demos with live reload, TypeScript transforms,\n"+
				"and import maps.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/reference/commands/serve/"))

		serveGate := func() bool { return configureServe }
		portFV.gate = serveGate
		groups = append(groups, portFV.Groups()...)

		groups = append(groups, huh.NewGroup(
			huh.NewSelect[string]().
				Title("Demo rendering mode").
				Options(
					huh.NewOption("Light -- full dev UI with knobs and event monitoring", "light"),
					huh.NewOption("Shadow -- demos in shadow DOM for CSS encapsulation testing (default)", "shadow"),
					huh.NewOption("Iframe -- full isolation, no style/tag leaks between demos", "iframe"),
					huh.NewOption("Chromeless -- minimal standalone pages for testing/screenshots", "chromeless"),
				).
				Value(&rendering),
		).Title("Demo Rendering Mode").
			Description("Controls how demo HTML is presented in the dev server.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/usage/rendering-modes/").
			WithHideFunc(func() bool { return !configureServe }))

		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Auto-generate import maps?").
				Value(&importMapGen),
		).Title("Import Maps").
			Description("Generates import maps from package.json so you can use bare specifiers\n"+
				"like `import { LitElement } from 'lit'` without bundling.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/usage/import-maps/").
			WithHideFunc(func() bool { return !configureServe }))

		cssIncludeFV := fieldValue{
			Title: "CSS include patterns",
			Description: "Glob patterns for CSS files to transform to JavaScript modules.\n" +
				"Learn more: https://bennypowers.dev/cem/docs/usage/buildless-development/",
			Existing: strings.Join(cfg.Serve.Transforms.CSS.Include, ", "),
			Detected: strings.Join(detectedCSSInclude, ", "),
		}
		cssExcludeFV := fieldValue{
			Title:       "CSS exclude patterns",
			Description: "Glob patterns for CSS files to skip during transformation.",
			Existing:    strings.Join(cfg.Serve.Transforms.CSS.Exclude, ", "),
			Detected:    strings.Join(detectedCSSExclude, ", "),
		}

		if len(detectedCSSInclude) > 0 || cfg.Serve.Transforms.CSS.Enabled {
			groups = append(groups, huh.NewGroup(
				huh.NewConfirm().
					Title("Enable CSS transforms?").
					Value(&enableCSS),
			).Title("CSS Transforms").
				Description("Transforms CSS files to JavaScript modules for use in web components.\n"+
					"Learn more: https://bennypowers.dev/cem/docs/usage/buildless-development/").
				WithHideFunc(func() bool { return !configureServe }))

			cssGate := func() bool { return configureServe && enableCSS }
			cssIncludeFV.gate = cssGate
			cssExcludeFV.gate = cssGate
			groups = append(groups, cssIncludeFV.Groups()...)
			groups = append(groups, cssExcludeFV.Groups()...)
		}

		// === Health ===
		existingFailBelow := ""
		if cfg.Health.FailBelow != 0 {
			existingFailBelow = strconv.Itoa(cfg.Health.FailBelow)
		}
		failBelowFV := fieldValue{
			Title: "Minimum health score",
			Description: "The `cem health` command exits with an error if the manifest health score\n" +
				"is below this threshold (0-100). Leave empty to skip.",
			Existing: existingFailBelow,
		}
		configureHealth := cfg.Health.FailBelow != 0 || len(cfg.Health.Disable) > 0
		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure health checks?").
				Value(&configureHealth),
		).Title("Health Checks").
			Description("CEM scores your manifest completeness and flags missing descriptions,\n"+
				"undocumented slots, and other quality issues.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/reference/configuration/"))

		failBelowFV.gate = func() bool { return configureHealth }
		groups = append(groups, failBelowFV.Groups()...)

		// === Export ===
		configureExport := len(cfg.Export) > 0
		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure framework exports?").
				Value(&configureExport),
		).Title("Framework Exports").
			Description("Generate framework-specific wrapper components (React, Vue, Angular)\n"+
				"from your custom elements manifest.\n"+
				"Learn more: https://bennypowers.dev/cem/docs/reference/configuration/"))

		availableFrameworks := []string{"react", "vue", "angular"}
		var selectedFrameworks []string
		for _, fw := range availableFrameworks {
			if _, ok := cfg.Export[fw]; ok {
				selectedFrameworks = append(selectedFrameworks, fw)
			}
		}
		groups = append(groups, huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Which frameworks?").
				Options(
					huh.NewOption("React", "react"),
					huh.NewOption("Vue", "vue"),
					huh.NewOption("Angular", "angular"),
				).
				Value(&selectedFrameworks),
		).Title("Framework Selection").
			Description("Select one or more frameworks to generate wrapper components for.").
			WithHideFunc(func() bool { return !configureExport }))

		type exportFVs struct {
			output      fieldValue
			stripPrefix fieldValue
			packageName fieldValue
		}
		exportData := map[string]*exportFVs{}
		for _, fw := range availableFrameworks {
			fwLabel := strings.ToUpper(fw[:1]) + fw[1:]
			var existingOutput, existingPrefix, existingPkg string
			if ec, ok := cfg.Export[fw]; ok {
				existingOutput = ec.Output
				existingPrefix = ec.StripPrefix
				existingPkg = ec.PackageName
			}

			fwGate := func() bool {
				return configureExport && slices.Contains(selectedFrameworks, fw)
			}

			efvs := &exportFVs{
				output: fieldValue{
					Title:       "Output directory",
					Description: "Directory for generated " + fw + " wrapper files.",
					Existing:    existingOutput,
					Fallback:    fw + "-wrappers",
					gate:        fwGate,
				},
				stripPrefix: fieldValue{
					Title: "Strip prefix",
					Description: "Prefix to remove from tag names when generating component names.\n" +
						"For example, 'my-' turns <my-button> into <Button>.",
					Existing: existingPrefix,
					gate:     fwGate,
				},
				packageName: fieldValue{
					Title:       "Package name",
					Description: "Override the npm package name used in import paths.",
					Existing:    existingPkg,
					gate:        fwGate,
				},
			}
			exportData[fw] = efvs

			groups = append(groups, huh.NewGroup(
				huh.NewNote().
					Title(fwLabel+" Export").
					Description("Configure "+fw+" wrapper generation.").
					Next(true),
			).WithHideFunc(func() bool { return !fwGate() }))

			groups = append(groups, efvs.output.Groups()...)
			groups = append(groups, efvs.stripPrefix.Groups()...)
			groups = append(groups, efvs.packageName.Groups()...)
		}

		// === MCP ===
		existingMaxDesc := ""
		if cfg.MCP.MaxDescriptionLength != 0 {
			existingMaxDesc = strconv.Itoa(cfg.MCP.MaxDescriptionLength)
		}
		maxDescFV := fieldValue{
			Title: "Max description length",
			Description: "Truncate component descriptions in MCP tool responses to this many characters.\n" +
				"Leave empty for no limit.",
			Existing: existingMaxDesc,
		}
		configureMCP := cfg.MCP.MaxDescriptionLength != 0
		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure MCP settings?").
				Value(&configureMCP),
		).Title("MCP").
			Description("CEM provides an MCP server for AI coding assistants.\n"+
				"Configure how component metadata is exposed to AI tools."))

		maxDescFV.gate = func() bool { return configureMCP }
		groups = append(groups, maxDescFV.Groups()...)

		// === Additional Packages ===
		existingAdditional := strings.Join(cfg.AdditionalPackages, ", ")
		additionalFV := fieldValue{
			Title: "Additional packages",
			Description: "Extra component packages to load for MCP and LSP,\n" +
				"beyond those listed in package.json.\n" +
				"Comma-separated absolute paths, URLs, npm:, or jsr: specifiers.\n" +
				"Learn more: https://bennypowers.dev/cem/docs/reference/configuration/",
			Existing:   existingAdditional,
			ValidateFn: validatePackageSpecifiers,
		}
		configureAdditional := len(cfg.AdditionalPackages) > 0
		groups = append(groups, huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to configure additional packages?").
				Value(&configureAdditional),
		).Title("Additional Packages").
			Description("Load extra component packages for MCP and LSP,\n"+
				"beyond those listed in package.json.\n"+
				"Useful for consuming third-party design system manifests\n"+
				"or working without a package.json."))

		additionalFV.gate = func() bool { return configureAdditional }
		groups = append(groups, additionalFV.Groups()...)

		// === Output Format ===
		if cfgPath == "" {
			groups = append(groups, huh.NewGroup(
				huh.NewSelect[string]().
					Title("Config file format").
					Options(
						huh.NewOption("YAML -- .config/cem.yaml", "yaml"),
						huh.NewOption("JSON -- .config/cem.json", "json"),
						huh.NewOption("JSONC -- .config/cem.jsonc (JSON with comments)", "jsonc"),
					).
					Value(&format),
			).Title("Output Format").
				Description("Choose the format for your config file."))
		}

		// Run the form
		km := huh.NewDefaultKeyMap()
		km.Quit = key.NewBinding(key.WithKeys("esc", "ctrl+c"), key.WithHelp("esc", "quit"))
		form := huh.
			NewForm(groups...).
			WithKeyMap(km).
			WithShowHelp(true)
		if formErr := tui.WrapAbort(form.Run()); formErr != nil {
			if errors.Is(formErr, tui.ErrCancelled) {
				return nil
			}
			return formErr
		}

		// Apply form results
		if configureGenerate {
			cfg.Generate.Files = splitCommaList(filesFV.Resolve())
			cfg.Generate.Output = outputFV.Resolve()
			cfg.SourceControlRootUrl = scurlFV.Resolve()
		}

		if configureDemos {
			cfg.Generate.DemoDiscovery.FileGlob = globFV.Resolve()
			cfg.Generate.DemoDiscovery.URLPattern = patternFV.Resolve()
			cfg.Generate.DemoDiscovery.URLTemplate = templateFV.Resolve()
		}
		if configureTokens {
			if tokenSpecFV.Resolve() != "" {
				cfg.Generate.DesignTokens.Spec = tokenSpecFV.Resolve()
				cfg.Generate.DesignTokens.Prefix = tokenPrefixFV.Resolve()
			}
		}

		if configureServe {
			portStr := portFV.Resolve()
			port, parseErr := strconv.Atoi(portStr)
			if parseErr != nil || port < 1 || port > 65535 {
				return fmt.Errorf("invalid port: %s", portStr)
			}
			cfg.Serve.Port = port
			cfg.Serve.Demos.Rendering = rendering
			cfg.Serve.ImportMap.Generate = importMapGen

			if enableCSS {
				cfg.Serve.Transforms.CSS.Enabled = true
				cfg.Serve.Transforms.CSS.Include = splitCommaList(cssIncludeFV.Resolve())
				cfg.Serve.Transforms.CSS.Exclude = splitCommaList(cssExcludeFV.Resolve())
			} else {
				cfg.Serve.Transforms.CSS.Enabled = false
			}
		}

		if configureExport {
			prevExport := cfg.Export
			cfg.Export = map[string]IC.FrameworkExportConfig{}
			for _, fw := range selectedFrameworks {
				efvs := exportData[fw]
				ec := IC.FrameworkExportConfig{
					Output:      efvs.output.Resolve(),
					StripPrefix: efvs.stripPrefix.Resolve(),
					PackageName: efvs.packageName.Resolve(),
				}
				if prev, ok := prevExport[fw]; ok {
					ec.ModuleName = prev.ModuleName
				}
				cfg.Export[fw] = ec
			}
		}

		if configureHealth {
			fb := failBelowFV.Resolve()
			if fb != "" {
				v, parseErr := strconv.Atoi(fb)
				if parseErr != nil {
					return fmt.Errorf("invalid health score %q: must be a number", fb)
				}
				if v < 0 || v > 100 {
					return fmt.Errorf("invalid health score %d: must be 0-100", v)
				}
				cfg.Health.FailBelow = v
			}
		}

		if configureMCP {
			md := maxDescFV.Resolve()
			if md != "" {
				v, parseErr := strconv.Atoi(md)
				if parseErr != nil {
					return fmt.Errorf("invalid max description length %q: must be a number", md)
				}
				if v < 0 {
					return fmt.Errorf("invalid max description length %d: must be non-negative", v)
				}
				cfg.MCP.MaxDescriptionLength = v
			}
		}

		if configureAdditional {
			pkgs := splitCommaList(additionalFV.Resolve())
			if len(pkgs) > 0 {
				cfg.AdditionalPackages = pkgs
			}
		}

		if hasTsConfig {
			logging.Info("tsconfig.json detected. URL rewrites (rootDir/outDir) are auto-detected at serve time.")
		}
	} else {
		// --yes mode: apply defaults from detection
		cfg.Generate.Files = splitCommaList(filesFV.Value())
		cfg.Generate.Output = outputFV.Value()
		if scurlFV.Value() != "" {
			cfg.SourceControlRootUrl = scurlFV.Value()
		}
		if globFV.Value() != "" {
			cfg.Generate.DemoDiscovery.FileGlob = globFV.Value()
			cfg.Generate.DemoDiscovery.URLPattern = patternFV.Value()
		}
		if cfg.Generate.DesignTokens.Spec == "" && len(detectedTokens) == 1 {
			spec := detectedTokens[0].Specifier
			if spec == "" {
				spec = detectedTokens[0].Path
			}
			cfg.Generate.DesignTokens.Spec = spec
		}
		hasExistingServe := existing != nil && existingData != nil && hasConfigKey(existingData, "serve")
		if !hasExistingServe {
			cfg.Serve.Port = 8000
			cfg.Serve.ImportMap.Generate = true
			cfg.Serve.Demos.Rendering = "shadow"
		}
	}

	if cfgPath != "" && existingFormat != "" {
		format = existingFormat
	}

	if W.IsWorkspaceMode(root) {
		logging.Info("This is a workspace root. Config will apply to all packages.")
		logging.Info("Package-level overrides can be added in each package's own .config/cem.yaml")
	}

	// Clear internal-only and deprecated fields before marshaling
	cfg.ProjectDir = ""
	cfg.ConfigFile = ""
	cfg.PackageName = ""
	cfg.Verbose = false
	cfg.LogLevel = ""

	output, marshalErr := marshalConfig(cfg, format, existingData)
	if marshalErr != nil {
		return marshalErr
	}

	lang := format
	if lang == "jsonc" {
		lang = "json"
	}

	if interactive && !yes {
		cmd.Println()
		if len(existingData) > 0 {
			diff := unifiedDiff(string(existingData), string(output), cfgPath)
			if diff != "" {
				printColorDiff(cmd, diff)
			} else {
				cmd.Println("No changes.")
			}
		} else {
			if err := tui.Highlight(cmd.OutOrStdout(), string(output), lang); err != nil {
				cmd.Println(string(output))
			}
		}
		cmd.Println()
	}

	if interactive && !yes {
		prompt := "Write this config?"
		if len(existingData) > 0 {
			prompt = "Apply these changes?"
		}
		confirmed, confirmErr := tui.Confirm(prompt, true)
		if confirmErr != nil || !confirmed {
			return nil
		}
	}

	outPath := cfgPath
	if outPath == "" {
		outPath = filepath.Join(root, ".config", "cem."+format)
	}

	if err := writeConfigAtomic(outPath, output); err != nil {
		return err
	}

	logging.Success("Config written to %s", outPath)

	if hasPkgJSON && interactive && !yes {
		if err := offerPackageJSONUpdate(root, cfg); err != nil {
			logging.Warning("package.json update: %v", err)
		}
	}

	if interactive && !yes {
		mcpConfirmed, mcpErr := tui.Confirm("Set up AI tool integration?", false)
		if mcpErr == nil && mcpConfirmed {
			logging.Info("Run: cem config mcp")
		}
	}

	return nil
}

func showWelcomePage() error {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}

	logo := centerBlock(strings.TrimRight(selectASCIILogo(width, height), "\n"), width)

	km := huh.NewDefaultKeyMap()
	km.Quit = key.NewBinding(key.WithKeys("esc", "ctrl+c"), key.WithHelp("esc", "quit"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title(logo).
				Description(`CEM - The Custom Elements Multitool

Use this interactive wizard to generate or update your cem config files
`).
				Next(true),
		),
	).WithKeyMap(km)

	return tui.WrapAbort(form.Run())
}
