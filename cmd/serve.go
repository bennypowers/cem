/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"golang.org/x/term"

	"bennypowers.dev/cem/cmd/config"
	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/serve"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware/routes"
	"bennypowers.dev/cem/serve/middleware/transform"
	"bennypowers.dev/cem/serve/middleware/types"
	servetui "bennypowers.dev/cem/serve/tui"
	W "bennypowers.dev/cem/internal/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveTUILogger *servetui.Logger

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a development server with live reload",
	Long: `Start a development server for custom element demos with:
- Live reload on file changes
- Automatic manifest regeneration
- WebSocket-based browser refresh
- Multiple rendering modes (full UI, shadow DOM, or chromeless)
- Static file serving with CORS`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if term.IsTerminal(int(os.Stdout.Fd())) {
			serveTUILogger = servetui.NewLogger()
			logging.SetMode(logging.ModeServe)
			logging.SetServeSink(serveTUILogger)
		} else {
			serveTUILogger = nil
		}
		return rootPersistentPreRunE(cmd, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}

		port := viper.GetInt("serve.port")
		reload := !viper.GetBool("serve.no-reload")
		targetStr := viper.GetString("serve.target")

		// Load transform configuration
		tsEnabled := viper.GetBool("serve.transforms.typescript.enabled")
		if !viper.IsSet("serve.transforms.typescript.enabled") {
			tsEnabled = true // Default to enabled
		}

		cssEnabled := viper.GetBool("serve.transforms.css.enabled")
		if !viper.IsSet("serve.transforms.css.enabled") {
			cssEnabled = true // Default to enabled
		}

		cssInclude := viper.GetStringSlice("serve.transforms.css.include")
		cssExclude := viper.GetStringSlice("serve.transforms.css.exclude")

		// Load import map configuration
		importMapGenerate := true
		// Config file sets the default value if present
		if viper.IsSet("serve.importMap.generate") {
			importMapGenerate = viper.GetBool("serve.importMap.generate")
		}
		// CLI flag overrides config if explicitly set
		if cmd.Flags().Changed("no-import-map-generate") {
			noGenerate, _ := cmd.Flags().GetBool("no-import-map-generate")
			importMapGenerate = !noGenerate
		}
		importMapOverrideFile := viper.GetString("serve.importMap.overrideFile")

		// Load config-based override (full import map structure)
		var importMapOverride types.ImportMapOverride
		if viper.IsSet("serve.importMap.override.imports") {
			importMapOverride.Imports = viper.GetStringMapString("serve.importMap.override.imports")
		}
		if viper.IsSet("serve.importMap.override.scopes") {
			scopes := viper.GetStringMap("serve.importMap.override.scopes")
			if len(scopes) > 0 {
				importMapOverride.Scopes = make(map[string]map[string]string)
				for scopeKey, scopeVal := range scopes {
					if scopeMap, ok := scopeVal.(map[string]any); ok {
						importMapOverride.Scopes[scopeKey] = make(map[string]string)
						for k, v := range scopeMap {
							if str, ok := v.(string); ok {
								importMapOverride.Scopes[scopeKey][k] = str
							}
						}
					}
				}
			}
		}

		// Get watch ignore patterns from config or flag
		watchIgnore := viper.GetStringSlice("serve.watchIgnore")

		// Get target from transforms config or fallback to --target flag
		configTarget := viper.GetString("serve.transforms.typescript.target")
		if targetStr == "" && configTarget != "" {
			targetStr = configTarget
		}

		var target transform.Target
		if targetStr != "" {
			if !IC.IsValidTarget(targetStr) {
				return fmt.Errorf("invalid target %q: must be one of %s", targetStr, strings.Join(IC.ValidTargets(), ", "))
			}
			target = transform.Target(targetStr)
		} else {
			target = transform.ES2022
		}

		// Get source control root URL from config
		sourceControlRootURL := viper.GetString("sourceControlRootUrl")

		// Compute demo URL prefix from urlTemplate for local route stripping
		demoURLPrefix := routes.DemoURLPrefixFromTemplate(
			viper.GetString("generate.demoDiscovery.urlTemplate"),
		)

		// Get URL rewrites from config
		var urlRewrites []config.URLRewrite
		if err := viper.UnmarshalKey("serve.urlRewrites", &urlRewrites); err != nil {
			return fmt.Errorf("failed to parse serve.urlRewrites: %w", err)
		}

		// Get demo rendering mode from config (default: "light")
		demoRendering := viper.GetString("serve.demos.rendering")
		if demoRendering == "" {
			demoRendering = "light"
		}
		if !IC.IsValidRenderingMode(demoRendering) {
			return fmt.Errorf("invalid demo rendering mode %q: must be one of %s", demoRendering, strings.Join(IC.ValidRenderingModes(), ", "))
		}
		// Create server config
		config := serve.Config{
			Port:                 port,
			Reload:               reload,
			Target:               target,
			WatchIgnore:          watchIgnore,
			SourceControlRootURL: sourceControlRootURL,
			DemoURLPrefix:        demoURLPrefix,
			URLRewrites:          urlRewrites,
			ConfigFile:           ctx.ConfigFile(),
			ImportMap: types.ImportMapConfig{
				Generate:     importMapGenerate,
				OverrideFile: importMapOverrideFile,
				Override:     importMapOverride,
			},
			Demos: serve.DemosConfig{
				Rendering: demoRendering,
			},
			Transforms: serve.TransformConfig{
				TypeScript: serve.TypeScriptConfig{
					Enabled: tsEnabled,
					Target:  target,
				},
				CSS: serve.CSSConfig{
					Enabled: cssEnabled,
					Include: cssInclude,
					Exclude: cssExclude,
				},
			},
		}

		if buildMode, _ := cmd.Flags().GetBool("build"); buildMode {
			return runBuild(config, ctx.Root(), cmd)
		}

		if serveTUILogger != nil {
			defer func() {
				logging.SetServeSink(nil)
				logging.SetMode(logging.ModeCLI)
			}()
			return runInteractive(serveTUILogger, config, ctx.Root(), reload)
		}
		return runNonInteractive(config, ctx.Root(), reload)
	},
}

