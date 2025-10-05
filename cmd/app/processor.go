package app

import (
	"fmt"
	"segregate-content/cmd/cameras"
	"segregate-content/cmd/utilities"
	"segregate-content/pkg/services"
	"sort"
	"strings"
	"time"

	"googlemaps.github.io/maps"
)

type Processor struct {
	app                 *App
	elevationApiService *services.ElevationApiService
}

type Project struct {
	Name        string
	Folder      string
	PlanFile    string
	HomeCountry string
}

func NewProcessor(app *App, elevationApiService *services.ElevationApiService) *Processor {
	return &Processor{
		app:                 app,
		elevationApiService: elevationApiService,
	}
}

func (p *Processor) doNeedTimezoneOffset(cameras []*cameras.Camera) bool {
	for _, camera := range cameras {
		if !(*camera).HasSyncedClock() {
			return true
		}
	}
	return false
}

func (p *Processor) CreateProject(project *Project) ([][][]*ProcessedFile, error) {
	fmt.Println("Processing project:", project)
	directoryExists, err := utilities.DirectoryExists(project.Folder)
	if err != nil {
		return nil, err
	}
	if !directoryExists {
		return nil, fmt.Errorf("directory does not exist")
	}
	directories, err := utilities.ListDirectories(project.Folder)
	if err != nil {
		return nil, err
	}
	fmt.Println("Directories:", directories)
	cameraMap := make(map[string]*cameras.Camera)
	for _, directory := range directories {
		directoryToCamera := cameras.DirectoryToCamera(directory, fmt.Sprintf("%s/%s", project.Folder, directory))
		if directoryToCamera == nil {
			fmt.Printf("Directory \"%s\" to camera: Unknown\n", directory)
			continue
		}
		cameraMap[directory] = directoryToCamera
		fmt.Printf("Directory \"%s\" to camera: %s\n", directory, directoryToCamera)
	}
	cameraList := make([]*cameras.Camera, 0, len(cameraMap))
	for _, camera := range cameraMap {
		cameraList = append(cameraList, camera)
	}
	if p.doNeedTimezoneOffset(cameraList) {
		fmt.Println("Need timezone offset")
	} else {
		fmt.Println("No need for timezone offset")
	}
	var cameraWithCoordinates *cameras.Camera
	var cameraWithCoordinatesPath string
	for directory, camera := range cameraMap {
		if (*camera).HasCoordinates() {
			cameraWithCoordinates = camera
			cameraWithCoordinatesPath = fmt.Sprintf("%s/%s", project.Folder, directory)
			fmt.Println("Camera with coordinates:", cameraWithCoordinates)
			break
		}
	}
	if cameraWithCoordinates == nil {
		return nil, fmt.Errorf("No camera with coordinates found")
	}
	timezone, homeCityTimezone, err := p.FindTimezone(cameraWithCoordinatesPath)
	if err != nil {
		return nil, err
	}
	diffSeconds := (homeCityTimezone.RawOffset + homeCityTimezone.DstOffset) - (timezone.RawOffset + timezone.DstOffset)
	fmt.Println("Timezone:", timezone)
	fmt.Println("Home city timezone:", homeCityTimezone)
	fmt.Println("Timezone difference:", diffSeconds)
	processedFiles, err := p.ProcessFiles(project, cameraMap, diffSeconds, timezone)
	if err != nil {
		return nil, err
	}
	for _, processedFile := range processedFiles {
		fmt.Printf("Processed file (%s) timestamp: %s (legacy: %s)\n", processedFile.Filename, processedFile.NormalizedTimestamp, processedFile.LegacyTimestamp)
	}

	// Grouping files in blocks
	groupedFiles := make([][]*ProcessedFile, 0)
	MAX_DIFF := time.Duration(20 * time.Minute)
	currentGroup := make([]*ProcessedFile, 0)
	currentGroup = append(currentGroup, processedFiles[0])

	for i := 1; i < len(processedFiles); i++ {
		if processedFiles[i].NormalizedTimestamp.Sub(currentGroup[len(currentGroup)-1].NormalizedTimestamp) > MAX_DIFF {
			groupedFiles = append(groupedFiles, currentGroup)
			currentGroup = make([]*ProcessedFile, 0)
		}
		currentGroup = append(currentGroup, processedFiles[i])
	}

	// Print grouped files
	for _, group := range groupedFiles {
		fmt.Printf("Group: %d\n", group)
		for _, file := range group {
			fmt.Printf("File: %s, %s\n", file.Filename, file.NormalizedTimestamp.Format("2006-01-02 15:04:05.000"))
		}
	}

	// Grouping blocks in days
	days := make([][][]*ProcessedFile, 0)
	day := make([][]*ProcessedFile, 0)
	currentDay := groupedFiles[0][0].NormalizedTimestamp.Format("2006-01-02")
	day = append(day, groupedFiles[0])

	for i := 1; i < len(groupedFiles); i++ {
		if groupedFiles[i][0].NormalizedTimestamp.Format("2006-01-02") != currentDay {
			days = append(days, day)
			day = make([][]*ProcessedFile, 0)
			currentDay = groupedFiles[i][0].NormalizedTimestamp.Format("2006-01-02")
		} else {
			day = append(day, groupedFiles[i])
		}
	}

	p.app.SetSelectedProjectPath(project.Folder)
	fmt.Println("Project processed")
	return days, nil
}

