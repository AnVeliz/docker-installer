package utils

import (
	"fmt"
	"os/exec"
)

// BashRunner is a commands runner which runs commands in Bash
type BashRunner struct {
}

// Run executes a command
func (runner BashRunner) Run(command Command) {
	if message := command.WelcomeMessage; message != "" {
		fmt.Println(message)
	}

	msg, err := execute(command.Command, command.Arguments...)
	fmt.Println(msg)
	if err != nil {
		fmt.Println("[===ERROR===]", err)
		if action := command.ErrorAction; action != nil {
			action()
		}
	} else {
		fmt.Println("[Done]")
		if action := command.SuccessAction; action != nil {
			action()
		}
	}

	if message := command.GoodbyMessage; message != "" {
		fmt.Println(message)
	}
}

// execute command functionality
func execute(command string, arguments ...string) (string, error) {
	resultBytes, err := exec.Command(command, arguments...).Output()
	return string(resultBytes), err
}
