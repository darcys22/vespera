package main

import (
	"embed"
  "log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/darcys22/vespera/vespera"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
      log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
}

func main() {
	// Create an instance of the app structure
  state := vespera.NewState()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "vespera",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        state.Startup,
    OnShutdown:       state.Shutdown,
		Bind: []interface{}{
			state,
		},
	})

	if err != nil {
    log.Fatal("Error:", err.Error())
	}
}

