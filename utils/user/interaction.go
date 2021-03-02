package user

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AnVeliz/docker-installer/installers"
)

// Interactor is for interaction with users
type Interactor interface {
	GetOperation() (installers.OperationType, error)
}

type interactor struct {
}

// NewInteractor returns a user interactor implementation
func NewInteractor() Interactor {
	return &interactor{}
}

// GetOperation asks for the operation user is going to perform
func (intrct interactor) GetOperation() (installers.OperationType, error) {
	fmt.Println("Select option: \n\t[1] Install Docker\n\t[2] Uninstall Docker")

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
		return intrct.GetOperation()
	}
}
