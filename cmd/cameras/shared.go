package cameras

import (
	"segregate-content/cmd/utilities"
	"strings"
)

func isIphone(directory string, fullPath string) bool {
	files, err := utilities.ListFiles(fullPath)
	if err != nil {
		return false
	}
	for _, file := range files {
		if strings.HasPrefix(file, "IMG_") && (strings.HasSuffix(file, ".MOV") || strings.HasSuffix(file, ".HEIC")) {
			return true
		}
	}
	return false
}
