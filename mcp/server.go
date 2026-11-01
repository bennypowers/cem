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
package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"reflect"
	"slices"
	"time"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/version"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/mcp/resources"
	"bennypowers.dev/cem/mcp/tools"
	MCPTypes "bennypowers.dev/cem/mcp/types"
	"bennypowers.dev/cem/types"
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ServerInfo represents basic server information
type ServerInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// ServerConfig contains configuration options for the MCP server
type ServerConfig struct {
	MaxDescriptionLength int
	// Additional packages to load manifests from (URLs, npm:, or jsr: specifiers)
	AdditionalPackages []string
	// FileSystem to use for filesystem operations. If nil, defaults to OSFileSystem.
	FileSystem platform.FileSystem
}

// Server implements an MCP server for custom elements
type Server struct {
	workspace types.WorkspaceContext
	registry  *MCPContext
	server    *mcp.Server
	config    ServerConfig
}

// NewServer creates a new CEM MCP server with default configuration
func NewServer(workspace types.WorkspaceContext) (*Server, error) {
	return NewServerWithConfig(workspace, ServerConfig{
		MaxDescriptionLength: 2000,
	})
}

// NewServerWithConfig creates a new CEM MCP server with custom configuration
func NewServerWithConfig(workspace types.WorkspaceContext, config ServerConfig) (*Server, error) {
	helpers.SafeDebugLog("Creating CEM MCP server for workspace: %s", workspace.Root())

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cem",
		Version: version.GetVersion(),
	}, nil)

	// Registration is shared with Register so both entry points expose the
	// same tools and resources. Unlike Register, this constructor defers
	// manifest loading to Run, preserving the existing lifecycle, and keeps
	// the watcher-enabled registry that Run starts. Run reloads that initial
	// context in place; the handle is intentionally not retained here.
	registration, err := mount(server, workspace, config, registerConfig{conflict: ConflictReplace}, mountOptions{watch: true})
	if err != nil {
		return nil, err
	}

	return &Server{
		workspace: workspace,
		registry:  registration.registry.load(),
		server:    server,
		config:    config,
	}, nil
}

// Register adds cem's tools and resources to an existing MCP server and
// returns a handle for reloading the workspace data behind them.
//
// Register must be called before server starts serving sessions: the SDK's
// registration methods mutate the server's handler tables. The caller owns
// the server and its lifetime: Register does not start a long-lived serving
// transport or file watching. It creates its own MCPContext and loads the
// workspace manifests and any configured additional packages before
// registering handlers, so the mounted tools and resources are usable as
// soon as Register returns. The workspace must already be initialized.
//
// Options select part of cem's surface and choose how to treat entries that
// already exist on server. With no options, Register exposes every tool and
// resource and replaces conflicting entries, which is the behavior before
// options existed. Register is additive: it never removes entries. Calling it
// again adds or replaces selected entries, and handlers installed by earlier
// calls keep their own registry.
//
// Filters match tool names, static resource URIs, or URI template strings
// exactly and case-sensitively. Include selects a subset; an empty Include
// selects everything. Exclude is applied after Include and wins on overlap.
// An unknown Include entry is an error, while an unknown Exclude entry is
// accepted so hosts can keep a stable deny list across cem versions. Empty
// strings in either list are rejected. Repeated filter options accumulate,
// and the last WithConflictPolicy wins. Options are validated before registry
// data is loaded, and unknown Include entries are validated before any
// manifest or additional-package load.
//
// ConflictError and ConflictSkip inspect the server's current tools and
// resources through a temporary in-memory client session before registering.
// That session performs a normal initialize handshake, which fires any
// ServerOptions.InitializedHandler on the host, and is closed before Register
// returns. ConflictError fails when any selected entry already exists and
// returns a *RegistrationConflictError. ConflictSkip leaves such entries in
// place and registers only the rest. All option validation and conflict
// detection happen before the first handler is registered, so an error leaves
// server unchanged.
//
// The returned Registration reloads the manifests, configuration, and
// additional packages behind the handlers installed by this call. Reload does
// not change the registered tools, resources, schemas, or URIs.
func Register(server *mcp.Server, workspace types.WorkspaceContext, config ServerConfig, opts ...RegisterOption) (*Registration, error) {
	if server == nil {
		return nil, fmt.Errorf("server must not be nil")
	}
	if isNilWorkspace(workspace) {
		return nil, fmt.Errorf("workspace must not be nil")
	}

	cfg, err := applyRegisterOptions(opts)
	if err != nil {
		return nil, err
	}

	return mount(server, workspace, config, cfg, mountOptions{load: true})
}

