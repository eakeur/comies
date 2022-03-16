package app

import (
	"gomies/app/core/wrappers"
	"gomies/app/start"
)

type App struct {
	wrappers.Workflows
}

func NewApp() *App {

	managers := start.NewManagers()
	actions := start.NewActions()
	workflows := start.NewWorkflows(actions, managers)

	return &App{
		Workflows: workflows,
	}
}
