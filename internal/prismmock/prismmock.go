// Package prismmock runs the Prism mock server that this module's
// specification-driven tests execute against.
//
// It is shared by every package whose tests speak to the mock so that the
// server is started the same way everywhere, and so that a machine without
// Prism installed skips those tests instead of failing them.
package prismmock

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// DefaultPort is the port the mock server listens on.
const DefaultPort = "4010"

// skipEnv disables the Prism-backed tests. Each test reads it directly, so
// setting it here is how an unavailable mock turns into a skip.
const skipEnv = "SKIP_INTEGRATION_TESTS"

// requireEnv makes an unavailable mock a failure instead of a skip. The
// release pipeline sets it so that a missing mock cannot quietly reduce the
// release gate to nothing.
const requireEnv = "AHASEND_REQUIRE_PRISM"

// startTimeout is generous because the npx fallback may download Prism first.
const startTimeout = 90 * time.Second

// Server is a running Prism mock server.
type Server struct {
	cmd  *exec.Cmd
	port string
}

// Start launches Prism against this module's OpenAPI specification and waits
// until it answers requests.
//
// The mock runs in static mode. Prism's --dynamic mode invents a fresh value
// for every field on every request, including values the specification permits
// but the API never emits, which makes it a generator of spurious failures
// rather than a contract check.
func Start(port string) (*Server, error) {
	specPath, err := findSpec()
	if err != nil {
		return nil, err
	}

	args := []string{"mock", specPath, "--host", "127.0.0.1", "--port", port}

	var cmd *exec.Cmd
	switch {
	case os.Getenv("PRISM_CMD") != "":
		cmd = exec.Command(os.Getenv("PRISM_CMD"), args...)
	case lookPathOK("prism"):
		cmd = exec.Command("prism", args...)
	case lookPathOK("npx"):
		cmd = exec.Command("npx", append([]string{"@stoplight/prism-cli"}, args...)...)
	default:
		return nil, fmt.Errorf("prism is not available: install it with 'npm install -g @stoplight/prism-cli', or make npx available")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start prism: %w", err)
	}

	server := &Server{cmd: cmd, port: port}
	deadline := time.Now().Add(startTimeout)
	for time.Now().Before(deadline) {
		if server.ready() {
			return server, nil
		}
		time.Sleep(time.Second)
	}

	_ = server.Stop()
	return nil, fmt.Errorf("prism did not become ready within %s", startTimeout)
}

// Stop terminates the mock server.
func (s *Server) Stop() error {
	if s == nil || s.cmd == nil || s.cmd.Process == nil {
		return nil
	}
	if err := s.cmd.Process.Kill(); err != nil {
		return err
	}
	_, _ = s.cmd.Process.Wait()
	return nil
}

// Addr is the host:port the mock server listens on.
func (s *Server) Addr() string {
	return "localhost:" + s.port
}

// ready reports whether the mock answers. An unauthenticated request is
// expected to be rejected, so any HTTP response at all means it is up.
func (s *Server) ready() bool {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://127.0.0.1:" + s.port + "/v2/ping")
	if err != nil {
		return false
	}
	_ = resp.Body.Close()
	return true
}

// RunTests is the TestMain body for a package whose tests run against the
// mock. It returns the exit code the caller should pass to os.Exit.
//
// When the mock cannot be started it sets SKIP_INTEGRATION_TESTS so the tests
// skip themselves, which keeps a plain `go test ./...` working on a machine
// without Prism. Setting AHASEND_REQUIRE_PRISM=true turns that skip into a
// failure, which is what the release pipeline wants: a gate that cannot pass
// by running nothing.
func RunTests(m *testing.M, port string) int {
	if os.Getenv(skipEnv) == "true" {
		fmt.Printf("prismmock: %s=true, running without the mock server\n", skipEnv)
		return m.Run()
	}

	server, err := Start(port)
	if err != nil {
		if os.Getenv(requireEnv) == "true" {
			fmt.Printf("prismmock: %s=true and the mock server could not start: %v\n", requireEnv, err)
			return 1
		}
		fmt.Printf("prismmock: skipping mock-backed tests, the mock server could not start: %v\n", err)
		if setErr := os.Setenv(skipEnv, "true"); setErr != nil {
			fmt.Printf("prismmock: could not set %s: %v\n", skipEnv, setErr)
			return 1
		}
		return m.Run()
	}

	fmt.Printf("prismmock: mock server listening on %s\n", server.Addr())
	code := m.Run()
	if stopErr := server.Stop(); stopErr != nil {
		fmt.Printf("prismmock: failed to stop the mock server: %v\n", stopErr)
	}
	return code
}

// findSpec walks up from the working directory to the module root and returns
// the absolute path of the OpenAPI specification, so that the caller's package
// directory does not matter.
func findSpec() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		candidate := filepath.Join(dir, "openapi", "openapi.yaml")
		if _, statErr := os.Stat(candidate); statErr == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("openapi/openapi.yaml not found above the working directory")
		}
		dir = parent
	}
}

func lookPathOK(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
