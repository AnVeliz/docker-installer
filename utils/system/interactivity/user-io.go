package interactivity

import (
	"bufio"
	"fmt"
	"os"
)

// IO is an interface for communication with user's input/output
type IO interface {
	PrintMessage(message ...string)
	GetRune() (rune, error)
}

type consoleIO struct{}

// NewIO is a factory method for creating console IO
func NewIO() IO {
	return &consoleIO{}
}

func (io consoleIO) PrintMessage(message ...string) {
	fmt.Println(message)
}

func (io consoleIO) GetRune() (rune, error) {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	return char, err
}
