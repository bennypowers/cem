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
	"fmt"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/set"
	W "bennypowers.dev/cem/internal/workspace"
	"charm.land/huh/v2"
)

const fieldValueCustom = "__custom__"

type fieldValue struct {
	Title       string
	Description string
	Placeholder string
	Existing    string
	Detected    string
	Fallback    string
	Selected    string
	Custom      string
	ValidateFn  func(string) error
	gate        func() bool
}

func (f fieldValue) Value() string {
	if f.Existing != "" {
		return f.Existing
	}
	if f.Detected != "" {
		return f.Detected
	}
	return f.Fallback
}

func (f *fieldValue) isGated() bool {
	return f.gate != nil && !f.gate()
}

// Groups builds huh form groups for the field value selection flow.
// When multiple distinct values exist (existing, detected, fallback), shows a
// select with all options plus "Enter a new value". When only one value exists,
// shows just the input. After form.Run(), call Resolve() for the final value.
func (f *fieldValue) Groups() []*huh.Group {
	f.Selected = f.Value()

	var opts []huh.Option[string]
	seen := set.NewSet[string]()

	if f.Existing != "" {
		norm := strings.Join(splitCommaList(f.Existing), ", ")
		if !seen.Has(norm) {
			seen.Add(norm)
			opts = append(opts, huh.NewOption(f.Existing+" (existing)", f.Existing))
		}
	}
	if f.Detected != "" {
		norm := strings.Join(splitCommaList(f.Detected), ", ")
		if !seen.Has(norm) {
			seen.Add(norm)
			opts = append(opts, huh.NewOption(f.Detected+" (detected)", f.Detected))
		}
	}
	if f.Fallback != "" {
		norm := strings.Join(splitCommaList(f.Fallback), ", ")
		if !seen.Has(norm) {
			seen.Add(norm)
			opts = append(opts, huh.NewOption(f.Fallback+" (default)", f.Fallback))
		}
	}

	if len(opts) == 0 {
		f.Custom = f.Value()
		input := huh.NewInput().
			Title(f.Title).
			Placeholder(f.Placeholder).
			Value(&f.Custom)
		if f.ValidateFn != nil {
			input.Validate(f.ValidateFn)
		}
		g := huh.NewGroup(input).
			Title(f.Title).
			Description(f.Description)
		if f.gate != nil {
			g.WithHideFunc(f.isGated)
		}
		return []*huh.Group{g}
	}

	opts = append(opts, huh.NewOption("Enter a new value", fieldValueCustom))
	f.Selected = f.Value()

	selectGroup := huh.NewGroup(
		huh.NewSelect[string]().
			Title("Choose a value").
			Options(opts...).
			Value(&f.Selected),
	).Title(f.Title).
		Description(f.Description)
	if f.gate != nil {
		selectGroup.WithHideFunc(f.isGated)
	}

	customInput := huh.NewInput().
		Title(f.Title).
		Placeholder(f.Placeholder).
		Value(&f.Custom)
	if f.ValidateFn != nil {
		customInput.Validate(f.ValidateFn)
	}
	inputGroup := huh.NewGroup(customInput).
		WithHideFunc(func() bool {
			return f.isGated() || f.Selected != fieldValueCustom
		})

	return []*huh.Group{selectGroup, inputGroup}
}

// Resolve returns the final value after the form has run.
func (f *fieldValue) Resolve() string {
	if f.Selected == fieldValueCustom || f.Selected == "" {
		if f.Custom != "" {
			return f.Custom
		}
		return f.Fallback
	}
	return f.Selected
}

func splitCommaList(s string) []string {
	var result []string
	for part := range strings.SplitSeq(s, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

func validatePackageSpecifiers(input string) error {
	if input == "" {
		return nil
	}
	for _, s := range splitCommaList(input) {
		if W.IsPackageSpecifier(s) || W.IsURLSpecifier(s) || filepath.IsAbs(s) {
			continue
		}
		return fmt.Errorf("%q is not a valid specifier (use npm:, jsr:, https://, or an absolute path)", s)
	}
	return nil
}