// mountOptions controls registry lifecycle for a mount call.
type mountOptions struct {
	// watch allocates a file watcher for the underlying LSP registry.
	watch bool
	// load loads workspace data before returning. NewServerWithConfig defers
	// this load to Run.
	load bool
}

// mount builds a registry, applies registration options, and adds the
// selected handlers to server. Filters and conflicts are validated before
// workspace data is loaded, so a bad option or conflict fails fast and leaves
// the server unchanged.
func mount(server *mcp.Server, workspace types.WorkspaceContext, config ServerConfig, cfg registerConfig, opts mountOptions) (*Registration, error) {
	initial, err := newMCPContext(workspace, config.FileSystem, opts.watch)
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	registry := newReloadableContext(initial)
	plan, err := buildRegistrationPlan(registry, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.conflict != ConflictReplace {
		existing, err := existingServerSurface(server)
		if err != nil {
			return nil, err
		}
		if err := plan.resolveConflicts(existing, cfg.conflict); err != nil {
			return nil, err
		}
	}

	if opts.load {
		if err := loadRegistryData(initial, config.AdditionalPackages, "config"); err != nil {
			return nil, err
		}
	}

	plan.register(server)

	return &Registration{
		workspace: workspace,
		config:    config,
		registry:  registry,
	}, nil
}

// isNilWorkspace reports whether workspace is nil, including a typed nil
// implementation stored in the interface.
func isNilWorkspace(workspace types.WorkspaceContext) bool {
	if workspace == nil {
		return true
	}

	value := reflect.ValueOf(workspace)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

// loadRegistryData loads manifests and additional packages into registry.
// Additional package failures are logged and do not fail the load.
func loadRegistryData(registry *MCPContext, additionalPackages []string, source string) error {
	return loadRegistryDataWithPolicy(registry, additionalPackages, source, false)
}

// loadRegistryDataStrict loads manifests and additional packages into registry
// and returns any load failure. Reload uses it so a partial load never
// replaces a working registry.
func loadRegistryDataStrict(registry *MCPContext, additionalPackages []string, source string) error {
	return loadRegistryDataWithPolicy(registry, additionalPackages, source, true)
}

func loadRegistryDataWithPolicy(registry *MCPContext, additionalPackages []string, source string, strict bool) error {
	var err error
	if strict {
		err = registry.LoadManifestsStrict()
	} else {
		err = registry.LoadManifests()
	}
	if err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	if len(additionalPackages) > 0 {
		helpers.SafeDebugLog("Loading %d additional packages from %s", len(additionalPackages), source)
		if strict {
			if err := registry.LSPRegistry().LoadAdditionalPackagesStrict(additionalPackages); err != nil {
				return fmt.Errorf("failed to load additional packages: %w", err)
			}
		} else {
			_ = registry.LSPRegistry().LoadAdditionalPackages(additionalPackages)
		}
	}

	return nil
}

// registeredTool is a tool definition paired with its handler.
type registeredTool struct {
	tool    *mcp.Tool
	handler mcp.ToolHandler
}

// registeredResource is a resource or resource template paired with its
// handler.
type registeredResource struct {
	resource *mcp.Resource
	template *mcp.ResourceTemplate
	handler  mcp.ResourceHandler
}

// registrationPlan holds the filtered definitions that registration will add.
type registrationPlan struct {
	tools     []registeredTool
	resources []registeredResource
}

// buildRegistrationPlan loads the static tool and resource definitions,
// applies filters, and converts every selected definition before any handler
// is registered.
func buildRegistrationPlan(registry MCPTypes.MCPContext, cfg registerConfig) (*registrationPlan, error) {
	toolDefs, err := tools.Tools(registry)
	if err != nil {
		return nil, fmt.Errorf("failed to load tool definitions: %w", err)
	}

	resourceDefs, err := resources.Resources(registry)
	if err != nil {
		return nil, fmt.Errorf("failed to load resource definitions: %w", err)
	}

	toolNames := make([]string, 0, len(toolDefs))
	for _, def := range toolDefs {
		toolNames = append(toolNames, def.Name)
	}
	selectedTools, err := selectNames("tool", cfg.tools, toolNames)
	if err != nil {
		return nil, err
	}

	var resourceURIs, templateURIs []string
	for _, def := range resourceDefs {
		if def.URITemplate {
			templateURIs = append(templateURIs, def.URI)
		} else {
			resourceURIs = append(resourceURIs, def.URI)
		}
	}
	selectedResources, err := selectNames("resource", cfg.resources, resourceURIs)
	if err != nil {
		return nil, err
	}
	selectedTemplates, err := selectNames("resource template", cfg.templates, templateURIs)
	if err != nil {
		return nil, err
	}

	plan := &registrationPlan{}
	for _, def := range toolDefs {
		if _, ok := selectedTools[def.Name]; !ok {
			continue
		}
		tool, err := toolForRegistration(def)
		if err != nil {
			return nil, err
		}
		plan.tools = append(plan.tools, registeredTool{tool: tool, handler: def.Handler})
	}

	for _, def := range resourceDefs {
		selected := selectedResources
		if def.URITemplate {
			selected = selectedTemplates
		}
		if _, ok := selected[def.URI]; !ok {
			continue
		}

		if def.URITemplate {
			plan.resources = append(plan.resources, registeredResource{
				template: &mcp.ResourceTemplate{
					URITemplate: def.URI,
					Name:        def.Name,
					MIMEType:    def.MimeType,
					Description: def.Description,
				},
				handler: def.Handler,
			})
			continue
		}

		plan.resources = append(plan.resources, registeredResource{
			resource: &mcp.Resource{
				URI:         def.URI,
				Name:        def.Name,
				MIMEType:    def.MimeType,
				Description: def.Description,
			},
			handler: def.Handler,
		})
	}

	return plan, nil
}

// selectNames validates a Selection against the available names and returns
// the selected set. Exclude is applied after Include and wins on overlap.
func selectNames(kind string, selection Selection, available []string) (map[string]struct{}, error) {
	all := make(map[string]struct{}, len(available))
	for _, name := range available {
		all[name] = struct{}{}
	}

	selected := make(map[string]struct{}, len(all))
	if len(selection.Include) == 0 {
		maps.Copy(selected, all)
	} else {
		for _, name := range selection.Include {
			if _, ok := all[name]; !ok {
				return nil, fmt.Errorf("%s include %q does not match any registered %s", kind, name, kind)
			}
			selected[name] = struct{}{}
		}
	}

	for _, name := range selection.Exclude {
		delete(selected, name)
	}

	return selected, nil
}

// toolForRegistration converts a tool definition's input schema.
func toolForRegistration(def MCPTypes.ToolDefinition) (*mcp.Tool, error) {
	var inputSchema *jsonschema.Schema
	if def.InputSchema != nil {
		schemaBytes, err := json.Marshal(def.InputSchema)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal input schema for tool %s: %w", def.Name, err)
		}

		var schema jsonschema.Schema
		if err := json.Unmarshal(schemaBytes, &schema); err != nil {
			return nil, fmt.Errorf("failed to unmarshal input schema for tool %s: %w", def.Name, err)
		}
		inputSchema = &schema
	}

	return &mcp.Tool{
		Name:        def.Name,
		Description: def.Description,
		InputSchema: inputSchema,
	}, nil
}

// serverSurface is the set of entries already registered on a server.
type serverSurface struct {
	tools     map[string]struct{}
	resources map[string]struct{}
	templates map[string]struct{}
}

// surfaceInspectionTimeout bounds the temporary session used to inspect the
// caller's server. The work is local, so a short bound is enough to avoid a
// stalled host middleware blocking Register indefinitely.
const surfaceInspectionTimeout = 10 * time.Second

// existingServerSurface lists the tools, resources, and resource templates
// already registered on server. The SDK does not expose its handler tables,
// so this uses a short-lived in-memory client session. The session sends the
// normal initialize handshake, which fires any host InitializedHandler, and is
// closed before this function returns.
func existingServerSurface(server *mcp.Server) (serverSurface, error) {
	ctx, cancel := context.WithTimeout(context.Background(), surfaceInspectionTimeout)
	defer cancel()

	serverTransport, clientTransport := mcp.NewInMemoryTransports()
	serverSession, err := server.Connect(ctx, serverTransport, nil)
	if err != nil {
		return serverSurface{}, fmt.Errorf("failed to inspect existing server surface: %w", err)
	}
	defer func() { _ = serverSession.Close() }()

	client := mcp.NewClient(&mcp.Implementation{
		Name:    "cem-registration-preflight",
		Version: version.GetVersion(),
	}, nil)
	clientSession, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		return serverSurface{}, fmt.Errorf("failed to inspect existing server surface: %w", err)
	}
	defer func() { _ = clientSession.Close() }()

	surface := serverSurface{
		tools:     make(map[string]struct{}),
		resources: make(map[string]struct{}),
		templates: make(map[string]struct{}),
	}

	for tool, err := range clientSession.Tools(ctx, nil) {
		if err != nil {
			return serverSurface{}, fmt.Errorf("failed to inspect existing server tools: %w", err)
		}
		surface.tools[tool.Name] = struct{}{}
	}
	for resource, err := range clientSession.Resources(ctx, nil) {
		if err != nil {
			return serverSurface{}, fmt.Errorf("failed to inspect existing server resources: %w", err)
		}
		surface.resources[resource.URI] = struct{}{}
	}
	for template, err := range clientSession.ResourceTemplates(ctx, nil) {
		if err != nil {
			return serverSurface{}, fmt.Errorf("failed to inspect existing server resource templates: %w", err)
		}
		surface.templates[template.URITemplate] = struct{}{}
	}

	return surface, nil
}

