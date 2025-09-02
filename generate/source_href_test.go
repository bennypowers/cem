/*
Copyright © 2025 Benny Powers

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
package generate

import (
	"fmt"
	"testing"
)

func TestBuildSourceHref(t *testing.T) {
	tests := []struct {
		name                  string
		sourceControlRootUrl  string
		filePath              string
		lineNumber            uint
		expected              string
		expectError           bool
	}{
		{
			name:                 "basic GitHub URL",
			sourceControlRootUrl: "https://github.com/user/repo/tree/main/",
			filePath:             "src/components.ts",
			lineNumber:           42,
			expected:             "https://github.com/user/repo/tree/main/src/components.ts#L42",
			expectError:          false,
		},
		{
			name:                 "URL without trailing slash",
			sourceControlRootUrl: "https://github.com/user/repo/tree/main",
			filePath:             "src/components.ts",
			lineNumber:           42,
			expected:             "https://github.com/user/repo/tree/main/src/components.ts#L42",
			expectError:          false,
		},
		{
			name:                 "file path with leading slash",
			sourceControlRootUrl: "https://github.com/user/repo/tree/main/",
			filePath:             "/src/components.ts",
			lineNumber:           1,
			expected:             "https://github.com/user/repo/tree/main/src/components.ts#L1",
			expectError:          false,
		},
		{
			name:                 "file path with leading ./",
			sourceControlRootUrl: "https://github.com/user/repo/tree/main/",
			filePath:             "./src/components.ts",
			lineNumber:           15,
			expected:             "https://github.com/user/repo/tree/main/src/components.ts#L15",
			expectError:          false,
		},
		{
			name:                 "empty sourceControlRootUrl",
			sourceControlRootUrl: "",
			filePath:             "src/components.ts",
			lineNumber:           42,
			expected:             "",
			expectError:          true,
		},
		{
			name:                 "empty filePath",
			sourceControlRootUrl: "https://github.com/user/repo/tree/main/",
			filePath:             "",
			lineNumber:           42,
			expected:             "",
			expectError:          true,
		},
		{
			name:                 "invalid URL",
			sourceControlRootUrl: "not-a-valid-url",
			filePath:             "src/components.ts",
			lineNumber:           42,
			expected:             "",
			expectError:          true,
		},
		{
			name:                 "URL missing scheme",
			sourceControlRootUrl: "//github.com/user/repo/tree/main/",
			filePath:             "src/components.ts",
			lineNumber:           42,
			expected:             "",
			expectError:          true,
		},
		{
			name:                 "GitLab URL",
			sourceControlRootUrl: "https://gitlab.com/user/repo/-/tree/main/",
			filePath:             "lib/index.js",
			lineNumber:           123,
			expected:             "https://gitlab.com/user/repo/-/tree/main/lib/index.js#L123",
			expectError:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := buildSourceHref(tt.sourceControlRootUrl, tt.filePath, tt.lineNumber)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("buildSourceHref() expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("buildSourceHref() unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("buildSourceHref() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGenerateSourceReferenceNilNode(t *testing.T) {
	mp := &ModuleProcessor{
		file: "test.ts",
		code: []byte("test content"),
	}
	
	sourceRef, err := mp.generateSourceReference(nil)
	if err != nil {
		t.Errorf("generateSourceReference(nil) should not return error, got: %v", err)
	}
	if sourceRef != nil {
		t.Errorf("generateSourceReference(nil) should return nil, got: %v", sourceRef)
	}
}

func TestByteOffsetToLineNumberPerformance(t *testing.T) {
	// Create a large file with many lines to test performance
	lines := 1000
	var content []byte
	for i := 0; i < lines; i++ {
		content = append(content, []byte(fmt.Sprintf("// Line %d\n", i+1))...)
	}

	mp := &ModuleProcessor{
		code: content,
	}

	// Test multiple lookups - this should be fast with caching
	for i := 0; i < 10; i++ {
		offset := uint(len(content) / 2) // Middle of file
		lineNum := mp.byteOffsetToLineNumber(offset)
		if lineNum < 400 || lineNum > 600 { // Should be around line 500
			t.Errorf("Expected line number around 500, got %d", lineNum)
		}
	}
}

func TestByteOffsetToLineNumber(t *testing.T) {
	tests := []struct {
		name     string
		code     []byte
		offset   uint
		expected uint
	}{
		{
			name:     "first line",
			code:     []byte("function foo() {\n  return 'bar';\n}"),
			offset:   0,
			expected: 1,
		},
		{
			name:     "second line",
			code:     []byte("function foo() {\n  return 'bar';\n}"),
			offset:   17,  // start of "return"
			expected: 2,
		},
		{
			name:     "third line",
			code:     []byte("function foo() {\n  return 'bar';\n}"),
			offset:   33,  // start of "}"
			expected: 3,
		},
		{
			name:     "middle of first line",
			code:     []byte("function foo() {\n  return 'bar';\n}"),
			offset:   9,   // start of "foo"
			expected: 1,
		},
		{
			name:     "empty file",
			code:     []byte(""),
			offset:   0,
			expected: 1,
		},
		{
			name:     "single line",
			code:     []byte("const x = 42;"),
			offset:   6,   // start of "x"
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock ModuleProcessor with the test code
			mp := &ModuleProcessor{
				code: tt.code,
			}
			result := mp.byteOffsetToLineNumber(tt.offset)
			if result != tt.expected {
				t.Errorf("byteOffsetToLineNumber() = %v, want %v", result, tt.expected)
			}
		})
	}
}