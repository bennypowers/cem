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
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/serve"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware/transform"
	"bennypowers.dev/cem/serve/middleware/types"
	W "bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a development server with live reload",
	Long: `Start a development server for custom element demos with:
- Live reload on file changes
- Automatic manifest regeneration
- WebSocket-based browser refresh
- Static file serving with CORS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get workspace context
		ctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}

		port := viper.GetInt("serve.port")
		reload := !viper.GetBool("serve.no-reload")
		verbose := viper.GetBool("verbose")
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
		noImportMapGenerate, _ := cmd.Flags().GetBool("no-import-map-generate")
		importMapGenerate := !noImportMapGenerate
		// Config file can also set this
		if viper.IsSet("serve.importMap.generate") {
			importMapGenerate = viper.GetBool("serve.importMap.generate")
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
					if scopeMap, ok := scopeVal.(map[string]interface{}); ok {
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

		// Validate and parse target (default to ES2022)
		var target transform.Target
		if targetStr != "" {
			if !transform.IsValidTarget(targetStr) {
				return fmt.Errorf("invalid target '%s': must be one of es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, or esnext", targetStr)
			}
			target = transform.Target(targetStr)
		} else {
			target = transform.ES2022
		}

		// Get source control root URL from config
		sourceControlRootURL := viper.GetString("sourceControlRootUrl")

		// Create server config
		config := serve.Config{
			Port:                 port,
			Reload:               reload,
			Target:               target,
			WatchIgnore:          watchIgnore,
			SourceControlRootURL: sourceControlRootURL,
			ImportMap: types.ImportMapConfig{
				Generate:     importMapGenerate,
				OverrideFile: importMapOverrideFile,
				Override:     importMapOverride,
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

		// Create pterm logger
		log := logger.NewPtermLogger(verbose)
		defer func() {
			if l, ok := log.(interface{ Stop() }); ok {
				l.Stop()
			}
		}()

		// Enable quiet mode on internal logger to suppress generation progress messages
		// These would otherwise appear below the status line
		logging.GetLogger().SetQuietEnabled(true)

		// Create server
		server, err := serve.NewServerWithConfig(config)
		if err != nil {
			return fmt.Errorf("failed to create server: %w", err)
		}
		defer func() {
			if err := server.Close(); err != nil {
				// Shutdown timeouts are expected with active browser sessions
				log.Warning("Server close: %v", err)
			}
		}()

		// Set pterm logger
		server.SetLogger(log)

		// Set watch directory to project root
		err = server.SetWatchDir(ctx.Root())
		if err != nil {
			return fmt.Errorf("failed to set watch directory: %w", err)
		}

		// Try workspace mode initialization first
		pterm.Info.Println("Initializing server...")
		err = server.InitializeWorkspaceMode()
		if err != nil {
			return fmt.Errorf("failed to initialize workspace mode: %w", err)
		}

		// If not workspace mode, try to load existing manifest or generate new one
		if !server.IsWorkspace() {
			// Try to load existing manifest from disk first for faster startup
			size, err := server.TryLoadExistingManifest()
			if err != nil {
				// Failed to load, generate fresh manifest
				pterm.Warning.Printf("Could not load cached manifest: %v\n", err)
				pterm.Info.Println("Generating initial manifest...")
				_, err = server.RegenerateManifest()
				if err != nil {
					pterm.Error.Printf("Failed to generate initial manifest: %v\n", err)
					pterm.Info.Println("Server will continue, but manifest may be unavailable")
				} else {
					pterm.Success.Println("Initial manifest generated")
				}
			} else if size > 0 {
				// Successfully loaded existing manifest
				pterm.Success.Printf("Loaded cached manifest from disk (%d bytes)\n", size)

				// Schedule background regeneration to ensure it's up-to-date
				go func(log logger.Logger, shutdownCh <-chan struct{}) {
					// Wait a moment for server to fully start and become idle, or cancel if shutting down
					select {
					case <-time.After(2 * time.Second):
						// Sleep completed, proceed with regeneration
					case <-shutdownCh:
						// Server shut down during sleep, cancel work
						return
					}

					log.Info("Regenerating manifest in background...")
					newSize, err := server.RegenerateManifest()
					if err != nil {
						log.Warning("Background manifest regeneration failed: %v", err)
					} else {
						log.Info("Background manifest regeneration complete (%d bytes)", newSize)
					}
				}(log, server.Done())
			} else {
				// No existing manifest found, generate fresh one
				pterm.Info.Println("Generating initial manifest...")
				_, err = server.RegenerateManifest()
				if err != nil {
					pterm.Error.Printf("Failed to generate initial manifest: %v\n", err)
					pterm.Info.Println("Server will continue, but manifest may be unavailable")
				} else {
					pterm.Success.Println("Initial manifest generated")
				}
			}
		} else {
			pterm.Success.Println("Workspace mode initialized")
		}

		// Start live rendering area AFTER initial setup
		if l, ok := log.(interface{ Start() }); ok {
			l.Start()
		}

		// Start server
		err = server.Start()
		if err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}

		// Single startup message
		reloadStatus := ""
		if reload {
			reloadStatus = " (live reload enabled)"
		}
		log.Info("Server started on http://localhost:%d%s", port, reloadStatus)

		// Update status with running info (with colors)
		reloadColor := pterm.FgRed.Sprint("false")
		if reload {
			reloadColor = pterm.FgGreen.Sprint("true")
		}
		statusMsg := fmt.Sprintf("Running on %s%s Live reload: %s %s Press %s for help, %s to quit",
			pterm.FgCyan.Sprintf("http://localhost:%d", port),
			pterm.FgGray.Sprint(" |"),
			reloadColor,
			pterm.FgGray.Sprint("|"),
			pterm.FgYellow.Sprint("h"),
			pterm.FgYellow.Sprint("q"),
		)
		if l, ok := log.(interface{ SetStatus(string) }); ok {
			l.SetStatus(statusMsg)
		}

		// Start keyboard input handler after a brief delay
		// This allows pterm's live area to stabilize before we enable raw mode
		quitChan := make(chan struct{})
		go func() {
			time.Sleep(100 * time.Millisecond)
			handleKeyboardInput(server, log, port, quitChan)
		}()

		// Wait for quit signal (keyboard or interrupt)
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		select {
		case <-quitChan:
			// User pressed 'q' or Ctrl+C in keyboard handler
		case <-sigChan:
			// Received system interrupt signal
		}

		if l, ok := log.(interface{ SetStatus(string) }); ok {
			l.SetStatus("Shutting down...")
		}
		log.Info("Shutting down server...")
		return nil
	},
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

// showHelp displays the keyboard shortcuts help menu
func showHelp(log logger.Logger) {
	log.Info(`Keyboard Shortcuts
	m - Force rebuild manifest
	v - Cycle log levels (normal/verbose/debug/quiet)
	o - Open in browser
	c - Clear console
	h - Show this help
	q - Quit server
	Ctrl+C - Also quits server                      `)
}

// logLevel tracks the current log verbosity level
type logLevel int

const (
	logLevelNormal logLevel = iota
	logLevelVerbose
	logLevelDebug
	logLevelQuiet
)

func (l logLevel) String() string {
	switch l {
	case logLevelNormal:
		return "normal"
	case logLevelVerbose:
		return "verbose"
	case logLevelDebug:
		return "debug"
	case logLevelQuiet:
		return "quiet"
	default:
		return "unknown"
	}
}

// handleKeyboardInput reads keyboard input and handles commands using atomicgo/keyboard
func handleKeyboardInput(server *serve.Server, log logger.Logger, port int, quitChan chan struct{}) {
	currentLogLevel := logLevelNormal

	// Handle all keyboard input
	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		// Handle Ctrl+C
		if key.Code == keys.CtrlC {
			close(quitChan)
			return true, nil // Stop listening
		}

		// Only handle rune key presses (regular characters)
		if key.Code != keys.RuneKey || len(key.Runes) == 0 {
			return false, nil
		}

		switch key.Runes[0] {
		case 'q', 'Q':
			log.Info("Quitting...")
			close(quitChan)
			return true, nil // Stop listening

		case 'm', 'M':
			log.Info("Force rebuilding manifest...")
			if _, err := server.RegenerateManifest(); err != nil {
				log.Warning("Failed to rebuild manifest: %v", err)
			} else {
				log.Info("Manifest rebuilt successfully")
			}

		case 'v', 'V':
			// Cycle through log levels
			currentLogLevel = (currentLogLevel + 1) % 4
			internalLogger := logging.GetLogger()

			switch currentLogLevel {
			case logLevelNormal:
				internalLogger.SetQuietEnabled(true) // Keep generation logs quiet
				internalLogger.SetDebugEnabled(false)
				if setter, ok := log.(interface{ SetVerbose(bool) }); ok {
					setter.SetVerbose(false)
				}
			case logLevelVerbose:
				internalLogger.SetQuietEnabled(false)
				internalLogger.SetDebugEnabled(false)
				if setter, ok := log.(interface{ SetVerbose(bool) }); ok {
					setter.SetVerbose(true)
				}
			case logLevelDebug:
				internalLogger.SetQuietEnabled(false)
				internalLogger.SetDebugEnabled(true)
				if setter, ok := log.(interface{ SetVerbose(bool) }); ok {
					setter.SetVerbose(true)
				}
			case logLevelQuiet:
				internalLogger.SetQuietEnabled(true)
				internalLogger.SetDebugEnabled(false)
				if setter, ok := log.(interface{ SetVerbose(bool) }); ok {
					setter.SetVerbose(false)
				}
			}
			log.Info("Log level: %s", currentLogLevel.String())

		case 'o', 'O':
			url := fmt.Sprintf("http://localhost:%d", port)
			log.Info("Opening %s in browser...", url)
			if err := openBrowser(url); err != nil {
				log.Warning("Failed to open browser: %v", err)
			}

		case 'c', 'C':
			// Clear console by clearing the log buffer
			if clearer, ok := log.(interface{ Clear() }); ok {
				clearer.Clear()
				log.Info("Console cleared")
			} else {
				log.Warning("Clear not supported by current logger")
			}

		case 'h', 'H', '?':
			showHelp(log)
		}

		return false, nil
	})

	if err != nil {
		log.Warning("Keyboard input disabled: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("port", 8000, "Port to serve on")
	serveCmd.Flags().Bool("no-reload", false, "Disable live reload")
	serveCmd.Flags().Bool("no-import-map-generate", false, "Disable automatic import map generation")
	serveCmd.Flags().String("import-map-override-file", "", "Path to JSON file with custom import map entries")
	serveCmd.Flags().String("target", "", "TypeScript/JavaScript transform target (es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, esnext)")
	serveCmd.Flags().StringSlice("watch-ignore", nil, "Glob patterns to ignore in file watcher (comma-separated, e.g., '_site/**,dist/**')")
	serveCmd.Flags().StringSlice("css-transform", nil, "Glob patterns for CSS files to transform to JavaScript modules (e.g., 'src/**/*.css,elements/**/*.css')")
	serveCmd.Flags().StringSlice("css-transform-exclude", nil, "Glob patterns for CSS files to exclude from transformation (e.g., 'demo/**/*.css')")

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
