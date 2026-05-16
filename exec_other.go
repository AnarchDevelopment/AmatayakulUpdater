//go:build !windows

package main

import (
	"os/exec"
)

func prepareHiddenCommand(cmd *exec.Cmd) {
	// Not needed on non-Windows platforms
}
