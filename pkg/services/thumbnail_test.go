package services

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitDirectories(t *testing.T) {
	getAppPath := func() string {
		return "/Volumes/Fifciuu SSD/Example"
	}
	thumbnailService := NewThumbnailService(getAppPath)

	t.Run("creates thumbnail in expected place", func(t *testing.T) {
		thumbnailService.InitDirectories()
		basePath := getAppPath()
		thumbnailPath := filepath.Join(basePath, ".separate-content", "thumbnails")
		if _, err := os.Stat(thumbnailPath); os.IsNotExist(err) {
			t.Errorf("Expected thumbnail path to exist, got %v", err)
		}
	})

}
