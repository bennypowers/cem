/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/tui"
	"github.com/tidwall/gjson"
)

func writeConfigAtomic(fsys platform.FileSystem, outPath string, output []byte) error {
	if err := fsys.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	tmp, err := fsys.CreateTemp(filepath.Dir(outPath), ".cem-config-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmp.Name()
	var cleanedUp sync.Once
	cleanup := func() {
		cleanedUp.Do(func() {
			_ = fsys.Remove(tmpPath)
		})
	}
	defer cleanup()

	sigCh := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-sigCh:
			cleanup()
			os.Exit(1)
		case <-done:
		}
	}()
	defer func() {
		signal.Stop(sigCh)
		close(done)
	}()

	if _, err := tmp.Write(output); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("failed to write config: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	if err := fsys.Rename(tmpPath, outPath); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func offerPackageJSONUpdate(fsys platform.FileSystem, root string, cfg *IC.CemConfig) error {
	pkgPath := filepath.Join(root, "package.json")
	data, err := fsys.ReadFile(pkgPath)
	if err != nil {
		return nil
	}

	if !json.Valid(data) {
		return nil
	}

	if gjson.GetBytes(data, "customElements").Exists() {
		return nil
	}

	output := cfg.Generate.Output
	if output == "" {
		output = "custom-elements.json"
	}

	confirmed, err := tui.Confirm(fmt.Sprintf("Add \"customElements\": \"%s\" to package.json?", output), true)
	if err != nil || !confirmed {
		return nil
	}

	entry, marshalErr := json.Marshal(output)
	if marshalErr != nil {
		return marshalErr
	}

	lastBrace := bytes.LastIndexByte(data, '}')
	if lastBrace < 0 {
		return nil
	}

	prefix := bytes.TrimSpace(data[:lastBrace])
	needsComma := len(prefix) > 0 && prefix[len(prefix)-1] != '{'
	var insertion string
	if needsComma {
		insertion = fmt.Sprintf(",\n  \"customElements\": %s", string(entry))
	} else {
		insertion = fmt.Sprintf("\n  \"customElements\": %s", string(entry))
	}

	var buf bytes.Buffer
	buf.Write(data[:lastBrace])
	buf.WriteString(insertion)
	buf.WriteByte('\n')
	buf.WriteByte('}')
	if lastBrace+1 < len(data) {
		buf.Write(data[lastBrace+1:])
	}

	return writeConfigAtomic(fsys, pkgPath, buf.Bytes())
}
