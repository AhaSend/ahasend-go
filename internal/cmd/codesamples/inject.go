package main

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	sampleLang  = "go"
	sampleLabel = "AhaSend Go SDK"
)

var (
	httpMethods = map[string]bool{
		"get": true, "put": true, "post": true, "delete": true,
		"options": true, "head": true, "patch": true, "trace": true,
	}

	operationIDLine = regexp.MustCompile(`^ {6}operationId:\s*(.+?)\s*$`)
	codeSamplesLine = regexp.MustCompile(`^ {6}x-code-samples:\s*$`)
	sampleItemLine  = regexp.MustCompile(`^ {8}- `)
	sampleLangLine  = regexp.MustCompile(`^ {8}(?:- | {2})lang:\s*(.+?)\s*$`)
)

type codeSample struct {
	Lang   string `yaml:"lang"`
	Label  string `yaml:"label"`
	Source string `yaml:"source"`
}

type operation struct {
	OperationID string       `yaml:"operationId"`
	CodeSamples []codeSample `yaml:"x-code-samples"`
}

// inject returns spec with each operation's go sample replaced by the program
// samples holds for its operation ID. The specification is edited as text, so
// only the lines of the go sample entries change; the result is then parsed
// back to prove that every operation carries exactly its own sample and that
// nothing else in the document moved.
func inject(spec []byte, samples map[string][]byte) ([]byte, error) {
	operations, err := parseOperations(spec)
	if err != nil {
		return nil, err
	}
	if err := matchSamples(operations, samples); err != nil {
		return nil, err
	}

	lines := strings.Split(string(spec), "\n")
	edits := make([]edit, 0, len(operations))
	for operationID := range operations {
		e, err := locateGoSample(lines, operationID, sampleLines(samples[operationID]))
		if err != nil {
			return nil, err
		}
		edits = append(edits, e)
	}
	// Apply from the bottom up so earlier line indexes stay valid.
	sort.Slice(edits, func(i, j int) bool { return edits[i].start > edits[j].start })
	for _, e := range edits {
		tail := append([]string{}, lines[e.end:]...)
		lines = append(append(lines[:e.start], e.replacement...), tail...)
	}
	updated := []byte(strings.Join(lines, "\n"))

	if err := verify(spec, updated, samples); err != nil {
		return nil, err
	}
	return updated, nil
}

// parseOperations returns every operation under paths keyed by operation ID.
func parseOperations(spec []byte) (map[string]operation, error) {
	var doc struct {
		Paths map[string]map[string]yaml.Node `yaml:"paths"`
	}
	if err := yaml.Unmarshal(spec, &doc); err != nil {
		return nil, err
	}
	operations := make(map[string]operation)
	for path, item := range doc.Paths {
		for method, node := range item {
			if !httpMethods[method] {
				continue
			}
			var op operation
			if err := node.Decode(&op); err != nil {
				return nil, fmt.Errorf("%s %s: %w", method, path, err)
			}
			if op.OperationID == "" {
				return nil, fmt.Errorf("%s %s has no operationId", method, path)
			}
			if _, dup := operations[op.OperationID]; dup {
				return nil, fmt.Errorf("operationId %s is used more than once", op.OperationID)
			}
			operations[op.OperationID] = op
		}
	}
	return operations, nil
}

// matchSamples requires exactly one sample per operation and no sample for an
// operation the specification does not have.
func matchSamples(operations map[string]operation, samples map[string][]byte) error {
	missing := map[string]bool{}
	for operationID := range operations {
		if _, ok := samples[operationID]; !ok {
			missing[operationID] = true
		}
	}
	orphaned := map[string]bool{}
	for operationID := range samples {
		if _, ok := operations[operationID]; !ok {
			orphaned[operationID] = true
		}
	}

	var problems []string
	if len(missing) > 0 {
		problems = append(problems, "no sample for operations: "+strings.Join(sortedKeys(missing), ", "))
	}
	if len(orphaned) > 0 {
		problems = append(problems, "samples name no operation: "+strings.Join(sortedKeys(orphaned), ", "))
	}
	if len(problems) > 0 {
		return fmt.Errorf("%s", strings.Join(problems, "; "))
	}
	return nil
}

// edit replaces lines[start:end] with replacement.
type edit struct {
	start, end  int
	replacement []string
}

