package api

import (
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go/internal/prismmock"
)

// TestMain starts the Prism mock server that the specification-driven
// Test_ahasend_* suites in this package execute against.
//
// Those suites address localhost:4010 directly. Nothing used to start a server
// there, so they could only ever run in the configuration that skips them, and
// a plain `go test ./api/` failed with connection errors. See prismmock for
// how an unavailable mock is handled.
func TestMain(m *testing.M) {
	os.Exit(prismmock.RunTests(m))
}
