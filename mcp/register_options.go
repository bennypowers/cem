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
	"errors"
	"fmt"
	"io/fs"
	"strings"
	"sync"
	"sync/atomic"

	cemworkspace "bennypowers.dev/cem/internal/workspace"
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
	// ConflictError fails Register before anything is registered and returns a
	// *RegistrationConflictError describing the conflicting entries.
	ConflictError
	// ConflictSkip keeps the existing entry and skips the conflicting one.
	ConflictSkip
)

// RegistrationConflictError reports the entries that caused a Register call
// with ConflictError to fail. Entries are grouped by kind and sorted.
type RegistrationConflictError struct {
	Tools             []string
	Resources         []string
	ResourceTemplates []string
}

// Error implements error.
func (e *RegistrationConflictError) Error() string {
	parts := make([]string, 0, len(e.Tools)+len(e.Resources)+len(e.ResourceTemplates))
	for _, name := range e.Tools {
		parts = append(parts, fmt.Sprintf("tool %q", name))
	}
	for _, uri := range e.Resources {
		parts = append(parts, fmt.Sprintf("resource %q", uri))
	}
	for _, uri := range e.ResourceTemplates {
		parts = append(parts, fmt.Sprintf("resource template %q", uri))
	}
	return fmt.Sprintf("registration conflicts with existing server entries: %s", strings.Join(parts, ", "))
}

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

	if err := validateSelection("tool", cfg.tools); err != nil {
		return registerConfig{}, err
	}
	if err := validateSelection("resource", cfg.resources); err != nil {
		return registerConfig{}, err
	}
	if err := validateSelection("resource template", cfg.templates); err != nil {
		return registerConfig{}, err
	}

	return cfg, nil
}

// validateSelection rejects empty names before any registry work begins.
func validateSelection(kind string, selection Selection) error {
	for _, name := range selection.Include {
		if name == "" {
			return fmt.Errorf("%s include entry must not be empty", kind)
		}
	}
	for _, name := range selection.Exclude {
		if name == "" {
			return fmt.Errorf("%s exclude entry must not be empty", kind)
		}
	}
	return nil
}

// Registration is the handle returned by Register.
//
// Reload refreshes the workspace data used by the handlers that this Register
// call installed. It does not re-register tools or resources, so names, URIs,
// and schemas stay stable for the lifetime of the server.
type Registration struct {
	workspace types.WorkspaceContext
	config    ServerConfig
	registry  *reloadableContext
	reloadMu  sync.Mutex
}

// Reload re-reads the workspace manifests, configuration, and additional
// packages, then atomically swaps the state used by the handlers installed by
// this Register call.
//
// Reload is explicit and synchronous, and safe to call while the server is
// serving. On failure the previous state stays active, including the
// configuration snapshot captured with it, and the error is returned.
//
// Workspace data is refreshed when the workspace implements
// types.ManifestInvalidator or types.ConfigInvalidator; types.CacheInvalidator
// implements both. For other workspaces, Reload rebuilds the registry from
// whatever the workspace returns, which may be cached; such workspaces are
// responsible for observing their own changes. Reload does not re-run
// WorkspaceContext.Init, so config-file discovery and package.json changes are
// not observed. A workspace without a manifest fails Reload if the registry
// cannot generate one from source.
//
// Additional package failures are returned as errors by Reload, unlike
// Register, where they are logged and ignored. Individual package manifests
// discovered under workspace packages or node_modules, and workspace-package
// pattern expansion, remain best effort, matching the registry's normal
// loading behavior. Workspace manifest, configuration, explicit
// additional-package, and node_modules inspection failures are returned.
// Reload does not start file watching; call it when the host knows the
// workspace changed. The context bounds the decision to start and commit the
// reload, but the synchronous load itself is not interruptible. Remote
// additional-package resolution may block on network I/O, and Reload
// serializes concurrent calls while loading.
//
// Register is additive: calling it again adds or replaces entries on the
// server, and the earlier Registration keeps the handlers it installed bound
// to its own registry. Server.Run has a separate, lenient in-place watcher
// reload; keep cache invalidation changes in sync with it.
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

	current := r.registry.load()
	manifestInvalidator, canInvalidateManifest := r.workspace.(types.ManifestInvalidator)
	if canInvalidateManifest {
		manifestInvalidator.InvalidateManifest()
	}

	// The LSP registry deliberately falls back to in-memory generation when a
	// workspace manifest cannot be loaded, which would turn a malformed
	// manifest into a silently empty reload. Fail first so the previous state
	// stays active.
	if err := validateReloadManifest(r.workspace); err != nil {
		return fmt.Errorf("failed to reload manifests: %w", err)
	}

	// Invalidate configuration only after the manifest is known to parse. The
	// replacement context snapshots the new configuration before the swap, so
	// a later failure cannot change what the active generation serves either.
	configInvalidator, canInvalidateConfig := r.workspace.(types.ConfigInvalidator)
	if canInvalidateConfig {
		configInvalidator.InvalidateConfig()
	}

	// Reuse the document manager: MCP handlers do not use it, and allocating a
	// new tree-sitter-backed manager on every reload would leak resources.
	next, err := newMCPContextWithDocumentManager(r.workspace, r.config.FileSystem, false, current.DocumentManager())
	if err != nil {
		return fmt.Errorf("failed to create registry: %w", err)
	}
	if err := loadRegistryDataStrict(next, r.config.AdditionalPackages, "reload"); err != nil {
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
		if errors.Is(err, fs.ErrNotExist) {
			// The registry may still build a manifest from source.
			return nil
		}
		if errors.Is(err, errors.ErrUnsupported) {
			// The workspace cannot inspect paths, but it may still expose a
			// readable manifest. Validate it when it is available.
			return validateOptionalManifest(workspace)
		}
		return fmt.Errorf("cannot inspect manifest %q: %w", path, err)
	}
	if _, err := workspace.Manifest(); err != nil {
		return err
	}
	return nil
}

// validateOptionalManifest accepts an unavailable manifest, which lets the
// registry fall back to generation, but rejects a manifest that is present and
// malformed.
func validateOptionalManifest(workspace types.WorkspaceContext) error {
	_, err := workspace.Manifest()
	if err == nil ||
		errors.Is(err, fs.ErrNotExist) ||
		errors.Is(err, errors.ErrUnsupported) ||
		errors.Is(err, cemworkspace.ErrNoManifest) ||
		errors.Is(err, cemworkspace.ErrManifestNotFound) ||
		errors.Is(err, cemworkspace.ErrNoPackageJSON) {
		return nil
	}
	return err
}

// reloadableContext gives handlers a stable view of the current MCPContext.
// Reload builds a complete replacement context and swaps it in atomically, so
// each registry access observes either the old state or the new one. A handler
// that performs several accesses during a concurrent Reload may observe both
// generations; the swap is atomic per access, not per request. The old context
// stays valid for requests already in flight.
type reloadableContext struct {
	*mcpContextAdapter
	current atomic.Pointer[MCPContext]
}

func newReloadableContext(initial *MCPContext) *reloadableContext {
	r := &reloadableContext{}
	r.current.Store(initial)
	r.mcpContextAdapter = newMCPContextAdapter(r.load)
	return r
}

func (r *reloadableContext) load() *MCPContext {
	return r.current.Load()
}

func (r *reloadableContext) store(next *MCPContext) {
	r.current.Store(next)
}