func runInteractive(tl *servetui.Logger, config serve.Config, root string, reload bool) error {

	config.Logger = tl

	var server *serve.Server
	var initErr error
	var mu sync.Mutex
	currentLogLevelIdx := 1
	for i, v := range verbosityOrder {
		if v == logging.CurrentVerbosity() {
			currentLogLevelIdx = i
			break
		}
	}

	callbacks := servetui.Callbacks{
		InitServer: func() tea.Msg {
			var err error
			s, err := initServer(config, tl, root, reload)
			if err != nil {
				mu.Lock()
				initErr = err
				mu.Unlock()
				return servetui.ServerInitErrorMsg{Err: err}
			}
			mu.Lock()
			server = s
			mu.Unlock()
			port := s.Port()
			tl.SetStatus(formatStatusLine(port, reload, logging.CurrentVerbosity().String()))
			return servetui.ServerReadyMsg{Port: port, Reload: reload, WatchDone: s.Done()}
		},
		RebuildManifest: func() (int, error) {
			mu.Lock()
			s := server
			mu.Unlock()
			if s == nil {
				return 0, fmt.Errorf("server not initialized")
			}
			tl.Info("Force rebuilding manifest...")
			size, err := s.RegenerateManifest()
			if err != nil {
				tl.Warning("Failed to rebuild manifest: %v", err)
			} else {
				tl.Info("Manifest rebuilt successfully")
			}
			return size, err
		},
		OpenBrowser: func() error {
			mu.Lock()
			s := server
			mu.Unlock()
			if s == nil {
				return fmt.Errorf("server not initialized")
			}
			port := s.Port()
			url := fmt.Sprintf("http://localhost:%d", port)
			tl.Info("Opening %s in browser...", url)
			return openBrowser(url)
		},
		CycleVerbosity: func() string {
			mu.Lock()
			currentLogLevelIdx = (currentLogLevelIdx + 1) % len(verbosityOrder)
			v := verbosityOrder[currentLogLevelIdx]
			s := server
			mu.Unlock()
			logging.SetVerbosity(v)
			tl.Info("Log level: %s", v)
			port := 0
			if s != nil {
				port = s.Port()
			}
			tl.SetStatus(formatStatusLine(port, reload, v.String()))
			return v.String()
		},
		Shutdown: func() {
			tl.Info("Shutting down server...")
		},
	}

	model := servetui.NewModel(callbacks, nil)
	p := tea.NewProgram(&model)
	model.SetPending(tl.SetProgram(p))

	if _, err := p.Run(); err != nil {
		if !errors.Is(err, tea.ErrInterrupted) {
			return err
		}
	}

	mu.Lock()
	s := server
	ie := initErr
	mu.Unlock()

	if s != nil {
		if err := s.Close(); err != nil {
			logging.Warning("Server close: %v", err)
		}
	}

	return ie
}

