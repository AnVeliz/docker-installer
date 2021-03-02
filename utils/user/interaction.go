package user

import (
	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/utils/system/interactivity"
)

// Interactor is for interaction with users
type Interactor interface {
	GetOperation() (installers.OperationType, error)
	IO() interactivity.IO
}

type interactor struct {
	io interactivity.IO
}

// NewInteractor returns a user interactor implementation
func NewInteractor(ioInteractor interactivity.IO) Interactor {
	return &interactor{
		io: ioInteractor,
	}
}

// GetOperation asks for the operation user is going to perform
func (intrct interactor) GetOperation() (installers.OperationType, error) {
	intrct.io.PrintMessage("Select option: \n\t[1] Install Docker\n\t[2] Uninstall Docker")

	char, err := intrct.io.GetRune()

	if err != nil {
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
		intrct.io.PrintMessage("Wrong input. You need to choose 1 or 2")
		return intrct.GetOperation()
	}
}

func (intrct interactor) IO() interactivity.IO {
	return intrct.io
}
