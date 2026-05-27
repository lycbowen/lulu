package main

import (
	"embed"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	PhotoPath    string `json:"photo_path"`
}

var (
	Cfg       = Config{}
	configDir = "."
)

func init() {
	defaultConfig := Config{
		FitMode:      "contain",
		PlaySpeed:    5000,
		AlwaysOnTop:  false,
		WindowWidth:  800,
		WindowHeight: 600,
		PhotoPath:    "photos",
	}

	exePath, err := os.Executable()
	if err != nil {
		exePath = "."
	}
	configDir = filepath.Dir(exePath)
	configPath := filepath.Join(configDir, "config.json")
	writeConfigGuide(filepath.Join(configDir, "config说明.txt"))

	f, err := os.ReadFile(configPath)
	if err != nil {
		Cfg = defaultConfig
		writeDefaultConfig(configPath, defaultConfig)
		_ = os.MkdirAll(resolvePhotoPath(Cfg.PhotoPath), 0755)
		return
	}

	Cfg = defaultConfig
	if err := json.Unmarshal(f, &Cfg); err != nil {
		Cfg = defaultConfig
		writeDefaultConfig(configPath, defaultConfig)
	}
	Cfg = normalizeConfig(Cfg, defaultConfig)
	_ = os.MkdirAll(resolvePhotoPath(Cfg.PhotoPath), 0755)
}

func normalizeConfig(cfg Config, defaults Config) Config {
	if cfg.FitMode == "" {
		cfg.FitMode = defaults.FitMode
	}
	if cfg.PlaySpeed <= 0 {
		cfg.PlaySpeed = defaults.PlaySpeed
	}
	if cfg.WindowWidth <= 0 {
		cfg.WindowWidth = defaults.WindowWidth
	}
	if cfg.WindowHeight <= 0 {
		cfg.WindowHeight = defaults.WindowHeight
	}
	if strings.TrimSpace(cfg.PhotoPath) == "" {
		cfg.PhotoPath = defaults.PhotoPath
	}
	return cfg
}

func writeDefaultConfig(path string, cfg Config) {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(path, data, 0644)
}

func writeConfigGuide(path string) {
	if _, err := os.Stat(path); err == nil {
		return
	}

	content := `lulu 照片播放器配置说明

本文件和 config.json 位于程序同一目录。
第一次运行时，程序会自动生成 config.json、config说明.txt 和默认 photos 文件夹。

config.json 示例：

{
  "fit_mode": "contain",
  "play_speed": 5000,
  "always_on_top": false,
  "window_width": 800,
  "window_height": 600,
  "photo_path": "photos"
}

字段说明：

- fit_mode：图片适配方式。
  - contain：完整显示图片，可能留黑边。
  - cover：铺满窗口，可能裁剪图片边缘。
  - fill：拉伸填满窗口，可能变形。

- play_speed：轮播间隔，单位毫秒。
  - 5000 表示每 5 秒切换一张。

- always_on_top：窗口是否置顶。
  - true：始终置顶。
  - false：不置顶。

- window_width / window_height：启动窗口宽高，单位像素。

- photo_path：照片文件夹路径。
  - 写相对路径时，相对于程序所在目录，例如 "photos"。
  - 写绝对路径也可以，例如 "D:\\Pictures\\album"。

支持的图片格式：

.jpg、.jpeg、.png、.gif、.webp、.bmp

注意：

- 默认只读取 photo_path 文件夹当前层的图片，不读取子文件夹。
- 图片按文件名排序播放，可以通过重命名控制顺序。
- GIF 动图会直接播放。
- 修改 config.json 后，需要重新启动程序才会生效。
`
	_ = os.WriteFile(path, []byte(content), 0644)
}

func resolvePhotoPath(path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	return filepath.Clean(filepath.Join(configDir, path))
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
			Assets:  assets,
			Handler: http.HandlerFunc(servePhoto),
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
