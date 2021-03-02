package main

import (
	"github.com/AnVeliz/docker-installer/utils/system"
	"github.com/AnVeliz/docker-installer/utils/system/interactivity"
	"github.com/AnVeliz/docker-installer/utils/user"
)

func main() {
	ioInteractor := interactivity.NewIO()
	userInteractor := user.NewInteractor(ioInteractor)

	userInteractor.IO().PrintMessage("Welcome to Docker installer!")
	err := Run(SetupRegistry(system.NewBashRunner(ioInteractor)), system.NewOsDetector(), userInteractor)
	if err != nil {
		userInteractor.IO().PrintMessage("Error while running: ", err.Error())
	}
}
