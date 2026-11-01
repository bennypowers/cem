/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

SPDX-License-Identifier: GPL-3.0-or-later

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
package mcp

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/platform"
	lspTypes "bennypowers.dev/cem/lsp/types"
	MCPTypes "bennypowers.dev/cem/mcp/types"
	"bennypowers.dev/cem/types"
)

// Selection filters a set of items by exact name.
//
// An empty Include list selects every item. Exclude is applied after Include
// and wins on overlap. Matching is exact and case-sensitive.
type Selection struct {
	Include []string
	Exclude []string
}

// RegisterOption configures a call to Register.
type RegisterOption func(*registerConfig)

// registerConfig accumulates the options passed to Register.
type registerConfig struct {
	tools     Selection
	resources Selection
	templates Selection
	conflict  ConflictPolicy
}

// WithTools limits the tools registered by Register to the given tool names.
func WithTools(sel Selection) RegisterOption {
	return func(c *registerConfig) {
		c.tools.Include = append(c.tools.Include, sel.Include...)
		c.tools.Exclude = append(c.tools.Exclude, sel.Exclude...)
	}
}

// WithResources limits the static resources registered by Register to the
// given resource URIs.
func WithResources(sel Selection) RegisterOption {
	return func(c *registerConfig) {
		c.resources.Include = append(c.resources.Include, sel.Include...)
		c.resources.Exclude = append(c.resources.Exclude, sel.Exclude...)
	}
}

// WithTemplates limits the resource templates registered by Register to the
// given URI template strings.
func WithTemplates(sel Selection) RegisterOption {
	return func(c *registerConfig) {
		c.templates.Include = append(c.templates.Include, sel.Include...)
		c.templates.Exclude = append(c.templates.Exclude, sel.Exclude...)
	}
}

// ConflictPolicy controls what Register does when a selected tool name or
// resource URI is already registered on the target server.
type ConflictPolicy int

const (
	// ConflictReplace overwrites the existing entry. This is the default and
	// matches the behavior before registration options existed.
	ConflictReplace ConflictPolicy = iota
	// ConflictError fails Register before anything is registered.
	ConflictError
	// ConflictSkip keeps the existing entry and skips the conflicting one.
	ConflictSkip
)

// WithConflictPolicy sets the policy for entries that already exist on the
// target server. The last WithConflictPolicy wins when repeated.
func WithConflictPolicy(p ConflictPolicy) RegisterOption {
	return func(c *registerConfig) {
		c.conflict = p
	}
}

// applyRegisterOptions applies opts and validates the resulting configuration.
func applyRegisterOptions(opts []RegisterOption) (registerConfig, error) {
	cfg := registerConfig{conflict: ConflictReplace}
	for _, opt := range opts {
		if opt == nil {
			return registerConfig{}, fmt.Errorf("register option must not be nil")
		}
		opt(&cfg)
	}

	switch cfg.conflict {
	case ConflictReplace, ConflictError, ConflictSkip:
	default:
		return registerConfig{}, fmt.Errorf("unknown conflict policy %d", cfg.conflict)
	}

	return cfg, nil
}

// Registration is the handle returned by Register.
//
// Reload refreshes the manifests and configuration used by the handlers that
// Register installed. It does not re-register tools or resources, so names,
// URIs, and schemas stay stable for the lifetime of the server.
type Registration struct {
	workspace types.WorkspaceContext
	config    ServerConfig
	registry  *reloadableContext
	reloadMu  sync.Mutex
}

