package app

import (
	"fmt"
	"segregate-content/cmd/cameras"
	"segregate-content/cmd/utilities"
	"strings"
)

type Processor struct {
	app *App
}

type Project struct {
	Name        string
	Folder      string
	PlanFile    string
	HomeCountry string
}

func NewProcessor(app *App) *Processor {
	return &Processor{
		app: app,
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

func (p *Processor) CreateProject(project *Project) error {
	fmt.Println("Processing project:", project)
	directoryExists, err := utilities.DirectoryExists(project.Folder)
	if err != nil {
		return err
	}
	if !directoryExists {
		return fmt.Errorf("directory does not exist")
	}
	directories, err := utilities.ListDirectories(project.Folder)
	if err != nil {
		return err
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
		return fmt.Errorf("No camera with coordinates found")
	}
	files, err := utilities.ListFiles(cameraWithCoordinatesPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasPrefix(file, "IMG_") && strings.HasSuffix(file, ".HEIC") {
			res := utilities.GetImageMetadata(fmt.Sprintf("%s/%s", cameraWithCoordinatesPath, file))
			if res {
				fmt.Println("Image has coordinates")
				break
			} else {
				fmt.Println("Image does not have coordinates")
			}
		}
	}
	fmt.Println("Project processed")
	return nil
}
