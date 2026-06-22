package tui

import (
	"io"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/term"
)

// Hex approximations of the app's ANSI palette.
var chromaStyle = chroma.MustNewStyle("cem", chroma.StyleEntries{
	chroma.NameTag:         "#00ffff", // Cyan (InfoStyle/CodeStyle)
	chroma.Literal:         "#5f87ff", // BrightBlue (KindStyle)
	chroma.LiteralString:   "#5f87ff",
	chroma.Keyword:         "#00ff00", // Green (SuccessStyle)
	chroma.KeywordConstant: "#00ff00",
	chroma.LiteralNumber:   "#5f87ff",
	chroma.Punctuation:     "#808080", // BrightBlack (MutedStyle)
	chroma.Comment:         "#808080 italic",
	chroma.Error:           "#ff0000", // Red (ErrorStyle)
})

func chromaFormatter(w io.Writer) chroma.Formatter {
	if f, ok := w.(term.File); ok {
		p := colorprofile.Detect(f, os.Environ())
		switch p {
		case colorprofile.TrueColor:
			return formatters.Get("terminal16m")
		case colorprofile.ANSI256:
			return formatters.Get("terminal256")
		case colorprofile.ANSI:
			return formatters.Get("terminal")
		}
	}
	return formatters.NoOp
}

// Highlight writes syntax-highlighted source to w.
// lang is a chroma lexer name (e.g. "yaml", "json").
// Falls back to plain text when the output is not a terminal.
func Highlight(w io.Writer, source, lang string) error {
	lexer := lexers.Get(lang)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	iterator, err := lexer.Tokenise(nil, source)
	if err != nil {
		_, err = io.WriteString(w, source)
		return err
	}

	return chromaFormatter(w).Format(w, chromaStyle, iterator)
}

