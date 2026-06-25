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
	"sync"
	"syscall"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

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

// Mirrors internal/config.CemConfig with omitempty tags for clean output.
// Update these types when the canonical config schema changes.
type initConfig struct {
	SourceControlRootUrl string              `yaml:"sourceControlRootUrl,omitempty" json:"sourceControlRootUrl,omitempty"`
	Generate             *initGenerateConfig `yaml:"generate,omitempty" json:"generate,omitempty"`
	Serve                *initServeConfig    `yaml:"serve,omitempty" json:"serve,omitempty"`
}

type initGenerateConfig struct {
	Files         []string                `yaml:"files,omitempty" json:"files,omitempty"`
	Exclude       []string                `yaml:"exclude,omitempty" json:"exclude,omitempty"`
	Output        string                  `yaml:"output,omitempty" json:"output,omitempty"`
	DesignTokens  *initDesignTokensConfig `yaml:"designTokens,omitempty" json:"designTokens,omitempty"`
	DemoDiscovery *initDemoDiscoveryConfig `yaml:"demoDiscovery,omitempty" json:"demoDiscovery,omitempty"`
}

type initDemoDiscoveryConfig struct {
	FileGlob    string `yaml:"fileGlob,omitempty" json:"fileGlob,omitempty"`
	URLPattern  string `yaml:"urlPattern,omitempty" json:"urlPattern,omitempty"`
	URLTemplate string `yaml:"urlTemplate,omitempty" json:"urlTemplate,omitempty"`
}

type initDesignTokensConfig struct {
	Spec   string `yaml:"spec,omitempty" json:"spec,omitempty"`
	Prefix string `yaml:"prefix,omitempty" json:"prefix,omitempty"`
}

type initServeConfig struct {
	Port       int                  `yaml:"port,omitempty" json:"port,omitempty"`
	ImportMap  *initImportMapConfig `yaml:"importMap,omitempty" json:"importMap,omitempty"`
	Transforms *initTransformsConfig `yaml:"transforms,omitempty" json:"transforms,omitempty"`
	Demos      *initDemosConfig     `yaml:"demos,omitempty" json:"demos,omitempty"`
}

type initImportMapConfig struct {
	Generate     bool   `yaml:"generate,omitempty" json:"generate,omitempty"`
	OverrideFile string `yaml:"overrideFile,omitempty" json:"overrideFile,omitempty"`
}

type initTransformsConfig struct {
	TypeScript *initTypeScriptConfig `yaml:"typescript,omitempty" json:"typescript,omitempty"`
	CSS        *initCSSConfig        `yaml:"css,omitempty" json:"css,omitempty"`
}

type initTypeScriptConfig struct {
	Enabled bool   `yaml:"enabled,omitempty" json:"enabled,omitempty"`
	Target  string `yaml:"target,omitempty" json:"target,omitempty"`
}

type initCSSConfig struct {
	Enabled bool     `yaml:"enabled,omitempty" json:"enabled,omitempty"`
	Include []string `yaml:"include,omitempty" json:"include,omitempty"`
	Exclude []string `yaml:"exclude,omitempty" json:"exclude,omitempty"`
}

type initDemosConfig struct {
	Rendering string `yaml:"rendering,omitempty" json:"rendering,omitempty"`
}

func runConfigInit(cmd *cobra.Command, args []string) error {
	yes, _ := cmd.Flags().GetBool("yes")
	format, _ := cmd.Flags().GetString("format")
	interactive := term.IsTerminal(int(os.Stdin.Fd()))

	if !interactive && !yes {
		return fmt.Errorf("no TTY detected; use --yes to accept defaults")
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

	var existing *IC.CemConfig
	var existingFormat string
	if cfgPath := IC.FindConfigFile(root); cfgPath != "" {
		existing, err = IC.LoadConfig(cfgPath)
		if err != nil {
			logging.Warning("Failed to load existing config: %v", err)
		} else {
			existingFormat = IC.FormatFromPath(cfgPath)
			logging.Info("Found existing config: %s", cfgPath)
		}
	}

	osFS := &platform.OSFileSystem{}
	hasPkgJSON := osFS.Exists(filepath.Join(root, "package.json"))
	dirFS := os.DirFS(root)

	cfg := &initConfig{}

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

	if existing != nil && existingFormat != "" {
		format = existingFormat
	} else if interactive && !yes {
		selected, selectErr := tui.Select("Output format", tui.StringOptions("yaml", "json", "jsonc"))
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

	output, marshalErr := marshalInitConfig(cfg, format)
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

	ext := format
	outPath := filepath.Join(root, ".config", "cem."+ext)
	if existing != nil && existing.ConfigFile != "" {
		outPath = existing.ConfigFile
	}

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

func initSourceFiles(cfg *initConfig, fsys fs.FS, existing *IC.CemConfig, interactive, yes bool) error {
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
		title := "Source file patterns (comma-separated)"
		if ann := fv.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, inputErr := tui.TextInput(title, fv.Value())
		if inputErr != nil {
			if errors.Is(inputErr, tui.ErrCancelled) {
				return inputErr
			}
			return inputErr
		}
		if input == "" {
			input = fv.Value()
		}
		files = splitCommaList(input)
	}

	if cfg.Generate == nil {
		cfg.Generate = &initGenerateConfig{}
	}
	cfg.Generate.Files = files
	return nil
}

func initOutput(cfg *initConfig, root string, hasPkgJSON bool, existing *IC.CemConfig, interactive, yes bool) error {
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
		title := "Output file"
		if ann := fv.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, err := tui.TextInput(title, fv.Value())
		if err != nil {
			if errors.Is(err, tui.ErrCancelled) {
				return err
			}
			return err
		}
		if input == "" {
			input = fv.Value()
		}
		output = input
	}

	if cfg.Generate == nil {
		cfg.Generate = &initGenerateConfig{}
	}
	cfg.Generate.Output = output
	return nil
}

func initSourceControlURL(cfg *initConfig, root string, existing *IC.CemConfig, interactive, yes bool) error {
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
		title := "Source control root URL"
		if ann := fv.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, err := tui.TextInput(title, val)
		if err != nil {
			if errors.Is(err, tui.ErrCancelled) {
				return err
			}
			return err
		}
		if input != "" {
			val = input
		}
	}

	cfg.SourceControlRootUrl = val
	return nil
}

