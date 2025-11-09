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

	"bennypowers.dev/cem/serve"
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

		// Validate and parse target (default to ES2022)
		var target serve.Target
		if targetStr != "" {
			if !serve.IsValidTarget(targetStr) {
				return fmt.Errorf("invalid target '%s': must be one of es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, or esnext", targetStr)
			}
			target = serve.Target(targetStr)
		} else {
			target = serve.ES2022
		}

		// Create server config
		config := serve.Config{
			Port:   port,
			Reload: reload,
			Target: target,
		}

		// Create pterm logger
		logger := serve.NewPtermLogger(verbose)
		defer func() {
			if l, ok := logger.(interface{ Stop() }); ok {
				l.Stop()
			}
		}()

		// Create server
		server, err := serve.NewServerWithConfig(config)
		if err != nil {
			return fmt.Errorf("failed to create server: %w", err)
		}
		defer func() {
			if err := server.Close(); err != nil {
				logger.Error("Failed to close server: %v", err)
			}
		}()

		// Set pterm logger
		server.SetLogger(logger)

		// Set watch directory to project root
		err = server.SetWatchDir(ctx.Root())
		if err != nil {
			return fmt.Errorf("failed to set watch directory: %w", err)
		}

		// Generate initial manifest (before starting live area)
		pterm.Info.Println("Generating initial manifest...")
		err = server.RegenerateManifest()
		if err != nil {
			pterm.Error.Printf("Failed to generate initial manifest: %v\n", err)
			pterm.Info.Println("Server will continue, but manifest may be unavailable")
		} else {
			pterm.Success.Println("Initial manifest generated")
		}

		// Start live rendering area AFTER initial setup
		if l, ok := logger.(interface{ Start() }); ok {
			l.Start()
		}

		// Start server
		err = server.Start()
		if err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}

		logger.Info("Server started on http://localhost:%d", port)
		if reload {
			logger.Info("Live reload enabled - watching for file changes")
		}

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
		if l, ok := logger.(interface{ SetStatus(string) }); ok {
			l.SetStatus(statusMsg)
		}

		// Wait for interrupt signal
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		if l, ok := logger.(interface{ SetStatus(string) }); ok {
			l.SetStatus("Shutting down...")
		}
		logger.Info("Shutting down server...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("port", 8000, "Port to serve on")
	serveCmd.Flags().Bool("no-reload", false, "Disable live reload")
	serveCmd.Flags().String("target", "es2022", "TypeScript/JavaScript transform target (es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, esnext)")

	if err := viper.BindPFlag("serve.port", serveCmd.Flags().Lookup("port")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.port: %v", err))
	}
	if err := viper.BindPFlag("serve.no-reload", serveCmd.Flags().Lookup("no-reload")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.no-reload: %v", err))
	}
	if err := viper.BindPFlag("serve.target", serveCmd.Flags().Lookup("target")); err != nil {
		panic(fmt.Sprintf("failed to bind flag serve.target: %v", err))
	}
}
