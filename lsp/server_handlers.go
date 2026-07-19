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
	"runtime/debug"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/lifecycle"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/methods/textDocument/codeAction"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/methods/textDocument/definition"
	"bennypowers.dev/cem/lsp/methods/textDocument/diagnostic"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/methods/textDocument/inlayHint"
	"bennypowers.dev/cem/lsp/methods/textDocument/references"
	"bennypowers.dev/cem/lsp/methods/workspace/configuration"
	workspaceDiag "bennypowers.dev/cem/lsp/methods/workspace/diagnostic"
	"bennypowers.dev/cem/lsp/methods/workspace/symbol"
	"go.lsp.dev/protocol"
)

func (s *Server) Initialize(_ context.Context, params *protocol.InitializeParams) (_ *protocol.InitializeResult, err error) {
	defer s.recover("initialize", &err)
	return lifecycle.Initialize(s, params)
}

func (s *Server) Initialized(_ context.Context, params *protocol.InitializedParams) (err error) {
	defer s.recover("initialized", &err)
	return lifecycle.Initialized(s, params)
}

func (s *Server) Shutdown(_ context.Context) (err error) {
	defer s.recover("shutdown", &err)
	return lifecycle.Shutdown(s)
}

func (s *Server) SetTrace(_ context.Context, params *protocol.SetTraceParams) (err error) {
	defer s.recover("$/setTrace", &err)
	return lifecycle.SetTrace(s, params)
}

func (s *Server) DidOpen(_ context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	defer s.recover("textDocument/didOpen", &err)
	return textDocument.DidOpen(s, params)
}

func (s *Server) DidChange(_ context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	defer s.recover("textDocument/didChange", &err)
	return textDocument.DidChange(s, params)
}

func (s *Server) DidClose(_ context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	defer s.recover("textDocument/didClose", &err)
	return textDocument.DidClose(s, params)
}

func (s *Server) Hover(_ context.Context, params *protocol.HoverParams) (_ *protocol.Hover, err error) {
	defer s.recover("textDocument/hover", &err)
	return hover.Hover(s, params)
}

func (s *Server) Completion(_ context.Context, params *protocol.CompletionParams) (_ protocol.CompletionResult, err error) {
	defer s.recover("textDocument/completion", &err)
	result, err := completion.Completion(s, params)
	if err != nil || result == nil {
		return nil, err
	}
	return protocol.CompletionItemSlice(result), nil
}

func (s *Server) CompletionResolve(_ context.Context, params *protocol.CompletionItem) (_ *protocol.CompletionItem, err error) {
	defer s.recover("completionItem/resolve", &err)
	return completion.Resolve(s, params)
}

func (s *Server) Definition(_ context.Context, params *protocol.DefinitionParams) (_ protocol.DefinitionResult, err error) {
	defer s.recover("textDocument/definition", &err)
	result, err := definition.Definition(s, params)
	if err != nil || result == nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) References(_ context.Context, params *protocol.ReferenceParams) (_ []protocol.Location, err error) {
	defer s.recover("textDocument/references", &err)
	return references.References(s, params)
}

func (s *Server) CodeAction(_ context.Context, params *protocol.CodeActionParams) (_ []protocol.CommandOrCodeAction, err error) {
	defer s.recover("textDocument/codeAction", &err)
	result, err := codeAction.CodeAction(s, params)
	if err != nil || result == nil {
		return nil, err
	}
	out := make([]protocol.CommandOrCodeAction, len(result))
	for i := range result {
		out[i] = &result[i]
	}
	return out, nil
}

func (s *Server) Diagnostic(_ context.Context, params *protocol.DocumentDiagnosticParams) (_ protocol.DocumentDiagnosticReport, err error) {
	defer s.recover("textDocument/diagnostic", &err)
	return diagnostic.DocumentDiagnostic(s, params)
}

func (s *Server) DiagnosticWorkspace(_ context.Context, params *protocol.WorkspaceDiagnosticParams) (_ *protocol.WorkspaceDiagnosticReport, err error) {
	defer s.recover("workspace/diagnostic", &err)
	return workspaceDiag.WorkspaceDiagnostic(s, params)
}

func (s *Server) InlayHint(_ context.Context, params *protocol.InlayHintParams) (_ []protocol.InlayHint, err error) {
	defer s.recover("textDocument/inlayHint", &err)
	return inlayHint.InlayHint(s, params)
}

func (s *Server) Symbols(_ context.Context, params *protocol.WorkspaceSymbolParams) (_ protocol.WorkspaceSymbolResult, err error) {
	defer s.recover("workspace/symbol", &err)
	result, err := symbol.Symbol(s, params)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return protocol.SymbolInformationSlice(result), nil
}

func (s *Server) DidChangeConfiguration(_ context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	defer s.recover("workspace/didChangeConfiguration", &err)
	return configuration.DidChangeConfiguration(s, params)
}

func (s *Server) recover(method string, err *error) {
	if r := recover(); r != nil {
		helpers.SafeDebugLog("[LSP] PANIC in %s: %v\nStack trace:\n%s",
			method, r, string(debug.Stack()))
		*err = fmt.Errorf("internal error in %s", method)
	}
}
