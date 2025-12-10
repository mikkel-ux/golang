package main

import (
	"embed"
	"golangfileExplore/test"

	files "golangfileExplore/backend/files"
	settings "golangfileExplore/backend/settings"
	tabs "golangfileExplore/backend/tabs"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	/* TODO implement settings check and create if settings don't exist */
	data, err := settings.CheckAndCreateSettings()
	if err != nil {
		println("Error checking/creating settings:", err.Error())
		return
	}
	settings.AppSettings = *data
	// Create an instance of the app structure
	app := NewApp()
	sampleTest := test.NewSampleTest()
	distanceTest := files.NewDistanceTest()
	tabsManager := tabs.NewTabsManager()

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "golangfileExplore",
		Width:     1024,
		Height:    768,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: true,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			sampleTest,
			distanceTest,
			tabsManager,
		},
		Logger: nil,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
