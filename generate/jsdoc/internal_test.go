package jsdoc

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

func TestStripTrailingSplat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "no splat", input: "hello world", want: "hello world"},
		{name: "trailing splat", input: "hello *", want: "hello"},
		{name: "trailing splat no space", input: "hello*", want: "hello"},
		{name: "trailing splat multiple spaces", input: "hello   *", want: "hello"},
		{name: "empty string", input: "", want: ""},
		{name: "just splat", input: "*", want: ""},
		{name: "splat in middle", input: "hello * world", want: "hello * world"},
		{name: "multiple trailing splats", input: "hello * *", want: "hello *"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, stripTrailingSplat(tt.input))
		})
	}
}

func TestNormalizeJsdocLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "plain text no stars",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "single line with leading star",
			input: " * hello world",
			want:  "hello world",
		},
		{
			name:  "multiline with leading stars",
			input: " * line one\n * line two",
			want:  "line one\nline two",
		},
		{
			name:  "tab-indented stars",
			input: "\t* hello\n\t* world",
			want:  "hello\nworld",
		},
		{
			name:  "trailing splat removal",
			input: " * description\n *",
			want:  "description\n",
		},
		{
			name:  "mixed indentation",
			input: "  *  indented\n * normal",
			want:  "indented\nnormal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, normalizeJsdocLines(tt.input))
		})
	}
}

func TestFindNamedMatches(t *testing.T) {
	re := regexp.MustCompile(`(?P<first>\w+)\s+(?P<second>\w+)?(?:\s+(?P<third>\w+))?`)

	tests := []struct {
		name                     string
		input                    string
		includeNotMatchedOptional bool
		want                     map[string]string
	}{
		{
			name:                     "all groups match",
			input:                    "hello world again",
			includeNotMatchedOptional: false,
			want:                     map[string]string{"first": "hello", "second": "world", "third": "again"},
		},
		{
			name:                     "optional group not matched, exclude",
			input:                    "hello world",
			includeNotMatchedOptional: false,
			want:                     map[string]string{"first": "hello", "second": "world"},
		},
		{
			name:                     "optional group not matched, include empty",
			input:                    "hello world",
			includeNotMatchedOptional: true,
			want:                     map[string]string{"first": "hello", "second": "world", "third": ""},
		},
		{
			name:                     "no match at all",
			input:                    "!!!",
			includeNotMatchedOptional: false,
			want:                     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findNamedMatches(re, tt.input, tt.includeNotMatchedOptional)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExtractExplicitCaption(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantCaption string
		wantCode    string
	}{
		{
			name:        "no caption",
			input:       "just some code",
			wantCaption: "",
			wantCode:    "just some code",
		},
		{
			name:        "with caption",
			input:       "<caption>My Caption</caption>\ncode here",
			wantCaption: "My Caption",
			wantCode:    "code here",
		},
		{
			name:        "caption with whitespace",
			input:       "<caption>  My Caption  </caption>\n  code here  ",
			wantCaption: "My Caption",
			wantCode:    "code here",
		},
		{
			name:        "multiline code after caption",
			input:       "<caption>Title</caption>\nline1\nline2",
			wantCaption: "Title",
			wantCode:    "line1\nline2",
		},
		{
			name:        "empty caption",
			input:       "<caption></caption>\ncode",
			wantCaption: "",
			wantCode:    "code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			caption, code := extractExplicitCaption(tt.input)
			assert.Equal(t, tt.wantCaption, caption)
			assert.Equal(t, tt.wantCode, code)
		})
	}
}

func TestExtractImplicitCaption(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantCaption string
		wantCode    string
	}{
		{
			name:        "no code block",
			input:       "just text",
			wantCaption: "",
			wantCode:    "just text",
		},
		{
			name:        "code block on first line",
			input:       "```js\nconsole.log('hello');\n```",
			wantCaption: "",
			wantCode:    "```js\nconsole.log('hello');\n```",
		},
		{
			name:        "text before code block",
			input:       "My example\n```js\nconsole.log('hello');\n```",
			wantCaption: "My example",
			wantCode:    "```js\nconsole.log('hello');\n```",
		},
		{
			name:        "multiple lines before code block",
			input:       "Line 1\nLine 2\n```js\ncode\n```",
			wantCaption: "Line 1\nLine 2",
			wantCode:    "```js\ncode\n```",
		},
		{
			name:        "indented code fence",
			input:       "Caption\n  ```js\ncode\n```",
			wantCaption: "Caption",
			wantCode:    "  ```js\ncode\n```",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			caption, code := extractImplicitCaption(tt.input)
			assert.Equal(t, tt.wantCaption, caption)
			assert.Equal(t, tt.wantCode, code)
		})
	}
}

func TestFormatFigure(t *testing.T) {
	tests := []struct {
		name    string
		caption string
		code    string
		want    string
	}{
		{
			name:    "basic figure",
			caption: "Example",
			code:    "console.log('hi');",
			want:    "<figure>\n<figcaption>Example</figcaption>\n\nconsole.log('hi');\n</figure>",
		},
		{
			name:    "multiline code",
			caption: "Usage",
			code:    "```js\nconst x = 1;\n```",
			want:    "<figure>\n<figcaption>Usage</figcaption>\n\n```js\nconst x = 1;\n```\n</figure>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, formatFigure(tt.caption, tt.code))
		})
	}
}

func TestAppendExample(t *testing.T) {
	tests := []struct {
		name        string
		description string
		example     string
		want        string
	}{
		{
			name:        "empty description",
			description: "",
			example:     "example code",
			want:        "example code",
		},
		{
			name:        "non-empty description",
			description: "Some desc",
			example:     "example code",
			want:        "Some desc\n\nexample code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, appendExample(tt.description, tt.example))
		})
	}
}

func TestHandleExampleTag(t *testing.T) {
	tests := []struct {
		name               string
		currentDescription string
		tagName            string
		content            string
		wantContains       string
	}{
		{
			name:               "plain example appended to empty description",
			currentDescription: "",
			tagName:            "@example",
			content:            "console.log('hello');",
			wantContains:       "console.log('hello');",
		},
		{
			name:               "example appended to existing description",
			currentDescription: "A thing.",
			tagName:            "@example",
			content:            "doThing();",
			wantContains:       "doThing();",
		},
		{
			name:               "example with caption tag",
			currentDescription: "",
			tagName:            "@example",
			content:            "<caption>Usage</caption>\ncode()",
			wantContains:       "<figcaption>Usage</figcaption>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleExampleTag(tt.currentDescription, tt.tagName, tt.content)
			assert.Contains(t, got, tt.wantContains)
		})
	}
}
