package app

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DirectorySelector struct {
	app       *App
	directory string
}

func (d *DirectorySelector) SelectDirectory() (string, error) {
	selection, err := runtime.OpenDirectoryDialog(d.app.GetCtx(), runtime.OpenDialogOptions{
		Title: "Wybierz folder",
	})
	if err != nil {
		return "", err
	}
	if selection == "" {
		return "", err
	}
	d.directory = selection
	return d.directory, nil
}

func (d *DirectorySelector) GetDirectory() string {
	return d.directory
}

func NewDirectorySelector(app *App) *DirectorySelector {
	return &DirectorySelector{
		app: app,
	}
}
