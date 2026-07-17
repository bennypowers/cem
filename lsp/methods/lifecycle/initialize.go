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
package lifecycle

import (
	"encoding/json"
	"os"

	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/version"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"go.lsp.dev/protocol"
)

// cemInitializationOptions represents the custom initialization options for CEM LSP
type cemInitializationOptions struct {
	AdditionalPackages []string `json:"additionalPackages,omitempty"`
	InlayHints         *bool    `json:"inlayHints,omitempty"`
}

// Initialize handles the LSP initialize request
func Initialize(ctx types.ServerContext, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	// Configure centralized logger for LSP mode
	logging.SetLSPClient(ctx.Client())

	// Send initialization message with version
	serverVersion := version.GetVersion()
	logging.Info("CEM LSP Server initializing (version %s)...", serverVersion)

	// Debug info about binary location
	if executable, err := os.Executable(); err == nil {
		logging.Debug("LSP server binary: %s", executable)
	}

	//nolint:staticcheck // RootURI still needed for older LSP clients
	if params.RootURI != nil {
		helpers.SafeDebugLog("[INITIALIZE] LSP client provided root URI: %s", *params.RootURI) //nolint:staticcheck
	}
	if folders, ok := params.WorkspaceFolders.Get(); ok {
		for i, folder := range folders {
			helpers.SafeDebugLog("[INITIALIZE] LSP client workspace folder %d: %s (%s)", i, folder.URI, folder.Name)
		}
	}

	// Update workspace context with information from LSP client
	var rootStr *string
	//nolint:staticcheck // RootURI still needed for older LSP clients
	if params.RootURI != nil {
		s := string(*params.RootURI) //nolint:staticcheck
		rootStr = &s
	}
	var folders []protocol.WorkspaceFolder
	if f, ok := params.WorkspaceFolders.Get(); ok {
		folders = f
	}
	if err := ctx.UpdateWorkspaceFromLSP(rootStr, folders); err != nil {
		helpers.SafeDebugLog("[INITIALIZE] Warning: Failed to update workspace from LSP parameters: %v", err)
		// Don't fail initialization, just log the warning
	}

	// Parse initializationOptions for additional packages and settings
	if len(params.InitializationOptions) > 0 {
		options := parseInitializationOptions(params.InitializationOptions)
		if len(options.AdditionalPackages) > 0 {
			helpers.SafeDebugLog("[INITIALIZE] Found %d additional packages in initializationOptions", len(options.AdditionalPackages))
			ctx.SetAdditionalPackages(options.AdditionalPackages)
		}
		if options.InlayHints != nil {
			config := ctx.Config()
			config.InlayHints = options.InlayHints
			ctx.SetConfig(config)
		}
	}

	// Detect pull diagnostics support
	if params.Capabilities.TextDocument != nil &&
		params.Capabilities.TextDocument.Diagnostic != nil {
		ctx.SetUsePullDiagnostics(true)
		helpers.SafeDebugLog("[INITIALIZE] Client supports pull diagnostics (LSP 3.17)")
	} else {
		ctx.SetUsePullDiagnostics(false)
		helpers.SafeDebugLog("[INITIALIZE] Client uses push diagnostics (LSP 3.16)")
	}

	openClose := true
	changeKind := protocol.TextDocumentSyncKindIncremental
	resolveProvider := true

	var capabilities protocol.ServerCapabilities
	capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{
		OpenClose: &openClose,
		Change:    &changeKind,
	}
	capabilities.HoverProvider = &protocol.HoverOptions{}
	capabilities.CompletionProvider = &protocol.CompletionOptions{
		TriggerCharacters: []string{"<", "=", "\"", "@", ".", "?"},
		ResolveProvider:   &resolveProvider,
	}
	capabilities.DefinitionProvider = &protocol.DefinitionOptions{}
	capabilities.ReferencesProvider = &protocol.ReferenceOptions{}
	capabilities.CodeActionProvider = &protocol.CodeActionOptions{
		CodeActionKinds: []protocol.CodeActionKind{
			protocol.CodeActionKindQuickFix,
		},
	}
	capabilities.WorkspaceSymbolProvider = &protocol.WorkspaceSymbolOptions{}
	capabilities.InlayHintProvider = &protocol.InlayHintOptions{}

	if ctx.UsePullDiagnostics() {
		identifier := "cem"
		capabilities.DiagnosticProvider = &protocol.DiagnosticOptions{
			Identifier:            &identifier,
			InterFileDependencies: false,
			WorkspaceDiagnostics:  true,
		}
	}

	return &protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: protocol.ServerInfo{
			Name:    "cem-lsp",
			Version: protocol.NewOptional(serverVersion),
		},
	}, nil
}

// parseInitializationOptions attempts to parse the initializationOptions into our expected format.
func parseInitializationOptions(raw protocol.LSPAny) cemInitializationOptions {
	result := cemInitializationOptions{}

	var opts map[string]any
	if err := json.Unmarshal(raw, &opts); err != nil {
		return result
	}

	if additionalPackages, ok := opts["additionalPackages"]; ok {
		result.AdditionalPackages = parseStringSlice(additionalPackages)
	}

	if cemSettings, ok := opts["cem"]; ok {
		if cemMap, ok := cemSettings.(map[string]any); ok {
			if inlayHints, ok := cemMap["inlayHints"]; ok {
				if b, ok := inlayHints.(bool); ok {
					result.InlayHints = &b
				}
			}
			if additionalPackages, ok := cemMap["additionalPackages"]; ok {
				if parsed := parseStringSlice(additionalPackages); len(parsed) > 0 {
					result.AdditionalPackages = parsed
				}
			}
		}
	}

	return result
}

// parseStringSlice attempts to convert various formats to []string
func parseStringSlice(v any) []string {
	switch arr := v.(type) {
	case []string:
		return arr
	case []any:
		// JSON arrays are typically decoded as []any
		result := make([]string, 0, len(arr))
		for _, item := range arr {
			if s, ok := item.(string); ok {
				result = append(result, s)
			}
		}
		return result
	}
	return nil
}
