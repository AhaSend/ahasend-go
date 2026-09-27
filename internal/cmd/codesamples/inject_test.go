package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const fixtureSpec = `openapi: 3.1.0
paths:
  /v2/ping:
    get:
      summary: Ping
      operationId: ping
      x-code-samples:
        - lang: go
          label: Go net/http
          source: |
            package main

            func main() {}



        - lang: javascript
          label: Node.js 22+ (AhaSend SDK)
          source: |
            await client.ping();

      responses:
        '200':
          description: Pong
  # Accounts
  /v2/accounts/{account_id}:
    get:
      x-code-samples:
        - label: Bootstrap
          lang: shell
          source: |
            curl https://api.ahasend.com
        - lang: go
          label: AhaSend Go SDK
          source: |
            package main
      operationId: getAccount
      responses:
        '200':
          description: Account
    put:
      operationId: updateAccount
      x-code-samples:
        - lang: javascript
          label: Node.js 22+ (AhaSend SDK)
          source: |
            await client.accounts.update();
      responses:
        '200':
          description: Account
`

func fixtureSamples() map[string][]byte {
	return map[string][]byte{
		"ping":          []byte("package main\n\nfunc main() {\n\tif true {\n\t\tprintln(\"pong\")\n\t}\n}\n"),
		"getAccount":    []byte("package main\n\nfunc main() {}\n"),
		"updateAccount": []byte("package main\n\nfunc main() {\n\tprintln(\"updated\")\n}\n"),
	}
}

func TestInjectReplacesEachGoSample(t *testing.T) {
	updated, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)

	operations, err := parseOperations(updated)
	require.NoError(t, err)
	for operationID, source := range map[string]string{
		"ping":          "package main\n\nfunc main() {\n  if true {\n    println(\"pong\")\n  }\n}\n",
		"getAccount":    "package main\n\nfunc main() {}\n",
		"updateAccount": "package main\n\nfunc main() {\n  println(\"updated\")\n}\n",
	} {
		var goSamples []codeSample
		for _, sample := range operations[operationID].CodeSamples {
			if sample.Lang == "go" {
				goSamples = append(goSamples, sample)
			}
		}
		require.Len(t, goSamples, 1, operationID)
		assert.Equal(t, "AhaSend Go SDK", goSamples[0].Label, operationID)
		assert.Equal(t, source, goSamples[0].Source, operationID)
	}
}

func TestInjectLeavesOtherLinesAlone(t *testing.T) {
	updated, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)

	want := strings.Replace(fixtureSpec, `        - lang: go
          label: Go net/http
          source: |
            package main

            func main() {}



`, `        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {
              if true {
                println("pong")
              }
            }

`, 1)
	want = strings.Replace(want, `        - lang: go
          label: AhaSend Go SDK
          source: |
            package main
      operationId: getAccount
`, `        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {}

      operationId: getAccount
`, 1)
	want = strings.Replace(want, `      operationId: updateAccount
      x-code-samples:
`, `      operationId: updateAccount
      x-code-samples:
        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {
              println("updated")
            }

`, 1)
	assert.Equal(t, want, string(updated))
}

func TestInjectIsIdempotent(t *testing.T) {
	once, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)
	twice, err := inject(once, fixtureSamples())
	require.NoError(t, err)
	assert.Equal(t, string(once), string(twice))
}

func TestInjectRequiresASamplePerOperation(t *testing.T) {
	samples := fixtureSamples()
	delete(samples, "getAccount")

	_, err := inject([]byte(fixtureSpec), samples)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no sample for operations: getAccount")
}

func TestInjectRejectsSamplesForUnknownOperations(t *testing.T) {
	samples := fixtureSamples()
	samples["getAccounts"] = []byte("package main\n")

	_, err := inject([]byte(fixtureSpec), samples)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "samples name no operation: getAccounts")
}

func TestInjectRejectsDuplicateGoSamples(t *testing.T) {
	spec := strings.Replace(fixtureSpec, `        - lang: javascript
          label: Node.js 22+ (AhaSend SDK)
          source: |
            await client.ping();
`, `        - lang: go
          label: Another
          source: |
            package main
`, 1)

	_, err := inject([]byte(spec), fixtureSamples())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ping: has more than one go sample")
}

func TestInjectRequiresACodeSamplesBlock(t *testing.T) {
	spec := strings.Replace(fixtureSpec, `      operationId: updateAccount
      x-code-samples:
        - lang: javascript
          label: Node.js 22+ (AhaSend SDK)
          source: |
            await client.accounts.update();
`, `      operationId: updateAccount
`, 1)

	_, err := inject([]byte(spec), fixtureSamples())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "updateAccount: has no x-code-samples block")
}

func TestRunCheckReportsAStaleSpecification(t *testing.T) {
	dir := t.TempDir()
	specPath := filepath.Join(dir, "openapi.yaml")
	samplesDir := filepath.Join(dir, "codesamples")
	require.NoError(t, os.WriteFile(specPath, []byte(fixtureSpec), 0o644))
	for operationID, source := range fixtureSamples() {
		require.NoError(t, os.MkdirAll(filepath.Join(samplesDir, operationID), 0o755))
		require.NoError(t, os.WriteFile(filepath.Join(samplesDir, operationID, "main.go"), source, 0o644))
	}

	err := run(specPath, samplesDir, true)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "run `make code-samples`")
	unchanged, err := os.ReadFile(specPath)
	require.NoError(t, err)
	assert.Equal(t, fixtureSpec, string(unchanged), "check mode must not write")

	require.NoError(t, run(specPath, samplesDir, false))
	require.NoError(t, run(specPath, samplesDir, true))
}

func TestRunRequiresMainGoInEachSampleDirectory(t *testing.T) {
	dir := t.TempDir()
	specPath := filepath.Join(dir, "openapi.yaml")
	require.NoError(t, os.WriteFile(specPath, []byte(fixtureSpec), 0o644))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "codesamples", "ping"), 0o755))

	err := run(specPath, filepath.Join(dir, "codesamples"), false)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "main.go")
}
