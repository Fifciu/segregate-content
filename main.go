package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

type FileLoader struct {
	http.Handler
	app *app.App
}

func NewFileLoader(app *app.App) *FileLoader {
	return &FileLoader{
		app: app,
	}
}

func (h *FileLoader) IsVideo(filename string) bool {
	return strings.HasSuffix(filename, ".MOV") || strings.HasSuffix(filename, ".mov") || strings.HasSuffix(filename, ".MP4")
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	if !strings.HasPrefix(requestedFilename, "current-project") {
		println("Invalid file path:", requestedFilename)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid file path"))
		return
	}
	requestedFilename = strings.TrimPrefix(requestedFilename, "current-project/")
	println("Requesting file:", requestedFilename)
	selectedProjectPath := h.app.GetSelectedProjectPath()
	println("Selected project path:", selectedProjectPath)
	println("Selected project path:", filepath.Join(selectedProjectPath, requestedFilename))
	requestedFilename = filepath.Join(selectedProjectPath, requestedFilename)
	if strings.HasSuffix(requestedFilename, ".insv") {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("This kind of file is not supported: %s", requestedFilename)))
		return
	}
	// if strings.HasSuffix(requestedFilename, ".MOV") || strings.HasSuffix(requestedFilename, ".mov") || strings.HasSuffix(requestedFilename, ".MP4") {
	// 	res.WriteHeader(http.StatusBadRequest)
	// 	res.Write([]byte(fmt.Sprintf("This kind of file is not supported: %s", requestedFilename)))
	// 	return
	// }
	if h.IsVideo(requestedFilename) {
		file, err := os.Open(requestedFilename)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(fmt.Sprintf("Could not get file stat for %s", requestedFilename)))
			return
		}
		http.ServeContent(res, req, requestedFilename, stat.ModTime(), file)
		return
	}
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}

	res.Write(fileData)
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
	fileLoader := NewFileLoader(appInstance)
	// services
	elevationApiService := services.NewElevationApiService(googleProjectKey)

	// main action
	processor := app.NewProcessor(appInstance, elevationApiService)

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "segregate-content",
		Width:             1280,
		Height:            800,
		MinWidth:          1280,
		MinHeight:         800,
		MaxWidth:          1280,
		MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: fileLoader,
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
