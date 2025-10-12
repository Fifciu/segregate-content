package app

import (
	"segregate-content/cmd/cameras"
	"segregate-content/pkg/services"
	"testing"
	"time"
)

func TestCreateThumbnail(t *testing.T) {
	application := NewApp()
	application.SetSelectedProjectPath("/Volumes/Fifciuu SSD/Example")
	processor := NewProcessor(application, nil, services.NewThumbnailService(application.GetSelectedProjectPath))

	t.Run("creates thumbnail in expected place", func(t *testing.T) {
		var camera cameras.Camera
		camera = (&cameras.CameraKomarek{}).New("Komarek")
		file, err := processor.CreateThumbnail("/Volumes/Fifciuu SSD/Example/Komarek", &ProcessedFile{
			Camera:              &camera,
			Filename:            "DJI_0627.MP4",
			CameraPath:          "Komarek",
			NormalizedTimestamp: time.Now(),
			LegacyTimestamp:     time.Now(),
			Thumbnail:           nil,
		})
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if file != "/Volumes/Fifciuu SSD/Example/.separate-content/thumbnails/komarek-dji_0627.jpg" {
			t.Errorf("Expected thumbnail to be created in /Volumes/Fifciuu SSD/Example/.separate-content/thumbnails/komarek-dji_0627.jpg, got %s", file)
		}
	})

	t.Run("uses camera directory for thumbnail name", func(t *testing.T) {
		var camera cameras.Camera
		camera = (&cameras.CameraKomarek{}).New("diffdirrectory")
		file, err := processor.CreateThumbnail("/Volumes/Fifciuu SSD/Example/Komarek", &ProcessedFile{
			Camera:              &camera,
			Filename:            "DJI_0627.MP4",
			CameraPath:          "Komarek",
			NormalizedTimestamp: time.Now(),
			LegacyTimestamp:     time.Now(),
			Thumbnail:           nil,
		})
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if file != "/Volumes/Fifciuu SSD/Example/.separate-content/thumbnails/diffdirrectory-dji_0627.jpg" {
			t.Errorf("Expected thumbnail to be created in /Volumes/Fifciuu SSD/Example/.separate-content/thumbnails/diffdirrectory-dji_0627.jpg, got %s", file)
		}
	})
}
