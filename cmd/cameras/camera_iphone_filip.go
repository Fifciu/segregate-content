package cameras

import (
	"strings"
	"time"
)

type CameraIphoneFilip struct {
	directory string
}

func (c *CameraIphoneFilip) GetCameraType() int {
	return CAMERA_IPHONE_FILIP
}

func (c *CameraIphoneFilip) GetCameraName() string {
	return "iPhone Filip"
}

func (c *CameraIphoneFilip) IsCamera(directory string, fullPath string) bool {
	return isIphoneFilip(directory, fullPath)
}

func (c *CameraIphoneFilip) HasSyncedClock() bool {
	return true
}

func (c *CameraIphoneFilip) HasCoordinates() bool {
	return true
}

func (c *CameraIphoneFilip) ShouldProcessFile(file string) bool {
	return strings.HasPrefix(file, "IMG_") && (strings.HasSuffix(file, ".MOV") || strings.HasSuffix(file, ".HEIC") || strings.HasSuffix(file, ".mov"))
}

func (c *CameraIphoneFilip) NormalizeDateTime(datetime *time.Time, timezoneOffsetSeconds int) *time.Time {
	return datetime
}

func (c *CameraIphoneFilip) New(directory string) Camera {
	return &CameraIphoneFilip{directory: directory}
}

func (c *CameraIphoneFilip) GetDirectory() string {
	return c.directory
}

func isIphoneFilip(directory string, fullPath string) bool {
	return isIphone(directory, fullPath) && strings.Contains(directory, "Fil")
}
