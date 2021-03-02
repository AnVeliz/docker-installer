package system

import (
	"bufio"
	"os/exec"

	"github.com/AnVeliz/docker-installer/utils/system/interactivity"
)

// IRunner is an interface of commands runner which runs commands
type IRunner interface {
	Run(command Command)
}

type bashRunner struct {
	io interactivity.IO
}

// NewBashRunner creates an instance of bashRunner
func NewBashRunner(userIO interactivity.IO) IRunner {
	return &bashRunner{
		io: userIO,
	}
}

// Run executes a command
func (runner bashRunner) Run(command Command) {
	if message := command.WelcomeMessage; message != "" {
		runner.io.PrintMessage("=====>", message)
	}

	err := runner.execute(command.Command, command.Arguments...)
	if err != nil {
		runner.io.PrintMessage("[===ERROR===]", err.Error())
		if action := command.ErrorAction; action != nil {
			action()
		}
	} else {
		runner.io.PrintMessage("[Done]")
		if action := command.SuccessAction; action != nil {
			action()
		}
	}

	if message := command.GoodbyeMessage; message != "" {
		runner.io.PrintMessage("=====>", message)
	}
}

// execute command functionality
func (runner bashRunner) execute(command string, arguments ...string) error {
	cmd := exec.Command(command, arguments...)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	commandOutputScanner := bufio.NewScanner(stdout)
	for commandOutputScanner.Scan() {
		message := commandOutputScanner.Text()
		runner.io.PrintMessage(message)
	}

	return cmd.Wait()
}