// Reload re-reads the workspace manifests, configuration, and additional
// packages, then atomically swaps the state used by the registered handlers.
//
// Reload is explicit and synchronous, and safe to call while the server is
// serving. On failure the previous state stays active and the error is
// returned. Reload does not start file watching; call it when the host knows
// the workspace changed.
//
// Calling Register again on the same server is a replacement operation, not a
// reload: it re-registers handlers and applies the conflict policy again.
func (r *Registration) Reload(ctx context.Context) error {
	if r == nil || r.registry == nil || r.workspace == nil {
		return fmt.Errorf("registration did not own a reloadable registry")
	}
	if ctx == nil {
		return fmt.Errorf("reload context must not be nil")
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	r.reloadMu.Lock()
	defer r.reloadMu.Unlock()

	if err := ctx.Err(); err != nil {
		return err
	}

	// Workspace implementations cache parsed config and manifests. Ask them to
	// drop those caches before building the replacement registry so the reload
	// observes changes on disk.
	if invalidator, ok := r.workspace.(interface{ InvalidateConfig() }); ok {
		invalidator.InvalidateConfig()
	}
	if invalidator, ok := r.workspace.(interface{ InvalidateManifest() }); ok {
		invalidator.InvalidateManifest()
	}

	// The LSP registry deliberately falls back to in-memory generation when a
	// workspace manifest cannot be loaded, which would turn a malformed
	// manifest into a silently empty reload. Fail first so the previous state
	// stays active.
	if err := validateReloadManifest(r.workspace); err != nil {
		return fmt.Errorf("failed to reload manifests: %w", err)
	}

	next, err := newMCPContext(r.workspace, r.config.FileSystem, false)
	if err != nil {
		return fmt.Errorf("failed to create registry: %w", err)
	}
	if err := loadRegistryData(next, r.config.AdditionalPackages, "reload"); err != nil {
		return err
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	r.registry.store(next)
	return nil
}

// validateReloadManifest parses the workspace manifest, when one exists, so a
// malformed manifest fails the reload before its replacement state is built.
func validateReloadManifest(workspace types.WorkspaceContext) error {
	path := workspace.CustomElementsManifestPath()
	if path == "" {
		return nil
	}
	if _, err := workspace.Stat(path); err != nil {
		// No manifest at the configured path: the registry may still build one
		// from source, which is not a reload failure.
		return nil
	}
	if _, err := workspace.Manifest(); err != nil {
		return err
	}
	return nil
}

// reloadableContext gives handlers a stable view of the current MCPContext.
// Reload builds a complete replacement context and swaps it in atomically, so
// a handler either observes the old state or the new one, never a partial
// update. The old context stays valid for requests already in flight.
type reloadableContext struct {
	current atomic.Pointer[MCPContext]
}

func newReloadableContext(initial *MCPContext) *reloadableContext {
	r := &reloadableContext{}
	r.store(initial)
	return r
}

func (r *reloadableContext) load() *MCPContext {
	return r.current.Load()
}

func (r *reloadableContext) store(next *MCPContext) {
	r.current.Store(next)
}

func (r *reloadableContext) ElementInfo(tagName string) (MCPTypes.ElementInfo, error) {
	return r.load().GetElementInfo(tagName)
}

func (r *reloadableContext) AllElements() map[string]MCPTypes.ElementInfo {
	return r.load().GetAllElements()
}

func (r *reloadableContext) LoadManifests() error {
	return r.load().LoadManifests()
}

func (r *reloadableContext) GetManifestSchemaVersions() []string {
	return r.load().GetManifestSchemaVersions()
}

func (r *reloadableContext) DocumentManager() lspTypes.DocumentManager {
	return r.load().DocumentManager()
}

func (r *reloadableContext) CommonPrefixes() []string {
	return r.load().CommonPrefixes()
}

func (r *reloadableContext) AllCSSProperties() []string {
	return r.load().AllCSSProperties()
}

func (r *reloadableContext) Config() (*IC.CemConfig, error) {
	return r.load().Config()
}

func (r *reloadableContext) ConfigFile() string {
	return r.load().ConfigFile()
}

func (r *reloadableContext) ConfigSchemaJSON() []byte {
	return r.load().ConfigSchemaJSON()
}

func (r *reloadableContext) Root() string {
	return r.load().Root()
}

func (r *reloadableContext) FileSystem() platform.FileSystem {
	return r.load().FileSystem()
}
