package user

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AnVeliz/docker-installer/installers"
)

// GetOperation asks for the operation user is going to perform
func GetOperation() (installers.OperationType, error) {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		return installers.NotSpecified, err
	}

	if char == '1' {
		return installers.Install, nil
	}

	switch char {
	case '1':
		return installers.Install, nil
	case '2':
		return installers.Uninstall, nil
	default:
		fmt.Println("Wrong input. You need to choose 1 or 2")
		return GetOperation()
	}
}
