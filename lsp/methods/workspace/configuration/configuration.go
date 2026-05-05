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
package configuration

import (
	"encoding/json"
	"fmt"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/bennypowers/glsp"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

// DidChangeConfiguration handles workspace/didChangeConfiguration notifications
func DidChangeConfiguration(ctx types.ServerContext, context *glsp.Context, params *protocol.DidChangeConfigurationParams) error {
	helpers.SafeDebugLog("[CONFIG] Configuration changed")

	config, found, err := parseConfiguration(params.Settings)
	if err != nil {
		helpers.SafeDebugLog("[CONFIG] Failed to parse configuration: %v", err)
		return nil
	}

	if !found {
		return nil
	}

	ctx.SetConfig(config)
	helpers.SafeDebugLog("[CONFIG] Updated configuration: inlayHints=%v", config.InlayHints)
	return nil
}

func parseConfiguration(settings any) (types.ServerConfig, bool, error) {
	config := types.DefaultConfig()

	if settings == nil {
		return config, false, nil
	}

	settingsMap, ok := settings.(map[string]any)
	if !ok {
		return config, false, fmt.Errorf("settings is not a map")
	}

	cemVal, exists := settingsMap["cem"]
	if !exists {
		return config, false, nil
	}

	settingsObj, ok := cemVal.(map[string]any)
	if !ok {
		return config, false, fmt.Errorf("cem settings is not an object")
	}

	jsonBytes, err := json.Marshal(settingsObj)
	if err != nil {
		return config, false, fmt.Errorf("failed to marshal settings: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, &config); err != nil {
		return config, false, fmt.Errorf("failed to unmarshal settings: %w", err)
	}

	return config, true, nil
}
