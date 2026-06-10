package config

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

// ValidationError describes a single config validation failure.
type ValidationError struct {
	Field   string
	Message string
	Value   string
}

func (e ValidationError) Error() string {
	if e.Value != "" {
		return fmt.Sprintf("%s: %s (got %q)", e.Field, e.Message, e.Value)
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateOptions controls which categories of validation run.
type ValidateOptions struct {
	CheckFilesystem bool
	Root            string
	FS              interface {
		Exists(string) bool
	}
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
	return slices.Contains(validTargets, strings.ToLower(target))
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

	if opts.CheckFilesystem && opts.FS != nil {
		errs = append(errs, validateFilesystem(cfg, opts)...)
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}

func validateFilesystem(cfg *CemConfig, opts ValidateOptions) []ValidationError {
	var errs []ValidationError

	if f := cfg.Serve.ImportMap.OverrideFile; f != "" {
		path := f
		if !filepath.IsAbs(path) && opts.Root != "" {
			path = filepath.Join(opts.Root, path)
		}
		if !opts.FS.Exists(path) {
			errs = append(errs, ValidationError{
				Field:   "serve.importMap.overrideFile",
				Message: "file not found",
				Value:   f,
			})
		}
	}

	if f := cfg.Generate.DesignTokens.Spec; f != "" {
		if !strings.HasPrefix(f, "npm:") && !strings.HasPrefix(f, "jsr:") && !strings.HasPrefix(f, "http://") && !strings.HasPrefix(f, "https://") {
			path := f
			if !filepath.IsAbs(path) && opts.Root != "" {
				path = filepath.Join(opts.Root, path)
			}
			if !opts.FS.Exists(path) {
				errs = append(errs, ValidationError{
					Field:   "generate.designTokens.spec",
					Message: "file not found",
					Value:   f,
				})
			}
		}
	}

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
	return validRenderingModes
}

// ValidTargets returns valid ES target values.
func ValidTargets() []string {
	return validTargets
}

var _ error = ValidationError{}
