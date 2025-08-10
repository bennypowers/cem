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
package lsp

import (
	"context"
	"fmt"
	"sync"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/lsp/helpers"
	W "bennypowers.dev/cem/workspace"
)

// InProcessGenerateWatcher manages in-process generate watching for the LSP server
// This replaces the subprocess-based approach to prevent hanging processes
type InProcessGenerateWatcher struct {
	watchSession *G.WatchSession
	ctx          context.Context
	cancel       context.CancelFunc
	mu           sync.RWMutex
	workspace    W.WorkspaceContext
	globs        []string
}

// NewInProcessGenerateWatcher creates a new in-process generate watcher
func NewInProcessGenerateWatcher(workspace W.WorkspaceContext, globs []string) (*InProcessGenerateWatcher, error) {
	// Create cancellable context for clean shutdown
	ctx, cancel := context.WithCancel(context.Background())

	helpers.SafeDebugLog("Creating InProcessGenerateWatcher for workspace: %s, globs: %v", workspace.Root(), globs)

	// Create the generate watch session
	watchSession, err := G.NewWatchSession(workspace, globs)
	if err != nil {
		cancel()
		helpers.SafeDebugLog("Failed to create watch session: %v", err)
		return nil, fmt.Errorf("failed to create watch session: %w", err)
	}

	helpers.SafeDebugLog("Successfully created watch session")

	return &InProcessGenerateWatcher{
		watchSession: watchSession,
		ctx:          ctx,
		cancel:       cancel,
		workspace:    workspace,
		globs:        globs,
	}, nil
}

// Start begins the file watching process in a goroutine
// This will watch for source file changes and regenerate the manifest automatically
func (w *InProcessGenerateWatcher) Start() error {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.watchSession == nil {
		return fmt.Errorf("watcher not initialized")
	}

	// Start watching in a goroutine to avoid blocking the LSP server
	go func() {
		defer func() {
			// Clean up watch session when done
			if w.watchSession != nil {
				w.watchSession.Close()
			}
		}()

		helpers.SafeDebugLog("Starting in-process generate watcher for workspace: %s", w.workspace.Root())

		// Run the watch session - this will block until cancelled or error
		if err := w.watchSession.RunWatch(); err != nil {
			// Check if this was a cancellation (normal shutdown)
			select {
			case <-w.ctx.Done():
				helpers.SafeDebugLog("Generate watcher stopped due to context cancellation")
			default:
				helpers.SafeDebugLog("Generate watcher error: %v", err)
			}
		} else {
			helpers.SafeDebugLog("Generate watcher completed successfully")
		}
	}()

	return nil
}

// Stop cleanly shuts down the file watcher
func (w *InProcessGenerateWatcher) Stop() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	helpers.SafeDebugLog("Stopping in-process generate watcher")

	// Cancel the context to signal shutdown
	if w.cancel != nil {
		w.cancel()
	}

	// Close the watch session
	if w.watchSession != nil {
		w.watchSession.Close()
		w.watchSession = nil
	}

	return nil
}

// IsRunning returns true if the watcher is currently active
func (w *InProcessGenerateWatcher) IsRunning() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	// Check if context is still active and session exists
	return w.ctx.Err() == nil && w.watchSession != nil
}
