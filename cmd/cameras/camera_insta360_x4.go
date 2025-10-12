package cameras

import (
	"segregate-content/cmd/utilities"
	"strings"
	"time"
)

type CameraInsta360X4 struct {
	directory string
}

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

func (c *CameraInsta360X4) ShouldProcessFile(file string) bool {
	return strings.HasSuffix(file, ".insv") && strings.HasPrefix(file, "VID_")
}

func (c *CameraInsta360X4) NormalizeDateTime(datetime *time.Time, timezoneOffsetSeconds int) *time.Time {
	return datetime
}

func (c *CameraInsta360X4) New(directory string) Camera {
	return &CameraInsta360X4{directory: directory}
}

func (c *CameraInsta360X4) GetDirectory() string {
	return c.directory
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
