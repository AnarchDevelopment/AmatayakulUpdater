package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx      context.Context
	isUpdate bool
	url      string
	path     string
	pid      int
	lang     string
}

func NewApp(isUpdate bool, url, path string, pid int, lang string) *App {
	return &App{
		isUpdate: isUpdate,
		url:      url,
		path:     path,
		pid:      pid,
		lang:     lang,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetArgs() map[string]interface{} {
	return map[string]interface{}{
		"isUpdate": a.isUpdate,
		"url":      a.url,
		"path":     a.path,
		"pid":      a.pid,
		"lang":     a.lang,
	}
}

func (a *App) RunUpdate() {
	go a.performUpdate()
}

type WriteCounter struct {
	Total      uint64
	TotalSize  uint64
	StartTime  time.Time
	LastUpdate time.Time
	App        *App
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)

	now := time.Now()
	if now.Sub(wc.LastUpdate) > 100*time.Millisecond {
		elapsed := now.Sub(wc.StartTime).Seconds()
		var speed float64
		if elapsed > 0 {
			speed = float64(wc.Total) / elapsed
		}
		mbps := (speed * 8) / 1000000.0

		downloadedMB := float64(wc.Total) / 1048576.0
		totalMB := float64(wc.TotalSize) / 1048576.0
		
		var remaining float64
		if speed > 0 {
			remaining = float64(wc.TotalSize-wc.Total) / speed
		}

		runtime.EventsEmit(wc.App.ctx, "update:progress", map[string]interface{}{
			"mbps":         mbps,
			"remainingSec": remaining,
			"downloadedMB": downloadedMB,
			"totalMB":      totalMB,
			"percent":      (float64(wc.Total) / float64(wc.TotalSize)) * 100,
		})
		wc.LastUpdate = now
	}
	return n, nil
}

func (a *App) performUpdate() {
	runtime.EventsEmit(a.ctx, "update:status", "Waiting for launcher to close...")
	
	if a.pid > 0 {
		attempts := 0
		for isProcessRunning(a.pid) {
			if attempts >= 10 { // 5 seconds
				cmd := exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprintf("%d", a.pid))
				prepareHiddenCommand(cmd)
				cmd.Run()
				break
			}
			time.Sleep(500 * time.Millisecond)
			attempts++
		}
	}
	
	if a.path != "" {
		exeName := filepath.Base(a.path)
		attempts := 0
		for isProcessNameRunning(exeName) {
			if attempts >= 10 { // 5 seconds
				cmd := exec.Command("taskkill", "/F", "/T", "/IM", exeName)
				prepareHiddenCommand(cmd)
				cmd.Run()
				break
			}
			time.Sleep(500 * time.Millisecond)
			attempts++
		}
	}

	runtime.EventsEmit(a.ctx, "update:status", "Renaming old executable...")
	
	oldPath := a.path
	if oldPath == "" {
		oldPath = "AmatayakulLauncher.exe"
	}
	
	renamedPath := oldPath + ".old"
	os.Remove(renamedPath)
	err := os.Rename(oldPath, renamedPath)
	if err != nil && !os.IsNotExist(err) {
		runtime.EventsEmit(a.ctx, "update:error", "Failed to rename executable: "+err.Error())
		return
	}

	runtime.EventsEmit(a.ctx, "update:status", "Downloading update...")

	resp, err := http.Get(a.url)
	if err != nil {
		runtime.EventsEmit(a.ctx, "update:error", "Failed to download update: "+err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		runtime.EventsEmit(a.ctx, "update:error", fmt.Sprintf("Failed to download update: HTTP %d", resp.StatusCode))
		return
	}

	out, err := os.OpenFile(oldPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		runtime.EventsEmit(a.ctx, "update:error", "Failed to create new executable: "+err.Error())
		return
	}

	counter := &WriteCounter{
		TotalSize:  uint64(resp.ContentLength),
		StartTime:  time.Now(),
		LastUpdate: time.Now(),
		App:        a,
	}

	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	out.Close()
	
	if err != nil {
		runtime.EventsEmit(a.ctx, "update:error", "Failed to save update: "+err.Error())
		return
	}

	runtime.EventsEmit(a.ctx, "update:progress", map[string]interface{}{
		"percent": 100,
	})

	runtime.EventsEmit(a.ctx, "update:done", oldPath)
}

func (a *App) LaunchAndExit(path string) {
	cmd := exec.Command(path)
	prepareHiddenCommand(cmd)
	cmd.Start()
	os.Exit(0)
}

func (a *App) Exit() {
	os.Exit(0)
}

func isProcessRunning(pid int) bool {
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid), "/NH")
	prepareHiddenCommand(cmd)
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	output := string(out)
	// If tasklist returns "INFO: No tasks are running...", it's not running
	if strings.Contains(output, "INFO:") {
		return false
	}
	return strings.Contains(output, fmt.Sprintf("%d", pid))
}

func isProcessNameRunning(name string) bool {
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("IMAGENAME eq %s", name), "/NH")
	prepareHiddenCommand(cmd)
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	output := string(out)
	// If tasklist returns "INFO: No tasks are running...", it's not running
	if strings.Contains(output, "INFO:") {
		return false
	}
	return strings.Contains(strings.ToLower(output), strings.ToLower(name))
}
