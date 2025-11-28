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
	"os/signal"
	"syscall"
	"time"

	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/serve"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware/transform"
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

		// Create server config
		config := serve.Config{
			Port:        port,
			Reload:      reload,
			Target:      target,
			WatchIgnore: watchIgnore,
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
				log.Error("Failed to close server: %v", err)
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
				go func(log logger.Logger) {
					// Wait a moment for server to fully start and become idle
					time.Sleep(2 * time.Second)
					log.Info("Regenerating manifest in background...")
					newSize, err := server.RegenerateManifest()
					if err != nil {
						log.Warning("Background manifest regeneration failed: %v", err)
					} else {
						log.Info("Background manifest regeneration complete (%d bytes)", newSize)
					}
				}(log)
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
		statusMsg := fmt.Sprintf("Running on %s%s Live reload: %s %s Press %s to stop",
			pterm.FgCyan.Sprintf("http://localhost:%d", port),
			pterm.FgGray.Sprint(" |"),
			reloadColor,
			pterm.FgGray.Sprint("|"),
			pterm.FgYellow.Sprint("Ctrl+C"),
		)
		if l, ok := log.(interface{ SetStatus(string) }); ok {
			l.SetStatus(statusMsg)
		}

		// Wait for interrupt signal
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		if l, ok := log.(interface{ SetStatus(string) }); ok {
			l.SetStatus("Shutting down...")
		}
		log.Info("Shutting down server...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("port", 8000, "Port to serve on")
	serveCmd.Flags().Bool("no-reload", false, "Disable live reload")
	serveCmd.Flags().String("target", "es2022", "TypeScript/JavaScript transform target (es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, esnext)")
	serveCmd.Flags().StringSlice("watch-ignore", nil, "Glob patterns to ignore in file watcher (comma-separated, e.g., '_site/**,dist/**')")
	serveCmd.Flags().StringSlice("css-transform", nil, "Glob patterns for CSS files to transform to JavaScript modules (e.g., 'src/**/*.css,elements/**/*.css')")
	serveCmd.Flags().StringSlice("css-transform-exclude", nil, "Glob patterns for CSS files to exclude from transformation (e.g., 'demo/**/*.css')")

	if err := viper.BindPFlag("serve.port", serveCmd.Flags().Lookup("port")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.port: %v", err))
	}
	if err := viper.BindPFlag("serve.no-reload", serveCmd.Flags().Lookup("no-reload")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.no-reload: %v", err))
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
