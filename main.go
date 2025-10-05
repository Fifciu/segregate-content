package main

import (
	"embed"
	"log"
	"os"
	"strconv"

	"segregate-content/cmd/app"
	"segregate-content/pkg/services"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

// getEnvOrDefault returns the value of an environment variable or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvOrDefaultInt returns the value of an environment variable as an integer or a default value
func getEnvOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Warning: Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
	}
	return defaultValue
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		// Continue execution even if .env file is not found
	}

	// Access GOOGLE_PROJECT_KEY environment variable
	googleProjectKey := os.Getenv("GOOGLE_PROJECT_KEY")
	if googleProjectKey == "" {
		log.Println("Warning: GOOGLE_PROJECT_KEY environment variable is not set")
	} else {
		log.Println("Google Project Key loaded") // Log only first 8 chars for security
	}

	// Create an instance of the app structure
	appInstance := app.NewApp()
	fileSelector := app.NewFileSelector(appInstance)
	directorySelector := app.NewDirectorySelector(appInstance)

	// services
	elevationApiService := services.NewElevationApiService(googleProjectKey)

	// main action
	processor := app.NewProcessor(appInstance, elevationApiService)

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "segregate-content",
		Width:             1024,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		MaxWidth:          1280,
		MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         logger.DEBUG,
		OnStartup:        appInstance.OnStartup,
		OnDomReady:       appInstance.OnDomReady,
		OnBeforeClose:    appInstance.OnBeforeClose,
		OnShutdown:       appInstance.OnShutdown,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			appInstance,
			processor,
			fileSelector,
			directorySelector,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			ZoomFactor:          1.0,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "segregate-content",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