func initServer(config serve.Config, log logger.Logger, root string, reload bool) (ret *serve.Server, retErr error) {
	server, err := serve.NewServerWithConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	defer func() {
		if retErr != nil {
			_ = server.Close()
		}
	}()

	server.SetLogger(log)

	if err := server.SetWatchDir(root); err != nil {
		return nil, fmt.Errorf("failed to set watch directory: %w", err)
	}

	log.Info("Initializing server...")
	if err := server.InitializeWorkspaceMode(); err != nil {
		return nil, fmt.Errorf("failed to initialize workspace mode: %w", err)
	}

	if !server.IsWorkspace() {
		size, err := server.TryLoadExistingManifest()
		if err != nil {
			log.Warning("Could not load cached manifest: %v", err)
			log.Info("Generating initial manifest...")
			if _, err = server.RegenerateManifest(); err != nil {
				log.Error("Failed to generate initial manifest: %v", err)
				log.Info("Server will continue, but manifest may be unavailable")
			} else {
				log.Success("Initial manifest generated")
			}
		} else if size > 0 {
			log.Success("Loaded cached manifest from disk (%d bytes)", size)
			go func() {
				select {
				case <-time.After(2 * time.Second):
				case <-server.Done():
					return
				}
				select {
				case <-server.Done():
					return
				default:
				}
				log.Info("Regenerating manifest in background...")
				newSize, err := server.RegenerateManifest()
				if err != nil {
					log.Warning("Background manifest regeneration failed: %v", err)
				} else {
					log.Info("Background manifest regeneration complete (%d bytes)", newSize)
				}
			}()
		} else {
			log.Info("Generating initial manifest...")
			if _, err = server.RegenerateManifest(); err != nil {
				log.Error("Failed to generate initial manifest: %v", err)
				log.Info("Server will continue, but manifest may be unavailable")
			} else {
				log.Success("Initial manifest generated")
			}
		}
	} else {
		log.Success("Workspace mode initialized")
	}

	if err := server.Start(); err != nil {
		return nil, fmt.Errorf("failed to start server: %w", err)
	}

	reloadStatus := ""
	if reload {
		reloadStatus = " (live reload enabled)"
	}
	log.Success("Server started on http://localhost:%d%s", server.Port(), reloadStatus)

	return server, nil
}

