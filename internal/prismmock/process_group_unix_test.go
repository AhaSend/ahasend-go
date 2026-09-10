//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package prismmock

import (
	"errors"
	"os/exec"
	"syscall"
	"testing"
)

func TestTerminateProcessGroupStopsDescendants(t *testing.T) {
	cmd := exec.Command("sh", "-c", "sleep 300 & wait")
	configureProcessGroup(cmd)
	if err := cmd.Start(); err != nil {
		t.Fatalf("start process tree: %v", err)
	}

	if err := terminateProcessGroup(cmd); err != nil {
		t.Fatalf("terminate process group: %v", err)
	}
	_ = cmd.Wait()

	if err := syscall.Kill(-cmd.Process.Pid, 0); !errors.Is(err, syscall.ESRCH) {
		t.Fatalf("process group still exists after termination: %v", err)
	}
}
