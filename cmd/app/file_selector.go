package app

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileSelector struct {
	app  *App
	file string
}

func (d *FileSelector) SelectFile() (string, error) {
	selection, err := runtime.OpenFileDialog(d.app.GetCtx(), runtime.OpenDialogOptions{
		Title: "Wybierz folder",
		Filters: []runtime.FileFilter{
			{
				Pattern:     "*.md",
				DisplayName: "Pliki MD",
			},
		},
	})
	if err != nil {
		return "", err
	}
	if selection == "" {
		return "", err
	}
	d.file = selection
	return d.file, nil
}

func (d *FileSelector) GetFile() string {
	return d.file
}

func NewFileSelector(app *App) *FileSelector {
	return &FileSelector{
		app: app,
	}
}
