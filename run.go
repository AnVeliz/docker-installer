package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/installers/docker"
	"github.com/AnVeliz/docker-installer/utils"
)

// Run is a start point
func Run() {
	osInfo, userInfo := utils.GetOsInfo()
	if userInfo.UserID != 0 {
		fmt.Println("Application should run as root")
		os.Exit(5)
	}

	fmt.Println("Select option: \n\t[1] Install Docker\n\t[2] Uninstall Docker")

	registry := setupRegistry()
	operation := getOperation()
	installer := registry.FindInstaller(osInfo)
	if installer == nil {
		fmt.Println("No compatible installer has been found!")
		os.Exit(3)
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
		fmt.Println("Unexpected operation requested")
		os.Exit(1)
	}
}

func setupRegistry() installers.IRegistry {
	commandsRunner := &utils.BashRunner{}
	registry := installers.CreateRegistry()
	registry.Register(docker.CreateInstaller(commandsRunner))
	return registry
}

func getOperation() installers.OperationType {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if char == '1' {
		return installers.Install
	}

	switch char {
	case '1':
		return installers.Install
	case '2':
		return installers.Uninstall
	default:
		fmt.Println("Wrong input. You need to choose 1 or 2")
		return getOperation()
	}
}
