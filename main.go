package main

import (
	"fmt"

	"github.com/AnVeliz/docker-installer/utils/system"
	"github.com/AnVeliz/docker-installer/utils/user"
)

func main() {
	fmt.Println("Welcome to Docker installer!")
	err := Run(system.NewOsDetector(), user.NewInteractor())
	if err != nil {
		fmt.Println("Error while running: ", err.Error())
	}
}
