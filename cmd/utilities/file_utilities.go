package utilities

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/barasher/go-exiftool"
)

type ImageGpsData struct {
	Latitude  string
	Longitude string
}

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

func GetImageGpsData(path string) *ImageGpsData {
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return nil
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		gpsData := &ImageGpsData{}
		for k, v := range fileInfo.Fields {
			// GPS Latitude
			// GPS Longitude
			if strings.Contains(k, "GPS") {
				fmt.Printf("[%v] %v\n", k, v)
			}
			if k == "GPSLongitude" {
				fmt.Printf("[%v] %v\n", k, v)
				gpsData.Longitude = v.(string)
			}
			if k == "GPSLatitude" {
				fmt.Printf("[%v] %v\n", k, v)
				gpsData.Latitude = v.(string)
			}
		}
		if gpsData.Latitude != "" && gpsData.Longitude != "" {
			return gpsData
		}
	}
	return nil
}

func GetImageDateTime(path string) (string, error) {
	et, err := exiftool.NewExiftool()
	if err != nil {
		return "", err
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Println(k)
			if k == "GPSDateTime" {
				fmt.Printf("[%v] %v\n", k, v)
				return v.(string), nil
			}
		}
	}
	return "", errors.New("GPSDateTime not found")
}

func GetFileCreateDate(path string) (string, error) {
	et, err := exiftool.NewExiftool()
	if err != nil {
		return "", err
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)
			if k == "CreateDate" {
				fmt.Printf("[%v] %v\n", k, v)
				return v.(string), nil
			}
		}
	}
	return "", errors.New("CreateDate not found")
}

func GetFileInodeChangeDate(path string) (string, error) {
	et, err := exiftool.NewExiftool()
	if err != nil {
		return "", err
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)
			if k == "FileInodeChangeDate" {
				fmt.Printf("[%v] %v\n", k, v)
				return v.(string), nil
			}
		}
	}
	return "", errors.New("FileInodeChangeDate not found")
}

// ParseDateTimeString converts a datetime string in format "2024:07:03 11:04:46" to time.Time
func ParseDateTimeString(dateTimeStr string) (time.Time, error) {
	// Remove extra whitespace and normalize the string
	dateTimeStr = strings.TrimSpace(dateTimeStr)

	// Regular expression to match the format: YYYY:MM:DD HH:MM:SS
	re := regexp.MustCompile(`^(\d{4}):(\d{2}):(\d{2})\s+(\d{2}):(\d{2}):(\d{2})$`)
	matches := re.FindStringSubmatch(dateTimeStr)

	if len(matches) != 7 {
		return time.Time{}, fmt.Errorf("invalid datetime format: %s", dateTimeStr)
	}

	// Parse year, month, day
	year, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year: %s", matches[1])
	}

	month, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month: %s", matches[2])
	}

	day, err := strconv.Atoi(matches[3])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %s", matches[3])
	}

	// Parse hour, minute, second
	hour, err := strconv.Atoi(matches[4])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid hour: %s", matches[4])
	}

	minute, err := strconv.Atoi(matches[5])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid minute: %s", matches[5])
	}

	second, err := strconv.Atoi(matches[6])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid second: %s", matches[6])
	}

	// Create time in UTC (assuming UTC if no timezone specified)
	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)

	return t, nil
}

// ParseDateTimeStringWithTimezone converts a datetime string with timezone info in format "2024:07:22 11:36:53+02:00" to time.Time
func ParseDateTimeStringWithTimezone(dateTimeStr string) (time.Time, error) {
	// Remove extra whitespace and normalize the string
	dateTimeStr = strings.TrimSpace(dateTimeStr)

	// Regular expression to match the format: YYYY:MM:DD HH:MM:SS+HH:MM or YYYY:MM:DD HH:MM:SS-HH:MM
	re := regexp.MustCompile(`^(\d{4}):(\d{2}):(\d{2})\s+(\d{2}):(\d{2}):(\d{2})([+-])(\d{2}):(\d{2})$`)
	matches := re.FindStringSubmatch(dateTimeStr)

	if len(matches) != 10 {
		return time.Time{}, fmt.Errorf("invalid datetime with timezone format: %s", dateTimeStr)
	}

	// Parse year, month, day
	year, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year: %s", matches[1])
	}

	month, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month: %s", matches[2])
	}

	day, err := strconv.Atoi(matches[3])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %s", matches[3])
	}

	// Parse hour, minute, second
	hour, err := strconv.Atoi(matches[4])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid hour: %s", matches[4])
	}

	minute, err := strconv.Atoi(matches[5])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid minute: %s", matches[5])
	}

	second, err := strconv.Atoi(matches[6])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid second: %s", matches[6])
	}

	// Parse timezone offset
	tzSign := matches[7] // + or -
	tzHour, err := strconv.Atoi(matches[8])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone hour: %s", matches[8])
	}

	tzMinute, err := strconv.Atoi(matches[9])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone minute: %s", matches[9])
	}

	// Calculate timezone offset in minutes
	tzOffsetMinutes := tzHour*60 + tzMinute
	if tzSign == "-" {
		tzOffsetMinutes = -tzOffsetMinutes
	}

	// Create timezone with the calculated offset
	timezone := time.FixedZone("", tzOffsetMinutes*60)

	// Create time with the specific timezone
	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, timezone)

	return t, nil
}
