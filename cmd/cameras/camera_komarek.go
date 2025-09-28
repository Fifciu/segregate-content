package cameras

import (
	"segregate-content/cmd/utilities"
	"strings"
)

type CameraKomarek struct{}

func (c *CameraKomarek) GetCameraType() int {
	return CAMERA_KOMAREK
}

func (c *CameraKomarek) GetCameraName() string {
	return "Komarek"
}

func (c *CameraKomarek) IsCamera(directory string, fullPath string) bool {
	return isDJI(directory, fullPath)
}

func (c *CameraKomarek) HasSyncedClock() bool {
	return false
}

func (c *CameraKomarek) HasCoordinates() bool {
	return false
}

func isDJI(directory string, fullPath string) bool {
	files, err := utilities.ListFiles(fullPath)
	if err != nil {
		return false
	}
	for _, file := range files {
		if strings.HasPrefix(file, "DJI_") {
			return true
		}
	}
	return false
}
