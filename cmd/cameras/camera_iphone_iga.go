package cameras

import "strings"

type CameraIphoneIga struct{}

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

func isIphoneIga(directory string, fullPath string) bool {
	return isIphone(directory, fullPath) && (strings.Contains(directory, "Iga") || strings.Contains(directory, "Igi"))
}