// locateGoSample returns the edit that gives an operation entry as its one go
// sample:
//
//   - a go entry, with the blank lines that follow it, is replaced in place.
//   - an x-code-samples block without a go entry gets one at its end, just
//     before the next line indented six spaces or less.
//   - an operation without x-code-samples gets the block, holding only the go
//     entry, at its end, just before the next line indented four spaces or
//     less (or at the end of the file).
func locateGoSample(lines []string, operationID string, entry []string) (edit, error) {
	idLine := -1
	for i, line := range lines {
		if m := operationIDLine.FindStringSubmatch(line); m != nil && unquote(m[1]) == operationID {
			if idLine >= 0 {
				return edit{}, fmt.Errorf("%s: operationId appears on more than one line", operationID)
			}
			idLine = i
		}
	}
	if idLine < 0 {
		return edit{}, fmt.Errorf("%s: cannot find its operationId line", operationID)
	}

	// An operation's keys sit at six spaces; its method line and everything
	// after the operation sit at four or fewer.
	opStart, opEnd := idLine, len(lines)
	for opStart > 0 && !endsOperation(lines[opStart]) {
		opStart--
	}
	for i := idLine + 1; i < len(lines); i++ {
		if endsOperation(lines[i]) {
			opEnd = i
			break
		}
	}

	samplesLine := -1
	for i := opStart + 1; i < opEnd; i++ {
		if codeSamplesLine.MatchString(lines[i]) {
			samplesLine = i
			break
		}
	}
	if samplesLine < 0 {
		block := append([]string{"      x-code-samples:"}, entry...)
		return edit{start: opEnd, end: opEnd, replacement: block}, nil
	}
	samplesEnd := opEnd
	for i := samplesLine + 1; i < opEnd; i++ {
		if strings.TrimSpace(lines[i]) != "" && indent(lines[i]) <= 6 {
			samplesEnd = i
			break
		}
	}

	var items []int
	for i := samplesLine + 1; i < samplesEnd; i++ {
		if sampleItemLine.MatchString(lines[i]) {
			items = append(items, i)
		}
	}
	goItem := -1
	for n, start := range items {
		end := samplesEnd
		if n+1 < len(items) {
			end = items[n+1]
		}
		if !isSampleLang(itemLang(lines[start:end])) {
			continue
		}
		if goItem >= 0 {
			return edit{}, fmt.Errorf("%s: has more than one go sample", operationID)
		}
		goItem = n
	}

	if goItem < 0 {
		return edit{start: samplesEnd, end: samplesEnd, replacement: entry}, nil
	}
	end := samplesEnd
	if goItem+1 < len(items) {
		end = items[goItem+1]
	}
	return edit{start: items[goItem], end: end, replacement: entry}, nil
}

// isSampleLang reports whether an x-code-samples lang names Go, in any case.
func isSampleLang(lang string) bool {
	return strings.EqualFold(lang, sampleLang)
}

func endsOperation(line string) bool {
	return strings.TrimSpace(line) != "" && indent(line) <= 4
}

// itemLang returns the lang of one x-code-samples entry, whether or not lang
// is its first key.
func itemLang(item []string) string {
	for _, line := range item {
		if m := sampleLangLine.FindStringSubmatch(line); m != nil {
			return unquote(m[1])
		}
	}
	return ""
}

// sampleLines renders a sample program as an x-code-samples entry whose
// source is a literal block. Leading tabs become two spaces each, the
// indentation the specification's samples use; a blank line separates the
// entry from whatever follows it.
func sampleLines(source []byte) []string {
	text := strings.TrimRight(strings.ReplaceAll(string(source), "\r\n", "\n"), "\n")
	out := []string{
		"        - lang: " + sampleLang,
		"          label: " + sampleLabel,
		"          source: |",
	}
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == "" {
			out = append(out, "")
			continue
		}
		body := strings.TrimLeft(line, "\t")
		out = append(out, "            "+strings.Repeat("  ", len(line)-len(body))+body)
	}
	return append(out, "")
}

// renderedSource is the source a sample carries once parsed back out of the
// specification.
func renderedSource(source []byte) string {
	lines := sampleLines(source)
	body := lines[3 : len(lines)-1]
	for i, line := range body {
		body[i] = strings.TrimPrefix(line, "            ")
	}
	return strings.Join(body, "\n") + "\n"
}

// verify parses the edited specification back and fails unless every
// operation carries exactly one go sample, that sample is the operation's
// program, and the document minus its go samples is unchanged.
func verify(before, after []byte, samples map[string][]byte) error {
	operations, err := parseOperations(after)
	if err != nil {
		return fmt.Errorf("edited specification does not parse: %w", err)
	}
	for operationID, op := range operations {
		var found []codeSample
		for _, sample := range op.CodeSamples {
			if isSampleLang(sample.Lang) {
				found = append(found, sample)
			}
		}
		if len(found) != 1 {
			return fmt.Errorf("%s: has %d go samples after the edit", operationID, len(found))
		}
		if found[0].Label != sampleLabel || found[0].Source != renderedSource(samples[operationID]) {
			return fmt.Errorf("%s: go sample does not read back as its program", operationID)
		}
	}

	var original, edited any
	if err := yaml.Unmarshal(before, &original); err != nil {
		return err
	}
	if err := yaml.Unmarshal(after, &edited); err != nil {
		return err
	}
	if !reflect.DeepEqual(withoutGoSamples(original), withoutGoSamples(edited)) {
		return fmt.Errorf("the edit changed more than the go samples")
	}
	return nil
}

// withoutGoSamples drops every go entry from each operation's x-code-samples.
func withoutGoSamples(doc any) any {
	root, _ := doc.(map[string]any)
	paths, _ := root["paths"].(map[string]any)
	for _, item := range paths {
		methods, _ := item.(map[string]any)
		for method, op := range methods {
			fields, ok := op.(map[string]any)
			if !httpMethods[method] || !ok {
				continue
			}
			entries, _ := fields["x-code-samples"].([]any)
			kept := make([]any, 0, len(entries))
			for _, entry := range entries {
				if sample, _ := entry.(map[string]any); isSampleLang(fmt.Sprint(sample["lang"])) {
					continue
				}
				kept = append(kept, entry)
			}
			fields["x-code-samples"] = kept
		}
	}
	return doc
}

func indent(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}

func unquote(value string) string {
	if len(value) >= 2 && (value[0] == '"' || value[0] == '\'') && value[len(value)-1] == value[0] {
		return value[1 : len(value)-1]
	}
	return value
}
