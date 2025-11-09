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

		// Create server config
		config := serve.Config{
			Port:   port,
			Reload: reload,
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
		defer server.Close()

		// Set pterm logger
		server.SetLogger(logger)

		// Set watch directory to project root
		err = server.SetWatchDir(ctx.Root())
		if err != nil {
			return fmt.Errorf("failed to set watch directory: %w", err)
		}

		// Update status
		if l, ok := logger.(interface{ SetStatus(string) }); ok {
			l.SetStatus("Generating initial manifest...")
		}

		// Generate initial manifest
		logger.Info("Generating initial manifest...")
		err = server.RegenerateManifest()
		if err != nil {
			logger.Error("Failed to generate initial manifest: %v", err)
			logger.Info("Server will continue, but manifest may be unavailable")
		} else {
			logger.Info("Initial manifest generated")
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

		// Update status with running info
		statusMsg := fmt.Sprintf("Running on http://localhost:%d | Live reload: %v | Press Ctrl+C to stop", port, reload)
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

	serveCmd.Flags().Int("port", 8080, "Port to serve on")
	serveCmd.Flags().Bool("no-reload", false, "Disable live reload")

	viper.BindPFlag("serve.port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("serve.no-reload", serveCmd.Flags().Lookup("no-reload"))
}
