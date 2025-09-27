package app

import (
	"fmt"
	"time"
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

func (p *Processor) CreateProject(project *Project) error {
	fmt.Println("Processing project:", project)
	time.Sleep(10 * time.Second)
	fmt.Println("Project processed")
	return nil
}
