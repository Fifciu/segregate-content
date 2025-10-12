package cameras

import (
	"segregate-content/cmd/utilities"
	"strings"
	"time"
)

type CameraLumix struct {
	directory string
}

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

func (c *CameraLumix) ShouldProcessFile(file string) bool {
	return strings.HasPrefix(file, "P") && (strings.HasSuffix(file, ".MOV") || strings.HasSuffix(file, ".JPG"))
	// return strings.HasPrefix(file, "P") && (strings.HasSuffix(file, ".MOV") || strings.HasSuffix(file, ".RW2") || strings.HasSuffix(file, ".JPG"))
}

func (c *CameraLumix) NormalizeDateTime(datetime *time.Time, timezoneOffsetSeconds int) *time.Time {
	t := datetime.Add(time.Duration(timezoneOffsetSeconds) * time.Second * -1)
	return &t
}

func (c *CameraLumix) New(directory string) Camera {
	return &CameraLumix{directory: directory}
}

func (c *CameraLumix) GetDirectory() string {
	return c.directory
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
