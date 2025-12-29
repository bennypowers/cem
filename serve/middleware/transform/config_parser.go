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

package transform

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	cfg "bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/internal/platform"
)

// ParseConfigFileURLRewrites reads a cem.yaml config file and extracts URL rewrites.
// Returns URL rewrites and the absolute path of the config file (for change tracking).
// Returns empty slice if config file doesn't exist or has no URL rewrites.
//
// Example config file:
//
//	serve:
//	  urlRewrites:
//	    - urlPattern: "/dist/:path*"
//	      urlTemplate: "/src/{{.path}}"
func ParseConfigFileURLRewrites(configPath string, fs platform.FileSystem) ([]cfg.URLRewrite, []string, error) {
	// Convert to absolute path for consistent tracking
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		absPath = configPath
	}

	// Read config file
	data, err := fs.ReadFile(configPath)
	if err != nil {
		// Missing config file is not an error - just return empty
		if os.IsNotExist(err) {
			return []cfg.URLRewrite{}, []string{}, nil
		}
		return nil, nil, fmt.Errorf("reading config file: %w", err)
	}

	// Parse YAML
	var config struct {
		Serve struct {
			URLRewrites []cfg.URLRewrite `yaml:"urlRewrites"`
		} `yaml:"serve"`
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, nil, fmt.Errorf("parsing config YAML: %w", err)
	}

	// Return rewrites and the absolute path for tracking
	return config.Serve.URLRewrites, []string{absPath}, nil
}
