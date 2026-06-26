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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"syscall"

	asimconfig "bennypowers.dev/asimonim/config"
	"bennypowers.dev/asimonim/specifier"
	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/set"
	"bennypowers.dev/cem/internal/tui"
	W "bennypowers.dev/cem/internal/workspace"
	"charm.land/bubbles/v2/key"
	"charm.land/huh/v2"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
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
	detectedFiles, _ := detectSourceFiles(dirFS)
	detectedGitRemote, _ := detectGitRemote(root)
	detectedGlob, detectedPattern, _ := detectDemoFiles(dirFS)
	var detectedTokens []*specifier.ResolvedFile
	if hasPkgJSON {
		detectedTokens, _ = detectDesignTokens(root, osFS)
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
		Existing: strings.Join(cfg.Generate.Files, ", "),
		Detected: strings.Join(detectedFiles, ", "),
		Fallback: "**/*.ts",
	}
	outputFV := fieldValue{
		Existing: cfg.Generate.Output,
		Detected: detectedPkgCE,
		Fallback: "custom-elements.json",
	}
	scurlFV := fieldValue{
		Existing: cfg.SourceControlRootUrl,
		Detected: detectedGitRemote,
	}
	globFV := fieldValue{Existing: cfg.Generate.DemoDiscovery.FileGlob, Detected: detectedGlob}
	patternFV := fieldValue{Existing: cfg.Generate.DemoDiscovery.URLPattern, Detected: detectedPattern}

	// Form field values (bound to huh fields)
	filesInput := filesFV.Value()
	outputInput := outputFV.Value()
	scurlInput := scurlFV.Value()

	if interactive && !yes {
		// Build the main form
		var groups []*huh.Group

		// Group 1: Generate settings
		filesDesc := "Glob patterns for files to scan for custom element definitions (comma-separated)"
		if ann := filesFV.Annotation(); ann != "" {
			filesDesc += "\n" + ann
		}
		outputDesc := "Where to write the generated custom-elements.json manifest"
		if ann := outputFV.Annotation(); ann != "" {
			outputDesc += "\n" + ann
		}
		scurlDesc := "Base URL for linking manifest entries to source files"
		if ann := scurlFV.Annotation(); ann != "" {
			scurlDesc += "\n" + ann
		}
		groups = append(groups, huh.NewGroup(
			huh.NewInput().
				Title("Source file patterns").
				Description(filesDesc).
				Prompt("?").
				Placeholder(filesFV.Fallback).
				Value(&filesInput),
			huh.NewInput().
				Title("Manifest output path").
				Description(outputDesc).
				Prompt("?").
				Placeholder(outputFV.Fallback).
				Value(&outputInput),
			huh.NewInput().
				Title("Source control root URL").
				Description(scurlDesc).
				Prompt("?").
				Placeholder("https://github.com/user/repo/tree/main/").
				Value(&scurlInput),
		).Title("Generate"))

		// Group 2: Demo discovery (conditional)
		hasDemos := globFV.Value() != "" || patternFV.Value() != ""
		if hasDemos {
			demoGlob := globFV.Value()
			demoPattern := patternFV.Value()
			demoTemplate := cfg.Generate.DemoDiscovery.URLTemplate

			globDesc := "Glob pattern to find demo HTML files"
			if ann := globFV.Annotation(); ann != "" {
				globDesc += "\n" + ann
			}
			patternDesc := "URL pattern with named params for matching demo paths to elements"
			if ann := patternFV.Annotation(); ann != "" {
				patternDesc += "\n" + ann
			}

			groups = append(groups, huh.NewGroup(
				huh.NewInput().
					Title("Demo file glob").
					Description(globDesc).
					Value(&demoGlob),
				huh.NewInput().
					Title("Demo URL pattern").
					Description(patternDesc).
					Placeholder("elements/:tag/demo/:demo.html").
					Value(&demoPattern),
				huh.NewInput().
					Title("Demo URL template").
					Description("Go template for generating public demo URLs from matched parameters (leave empty to skip)").
					Placeholder("https://example.com/{{.tag}}/{{.demo}}/").
					Value(&demoTemplate),
			).Title("Demo Discovery").
				Description("Demos are HTML partials that showcase your elements. "+
					"The dev server discovers and serves them with live reload."))

			defer func() {
				cfg.Generate.DemoDiscovery.FileGlob = demoGlob
				cfg.Generate.DemoDiscovery.URLPattern = demoPattern
				cfg.Generate.DemoDiscovery.URLTemplate = demoTemplate
			}()
		}

		// Group 3: Design tokens (conditional)
		var tokenSpec, tokenPrefix string
		if len(detectedTokens) > 0 || cfg.Generate.DesignTokens.Spec != "" {
			tokenSpec = cfg.Generate.DesignTokens.Spec
			if tokenSpec == "" && len(detectedTokens) == 1 {
				tokenSpec = detectedTokens[0].Specifier
				if tokenSpec == "" {
					tokenSpec = detectedTokens[0].Path
				}
			}
			tokenPrefix = cfg.Generate.DesignTokens.Prefix

			groups = append(groups, huh.NewGroup(
				huh.NewInput().
					Title("Design token spec").
					Description("Path or specifier for the DTCG design token file").
					Value(&tokenSpec),
				huh.NewInput().
					Title("Token prefix").
					Description("CSS custom property prefix (e.g. 'my-ds' for --my-ds-color-primary)").
					Value(&tokenPrefix),
			).Title("Design Tokens").
				Description("Document CSS custom properties from DTCG design token files."))

			defer func() {
				cfg.Generate.DesignTokens.Spec = tokenSpec
				cfg.Generate.DesignTokens.Prefix = tokenPrefix
			}()
		}

		// Group 4: Serve settings
		portStr := "8000"
		if cfg.Serve.Port != 0 {
			portStr = strconv.Itoa(cfg.Serve.Port)
		}
		rendering := cfg.Serve.Demos.Rendering
		if rendering == "" {
			rendering = "shadow"
		}
		importMapGen := cfg.Serve.ImportMap.Generate
		if existing == nil {
			importMapGen = true
		}
		enableCSS := cfg.Serve.Transforms.CSS.Enabled || len(detectedCSSInclude) > 0

		serveFields := []huh.Field{
			huh.NewInput().
				Title("Dev server port").
				Value(&portStr),
			huh.NewSelect[string]().
				Title("Demo rendering mode").
				Options(huh.NewOption("Shadow DOM", "shadow"), huh.NewOption("Light DOM", "light")).
				Value(&rendering),
			huh.NewConfirm().
				Title("Auto-generate import maps?").
				Value(&importMapGen),
		}

		if len(detectedCSSInclude) > 0 || cfg.Serve.Transforms.CSS.Enabled {
			serveFields = append(serveFields,
				huh.NewConfirm().
					Title("Enable CSS transforms?").
					Description("Transform CSS files to JavaScript modules ("+strings.Join(detectedCSSInclude, ", ")+")").
					Value(&enableCSS),
			)
		}

		groups = append(groups, huh.NewGroup(serveFields...).
			Title("Dev Server").
			Description("The cem dev server serves demos with live reload, TypeScript transforms, and import maps."))

		// Group 5: Output format (if no existing config)
		if cfgPath == "" {
			groups = append(groups, huh.NewGroup(
				huh.NewSelect[string]().
					Title("Output format").
					Options(
						huh.NewOption("YAML", "yaml"),
						huh.NewOption("JSON", "json"),
						huh.NewOption("JSON with comments", "jsonc"),
					).
					Value(&format),
			).Title("Output"))
		}

		// Run the form
		km := huh.NewDefaultKeyMap()
		km.Quit = key.NewBinding(key.WithKeys("esc", "ctrl+c"), key.WithHelp("esc", "quit"))
		form := huh.NewForm(groups...).WithKeyMap(km)
		if formErr := tui.WrapAbort(form.Run()); formErr != nil {
			if errors.Is(formErr, tui.ErrCancelled) {
				return nil
			}
			return formErr
		}

		// Apply form results
		cfg.Generate.Files = splitCommaList(filesInput)
		cfg.Generate.Output = outputInput
		cfg.SourceControlRootUrl = scurlInput

		port, parseErr := strconv.Atoi(portStr)
		if parseErr != nil || port < 1 || port > 65535 {
			return fmt.Errorf("invalid port: %s", portStr)
		}
		cfg.Serve.Port = port
		cfg.Serve.Demos.Rendering = rendering
		cfg.Serve.ImportMap.Generate = importMapGen

		if enableCSS {
			cssInclude := detectedCSSInclude
			cssExclude := detectedCSSExclude
			if cfg.Serve.Transforms.CSS.Enabled {
				if len(cfg.Serve.Transforms.CSS.Include) > 0 {
					cssInclude = cfg.Serve.Transforms.CSS.Include
				}
				if len(cfg.Serve.Transforms.CSS.Exclude) > 0 {
					cssExclude = cfg.Serve.Transforms.CSS.Exclude
				}
			}
			cfg.Serve.Transforms.CSS.Enabled = true
			cfg.Serve.Transforms.CSS.Include = cssInclude
			cfg.Serve.Transforms.CSS.Exclude = cssExclude
		} else {
			cfg.Serve.Transforms.CSS.Enabled = false
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
		hasExistingServe := existing != nil && (existing.Serve.Port != 0 ||
			existing.Serve.ImportMap.Generate)
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

	// Clear internal-only fields before marshaling
	cfg.ProjectDir = ""
	cfg.ConfigFile = ""
	cfg.PackageName = ""

	output, marshalErr := marshalConfig(cfg, format, existingData)
	if marshalErr != nil {
		return marshalErr
	}

	lang := format
	if lang == "jsonc" {
		lang = "json"
	}

	cmd.Println()
	if err := tui.Highlight(cmd.OutOrStdout(), string(output), lang); err != nil {
		cmd.Println(string(output))
	}
	cmd.Println()

	if interactive && !yes {
		confirmed, confirmErr := tui.Confirm("Write this config?", true)
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

func writeConfigAtomic(outPath string, output []byte) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	tmp, err := os.CreateTemp(filepath.Dir(outPath), ".cem-config-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmp.Name()
	var cleanedUp sync.Once
	cleanup := func() {
		cleanedUp.Do(func() {
			_ = os.Remove(tmpPath)
		})
	}
	defer cleanup()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		cleanup()
		os.Exit(1)
	}()
	defer signal.Stop(sigCh)

	if _, err := tmp.Write(output); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("failed to write config: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	if err := os.Rename(tmpPath, outPath); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func marshalConfig(cfg *IC.CemConfig, format string, existingData []byte) ([]byte, error) {
	switch format {
	case "json":
		data, err := marshalConfigJSON(cfg)
		if err != nil {
			return nil, err
		}
		return data, nil
	case "jsonc":
		data, err := marshalConfigJSON(cfg)
		if err != nil {
			return nil, err
		}
		return append([]byte("// Generated by cem config init\n"), data...), nil
	default:
		return marshalConfigYAML(cfg, existingData)
	}
}

func marshalConfigJSON(cfg *IC.CemConfig) ([]byte, error) {
	var intermediate map[string]any
	raw, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	if json.Unmarshal(raw, &intermediate) != nil {
		return raw, nil
	}
	pruneJSONMap(intermediate)
	return json.MarshalIndent(intermediate, "", "  ")
}

func pruneJSONMap(m map[string]any) {
	for k, v := range m {
		if isZeroJSON(v) {
			delete(m, k)
			continue
		}
		if sub, ok := v.(map[string]any); ok {
			pruneJSONMap(sub)
			if len(sub) == 0 {
				delete(m, k)
			}
		}
	}
}

func isZeroJSON(v any) bool {
	switch val := v.(type) {
	case nil:
		return true
	case string:
		return val == ""
	case float64:
		return val == 0
	case bool:
		return !val
	case []any:
		return len(val) == 0
	case map[string]any:
		return len(val) == 0
	}
	return false
}

func marshalConfigYAML(cfg *IC.CemConfig, existingData []byte) ([]byte, error) {
	var node yaml.Node
	if err := node.Encode(cfg); err != nil {
		return nil, err
	}

	if node.Kind == yaml.MappingNode {
		pruneYAMLNode(&node)
	} else if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		pruneYAMLNode(node.Content[0])
	}

	indent := detectIndent(existingData)

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(indent)
	if err := enc.Encode(&node); err != nil {
		return nil, err
	}
	if err := enc.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func pruneYAMLNode(node *yaml.Node) {
	if node.Kind != yaml.MappingNode {
		return
	}

	var kept []*yaml.Node
	for i := 0; i+1 < len(node.Content); i += 2 {
		key := node.Content[i]
		val := node.Content[i+1]

		if isZeroYAMLNode(val) {
			continue
		}

		if val.Kind == yaml.MappingNode {
			pruneYAMLNode(val)
			if len(val.Content) == 0 {
				continue
			}
		}

		kept = append(kept, key, val)
	}
	node.Content = kept
}

func isZeroYAMLNode(n *yaml.Node) bool {
	switch n.Kind {
	case yaml.ScalarNode:
		if n.Tag == "!!null" {
			return true
		}
		return n.Value == "" || n.Value == "0" || n.Value == "false"
	case yaml.SequenceNode:
		return len(n.Content) == 0
	case yaml.MappingNode:
		return len(n.Content) == 0
	case yaml.AliasNode:
		return n.Alias != nil && isZeroYAMLNode(n.Alias)
	}
	return n.Tag == "!!null"
}

func detectIndent(data []byte) int {
	if len(data) == 0 {
		return 2
	}
	for line := range strings.SplitSeq(string(data), "\n") {
		trimmed := strings.TrimLeft(line, " ")
		indent := len(line) - len(trimmed)
		if indent > 0 {
			return indent
		}
	}
	return 2
}

func offerPackageJSONUpdate(root string, cfg *IC.CemConfig) error {
	pkgPath := filepath.Join(root, "package.json")
	data, err := os.ReadFile(pkgPath)
	if err != nil {
		return nil
	}

	if !json.Valid(data) {
		return nil
	}

	if gjson.GetBytes(data, "customElements").Exists() {
		return nil
	}

	output := cfg.Generate.Output
	if output == "" {
		output = "custom-elements.json"
	}

	confirmed, err := tui.Confirm(fmt.Sprintf("Add \"customElements\": \"%s\" to package.json?", output), true)
	if err != nil || !confirmed {
		return nil
	}

	entry, marshalErr := json.Marshal(output)
	if marshalErr != nil {
		return marshalErr
	}
	insertion := fmt.Sprintf(",\n  \"customElements\": %s", string(entry))

	lastBrace := bytes.LastIndexByte(data, '}')
	if lastBrace < 0 {
		return nil
	}
	var buf bytes.Buffer
	buf.Write(data[:lastBrace])
	buf.WriteString(insertion)
	buf.WriteByte('\n')
	buf.WriteByte('}')
	if lastBrace+1 < len(data) {
		buf.Write(data[lastBrace+1:])
	}

	return os.WriteFile(pkgPath, buf.Bytes(), 0o644)
}

func splitCommaList(s string) []string {
	var result []string
	for part := range strings.SplitSeq(s, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

type fieldValue struct {
	Existing string
	Detected string
	Fallback string
}

func (f fieldValue) Value() string {
	if f.Existing != "" {
		return f.Existing
	}
	if f.Detected != "" {
		return f.Detected
	}
	return f.Fallback
}

func (f fieldValue) Annotation() string {
	if f.Existing != "" && f.Detected != "" && f.Existing != f.Detected {
		return "detected: " + f.Detected
	}
	return ""
}

func normalizeGitURL(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return ""
	}

	var host, path string

	if after, ok := strings.CutPrefix(rawURL, "ssh://"); ok {
		after = strings.TrimPrefix(after, "git@")
		host, path, ok = strings.Cut(after, "/")
		if !ok {
			return rawURL
		}
		// strip optional :port from host
		if h, _, hasPort := strings.Cut(host, ":"); hasPort {
			host = h
		}
	} else if after, ok := strings.CutPrefix(rawURL, "git@"); ok {
		host, path, ok = strings.Cut(after, ":")
		if !ok {
			return rawURL
		}
	} else if after, ok := strings.CutPrefix(rawURL, "https://"); ok {
		host, path, ok = strings.Cut(after, "/")
		if !ok {
			return rawURL
		}
	} else if after, ok := strings.CutPrefix(rawURL, "http://"); ok {
		host, path, ok = strings.Cut(after, "/")
		if !ok {
			return rawURL
		}
	} else {
		return rawURL
	}

	path = strings.TrimSuffix(path, ".git")
	path = strings.TrimSuffix(path, "/")

	return fmt.Sprintf("https://%s/%s/", host, path)
}

var skipDirs = set.NewSet("node_modules", "dist", "test", "demo", ".git", "build", "_site")

func detectSourceFiles(fsys fs.FS) ([]string, error) {
	topDirs := set.NewSet[string]()

	err := platform.WalkDir(fsys, ".", skipDirs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".ts") {
			return nil
		}
		if strings.HasSuffix(path, ".d.ts") {
			return nil
		}

		parts := strings.SplitN(filepath.ToSlash(path), "/", 2)
		if len(parts) >= 2 {
			topDirs.Add(parts[0])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(topDirs) == 0 {
		return nil, nil
	}

	dirs := topDirs.Members()
	slices.Sort(dirs)

	patterns := make([]string, len(dirs))
	for i, dir := range dirs {
		patterns[i] = dir + "/**/*.ts"
	}
	return patterns, nil
}

func detectDemoFiles(fsys fs.FS) (fileGlob string, urlPattern string, err error) {
	topDirs := set.NewSet[string]()

	walkErr := platform.WalkDir(fsys, ".", set.NewSet("node_modules", ".git"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".html") {
			return nil
		}

		parts := strings.Split(filepath.ToSlash(path), "/")
		for i, part := range parts {
			if part == "demo" && i > 0 && i < len(parts)-1 {
				topDirs.Add(parts[0])
				break
			}
		}
		return nil
	})
	if walkErr != nil {
		return "", "", walkErr
	}

	if len(topDirs) == 0 {
		return "", "", nil
	}

	dirs := topDirs.Members()
	slices.Sort(dirs)
	topDir := dirs[0]

	fileGlob = topDir + "/**/demo/*.html"
	urlPattern = topDir + "/:tag/demo/:demo.html"
	return fileGlob, urlPattern, nil
}

func detectCSSPatterns(sourcePatterns []string, fsys fs.FS) (include []string, exclude []string) {
	cssDirs := set.NewSet[string]()

	_ = platform.WalkDir(fsys, ".", skipDirs, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".css") {
			return nil
		}

		parts := strings.SplitN(filepath.ToSlash(path), "/", 2)
		if len(parts) >= 2 {
			topDir := parts[0]
			for _, sp := range sourcePatterns {
				if strings.HasPrefix(sp, topDir+"/") {
					cssDirs.Add(topDir)
					break
				}
			}
		}
		return nil
	})

	if len(cssDirs) == 0 {
		return nil, nil
	}

	dirs := cssDirs.Members()
	slices.Sort(dirs)

	include = make([]string, len(dirs))
	for i, dir := range dirs {
		include[i] = dir + "/**/*.css"
	}

	exclude = []string{"demo/**/*.css", "**/*.min.css"}
	return include, exclude
}

type gitRemoteResult struct {
	url    string
	branch string
}

func detectGitRemote(root string) (string, error) {
	gitCmd := exec.Command("git", "config", "--get", "remote.origin.url")
	gitCmd.Dir = root
	out, err := gitCmd.Output()
	if err != nil {
		return "", nil
	}
	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return "", nil
	}

	result := gitRemoteResult{
		url:    normalizeGitURL(raw),
		branch: detectDefaultBranch(root),
	}

	if ghPath, ghErr := exec.LookPath("gh"); ghErr == nil && ghPath != "" {
		result = detectForkUpstream(root, result)
	}

	return result.url + "tree/" + result.branch + "/", nil
}

func detectForkUpstream(root string, fallback gitRemoteResult) gitRemoteResult {
	ghCmd := exec.Command("gh", "repo", "view", "--json", "parent,defaultBranchRef")
	ghCmd.Dir = root
	out, err := ghCmd.Output()
	if err != nil {
		return fallback
	}

	var result struct {
		Parent struct {
			URL string `json:"url"`
		} `json:"parent"`
		DefaultBranchRef struct {
			Name string `json:"name"`
		} `json:"defaultBranchRef"`
	}
	if json.Unmarshal(out, &result) != nil || result.Parent.URL == "" {
		return fallback
	}

	upstream := normalizeGitURL(result.Parent.URL)
	if upstream == "" {
		return fallback
	}

	branch := result.DefaultBranchRef.Name
	if branch == "" {
		branch = fallback.branch
	}

	return gitRemoteResult{url: upstream, branch: branch}
}

func detectDefaultBranch(root string) string {
	refCmd := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD")
	refCmd.Dir = root
	out, err := refCmd.Output()
	if err == nil {
		ref := strings.TrimSpace(string(out))
		if _, branch, ok := strings.Cut(ref, "refs/remotes/origin/"); ok && branch != "" {
			return branch
		}
	}

	return "main"
}

func detectDesignTokens(root string, fsys platform.FileSystem) ([]*specifier.ResolvedFile, error) {
	return asimconfig.DiscoverDesignTokens(fsys, root)
}

func detectTypeScript(root string, fsys platform.FileSystem) (hasTsConfig bool, target string) {
	_, err := fsys.Stat(filepath.Join(root, "tsconfig.json"))
	if err != nil {
		return false, ""
	}
	return true, ""
}
