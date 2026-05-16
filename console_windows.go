//go:build windows

package main

import (
	"os"
	"syscall"
)

func attachConsole() {
	_, _, _ = syscall.NewLazyDLL("kernel32.dll").NewProc("AttachConsole").Call(uintptr(0xFFFFFFFF))
	if h, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE); err == nil {
		os.Stdout = os.NewFile(uintptr(h), "/dev/stdout")
	}
	if h, err := syscall.GetStdHandle(syscall.STD_ERROR_HANDLE); err == nil {
		os.Stderr = os.NewFile(uintptr(h), "/dev/stderr")
	}
}
