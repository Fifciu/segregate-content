package services

import (
	"os"
	"os/exec"
	"path/filepath"
)

type ThumbnailService struct {
	getAppPath func() string
}

func NewThumbnailService(getAppPath func() string) *ThumbnailService {
	return &ThumbnailService{
		getAppPath: getAppPath,
	}
}

// func (s *ThumbnailService) getAppPath() string {
// 	return filepath.Join(s.app.GetSelectedProjectPath())
// }

func (s *ThumbnailService) InitDirectories() error {
	// Make a directory inside base path
	basePath := s.getAppPath()
	appPath := filepath.Join(basePath, ".separate-content")
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		err := os.Mkdir(appPath, 0755)
		if err != nil {
			return err
		}
	}

	// Make a directory inside thumbnail path
	thumbnailPath := filepath.Join(basePath, ".separate-content", "thumbnails")
	if _, err := os.Stat(thumbnailPath); os.IsNotExist(err) {
		err := os.Mkdir(thumbnailPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ThumbnailService) CreateThumbnail(videoPath string, filePath string) error {
	thumbnailPath := filepath.Join(s.getAppPath(), ".separate-content", "thumbnails", filepath.Base(filePath))
	cmd := exec.Command("ffmpeg",
		"-i", videoPath,
		"-vf", "scale=320:-2",
		"-q:v", "2",
		"-frames:v", "1",
		"-y",
		thumbnailPath,
	)

	err := cmd.Run()
	if err != nil {
		// Check if the thumbnail was actually created despite the error
		if _, statErr := os.Stat(thumbnailPath); statErr == nil {
			// Thumbnail exists, so we can ignore the error
			return nil
		}
		return err
	}
	return nil
}