// resolveConflicts applies policy to any plan entries that already exist on
// the server. mount calls it only for ConflictSkip and ConflictError; the
// ConflictReplace and default arms keep the function exhaustive for future
// callers.
func (p *registrationPlan) resolveConflicts(existing serverSurface, policy ConflictPolicy) error {
	switch policy {
	case ConflictReplace:
		return nil

	case ConflictSkip:
		p.tools = slices.DeleteFunc(p.tools, func(tool registeredTool) bool {
			_, ok := existing.tools[tool.tool.Name]
			return ok
		})
		p.resources = slices.DeleteFunc(p.resources, func(resource registeredResource) bool {
			if resource.template != nil {
				_, ok := existing.templates[resource.template.URITemplate]
				return ok
			}
			_, ok := existing.resources[resource.resource.URI]
			return ok
		})
		return nil

	case ConflictError:
		conflict := &RegistrationConflictError{}
		for _, tool := range p.tools {
			if _, ok := existing.tools[tool.tool.Name]; ok {
				conflict.Tools = append(conflict.Tools, tool.tool.Name)
			}
		}
		for _, resource := range p.resources {
			if resource.template != nil {
				if _, ok := existing.templates[resource.template.URITemplate]; ok {
					conflict.ResourceTemplates = append(conflict.ResourceTemplates, resource.template.URITemplate)
				}
				continue
			}
			if _, ok := existing.resources[resource.resource.URI]; ok {
				conflict.Resources = append(conflict.Resources, resource.resource.URI)
			}
		}
		slices.Sort(conflict.Tools)
		slices.Sort(conflict.Resources)
		slices.Sort(conflict.ResourceTemplates)
		if len(conflict.Tools)+len(conflict.Resources)+len(conflict.ResourceTemplates) > 0 {
			return conflict
		}
		return nil

	default:
		return fmt.Errorf("unknown conflict policy %d", policy)
	}
}

