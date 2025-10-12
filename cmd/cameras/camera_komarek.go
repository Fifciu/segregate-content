package cameras

import (
	"segregate-content/cmd/utilities"
	"strings"
	"time"

	"googlemaps.github.io/maps"
)

type CameraKomarek struct {
	directory string
}

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

func (c *CameraKomarek) ShouldProcessFile(file string) bool {
	// return strings.HasPrefix(file, "DJI_") && (strings.HasSuffix(file, ".MP4") || strings.HasSuffix(file, ".JPG") || strings.HasSuffix(file, ".DNG"))
	return strings.HasPrefix(file, "DJI_") && (strings.HasSuffix(file, ".MP4") || strings.HasSuffix(file, ".JPG"))
}

func (c *CameraKomarek) NormalizeDateTime(datetime *time.Time, timezoneOffsetSeconds int) *time.Time {
	t := datetime.Add(time.Duration(timezoneOffsetSeconds) * time.Second * -1)
	t = t.Add(24 * time.Hour * -1)
	return &t
}

func (c *CameraKomarek) NormalizeDateTimeWithTimezone(datetime *time.Time, timezone *maps.TimezoneResult) *time.Time {
	// Convert the datetime to the target timezone
	// The input datetime already contains timezone information in the time.Time object
	// The timezone parameter contains the target timezone information

	// First, convert the input datetime to UTC
	utcTime := datetime.UTC()

	// Then apply the target timezone offset (this converts UTC to the target timezone)
	targetOffsetSeconds := timezone.RawOffset + timezone.DstOffset
	targetTime := utcTime.Add(time.Duration(targetOffsetSeconds) * time.Second)

	// Apply the camera-specific adjustment (24 hours back for Komarek)
	targetTime = targetTime.Add(24 * time.Hour * -1)

	return &targetTime
}

func (c *CameraKomarek) New(directory string) Camera {
	return &CameraKomarek{directory: directory}
}

func (c *CameraKomarek) GetDirectory() string {
	return c.directory
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
