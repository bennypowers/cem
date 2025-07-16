package demodiscovery

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
)

func extractDemoDescription(path string) (string, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read demo file: %w", err)
	}
	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()
	cursor := tree.Walk()
	for _, node := range root.NamedChildren(cursor) {
		if node.GrammarName() == "comment" {
			commentText := node.Utf8Text(code)
			// Accept only leading block comments (<!--- ... -->)
			// FIXME: quick hack to avoid consuming magic tag names comments.
			if strings.HasPrefix(commentText, "<!--") && !strings.Contains(commentText, "@tag") {
				desc := strings.TrimPrefix(commentText, "<!--")
				desc = strings.TrimSuffix(desc, "-->")
				return strings.TrimSpace(desc), nil
			}
		}
	}
	return "", nil
}

// DiscoverDemos attaches demos (indexed by tag name) to custom element declarations.
func DiscoverDemos(
	cfg *C.CemConfig,
	tagAliases map[string]string,
	module *M.Module,
	qm *Q.QueryManager,
	demoMap DemoMap,
) (errs error) {
	urlPattern := cfg.Generate.DemoDiscovery.URLPattern
	if urlPattern == "" {
		return nil
	}
	rx, err := regexp.Compile(urlPattern)
	if err != nil {
		return err
	}
	for _, decl := range module.Declarations {
		// Only attach to custom element classes with a tag name
		ce, ok := decl.(*M.CustomElementDeclaration)
		tagName := ce.CustomElement.TagName
		if !ok || ce.Kind != "class" || ce.TagName == "" {
			continue
		}
		demoFiles := demoMap[tagName]
		for _, demoPath := range demoFiles {
			m := rx.FindStringSubmatch(demoPath)
			if m == nil {
				pterm.Warning.Printfln("Demo path %q did not match URLPattern %q for tagName %q", demoPath, urlPattern, tagName)
				continue
			}
			groupNames := rx.SubexpNames()
			params := map[string]string{}
			for i, name := range groupNames {
				if i > 0 && name != "" {
					slug := m[i]
					// WARN: this is a simple heuristic, but since we provide tremendous flexibility
					// by allowing users to define capture names i.e. template variables themselves,
					// we can't use well-known keys.
					// This might fail if for some reason the user wants to use the aliased tag name
					// in multiple captures, but only use the alias in one template variable
					if alias, ok := tagAliases[slug]; ok {
						slug = alias
					}
					params[name] = slug
				}
			}
			demoUrl := strings.Clone(cfg.Generate.DemoDiscovery.URLTemplate)
			for k, v := range params {
				demoUrl = strings.ReplaceAll(demoUrl, "{"+k+"}", v)
			}
			description, err := extractDemoDescription(demoPath)
			if err != nil {
				errs = errors.Join(errs, err)
			}
			base, _ := url.Parse(cfg.SourceControlRootUrl)
			rel, _ := url.Parse(demoPath)
			demo := M.Demo{
				URL:         demoUrl,
				Description: description,
				Source: &M.SourceReference{
					Href: base.ResolveReference(rel).String(),
				},
			}
			ce.Demos = append(ce.Demos, demo)
		}
	}
	return errs
}
