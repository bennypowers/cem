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

	if err := initSourceFiles(cfg, dirFS, existing, interactive, yes); err != nil {
		if errors.Is(err, tui.ErrCancelled) {
			return nil
		}
		return err
	}

	if err := initOutput(cfg, root, hasPkgJSON, existing, interactive, yes); err != nil {
		if errors.Is(err, tui.ErrCancelled) {
			return nil
		}
		return err
	}

	if err := initSourceControlURL(cfg, root, existing, interactive, yes); err != nil {
		if errors.Is(err, tui.ErrCancelled) {
			return nil
		}
		return err
	}

	if err := initDemoDiscovery(cfg, dirFS, existing, interactive, yes); err != nil {
		if errors.Is(err, tui.ErrCancelled) {
			return nil
		}
		return err
	}

	if hasPkgJSON {
		if err := initDesignTokens(cfg, root, osFS, existing, interactive, yes); err != nil {
			if errors.Is(err, tui.ErrCancelled) {
				return nil
			}
			return err
		}
	}

	if err := initServe(cfg, root, dirFS, existing, interactive, yes); err != nil {
		if errors.Is(err, tui.ErrCancelled) {
			return nil
		}
		return err
	}

	if cfgPath != "" && existingFormat != "" {
		format = existingFormat
	} else if interactive && !yes {
		selected, selectErr := tui.Select("Output format", tui.StringOptions(validFormats...))
		if selectErr != nil {
			if errors.Is(selectErr, tui.ErrCancelled) {
				return nil
			}
			return selectErr
		}
		format = selected
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

func initSourceFiles(cfg *IC.CemConfig, fsys fs.FS, existing *IC.CemConfig, interactive, yes bool) error {
	detected, err := detectSourceFiles(fsys)
	if err != nil {
		return err
	}

	var existingVal string
	if existing != nil && len(existing.Generate.Files) > 0 {
		existingVal = strings.Join(existing.Generate.Files, ", ")
	}

	fv := fieldValue{
		Existing: existingVal,
		Detected: strings.Join(detected, ", "),
		Fallback: "**/*.ts",
	}

	var files []string
	if yes {
		files = splitCommaList(fv.Value())
	} else if interactive {
		title := "Glob patterns for source files to scan for custom element definitions (comma-separated)"
		if ann := fv.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, inputErr := tui.TextInput(title, fv.Value())
		if inputErr != nil {
			return inputErr
		}
		if input == "" {
			input = fv.Value()
		}
		files = splitCommaList(input)
	}

	cfg.Generate.Files = files
	return nil
}

func initOutput(cfg *IC.CemConfig, root string, hasPkgJSON bool, existing *IC.CemConfig, interactive, yes bool) error {
	var detected string
	if hasPkgJSON {
		data, err := os.ReadFile(filepath.Join(root, "package.json"))
		if err == nil {
			var pkg struct {
				CustomElements string `json:"customElements"`
			}
			if json.Unmarshal(data, &pkg) == nil && pkg.CustomElements != "" {
				detected = pkg.CustomElements
			}
		}
	}

	var existingVal string
	if existing != nil && existing.Generate.Output != "" {
		existingVal = existing.Generate.Output
	}

	fv := fieldValue{
		Existing: existingVal,
		Detected: detected,
		Fallback: "custom-elements.json",
	}

	var output string
	if yes {
		output = fv.Value()
	} else if interactive {
		title := "Output path for the generated custom-elements.json manifest"
		if ann := fv.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, err := tui.TextInput(title, fv.Value())
		if err != nil {
			return err
		}
		if input == "" {
			input = fv.Value()
		}
		output = input
	}

	cfg.Generate.Output = output
	return nil
}

func initSourceControlURL(cfg *IC.CemConfig, root string, existing *IC.CemConfig, interactive, yes bool) error {
	detected, _ := detectGitRemote(root)

	var existingVal string
	if existing != nil {
		existingVal = existing.SourceControlRootUrl
	}

	fv := fieldValue{
		Existing: existingVal,
		Detected: detected,
	}

	val := fv.Value()
	if val == "" && yes {
		return nil
	}

	if interactive && !yes {
		title := "Base URL for linking manifest entries to source files (e.g. GitHub tree URL)"
		if ann := fv.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, err := tui.TextInput(title, val)
		if err != nil {
			return err
		}
		if input != "" {
			val = input
		}
	}

	cfg.SourceControlRootUrl = val
	return nil
}

func initDemoDiscovery(cfg *IC.CemConfig, fsys fs.FS, existing *IC.CemConfig, interactive, yes bool) error {
	detectedGlob, detectedPattern, err := detectDemoFiles(fsys)
	if err != nil {
		return err
	}

	var existingGlob, existingPattern, existingTemplate string
	if existing != nil {
		existingGlob = existing.Generate.DemoDiscovery.FileGlob
		existingPattern = existing.Generate.DemoDiscovery.URLPattern
		existingTemplate = existing.Generate.DemoDiscovery.URLTemplate
	}

	globFV := fieldValue{Existing: existingGlob, Detected: detectedGlob}
	patternFV := fieldValue{Existing: existingPattern, Detected: detectedPattern}

	hasDemos := globFV.Value() != "" || patternFV.Value() != ""

	if !hasDemos && interactive && !yes {
		logging.Info("No demo files detected.")
		return nil
	}

	if !hasDemos {
		return nil
	}

	if interactive && !yes {
		want, confirmErr := tui.Confirm(
			"Demos are HTML partials that showcase your elements. "+
				"The dev server discovers and serves them with live reload. "+
				"Configure demo discovery?", true)
		if confirmErr != nil {
			return confirmErr
		}
		if !want {
			return nil
		}
	}

	if yes {
		cfg.Generate.DemoDiscovery.FileGlob = globFV.Value()
		cfg.Generate.DemoDiscovery.URLPattern = patternFV.Value()
		if existingTemplate != "" {
			cfg.Generate.DemoDiscovery.URLTemplate = existingTemplate
		}
		return nil
	}

	title := "Glob pattern to find demo HTML files"
	if ann := globFV.Annotation(); ann != "" {
		title += " (" + ann + ")"
	}
	input, inputErr := tui.TextInput(title, globFV.Value())
	if inputErr != nil {
		return inputErr
	}
	if input != "" {
		cfg.Generate.DemoDiscovery.FileGlob = input
	} else {
		cfg.Generate.DemoDiscovery.FileGlob = globFV.Value()
	}

	title = "URL pattern with named parameters for matching demo file paths to elements (e.g. elements/:tag/demo/:demo.html)"
	if ann := patternFV.Annotation(); ann != "" {
		title += " (" + ann + ")"
	}
	input, inputErr = tui.TextInput(title, patternFV.Value())
	if inputErr != nil {
		return inputErr
	}
	if input != "" {
		cfg.Generate.DemoDiscovery.URLPattern = input
	} else {
		cfg.Generate.DemoDiscovery.URLPattern = patternFV.Value()
	}

	templatePlaceholder := existingTemplate
	if templatePlaceholder == "" {
		templatePlaceholder = "https://example.com/{{.tag}}/{{.demo}}/"
	}
	input, inputErr = tui.TextInput(
		"Go template for generating public demo URLs from matched parameters (leave empty to skip)",
		templatePlaceholder)
	if inputErr != nil {
		return inputErr
	}
	cfg.Generate.DemoDiscovery.URLTemplate = input
	return nil
}

func initDesignTokens(cfg *IC.CemConfig, root string, fsys platform.FileSystem, existing *IC.CemConfig, interactive, yes bool) error {
	tokens, _ := detectDesignTokens(root, fsys)

	var existingSpec, existingPrefix string
	if existing != nil {
		existingSpec = existing.Generate.DesignTokens.Spec
		existingPrefix = existing.Generate.DesignTokens.Prefix
	}

	var spec string
	switch {
	case existingSpec != "":
		spec = existingSpec
	case len(tokens) == 1:
		spec = tokens[0].Specifier
		if spec == "" {
			spec = tokens[0].Path
		}
	case len(tokens) > 1 && interactive && !yes:
		options := make([]huh.Option[string], len(tokens))
		for i, t := range tokens {
			label := t.Specifier
			if label == "" {
				label = t.Path
			}
			options[i] = huh.NewOption(label, label)
		}
		selected, selectErr := tui.Select("Design tokens spec", options)
		if selectErr != nil {
			if errors.Is(selectErr, tui.ErrCancelled) {
				return nil
			}
			return selectErr
		}
		spec = selected
	}

	if spec == "" && interactive && !yes {
		wantTokens, confirmErr := tui.Confirm(
			"Document CSS custom properties from DTCG design token files?", false)
		if confirmErr != nil || !wantTokens {
			return nil
		}
		input, inputErr := tui.TextInput("Path or specifier for the design token spec file", "")
		if inputErr != nil || input == "" {
			return nil
		}
		spec = input
	}

	if spec == "" && yes {
		if existingSpec != "" {
			spec = existingSpec
		} else {
			return nil
		}
	}

	if spec == "" {
		return nil
	}

	prefix := existingPrefix
	if interactive && !yes {
		title := "CSS custom property prefix for design tokens (e.g. 'my-ds' for --my-ds-color-primary)"
		if existingPrefix != "" {
			title += " (existing: " + existingPrefix + ")"
		}
		input, inputErr := tui.TextInput(title, existingPrefix)
		if inputErr != nil {
			return inputErr
		}
		if input != "" {
			prefix = input
		}
	}

	cfg.Generate.DesignTokens.Spec = spec
	cfg.Generate.DesignTokens.Prefix = prefix
	return nil
}

func initServe(cfg *IC.CemConfig, root string, fsys fs.FS, existing *IC.CemConfig, interactive, yes bool) error {
	hasExistingServe := existing != nil && (existing.Serve.Port != 0 ||
		existing.Serve.ImportMap.Generate ||
		existing.Serve.ImportMap.OverrideFile != "" ||
		existing.Serve.Demos.Rendering != "" ||
		existing.Serve.Transforms.CSS.Enabled ||
		existing.Serve.Transforms.TypeScript.Enabled ||
		len(existing.Serve.URLRewrites) > 0 ||
		existing.Serve.OpenBrowser != nil)

	if !hasExistingServe && yes {
		cfg.Serve.Port = 8000
		cfg.Serve.ImportMap.Generate = true
		cfg.Serve.Demos.Rendering = "shadow"
		return nil
	}

	if !hasExistingServe && interactive {
		want, err := tui.Confirm(
			"Configure the cem dev server? "+
				"It serves demos with live reload, TypeScript transforms, and import maps.", true)
		if err != nil || !want {
			return nil
		}
	}

	existingPort := 8000
	if existing != nil && existing.Serve.Port != 0 {
		existingPort = existing.Serve.Port
	}

	if interactive && !yes {
		portStr, err := tui.TextInput("Dev server port", strconv.Itoa(existingPort))
		if err != nil {
			return err
		}
		if portStr == "" {
			portStr = strconv.Itoa(existingPort)
		}
		port, parseErr := strconv.Atoi(portStr)
		if parseErr != nil {
			return fmt.Errorf("invalid port: %s", portStr)
		}
		if port < 1 || port > 65535 {
			return fmt.Errorf("port must be between 1 and 65535, got %d", port)
		}
		cfg.Serve.Port = port
	} else {
		cfg.Serve.Port = existingPort
	}

	existingRendering := "shadow"
	if existing != nil && existing.Serve.Demos.Rendering != "" {
		existingRendering = existing.Serve.Demos.Rendering
	}

	if interactive && !yes {
		rendering, err := tui.Select("Demo rendering mode", tui.StringOptions("shadow", "light"))
		if err != nil {
			return err
		}
		cfg.Serve.Demos.Rendering = rendering
	} else {
		cfg.Serve.Demos.Rendering = existingRendering
	}

	existingImportMapGen := true
	if existing != nil {
		existingImportMapGen = existing.Serve.ImportMap.Generate
	}

	if interactive && !yes {
		gen, err := tui.Confirm("Auto-generate import maps?", existingImportMapGen)
		if err != nil {
			return err
		}
		cfg.Serve.ImportMap.Generate = gen
	} else {
		cfg.Serve.ImportMap.Generate = existingImportMapGen
	}

	sourcePatterns := cfg.Generate.Files
	cssInclude, cssExclude := detectCSSPatterns(sourcePatterns, fsys)

	if existing != nil && existing.Serve.Transforms.CSS.Enabled {
		if len(existing.Serve.Transforms.CSS.Include) > 0 {
			cssInclude = existing.Serve.Transforms.CSS.Include
		}
		if len(existing.Serve.Transforms.CSS.Exclude) > 0 {
			cssExclude = existing.Serve.Transforms.CSS.Exclude
		}
	}

	if len(cssInclude) > 0 {
		enableCSS := true
		if interactive && !yes {
			logging.Info("Detected CSS files: %s", strings.Join(cssInclude, ", "))
			var confirmErr error
			enableCSS, confirmErr = tui.Confirm("Enable CSS transforms?", true)
			if confirmErr != nil {
				return confirmErr
			}
		}
		if enableCSS {
			cfg.Serve.Transforms.CSS.Enabled = true
			cfg.Serve.Transforms.CSS.Include = cssInclude
			cfg.Serve.Transforms.CSS.Exclude = cssExclude
		}
	}

	hasTsConfig, _ := detectTypeScript(root, &platform.OSFileSystem{})
	if hasTsConfig && interactive && !yes {
		logging.Info("tsconfig.json detected. URL rewrites (rootDir/outDir) are auto-detected at serve time.")
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
