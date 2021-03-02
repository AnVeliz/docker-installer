package main

import (
	"errors"

	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/installers/docker"
	"github.com/AnVeliz/docker-installer/utils/system"
	"github.com/AnVeliz/docker-installer/utils/user"
)

// Run is a start point
func Run(registry installers.IRegistry, osDetector system.IOsDetector, userInteractor user.Interactor) error {
	osInfo, userInfo, err := osDetector.GetOsInfo()
	if err != nil {
		return err
	}
	if userInfo.UserID != 0 {
		return errors.New("Application should run as root")
	}

	operation, err := userInteractor.GetOperation()
	if err != nil {
		return err
	}

	installer := registry.FindInstaller(osInfo)
	if installer == nil {
		return errors.New("No compatible installer has been found")
	}
	userInteractor.IO().PrintMessage("Installer has been found")

	switch operation {
	case installers.Install:
		installer.Install()
		break
	case installers.Uninstall:
		installer.Uninstall()
		break
	default:
		return errors.New("Unexpected operation requested")
	}

	return nil
}

// SetupRegistry is for creating a registry with real installers
func SetupRegistry(commandsRunner system.IRunner) installers.IRegistry {
	registry := installers.CreateRegistry()
	registry.Register(docker.CreateInstaller(commandsRunner))
	return registry
}
