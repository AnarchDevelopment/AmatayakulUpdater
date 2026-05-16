package main

import (
	"embed"
	"flag"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	isUpdate bool
	urlStr   string
	pathStr  string
	pidVal   int
	langStr  string
)

func init() {
	flag.BoolVar(&isUpdate, "update", false, "Start an update")
	flag.StringVar(&urlStr, "url", "", "URL for the launcher release")
	flag.StringVar(&pathStr, "path", "", "Path to the current launcher version executable")
	flag.IntVar(&pidVal, "pid", 0, "PID of the current launcher session")
	flag.StringVar(&langStr, "lang", "en", "Language to use (en, es, pt)")
}

func main() {
	for _, arg := range os.Args {
		if arg == "-h" || arg == "-help" || arg == "--help" {
			attachConsole()
		}
	}

	flag.Parse()

	// Wails bindings generator runs a temporary binary, so we must allow it to run without args
	if len(os.Args) < 2 && !strings.Contains(os.Args[0], "wailsbindings") {
		attachConsole()
		flag.Usage()
		os.Exit(1)
	}

	// Create an instance of the app structure
	app := NewApp(isUpdate, urlStr, pathStr, pidVal, langStr)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Amatayakul Updater",
		Width:  800,
		Height: 500,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 20, G: 20, B: 20, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
