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

package inject

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

// InjectScript injects a script into HTML by parsing and manipulating the DOM.
// It attempts to inject before </head>, or at the start of <body> if no head exists.
// Falls back to string replacement if DOM parsing fails.
func InjectScript(htmlStr, script string) string {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return fallbackInject(htmlStr, script)
	}

	// Parse the script string into nodes
	scriptNodes, err := html.ParseFragment(strings.NewReader(script), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: 0,
	})
	if err != nil || len(scriptNodes) == 0 {
		return fallbackInject(htmlStr, script)
	}

	// Find the <head> element and inject before its closing tag
	head := findElement(doc, "head")
	if head != nil {
		// Append script nodes to head
		for _, scriptNode := range scriptNodes {
			head.AppendChild(scriptNode)
		}
	} else {
		// No head found, try to find body and prepend
		body := findElement(doc, "body")
		if body != nil {
			// Prepend script nodes to body
			for i := len(scriptNodes) - 1; i >= 0; i-- {
				if body.FirstChild != nil {
					body.InsertBefore(scriptNodes[i], body.FirstChild)
				} else {
					body.AppendChild(scriptNodes[i])
				}
			}
		}
	}

	// Render the modified DOM back to HTML
	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		return fallbackInject(htmlStr, script)
	}

	return buf.String()
}

// findElement recursively searches for an element with the given tag name
func findElement(n *html.Node, tag string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findElement(c, tag); result != nil {
			return result
		}
	}
	return nil
}

// fallbackInject uses string replacement as a fallback when DOM parsing fails
func fallbackInject(htmlStr, script string) string {
	if strings.Contains(htmlStr, "</head>") {
		return strings.Replace(htmlStr, "</head>", script+"\n</head>", 1)
	}
	return htmlStr + "\n" + script
}