func runBuild(config serve.Config, root string, cmd *cobra.Command) error {
	log := logger.NewDefaultLogger()
	config.Logger = log

	server, err := serve.NewServerWithConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	server.SetLogger(log)
	if err := server.SetWatchDir(root); err != nil {
		return fmt.Errorf("failed to set watch directory: %w", err)
	}
	log.Info("Initializing server...")
	if err := server.InitializeWorkspaceMode(); err != nil {
		return fmt.Errorf("failed to initialize workspace mode: %w", err)
	}
	if !server.IsWorkspace() {
		if _, err := server.TryLoadExistingManifest(); err != nil {
			if _, err = server.RegenerateManifest(); err != nil {
				return fmt.Errorf("failed to generate manifest: %w", err)
			}
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}
	outputDir, _ := cmd.Flags().GetString("output")
	basePath, _ := cmd.Flags().GetString("base-path")
	importMode, _ := cmd.Flags().GetString("import")
	return server.Build(serve.BuildConfig{
		OutputDir:  outputDir,
		BasePath:   basePath,
		ImportMode: importMode,
		WorkDir:    cwd,
	})
}

func runNonInteractive(config serve.Config, root string, reload bool) error {
	log := logger.NewDefaultLogger()
	config.Logger = log

	server, err := initServer(config, log, root, reload)
	if err != nil {
		return err
	}
	defer func() {
		if err := server.Close(); err != nil {
			logging.Warning("Server close: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigChan:
	case <-server.Done():
	}
	return nil
}

// openBrowser opens the given URL in the default browser
func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Start()
}

var verbosityOrder = []logging.Verbosity{
	logging.VerbosityQuiet,
	logging.VerbosityNormal,
	logging.VerbosityVerbose,
	logging.VerbosityDebug,
	logging.VerbosityTrace,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("port", 8000, "Port to serve on")
	serveCmd.Flags().Bool("no-reload", false, "Disable live reload")
	serveCmd.Flags().Bool("no-import-map-generate", false, "Disable automatic import map generation")
	serveCmd.Flags().String("import-map-override-file", "", "Path to JSON file with custom import map entries")
	serveCmd.Flags().String("target", "", "TypeScript/JavaScript transform target (es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, esnext)")
	serveCmd.Flags().String("rendering", "", "Demo rendering mode: light (full UI), shadow (Shadow DOM), or chromeless (minimal, no UI)")
	serveCmd.Flags().StringSlice("watch-ignore", nil, "Glob patterns to ignore in file watcher (comma-separated, e.g., '_site/**,dist/**')")
	serveCmd.Flags().StringSlice("css-transform", nil, "Glob patterns for CSS files to transform to JavaScript modules (e.g., 'src/**/*.css,elements/**/*.css')")
	serveCmd.Flags().StringSlice("css-transform-exclude", nil, "Glob patterns for CSS files to exclude from transformation (e.g., 'demo/**/*.css')")
	serveCmd.Flags().Bool("build", false, "Build a static site instead of starting a dev server")
	serveCmd.Flags().StringP("output", "o", "dist", "Output directory for static build")
	serveCmd.Flags().String("base-path", "", "URL base path for static build deployment (e.g., /docs/components/)")
	serveCmd.Flags().String("import", "vendor", "Dependency resolution for static builds: vendor, esm, jspm, unpkg")

	if err := viper.BindPFlag("serve.port", serveCmd.Flags().Lookup("port")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.port: %v", err))
	}
	if err := viper.BindPFlag("serve.no-reload", serveCmd.Flags().Lookup("no-reload")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.no-reload: %v", err))
	}
	// Bind import map flags (note: --no-import-map-generate is handled specially in RunE to invert to serve.importMap.generate)
	if err := viper.BindPFlag("serve.importMap.overrideFile", serveCmd.Flags().Lookup("import-map-override-file")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.importMap.overrideFile: %v", err))
	}
	if err := viper.BindPFlag("serve.target", serveCmd.Flags().Lookup("target")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.target: %v", err))
	}
	if err := viper.BindPFlag("serve.demos.rendering", serveCmd.Flags().Lookup("rendering")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.demos.rendering: %v", err))
	}
	if err := viper.BindPFlag("serve.watchIgnore", serveCmd.Flags().Lookup("watch-ignore")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.watchIgnore: %v", err))
	}
	if err := viper.BindPFlag("serve.transforms.css.include", serveCmd.Flags().Lookup("css-transform")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.transforms.css.include: %v", err))
	}
	if err := viper.BindPFlag("serve.transforms.css.exclude", serveCmd.Flags().Lookup("css-transform-exclude")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.transforms.css.exclude: %v", err))
	}
}
