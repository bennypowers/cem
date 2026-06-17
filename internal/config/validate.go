package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"text/template"
	"text/template/parse"
)

// ValidationSeverity indicates how serious a validation finding is.
type ValidationSeverity string

const (
	SeverityError   ValidationSeverity = "error"
	SeverityWarning ValidationSeverity = "warning"
)

func (s ValidationSeverity) MarshalJSON() ([]byte, error) {
	if s == "" {
		return []byte(`"error"`), nil
	}
	return []byte(`"` + string(s) + `"`), nil
}

// ValidationError describes a single config validation failure.
type ValidationError struct {
	Field    string             `json:"field"`
	Message  string             `json:"message"`
	Value    string             `json:"value,omitempty"`
	Severity ValidationSeverity `json:"severity"`
	Line     int                `json:"line,omitempty"`
	Column   int                `json:"column,omitempty"`
}

func (e ValidationError) Error() string {
	if e.Value != "" {
		return fmt.Sprintf("%s: %s (got %q)", e.Field, e.Message, e.Value)
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// IOChecker abstracts filesystem and future fetch-based I/O for validation.
type IOChecker interface {
	Exists(string) bool
	ReadFile(string) ([]byte, error)
}

// ValidateOptions controls which categories of validation run.
type ValidateOptions struct {
	CheckIO bool
	Root    string
	IO      IOChecker
	ValidateURLRewrites   func([]URLRewrite) error
	ValidateDemoDiscovery func(*CemConfig) error
}

const (
	RenderingLight      = "light"
	RenderingShadow     = "shadow"
	RenderingIframe     = "iframe"
	RenderingChromeless = "chromeless"
)

var validRenderingModes = []string{
	RenderingLight,
	RenderingShadow,
	RenderingIframe,
	RenderingChromeless,
}

func IsValidRenderingMode(mode string) bool {
	return slices.Contains(validRenderingModes, mode)
}

var validTargets = []string{
	"es2015", "es2016", "es2017", "es2018", "es2019",
	"es2020", "es2021", "es2022", "es2023", "esnext",
}

func IsValidTarget(target string) bool {
	return slices.Contains(validTargets, target)
}

// Validate checks a CemConfig for invalid values and returns all errors found.
func Validate(cfg *CemConfig, opts ValidateOptions) []ValidationError {
	if cfg == nil {
		return nil
	}

	var errs []ValidationError

	if r := cfg.Serve.Demos.Rendering; r != "" && !IsValidRenderingMode(r) {
		errs = append(errs, ValidationError{
			Field:   "serve.demos.rendering",
			Message: fmt.Sprintf("must be one of: %s", strings.Join(validRenderingModes, ", ")),
			Value:   r,
		})
	}

	if t := cfg.Serve.Transforms.TypeScript.Target; t != "" && !IsValidTarget(t) {
		errs = append(errs, ValidationError{
			Field:   "serve.transforms.typescript.target",
			Message: fmt.Sprintf("must be one of: %s", strings.Join(validTargets, ", ")),
			Value:   t,
		})
	}

	if p := cfg.Serve.Port; p < 0 || p > 65535 {
		errs = append(errs, ValidationError{
			Field:   "serve.port",
			Message: "must be between 0 and 65535",
			Value:   fmt.Sprintf("%d", p),
		})
	}

	if prefix := cfg.Generate.DesignTokens.Prefix; prefix != "" && strings.HasPrefix(prefix, "-") {
		errs = append(errs, ValidationError{
			Field:   "generate.designTokens.prefix",
			Message: fmt.Sprintf("should not start with dashes (use %q instead)", strings.TrimLeft(prefix, "-")),
			Value:   prefix,
		})
	}

	if opts.ValidateURLRewrites != nil && len(cfg.Serve.URLRewrites) > 0 {
		if err := opts.ValidateURLRewrites(cfg.Serve.URLRewrites); err != nil {
			errs = append(errs, ValidationError{
				Field:   "serve.urlRewrites",
				Message: err.Error(),
			})
		}
	}

	if opts.ValidateDemoDiscovery != nil {
		if err := opts.ValidateDemoDiscovery(cfg); err != nil {
			errs = append(errs, ValidationError{
				Field:   "generate.demoDiscovery",
				Message: err.Error(),
			})
		}
	}

	errs = append(errs, validateDemoDiscoveryParams(cfg)...)

	if opts.CheckIO && opts.IO != nil {
		errs = append(errs, validateIO(cfg, opts)...)
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}

func validateIO(cfg *CemConfig, opts ValidateOptions) []ValidationError {
	var errs []ValidationError

	if f := cfg.Serve.ImportMap.OverrideFile; f != "" {
		path := f
		if !filepath.IsAbs(path) && opts.Root != "" {
			path = filepath.Join(opts.Root, path)
		}
		if !opts.IO.Exists(path) {
			errs = append(errs, ValidationError{
				Field:   "serve.importMap.overrideFile",
				Message: "file not found",
				Value:   f,
			})
		}
	}

	if f := cfg.Generate.DesignTokens.Spec; f != "" {
		if u, err := url.Parse(f); err != nil || u.Scheme == "" {
			path := f
			if !filepath.IsAbs(path) && opts.Root != "" {
				path = filepath.Join(opts.Root, path)
			}
			if !opts.IO.Exists(path) {
				errs = append(errs, ValidationError{
					Field:   "generate.designTokens.spec",
					Message: "file not found",
					Value:   f,
				})
			}
		}
	}

	errs = append(errs, validateOutputMismatch(cfg, opts)...)

	return errs
}

// ValidationErrors converts a slice of ValidationError to a single error.
func ValidationErrors(errs []ValidationError) error {
	if len(errs) == 0 {
		return nil
	}
	msgs := make([]string, len(errs))
	for i, e := range errs {
		msgs[i] = e.Error()
	}
	return fmt.Errorf("config validation failed:\n  %s", strings.Join(msgs, "\n  "))
}

// ValidRenderingModes returns valid rendering mode values.
func ValidRenderingModes() []string {
	return slices.Clone(validRenderingModes)
}

// ValidTargets returns valid ES target values.
func ValidTargets() []string {
	return slices.Clone(validTargets)
}

func validateDemoDiscoveryParams(cfg *CemConfig) []ValidationError {
	var errs []ValidationError
	dd := cfg.Generate.DemoDiscovery

	hasPattern := dd.URLPattern != ""
	hasTemplate := dd.URLTemplate != ""

	if hasPattern && !hasTemplate {
		errs = append(errs, ValidationError{
			Field:    "generate.demoDiscovery",
			Message:  "urlPattern is set without urlTemplate",
			Severity: SeverityWarning,
		})
		return errs
	}
	if hasTemplate && !hasPattern {
		errs = append(errs, ValidationError{
			Field:    "generate.demoDiscovery",
			Message:  "urlTemplate is set without urlPattern",
			Severity: SeverityWarning,
		})
		return errs
	}
	if !hasPattern || !hasTemplate {
		return nil
	}

	patternParams := extractURLPatternParams(dd.URLPattern)
	templateParams := extractTemplateParams(dd.URLTemplate)

	for _, p := range templateParams {
		if !slices.Contains(patternParams, p) {
			errs = append(errs, ValidationError{
				Field:    "generate.demoDiscovery.urlTemplate",
				Message:  fmt.Sprintf("references param %q not defined in urlPattern", p),
				Value:    dd.URLTemplate,
				Severity: SeverityError,
			})
		}
	}

	for _, p := range patternParams {
		if !slices.Contains(templateParams, p) {
			errs = append(errs, ValidationError{
				Field:    "generate.demoDiscovery.urlPattern",
				Message:  fmt.Sprintf("defines param %q not referenced in urlTemplate", p),
				Value:    dd.URLPattern,
				Severity: SeverityWarning,
			})
		}
	}

	return errs
}

func extractURLPatternParams(pattern string) []string {
	var params []string
	for part := range strings.SplitSeq(pattern, "/") {
		if name, ok := strings.CutPrefix(part, ":"); ok {
			name = strings.TrimRight(name, "?*+")
			if idx := strings.Index(name, "."); idx >= 0 {
				name = name[:idx]
			}
			if name != "" {
				params = append(params, name)
			}
		}
	}
	sort.Strings(params)
	return params
}

// TemplateFuncNames lists the template functions allowed in urlTemplate fields.
// demodiscovery registers real implementations; validation uses no-op stubs.
var TemplateFuncNames = []string{"alias", "slug", "lower", "upper"}

var templateStubFuncs = makeStubFuncMap(TemplateFuncNames)

func makeStubFuncMap(names []string) template.FuncMap {
	m := make(template.FuncMap, len(names))
	for _, name := range names {
		m[name] = func(s string) string { return s }
	}
	return m
}

func extractTemplateParams(tmplStr string) []string {
	t, err := template.New("").Funcs(templateStubFuncs).Parse(tmplStr)
	if err != nil {
		return nil
	}
	var params []string
	for _, node := range t.Root.Nodes {
		collectFieldNames(node, &params)
	}
	sort.Strings(params)
	return slices.Compact(params)
}

func collectFieldNames(node parse.Node, params *[]string) {
	switch n := node.(type) {
	case *parse.ActionNode:
		if n.Pipe != nil {
			for _, cmd := range n.Pipe.Cmds {
				for _, arg := range cmd.Args {
					collectFieldNames(arg, params)
				}
			}
		}
	case *parse.FieldNode:
		if len(n.Ident) > 0 {
			*params = append(*params, n.Ident[0])
		}
	case *parse.PipeNode:
		for _, cmd := range n.Cmds {
			for _, arg := range cmd.Args {
				collectFieldNames(arg, params)
			}
		}
	case *parse.ListNode:
		for _, child := range n.Nodes {
			collectFieldNames(child, params)
		}
	case *parse.IfNode:
		collectBranchFields(&n.BranchNode, params)
	case *parse.RangeNode:
		collectBranchFields(&n.BranchNode, params)
	case *parse.WithNode:
		collectBranchFields(&n.BranchNode, params)
	}
}

func collectBranchFields(n *parse.BranchNode, params *[]string) {
	if n.Pipe != nil {
		collectFieldNames(n.Pipe, params)
	}
	if n.List != nil {
		collectFieldNames(n.List, params)
	}
	if n.ElseList != nil {
		collectFieldNames(n.ElseList, params)
	}
}

func validateOutputMismatch(cfg *CemConfig, opts ValidateOptions) []ValidationError {
	output := cfg.Generate.Output
	if output == "" {
		return nil
	}

	pkgJSONPath := "package.json"
	if opts.Root != "" {
		pkgJSONPath = filepath.Join(opts.Root, pkgJSONPath)
	}

	if !opts.IO.Exists(pkgJSONPath) {
		return nil
	}

	data, err := opts.IO.ReadFile(pkgJSONPath)
	if err != nil {
		return nil
	}

	var pkg struct {
		CustomElements string `json:"customElements"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil
	}

	if pkg.CustomElements == "" {
		return nil
	}

	relOutput := output
	if opts.Root != "" {
		if rel, err := filepath.Rel(opts.Root, output); err == nil {
			relOutput = rel
		}
	}
	normalizedOutput := filepath.Clean(relOutput)
	normalizedPkg := filepath.Clean(pkg.CustomElements)
	if normalizedOutput != normalizedPkg {
		return []ValidationError{{
			Field:    "generate.output",
			Message:  fmt.Sprintf("does not match package.json customElements field (%q vs %q)", relOutput, pkg.CustomElements),
			Severity: SeverityWarning,
		}}
	}

	return nil
}

var _ error = ValidationError{}
