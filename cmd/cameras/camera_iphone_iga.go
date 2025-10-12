package cameras

import (
	"strings"
	"time"
)

type CameraIphoneIga struct {
	directory string
}

func (c *CameraIphoneIga) GetCameraType() int {
	return CAMERA_IPHONE_IGA
}

func (c *CameraIphoneIga) GetCameraName() string {
	return "iPhone Iga"
}

func (c *CameraIphoneIga) IsCamera(directory string, fullPath string) bool {
	return isIphoneIga(directory, fullPath)
}

func (c *CameraIphoneIga) HasSyncedClock() bool {
	return true
}

func (c *CameraIphoneIga) HasCoordinates() bool {
	return false
}

func (c *CameraIphoneIga) ShouldProcessFile(file string) bool {
	return strings.HasPrefix(file, "IMG_") && (strings.HasSuffix(file, ".MOV") || strings.HasSuffix(file, ".HEIC") || strings.HasSuffix(file, ".mov"))
}

func (c *CameraIphoneIga) NormalizeDateTime(datetime *time.Time, timezoneOffsetSeconds int) *time.Time {
	return datetime
}

func (c *CameraIphoneIga) New(directory string) Camera {
	return &CameraIphoneIga{directory: directory}
}

func (c *CameraIphoneIga) GetDirectory() string {
	return c.directory
}

func isIphoneIga(directory string, fullPath string) bool {
	return isIphone(directory, fullPath) && (strings.Contains(directory, "Iga") || strings.Contains(directory, "Igi"))
}
