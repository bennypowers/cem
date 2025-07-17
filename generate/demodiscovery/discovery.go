package demodiscovery

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	"github.com/gosimple/slug"
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

func resolveDemoSourceURL(cfg *C.CemConfig, demoPath string) (string, error) {
	base, err := url.Parse(cfg.SourceControlRootUrl)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(demoPath)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(rel).String(), nil
}

// DiscoverDemos attaches demos (indexed by tag name) to custom element declarations.
func DiscoverDemos(
	ctx W.WorkspaceContext,
	tagAliases map[string]string,
	module *M.Module,
	qm *Q.QueryManager,
	demoMap DemoMap,
) (errs error) {
	cfg, err := ctx.Config()
	if err != nil {
		return err
	}
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
		moduleDir := filepath.Dir(module.Path)
		sort.SliceStable(demoFiles, func(i, j int) bool {
			iDir := filepath.Dir(demoFiles[i])
			jDir := filepath.Dir(demoFiles[j])
			// module path is relative to the package, but demo path is relative to the project
			// so we need to account for that
			iMatch := strings.HasPrefix(iDir, filepath.Join(ctx.Root(), moduleDir))
			jMatch := strings.HasPrefix(jDir, filepath.Join(ctx.Root(), moduleDir))
			return iMatch && !jMatch
		})
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
					match := m[i]
					// WARN: this is a simple heuristic, but since we provide tremendous flexibility
					// by allowing users to define capture names i.e. template variables themselves,
					// we can't use well-known keys.
					// This might fail if for some reason the user wants to use the aliased tag name
					// in multiple captures, but only use the alias in one template variable
					if alias, ok := tagAliases[match]; ok {
						match = slug.Make(alias)
					}
					params[name] = match
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
			href, err := resolveDemoSourceURL(cfg, demoPath)
			if err != nil {
				errs = errors.Join(errs, err)
				continue
			}
			demo := M.Demo{
				URL:         demoUrl,
				Description: description,
				Source: &M.SourceReference{
					Href: href,
				},
			}
			ce.Demos = append(ce.Demos, demo)
		}
	}
	return errs
}
