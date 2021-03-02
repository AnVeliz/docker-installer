package main

import (
	"errors"
	"fmt"

	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/installers/docker"
	"github.com/AnVeliz/docker-installer/utils/system"
	"github.com/AnVeliz/docker-installer/utils/user"
)

// Run is a start point
func Run() error {
	osInfo, userInfo := system.GetOsInfo()
	if userInfo.UserID != 0 {
		return errors.New("Application should run as root")
	}

	fmt.Println("Select option: \n\t[1] Install Docker\n\t[2] Uninstall Docker")

	registry := setupRegistry()
	operation, err := user.GetOperation()
	if err != nil {
		return err
	}

	installer := registry.FindInstaller(osInfo)
	if installer == nil {
		return errors.New("No compatible installer has been found")
	}
	fmt.Println("Installer has been found")

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

func setupRegistry() installers.IRegistry {
	commandsRunner := &system.BashRunner{}
	registry := installers.CreateRegistry()
	registry.Register(docker.CreateInstaller(commandsRunner))
	return registry
}