func initDemoDiscovery(cfg *initConfig, fsys fs.FS, existing *IC.CemConfig, interactive, yes bool) error {
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

	if globFV.Value() == "" && patternFV.Value() == "" {
		if interactive && !yes {
			logging.Info("No demo files detected.")
		}
		return nil
	}

	if cfg.Generate == nil {
		cfg.Generate = &initGenerateConfig{}
	}

	dd := &initDemoDiscoveryConfig{}

	if yes {
		dd.FileGlob = globFV.Value()
		dd.URLPattern = patternFV.Value()
		dd.URLTemplate = existingTemplate
	} else if interactive {
		title := "Demo file glob"
		if ann := globFV.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, inputErr := tui.TextInput(title, globFV.Value())
		if inputErr != nil {
			if errors.Is(inputErr, tui.ErrCancelled) {
				return inputErr
			}
			return inputErr
		}
		if input != "" {
			dd.FileGlob = input
		} else {
			dd.FileGlob = globFV.Value()
		}

		title = "Demo URL pattern"
		if ann := patternFV.Annotation(); ann != "" {
			title += " (" + ann + ")"
		}
		input, inputErr = tui.TextInput(title, patternFV.Value())
		if inputErr != nil {
			if errors.Is(inputErr, tui.ErrCancelled) {
				return inputErr
			}
			return inputErr
		}
		if input != "" {
			dd.URLPattern = input
		} else {
			dd.URLPattern = patternFV.Value()
		}

		templatePlaceholder := existingTemplate
		if templatePlaceholder == "" {
			templatePlaceholder = "https://example.com/{{.tag}}/{{.demo}}/"
		}
		input, inputErr = tui.TextInput("Demo URL template (or leave empty to skip)", templatePlaceholder)
		if inputErr != nil {
			if errors.Is(inputErr, tui.ErrCancelled) {
				return inputErr
			}
			return inputErr
		}
		dd.URLTemplate = input
	}

	cfg.Generate.DemoDiscovery = dd
	return nil
}

func initDesignTokens(cfg *initConfig, root string, fsys platform.FileSystem, existing *IC.CemConfig, interactive, yes bool) error {
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
		wantTokens, confirmErr := tui.Confirm("Configure design tokens?", false)
		if confirmErr != nil || !wantTokens {
			return nil
		}
		input, inputErr := tui.TextInput("Design tokens spec path", "")
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
		title := "Token prefix"
		if existingPrefix != "" {
			title += " (existing: " + existingPrefix + ")"
		}
		input, inputErr := tui.TextInput(title, existingPrefix)
		if inputErr != nil {
			if errors.Is(inputErr, tui.ErrCancelled) {
				return inputErr
			}
			return inputErr
		}
		if input != "" {
			prefix = input
		}
	}

	if cfg.Generate == nil {
		cfg.Generate = &initGenerateConfig{}
	}
	cfg.Generate.DesignTokens = &initDesignTokensConfig{
		Spec:   spec,
		Prefix: prefix,
	}
	return nil
}