type ProcessedFile struct {
	Camera              *cameras.Camera
	Filename            string
	CameraPath          string
	NormalizedTimestamp time.Time
	LegacyTimestamp     time.Time
}

func (p *Processor) ProcessFiles(project *Project, cameraMap map[string]*cameras.Camera, diffSeconds int, timezone *maps.TimezoneResult) ([]*ProcessedFile, error) {
	startTime := time.Now()
	fmt.Printf("ProcessFiles started at: %s\n", startTime.Format("2006-01-02 15:04:05.000"))

	processedFiles := make([]*ProcessedFile, 0)
	totalFiles := 0
	processedCount := 0

	for directory, camera := range cameraMap {
		cameraPath := fmt.Sprintf("%s/%s", project.Folder, directory)
		fmt.Printf("Processing directory: %s\n", directory)

		dirStartTime := time.Now()
		files, err := utilities.ListFiles(cameraPath)
		if err != nil {
			return nil, err
		}
		fmt.Println(files)
		dirListTime := time.Since(dirStartTime)
		fmt.Printf("Directory %s: Listed %d files in %v\n", directory, len(files), dirListTime)

		for _, file := range files {
			totalFiles++
			fileStartTime := time.Now()

			if !(*camera).ShouldProcessFile(file) {
				fileSkipTime := time.Since(fileStartTime)
				fmt.Printf("Skipped file %s (not processable) in %v\n", file, fileSkipTime)
				continue
			}

			fmt.Printf("Processing file: %s\n", file)

			if (*camera).GetCameraType() == cameras.CAMERA_KOMAREK {
				// Get file create date
				createDateStartTime := time.Now()
				inodeChangeDate, err := utilities.GetFileInodeChangeDate(fmt.Sprintf("%s/%s", cameraPath, file))
				var parsedTimestampWithTimezone time.Time
				var createDate string
				if err == nil {
					parsedTimestampWithTimezone, err = utilities.ParseDateTimeStringWithTimezone(inodeChangeDate)
					if err != nil {
						fmt.Printf("Error parsing datetime with timezone for %s, %v\n", file, err)
					}
				}
				// if err != nil {
				fmt.Printf("Error getting inode change date for %s, %v\nLooking for CreateDate then...", file, err)
				createDate, err = utilities.GetFileCreateDate(fmt.Sprintf("%s/%s", cameraPath, file))
				fileErrorTime := time.Since(fileStartTime)
				fmt.Printf("Error getting create date for %s after %v: %v\n", file, fileErrorTime, err)
				// return nil, err
				// }

				createDateTime := time.Since(createDateStartTime)
				fmt.Printf("File %s: Retrieved create date in %v\n", file, createDateTime)

				// Parse datetime string
				parseStartTime := time.Now()
				parsedTimestamp, err := utilities.ParseDateTimeString(createDate)
				if err != nil {
					fileErrorTime := time.Since(fileStartTime)
					fmt.Printf("Error parsing datetime for %s after %v: %v\n", file, fileErrorTime, err)
					return nil, err
				}
				parseTime := time.Since(parseStartTime)
				fmt.Printf("File %s: Parsed datetime in %v\n", file, parseTime)

				fmt.Printf("Create date: %s\n", createDate)

				// Normalize timezone using timezone data when available
				var normalizedTimestamp *time.Time
				if komarekCamera, ok := (*camera).(*cameras.CameraKomarek); ok {
					if parsedTimestampWithTimezone != (time.Time{}) {
						normalizedTimestamp = komarekCamera.NormalizeDateTimeWithTimezone(&parsedTimestampWithTimezone, timezone)
					} else {
						normalizedTimestamp = komarekCamera.NormalizeDateTimeWithTimezone(&parsedTimestamp, timezone)
					}
				} else {
					normalizedTimestamp = (*camera).NormalizeDateTime(&parsedTimestampWithTimezone, diffSeconds)
				}

				processedFile := &ProcessedFile{
					Camera:              camera,
					Filename:            file,
					CameraPath:          directory,
					NormalizedTimestamp: *normalizedTimestamp,
					LegacyTimestamp:     parsedTimestamp,
				}
				processedFiles = append(processedFiles, processedFile)
				fileTotalTime := time.Since(fileStartTime)
				fmt.Printf("File %s: Total processing time %v (create date: %v, parse: %v)\n",
					file, fileTotalTime, createDateTime, parseTime)
				processedCount++
			} else {
				// Get file create date
				createDateStartTime := time.Now()
				createDate, err := utilities.GetFileCreateDate(fmt.Sprintf("%s/%s", cameraPath, file))
				if err != nil {
					fileErrorTime := time.Since(fileStartTime)
					fmt.Printf("Error getting create date for %s after %v: %v\n", file, fileErrorTime, err)
					return nil, err
				}
				createDateTime := time.Since(createDateStartTime)
				fmt.Printf("File %s: Retrieved create date in %v\n", file, createDateTime)

				// Parse datetime string
				parseStartTime := time.Now()
				parsedTimestamp, err := utilities.ParseDateTimeString(createDate)
				if err != nil {
					fileErrorTime := time.Since(fileStartTime)
					fmt.Printf("Error parsing datetime for %s after %v: %v\n", file, fileErrorTime, err)
					return nil, err
				}
				parseTime := time.Since(parseStartTime)
				fmt.Printf("File %s: Parsed datetime in %v\n", file, parseTime)

				fmt.Printf("Create date: %s\n", createDate)

				// Normalize timezone using timezone data when available
				// Use timezone-aware normalization for Komarek camera
				var normalizedTimestamp *time.Time
				if komarekCamera, ok := (*camera).(*cameras.CameraKomarek); ok {
					normalizedTimestamp = komarekCamera.NormalizeDateTimeWithTimezone(&parsedTimestamp, timezone)
				} else {
					normalizedTimestamp = (*camera).NormalizeDateTime(&parsedTimestamp, diffSeconds)
				}

				processedFile := &ProcessedFile{
					Camera:              camera,
					Filename:            file,
					CameraPath:          directory,
					NormalizedTimestamp: *normalizedTimestamp,
					LegacyTimestamp:     parsedTimestamp,
				}
				processedFiles = append(processedFiles, processedFile)
				fileTotalTime := time.Since(fileStartTime)
				fmt.Printf("File %s: Total processing time %v (create date: %v, parse: %v)\n",
					file, fileTotalTime, createDateTime, parseTime)
				processedCount++
			}
		}
	}

	totalTime := time.Since(startTime)
	fmt.Printf("ProcessFiles completed at: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))
	fmt.Printf("ProcessFiles summary: %d total files, %d processed, %d skipped in %v\n",
		totalFiles, processedCount, totalFiles-processedCount, totalTime)

	// Sort processedFiles by NormalizedTimestamp in ascending order
	sort.Slice(processedFiles, func(i, j int) bool {
		return processedFiles[i].NormalizedTimestamp.Before(processedFiles[j].NormalizedTimestamp)
	})

	return processedFiles, nil
}

func (p *Processor) FindTimezone(cameraWithCoordinatesPath string) (*maps.TimezoneResult, *maps.TimezoneResult, error) {
	files, err := utilities.ListFiles(cameraWithCoordinatesPath)
	if err != nil {
		return nil, nil, err
	}
	// TODO: Add support for multiple timezones
	for _, file := range files {
		if strings.HasPrefix(file, "IMG_") && strings.HasSuffix(file, ".HEIC") {
			fmt.Println("Checking file with name:", file)
			gpsData := utilities.GetImageGpsData(fmt.Sprintf("%s/%s", cameraWithCoordinatesPath, file))
			imageDateTime, err := utilities.GetImageDateTime(fmt.Sprintf("%s/%s", cameraWithCoordinatesPath, file))
			if err != nil {
				fmt.Println("Error getting image date time:", err)
				continue
			}
			fmt.Println("Image date time:", imageDateTime)
			parsedDateTime, err := p.elevationApiService.ParseDateTimeString(imageDateTime)
			if err != nil {
				fmt.Println("Error parsing date time:", err)
				continue
			}
			fmt.Println("Parsed date time:", parsedDateTime)
			if gpsData != nil {
				fmt.Println("Image has coordinates")
				fmt.Printf("%s, %s\n", gpsData.Latitude, gpsData.Longitude)
				latitude, err := p.elevationApiService.ParseCoordinateString(gpsData.Latitude)
				if err != nil {
					return nil, nil, err
				}
				longitude, err := p.elevationApiService.ParseCoordinateString(gpsData.Longitude)
				if err != nil {
					return nil, nil, err
				}
				fmt.Printf("%f, %f\n", latitude, longitude)
				timezone, err := p.elevationApiService.GetTimezone(p.app.GetCtx(), latitude, longitude, parsedDateTime)
				if err != nil {
					return nil, nil, err
				}
				homeCountryTimezone, err := p.elevationApiService.GetTimezone(p.app.GetCtx(), 51.403145, 21.088193, parsedDateTime)
				if err != nil {
					return nil, nil, err
				}
				return timezone, homeCountryTimezone, nil
			} else {
				fmt.Println("Image does not have coordinates")
			}
		}
	}
	return nil, nil, fmt.Errorf("No timezone found")
}
