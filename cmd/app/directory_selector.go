package app

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DirectorySelector struct {
	app                  *App
	sourceDirectory      string
	destinationDirectory string
}

func (d *DirectorySelector) SelectDirectory(isSource bool) (string, error) {
	selection, err := runtime.OpenDirectoryDialog(d.app.GetCtx(), runtime.OpenDialogOptions{
		Title: "Wybierz folder",
	})
	if err != nil {
		return "", err
	}
	if selection == "" {
		return "", err
	}
	if isSource {
		d.sourceDirectory = selection
	} else {
		d.destinationDirectory = selection
	}
	return selection, nil
}

func (d *DirectorySelector) GetSourceDirectory() string {
	return d.sourceDirectory
}

func (d *DirectorySelector) GetDestinationDirectory() string {
	return d.destinationDirectory
}

func NewDirectorySelector(app *App) *DirectorySelector {
	return &DirectorySelector{
		app: app,
	}
}
