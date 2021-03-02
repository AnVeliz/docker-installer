package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Docker installer!")
	err := Run()
	if err != nil {
		fmt.Println("Error while running: ", err.Error())
	}
}
