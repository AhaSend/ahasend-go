//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package prismmock

import (
	"errors"
	"io"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestTerminateProcessGroupStopsDescendants(t *testing.T) {
	// The shell announces itself only once the descendant exists: Start
	// returns as soon as the shell is exec'd, which is before it has forked
	// anything, and a kill delivered in that window would leave a "sleep" the
	// signal never reached.
	cmd := exec.Command("sh", "-c", "sleep 300 & echo ready && wait")
	configureProcessGroup(cmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("open the shell's stdout: %v", err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start process tree: %v", err)
	}

	announcement := make([]byte, len("ready\n"))
	if _, err := io.ReadFull(stdout, announcement); err != nil {
		t.Fatalf("wait for the descendant to start: %v", err)
	}

	if err := terminateProcessGroup(cmd); err != nil {
		t.Fatalf("terminate process group: %v", err)
	}
	_ = cmd.Wait()

	// A killed process keeps its group alive until it is reaped, and the
	// descendant is reparented to init rather than to this test, so the group
	// can outlive the kill by however long init takes to get to it. Poll for
	// the group to go rather than reading it once: the single read is what
	// made this pass on a developer's machine and fail on a busy runner.
	deadline := time.Now().Add(10 * time.Second)
	for {
		err := syscall.Kill(-cmd.Process.Pid, 0)
		if errors.Is(err, syscall.ESRCH) {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("process group still exists after termination: %v", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
