package cameras

import "strings"

type CameraIphoneFilip struct{}

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

func isIphoneFilip(directory string, fullPath string) bool {
	return isIphone(directory, fullPath) && strings.Contains(directory, "Fil")
}
