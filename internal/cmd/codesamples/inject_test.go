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
    delete:
      operationId: deleteAccount
      responses:
        '200':
          description: Deleted


  /v2/accounts/{account_id}/members:
    get:
      operationId: getAccountMembers
      x-code-samples:
        - lang: Go
          label: Go net/http
          source: |
            package main
      responses:
        '200':
          description: Members
    delete:
      operationId: removeAccountMember
      responses:
        '200':
          description: Removed
`

func fixtureSamples() map[string][]byte {
	return map[string][]byte{
		"ping":                []byte("package main\n\nfunc main() {\n\tif true {\n\t\tprintln(\"pong\")\n\t}\n}\n"),
		"getAccount":          []byte("package main\n\nfunc main() {}\n"),
		"updateAccount":       []byte("package main\n\nfunc main() {\n\tprintln(\"updated\")\n}\n"),
		"deleteAccount":       []byte("package main\n\nfunc main() {\n\tprintln(\"deleted\")\n}\n"),
		"getAccountMembers":   []byte("package main\n\nfunc main() {\n\tprintln(\"members\")\n}\n"),
		"removeAccountMember": []byte("package main\n\nfunc main() {}\n"),
	}
}

// fixtureWant is fixtureSpec after inject. A replaced go entry takes the
// blank lines that followed it; a new entry goes at the end of its block, and
// a new block at the end of its operation, after any blank lines there.
const fixtureWant = `openapi: 3.1.0
paths:
  /v2/ping:
    get:
      summary: Ping
      operationId: ping
      x-code-samples:
        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {
              if true {
                println("pong")
              }
            }

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

            func main() {}

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

        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {
              println("updated")
            }

      responses:
        '200':
          description: Account
    delete:
      operationId: deleteAccount
      responses:
        '200':
          description: Deleted


      x-code-samples:
        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {
              println("deleted")
            }

  /v2/accounts/{account_id}/members:
    get:
      operationId: getAccountMembers
      x-code-samples:
        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {
              println("members")
            }

      responses:
        '200':
          description: Members
    delete:
      operationId: removeAccountMember
      responses:
        '200':
          description: Removed

      x-code-samples:
        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {}
`

func TestInjectReplacesEachGoSample(t *testing.T) {
	updated, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)

	operations, err := parseOperations(updated)
	require.NoError(t, err)
	for operationID, source := range map[string]string{
		"ping":                "package main\n\nfunc main() {\n  if true {\n    println(\"pong\")\n  }\n}\n",
		"getAccount":          "package main\n\nfunc main() {}\n",
		"updateAccount":       "package main\n\nfunc main() {\n  println(\"updated\")\n}\n",
		"deleteAccount":       "package main\n\nfunc main() {\n  println(\"deleted\")\n}\n",
		"getAccountMembers":   "package main\n\nfunc main() {\n  println(\"members\")\n}\n",
		"removeAccountMember": "package main\n\nfunc main() {}\n",
	} {
		var goSamples []codeSample
		for _, sample := range operations[operationID].CodeSamples {
			if strings.EqualFold(sample.Lang, "go") {
				goSamples = append(goSamples, sample)
			}
		}
		require.Len(t, goSamples, 1, operationID)
		assert.Equal(t, "go", goSamples[0].Lang, operationID)
		assert.Equal(t, "AhaSend Go SDK", goSamples[0].Label, operationID)
		assert.Equal(t, source, goSamples[0].Source, operationID)
	}
}

func TestInjectPlacesEachSampleExactly(t *testing.T) {
	updated, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)
	assert.Equal(t, fixtureWant, string(updated))
}

func TestInjectCreatesABlockAtTheEndOfAnOperation(t *testing.T) {
	updated, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)

	assert.Contains(t, string(updated), `          description: Deleted


      x-code-samples:
        - lang: go
`, "the block goes after the operation's blank lines")
	assert.Contains(t, string(updated), `              println("deleted")
            }

  /v2/accounts/{account_id}/members:
`, "the block ends just before the next path")
	assert.True(t, strings.HasSuffix(string(updated), `          description: Removed

      x-code-samples:
        - lang: go
          label: AhaSend Go SDK
          source: |
            package main

            func main() {}
`), "the last operation's block goes at the end of the file")
}

func TestInjectAppendsAtTheEndOfTheBlock(t *testing.T) {
	updated, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)

	assert.Contains(t, string(updated), `            await client.accounts.update();

        - lang: go
`, "the go entry goes after the block's blank lines")
	assert.Contains(t, string(updated), `              println("updated")
            }

      responses:
`, "the go entry ends just before the next key")
}

func TestInjectReplacesAGoEntryWithTheBlankLinesAfterIt(t *testing.T) {
	updated, err := inject([]byte(fixtureSpec), fixtureSamples())
	require.NoError(t, err)

	assert.Contains(t, string(updated), `                println("pong")
              }
            }

        - lang: javascript
`, "the three blank lines after the old entry become the new entry's one")
	assert.Contains(t, string(updated), `            func main() {}

      operationId: getAccount
`, "an entry directly followed by a key gains its blank line")
}

func TestInjectMatchesGoInAnyCase(t *testing.T) {
	spec := strings.Replace(fixtureSpec, `        - lang: javascript
          label: Node.js 22+ (AhaSend SDK)
          source: |
            await client.ping();
`, `        - lang: GO
          label: Another
          source: |
            package main
`, 1)

	_, err := inject([]byte(spec), fixtureSamples())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ping: has more than one go sample")
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
