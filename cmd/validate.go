package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"bennypowers.dev/cem/workspace"
	"github.com/adrg/xdg"
	"github.com/pterm/pterm"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"

	"golang.org/x/mod/semver"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a custom-elements.json manifest",
	Long:  `Validate a custom-elements.json manifest against its JSON schema.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, err := workspace.GetWorkspaceContext(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting workspace context: %v\n", err)
			os.Exit(1)
		}

		manifestPath := ctx.CustomElementsManifestPath()
		if len(args) > 0 {
			manifestPath = args[0]
		}

		if manifestPath == "" {
			fmt.Fprintln(os.Stderr, "Could not find custom-elements.json")
			os.Exit(1)
		}

		manifestData, err := os.ReadFile(manifestPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading manifest file: %v\n", err)
			os.Exit(1)
		}

		var manifest struct {
			SchemaVersion string `json:"schemaVersion"`
		}
		if err := json.Unmarshal(manifestData, &manifest); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing manifest to find schemaVersion: %v\n", err)
			os.Exit(1)
		}

		if manifest.SchemaVersion == "" {
			fmt.Fprintln(os.Stderr, "Error: schemaVersion not found in manifest.")
			os.Exit(1)
		}

		if semver.Compare(
			semver.Canonical("v"+manifest.SchemaVersion),
			semver.Canonical("v2.1.1"),
		) < 0 {
			pterm.Warning.Printfln("validation for manifests with schemaVersion <= 2.1.0 may not produce accurate results (version: %s)", manifest.SchemaVersion)
		}

		schemaData, err := getSchema(manifest.SchemaVersion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting schema: %v\n", err)
			os.Exit(1)
		}

		compiler := jsonschema.NewCompiler()
		if err := compiler.AddResource("schema.json", bytes.NewReader(schemaData)); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding schema resource: %v\n", err)
			os.Exit(1)
		}
		schema, err := compiler.Compile("schema.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error compiling schema: %v\n", err)
			os.Exit(1)
		}

		var v any
		if err := json.Unmarshal(manifestData, &v); err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshaling manifest for validation: %v\n", err)
			os.Exit(1)
		}

		if err := schema.Validate(v); err != nil {
			validationError := err.(*jsonschema.ValidationError)
			fmt.Fprintf(os.Stderr, "Manifest is invalid (schema version %s):\n", manifest.SchemaVersion)
			printAllErrors(os.Stderr, validationError)
			os.Exit(1)
		}

		fmt.Println("Manifest is valid.")
	},
}

func printAllErrors(w io.Writer, err *jsonschema.ValidationError) {
	for _, cause := range err.Causes {
		if len(cause.Causes) == 0 {
			fmt.Fprintf(w, "- %s: %s\n", cause.InstanceLocation, cause.Message)
		} else {
			printAllErrors(w, cause)
		}
	}
}

func getSchema(version string) ([]byte, error) {
	cacheDir, err := xdg.CacheFile(filepath.Join("cem", "schemas"))
	if err != nil {
		return nil, fmt.Errorf("could not get cache directory: %w", err)
	}

	schemaPath := filepath.Join(cacheDir, version+".json")

	if _, err := os.Stat(schemaPath); err == nil {
		return os.ReadFile(schemaPath)
	}

	url := fmt.Sprintf("https://unpkg.com/custom-elements-manifest@%s/schema.json", version)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch schema from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status fetching schema from %s: %s", url, resp.Status)
	}

	schemaData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read schema from response body: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(schemaPath), 0755); err != nil {
		return nil, fmt.Errorf("could not create cache directory: %w", err)
	}

	if err := os.WriteFile(schemaPath, schemaData, 0644); err != nil {
		return nil, fmt.Errorf("could not write schema to cache: %w", err)
	}

	return schemaData, nil
}
