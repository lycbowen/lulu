package main

import (
	"embed"
	"encoding/json"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type Config struct {
	FitMode      string `json:"fit_mode"`
	PlaySpeed    int    `json:"play_speed"`
	AlwaysOnTop  bool   `json:"always_on_top"`
	WindowWidth  int    `json:"window_width"`
	WindowHeight int    `json:"window_height"`
}

var Cfg = Config{}

func init() {
	f, err := os.ReadFile("config")
	if err != nil {
		defaultConfig := Config{
			FitMode:      "contain",
			PlaySpeed:    5000,
			AlwaysOnTop:  false,
			WindowWidth:  800,
			WindowHeight: 600,
		}
		Cfg = defaultConfig
		cfg, _ := json.Marshal(defaultConfig)
		newfile, _ := os.OpenFile("config", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		defer newfile.Close()
		newfile.Write(cfg)
		return
	}
	json.Unmarshal(f, &Cfg)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "lulu",
		Width:  Cfg.WindowWidth,
		Height: Cfg.WindowHeight,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		AlwaysOnTop: Cfg.AlwaysOnTop,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
