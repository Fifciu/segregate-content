package cameras

import (
	"segregate-content/cmd/utilities"
	"strings"
)

type CameraLumix struct{}

func (c *CameraLumix) GetCameraType() int {
	return CAMERA_LUMIX
}

func (c *CameraLumix) GetCameraName() string {
	return "Lumix"
}

func (c *CameraLumix) IsCamera(directory string, fullPath string) bool {
	return isLumix(directory, fullPath)
}

func (c *CameraLumix) HasSyncedClock() bool {
	return false
}

func (c *CameraLumix) HasCoordinates() bool {
	return false
}

func isLumix(directory string, fullPath string) bool {
	files, err := utilities.ListFiles(fullPath)
	if err != nil {
		return false
	}
	for _, file := range files {
		if strings.HasPrefix(file, "P") && (strings.HasSuffix(file, ".MOV") || strings.HasSuffix(file, ".RW2") || strings.HasSuffix(file, ".JPG")) {
			return true
		}
	}
	return false
}