// register adds every planned tool and resource to server.
func (p *registrationPlan) register(server *mcp.Server) {
	for _, tool := range p.tools {
		server.AddTool(tool.tool, tool.handler)
	}
	for _, resource := range p.resources {
		if resource.template != nil {
			server.AddResourceTemplate(resource.template, resource.handler)
			continue
		}
		server.AddResource(resource.resource, resource.handler)
	}
}

// GetInfo returns server information
func (s *Server) GetInfo() ServerInfo {
	return ServerInfo{
		Name:        "cem",
		Version:     version.GetVersion(),
		Description: "Custom Elements Manifest MCP Server - provides intelligent context about custom elements for AI systems",
	}
}

// Registry returns the MCP registry for testing
// This is used by tests to verify stdout cleanliness
func (s *Server) Registry() *MCPContext {
	return s.registry
}

// reloadFromWatcher refreshes the in-place registry used by Run. It is the
// lenient counterpart to Registration.Reload, which validates the manifest and
// swaps a replacement generation. Keep cache invalidation changes in sync with
// both paths.
func (s *Server) reloadFromWatcher() {
	helpers.SafeDebugLog("File change detected, reloading manifests and config")
	s.registry.InvalidateConfig()
	s.registry.InvalidateManifest()
	if err := s.registry.LoadManifests(); err != nil {
		helpers.SafeDebugLog("Warning: Failed to reload manifests: %v", err)
	}
}

// Run starts the MCP server with stdio transport
func (s *Server) Run(ctx context.Context) error {
	helpers.SafeDebugLog("Starting CEM MCP server with stdio transport")

	// Initialize the registry
	if err := loadRegistryData(s.registry, s.config.AdditionalPackages, "CLI args"); err != nil {
		return err
	}

	// Start file watching for manifest and config changes. This in-place
	// recovery path is intentionally lenient. Registration.Reload is the strict
	// path: it validates the manifest and swaps a replacement generation. Keep
	// cache invalidation changes in sync with that path.
	if err := s.registry.LSPRegistry().StartFileWatching(s.reloadFromWatcher); err != nil {
		helpers.SafeDebugLog("Warning: Could not start file watching: %v", err)
	}

	// Run the MCP server
	return s.server.Run(ctx, &mcp.StdioTransport{})
}
