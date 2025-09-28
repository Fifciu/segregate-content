package cameras

type Camera interface {
	GetCameraType() int
	GetCameraName() string
	IsCamera(directory string, fullPath string) bool
	HasSyncedClock() bool
	HasCoordinates() bool
}

const (
	CAMERA_UNDEFINED = iota
	CAMERA_IPHONE_FILIP
	CAMERA_IPHONE_IGA
	CAMERA_LUMIX
	CAMERA_KOMAREK
	// CAMERA_AVATA // Disabled for now, as I am not flying it abroad much
	CAMERA_INSTA360_X4
)

var cameras = []Camera{
	&CameraIphoneFilip{},
	&CameraIphoneIga{},
	&CameraLumix{},
	&CameraKomarek{},
	&CameraInsta360X4{},
}

func DirectoryToCamera(directory string, fullPath string) *Camera {
	for _, camera := range cameras {
		if camera.IsCamera(directory, fullPath) {
			return &camera
		}
	}
	return nil
}
