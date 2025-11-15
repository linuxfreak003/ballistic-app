package main

import (
	"embed"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "ballistic-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Logger:           NewCustomLogger(),
		Bind: []any{
			app,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "Ballistic Calculator",
				Message: "© 2025 Jared Chapman",
				//Icon:    icon,
			},
			//TitleBar: mac.TitleBarHiddenInset(),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

type CustomLogger struct {
	*logrus.Logger
}

func NewCustomLogger() *CustomLogger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return &CustomLogger{
		Logger: log,
	}
}

func (c *CustomLogger) Trace(msg string)   { c.Logger.Tracef(msg) }
func (c *CustomLogger) Debug(msg string)   { c.Logger.Debugf(msg) }
func (c *CustomLogger) Info(msg string)    { c.Logger.Infof(msg) }
func (c *CustomLogger) Warning(msg string) { c.Logger.Warnf(msg) }
func (c *CustomLogger) Error(msg string)   { c.Logger.Errorf(msg) }
func (c *CustomLogger) Fatal(msg string)   { c.Logger.Fatalf(msg) }
func (c *CustomLogger) Print(msg string)   { c.Logger.Printf(msg) }