func initServe(cfg *initConfig, root string, fsys fs.FS, existing *IC.CemConfig, interactive, yes bool) error {
	hasExistingServe := existing != nil && (existing.Serve.Port != 0 ||
		existing.Serve.ImportMap.Generate ||
		existing.Serve.ImportMap.OverrideFile != "" ||
		existing.Serve.Demos.Rendering != "" ||
		existing.Serve.Transforms.CSS.Enabled ||
		existing.Serve.Transforms.TypeScript.Enabled ||
		len(existing.Serve.URLRewrites) > 0 ||
		existing.Serve.OpenBrowser != nil)

	if !hasExistingServe && yes {
		cfg.Serve = &initServeConfig{
			Port:      8000,
			ImportMap: &initImportMapConfig{Generate: true},
			Demos:     &initDemosConfig{Rendering: "shadow"},
		}
		return nil
	}

	if !hasExistingServe && interactive {
		want, err := tui.Confirm("Configure serve settings?", true)
		if err != nil || !want {
			return nil
		}
	}

	serve := &initServeConfig{}

	existingPort := 8000
	if existing != nil && existing.Serve.Port != 0 {
		existingPort = existing.Serve.Port
	}

	if interactive && !yes {
		portStr, err := tui.TextInput("Port", strconv.Itoa(existingPort))
		if err != nil {
			if errors.Is(err, tui.ErrCancelled) {
				return err
			}
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
		serve.Port = port
	} else {
		serve.Port = existingPort
	}

	existingRendering := "shadow"
	if existing != nil && existing.Serve.Demos.Rendering != "" {
		existingRendering = existing.Serve.Demos.Rendering
	}

	if interactive && !yes {
		rendering, err := tui.Select("Demo rendering mode", tui.StringOptions("shadow", "light"))
		if err != nil {
			if errors.Is(err, tui.ErrCancelled) {
				return err
			}
			return err
		}
		serve.Demos = &initDemosConfig{Rendering: rendering}
	} else {
		serve.Demos = &initDemosConfig{Rendering: existingRendering}
	}

	existingImportMapGen := true
	if existing != nil {
		existingImportMapGen = existing.Serve.ImportMap.Generate
	}

	if interactive && !yes {
		gen, err := tui.Confirm("Auto-generate import maps?", existingImportMapGen)
		if err != nil {
			if errors.Is(err, tui.ErrCancelled) {
				return err
			}
			return err
		}
		if gen {
			serve.ImportMap = &initImportMapConfig{Generate: true}
		}
	} else if existingImportMapGen {
		serve.ImportMap = &initImportMapConfig{Generate: true}
	}

	var sourcePatterns []string
	if cfg.Generate != nil {
		sourcePatterns = cfg.Generate.Files
	}
	cssInclude, cssExclude := detectCSSPatterns(sourcePatterns, fsys)

	var existingCSSInclude []string
	var existingCSSExclude []string
	if existing != nil && existing.Serve.Transforms.CSS.Enabled {
		existingCSSInclude = existing.Serve.Transforms.CSS.Include
		existingCSSExclude = existing.Serve.Transforms.CSS.Exclude
	}

	if len(existingCSSInclude) > 0 {
		cssInclude = existingCSSInclude
	}
	if len(existingCSSExclude) > 0 {
		cssExclude = existingCSSExclude
	}

	if len(cssInclude) > 0 {
		enableCSS := true
		if interactive && !yes {
			logging.Info("Detected CSS files: %s", strings.Join(cssInclude, ", "))
			var confirmErr error
			enableCSS, confirmErr = tui.Confirm("Enable CSS transforms?", true)
			if confirmErr != nil {
				if errors.Is(confirmErr, tui.ErrCancelled) {
					return confirmErr
				}
				return confirmErr
			}
		}
		if enableCSS {
			serve.Transforms = &initTransformsConfig{
				CSS: &initCSSConfig{
					Enabled: true,
					Include: cssInclude,
					Exclude: cssExclude,
				},
			}
		}
	}

	hasTsConfig, _ := detectTypeScript(root, &platform.OSFileSystem{})
	if hasTsConfig && interactive && !yes {
		logging.Info("tsconfig.json detected. URL rewrites (rootDir/outDir) are auto-detected at serve time.")
	}

	cfg.Serve = serve
	return nil
}

func marshalInitConfig(cfg *initConfig, format string) ([]byte, error) {
	switch format {
	case "json":
		return json.MarshalIndent(cfg, "", "  ")
	case "jsonc":
		data, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return nil, err
		}
		return append([]byte("// Generated by cem config init\n"), data...), nil
	default:
		return yaml.Marshal(cfg)
	}
}

func offerPackageJSONUpdate(root string, cfg *initConfig) error {
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

	output := "custom-elements.json"
	if cfg.Generate != nil && cfg.Generate.Output != "" {
		output = cfg.Generate.Output
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

	// Insert before the last closing brace, preserving existing formatting
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

	if after, ok := strings.CutPrefix(rawURL, "git@"); ok {
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
		// look for demo/ directory in path
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

	repoURL := normalizeGitURL(raw)

	if ghPath, ghErr := exec.LookPath("gh"); ghErr == nil && ghPath != "" {
		repoURL = detectForkUpstream(root, repoURL)
	}

	branch := detectDefaultBranch(root)
	return repoURL + "tree/" + branch + "/", nil
}

func detectForkUpstream(root, fallback string) string {
	cmd := exec.Command("gh", "repo", "view", "--json", "parent")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return fallback
	}

	var result struct {
		Parent struct {
			URL string `json:"url"`
		} `json:"parent"`
	}
	if json.Unmarshal(out, &result) != nil || result.Parent.URL == "" {
		return fallback
	}

	upstream := normalizeGitURL(result.Parent.URL)
	if upstream == "" {
		return fallback
	}
	return upstream
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
