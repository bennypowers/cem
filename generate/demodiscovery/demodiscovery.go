package demodiscovery

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func extractDemoDescription(path string) (string, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(Q.Languages.Html)
	tree := parser.Parse(code, nil)
	root := tree.RootNode()
	cursor := tree.Walk()
	for _, node := range root.NamedChildren(cursor) {
		if node.GrammarName() == "comment" {
			commentText := node.Utf8Text(code)
			// Accept only leading block comments (<!--- ... -->)
			if strings.HasPrefix(commentText, "<!--") {
				desc := strings.TrimPrefix(commentText, "<!--")
				desc = strings.TrimSuffix(desc, "-->")
				return strings.TrimSpace(desc), nil
			}
		}
	}
	return "", nil
}

// demoPathMatchesModulePath returns true if the given demoPath "belongs to" the given module.
// This is typically true if the demo file is in the same directory tree as the module's source file,
// or if the demoPath shares a directory prefix with the module.Path.
//
// The match is case-insensitive and is normalized for OS-specific path separators.
func demoPathMatchesModulePath(demoPath string, module *M.Module) bool {
	if module == nil || module.Path == "" {
		return false
	}

	// Normalize both paths to absolute and clean them
	moduleDir := filepath.Dir(filepath.Clean(module.Path))
	demoPathClean := filepath.Clean(demoPath)

	// Check if the demo path is under the module's directory
	rel, err := filepath.Rel(moduleDir, demoPathClean)
	if err != nil {
		return false
	}

	// If rel does not start with "..", then demoPath is in or under moduleDir
	return !strings.HasPrefix(rel, "..")
}

func DiscoverDemos(
	cfg *C.CemConfig,
	tagAliases map[string]string,
	module *M.Module,
	qm *Q.QueryManager,
	demoFiles []string,
) (errs error) {
	for _, demoPath := range demoFiles {
		if demoPathMatchesModulePath(demoPath, module) {
			for _, decl := range module.Declarations {
				// Only attach to custom element classes with a tag name
				ce, ok := decl.(*M.CustomElementDeclaration)
				if ok && ce.Kind == "class" && ce.TagName != "" {
					rx := regexp.MustCompile(cfg.Generate.DemoDiscovery.URLPattern)
					m := rx.FindStringSubmatch(demoPath)
					if m == nil {
						continue
					}
					groupNames := rx.SubexpNames()
					params := map[string]string{}
					for i, name := range groupNames {
						if i > 0 && name != "" {
							params[name] = m[i]
						}
					}
					// Build the demo URL
					url := cfg.Generate.DemoDiscovery.URLTemplate
					for k, v := range params {
						if k == "slug" {
							if alias, ok := tagAliases[ce.TagName]; ok {
								v = alias
							}
						}
						url = strings.ReplaceAll(url, "{"+k+"}", v)
					}
					description, err := extractDemoDescription(demoPath)
					if err != nil {
						errs = errors.Join(errs, err)
					}
					// Attach demo to declaration
					ce.Demos = append(ce.Demos, M.Demo{
						URL: url,
						Description: description,
						Source: &M.SourceReference{
							Href: path.Join(cfg.Generate.DemoDiscovery.SourceControlUrl, demoPath),
						},
					})
				}
			}
		}
	}
	return errs
}
