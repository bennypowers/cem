/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package resources

import (
	"testing"
)

func TestGetStringArg(t *testing.T) {
	tests := []struct {
		name         string
		args         map[string]any
		key          string
		defaultValue string
		want         string
	}{
		{
			name:         "key exists with string value",
			args:         map[string]any{"tagName": "my-element"},
			key:          "tagName",
			defaultValue: "default",
			want:         "my-element",
		},
		{
			name:         "key missing returns default",
			args:         map[string]any{"other": "value"},
			key:          "tagName",
			defaultValue: "fallback",
			want:         "fallback",
		},
		{
			name:         "wrong type returns default",
			args:         map[string]any{"count": 42},
			key:          "count",
			defaultValue: "zero",
			want:         "zero",
		},
		{
			name:         "nil map returns default",
			args:         nil,
			key:          "anything",
			defaultValue: "safe",
			want:         "safe",
		},
		{
			name:         "empty string value is returned",
			args:         map[string]any{"key": ""},
			key:          "key",
			defaultValue: "notempty",
			want:         "",
		},
		{
			name:         "empty default with missing key",
			args:         map[string]any{},
			key:          "missing",
			defaultValue: "",
			want:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getStringArg(tt.args, tt.key, tt.defaultValue)
			if got != tt.want {
				t.Errorf("getStringArg() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParseResourceURI(t *testing.T) {
	tests := []struct {
		name        string
		actualURI   string
		templateURI string
		isTemplate  bool
		wantArgs    map[string]any
		wantErr     bool
	}{
		{
			name:        "non-template returns empty args",
			actualURI:   "cem://elements",
			templateURI: "cem://elements",
			isTemplate:  false,
			wantArgs:    map[string]any{},
		},
		{
			name:        "single template parameter",
			actualURI:   "cem://element/my-button",
			templateURI: "cem://element/{tagName}",
			isTemplate:  true,
			wantArgs:    map[string]any{"tagName": "my-button"},
		},
		{
			name:        "multiple template parameters",
			actualURI:   "cem://element/my-button/attributes",
			templateURI: "cem://element/{tagName}/{section}",
			isTemplate:  true,
			wantArgs:    map[string]any{"tagName": "my-button", "section": "attributes"},
		},
		{
			name:        "sub-resource beyond template",
			actualURI:   "cem://element/my-button/css/parts",
			templateURI: "cem://element/{tagName}",
			isTemplate:  true,
			wantArgs: map[string]any{
				"tagName":     "my-button",
				"subResource": "css/parts",
			},
		},
		{
			name:        "actual URI shorter than template",
			actualURI:   "cem://element",
			templateURI: "cem://element/{tagName}/{section}",
			isTemplate:  true,
			wantArgs:    map[string]any{},
		},
		{
			name:        "non-parameter segment is ignored",
			actualURI:   "cem://element/my-button",
			templateURI: "cem://element/{tagName}",
			isTemplate:  true,
			wantArgs:    map[string]any{"tagName": "my-button"},
		},
		{
			name:        "static segment in template is not extracted",
			actualURI:   "cem://api/v2/elements",
			templateURI: "cem://api/v2/elements",
			isTemplate:  true,
			wantArgs:    map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseResourceURI(tt.actualURI, tt.templateURI, tt.isTemplate)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseResourceURI() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(got) != len(tt.wantArgs) {
				t.Fatalf("parseResourceURI() returned %d args, want %d: got=%v want=%v", len(got), len(tt.wantArgs), got, tt.wantArgs)
			}
			for key, wantVal := range tt.wantArgs {
				gotVal, ok := got[key]
				if !ok {
					t.Errorf("missing key %q in args", key)
					continue
				}
				if gotVal != wantVal {
					t.Errorf("args[%q] = %v, want %v", key, gotVal, wantVal)
				}
			}
		})
	}
}
