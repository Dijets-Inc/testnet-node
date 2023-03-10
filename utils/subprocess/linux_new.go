//go:build linux
// +build linux

// ^ syscall.SysProcAttr only has field Pdeathsig on Linux

// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package subprocess

import (
	"os/exec"
	"syscall"
)

func New(path string, args ...string) *exec.Cmd {
	cmd := exec.Command(path, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGTERM}
	return cmd
}
