package cameras

import (
	"segregate-content/cmd/utilities"
	"strings"
)

type CameraInsta360X4 struct{}

func (c *CameraInsta360X4) GetCameraType() int {
	return CAMERA_INSTA360_X4
}

func (c *CameraInsta360X4) GetCameraName() string {
	return "Insta360 X4"
}

func (c *CameraInsta360X4) IsCamera(directory string, fullPath string) bool {
	return isInsta360_X4(directory, fullPath)
}

func (c *CameraInsta360X4) HasSyncedClock() bool {
	return false
}

func (c *CameraInsta360X4) HasCoordinates() bool {
	return false
}

func isInsta360_X4(directory string, fullPath string) bool {
	files, err := utilities.ListFiles(fullPath)
	if err != nil {
		return false
	}
	for _, file := range files {
		if strings.HasSuffix(file, ".insv") {
			return true
		}
	}
	return false
}
