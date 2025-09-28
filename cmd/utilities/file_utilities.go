package utilities

import (
	"fmt"
	"os"

	"strings"

	"github.com/barasher/go-exiftool"
)

// DirectoryExists checks if the specified directory exists
func DirectoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// ListDirectories returns a list of directory names in the specified directory
func ListDirectories(path string) ([]string, error) {
	var directories []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			directories = append(directories, entry.Name())
		}
	}

	return directories, nil
}

// ListFiles returns a list of file names (with extensions) in the specified directory
func ListFiles(path string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func GetImageMetadata(path string) bool {
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return false
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			if strings.Contains(k, "GPS") {
				fmt.Printf("[%v] %v\n", k, v)
				return true
			}
		}
	}
	return false
}
