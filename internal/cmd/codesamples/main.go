// Command codesamples writes the SDK's Go code samples into the OpenAPI
// specification.
//
// Every operation in the specification has one sample program at
// codesamples/<operationId>/main.go. The command replaces each operation's
// single `go` entry under x-code-samples with that program, labelled
// "AhaSend Go SDK", and leaves every other byte of the file alone: the server
// owns the rest of the specification, and a later sync pulls these samples
// back into it.
//
// Usage:
//
//	go run ./internal/cmd/codesamples          # rewrite openapi/openapi.yaml
//	go run ./internal/cmd/codesamples -check   # fail if it is out of date
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	specPath := flag.String("spec", "openapi/openapi.yaml", "OpenAPI specification to update")
	samplesDir := flag.String("samples", "codesamples", "directory holding one <operationId>/main.go per operation")
	check := flag.Bool("check", false, "report whether the specification is up to date instead of writing it")
	flag.Parse()

	if err := run(*specPath, *samplesDir, *check); err != nil {
		fmt.Fprintln(os.Stderr, "codesamples:", err)
		os.Exit(1)
	}
}

func run(specPath, samplesDir string, check bool) error {
	spec, err := os.ReadFile(specPath)
	if err != nil {
		return err
	}
	samples, err := readSamples(samplesDir)
	if err != nil {
		return err
	}
	updated, err := inject(spec, samples)
	if err != nil {
		return fmt.Errorf("%s: %w", specPath, err)
	}

	if check {
		if !bytes.Equal(spec, updated) {
			return fmt.Errorf("%s does not carry the current Go samples; run `make code-samples`", specPath)
		}
		return nil
	}
	if bytes.Equal(spec, updated) {
		return nil
	}
	return os.WriteFile(specPath, updated, 0o644)
}

// readSamples returns each sample program's source keyed by the operation ID
// its directory names.
func readSamples(dir string) (map[string][]byte, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	samples := make(map[string][]byte, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			return nil, fmt.Errorf("%s: expected only operation directories, found file %s", dir, entry.Name())
		}
		source, err := os.ReadFile(filepath.Join(dir, entry.Name(), "main.go"))
		if err != nil {
			return nil, err
		}
		samples[entry.Name()] = source
	}
	return samples, nil
}

func sortedKeys(set map[string]bool) []string {
	keys := make([]string, 0, len(set))
	for key := range set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
