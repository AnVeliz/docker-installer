package user

import (
	"errors"
	"testing"

	"github.com/AnVeliz/docker-installer/installers"
)

var fakeIoInteractorExecutionNumber int = 0

type fakeIoInteractor struct {
	rn  rune
	err error

	getRuneSecondTime func() rune
}

func (fakeInteractor fakeIoInteractor) PrintMessage(message ...string) {
}

func (fakeInteractor fakeIoInteractor) GetRune() (rune, error) {
	fakeIoInteractorExecutionNumber++
	if fakeInteractor.getRuneSecondTime != nil && fakeIoInteractorExecutionNumber == 2 {
		return fakeInteractor.getRuneSecondTime(), fakeInteractor.err
	}
	return fakeInteractor.rn, fakeInteractor.err
}

func TestUserInteractorHasIoInteractor(t *testing.T) {
	interactor := NewInteractor(fakeIoInteractor{})

	if interactor == nil {
		t.Error("Users interactor can't be created")
	}

	if interactor.IO() == nil {
		t.Error("IO interactor has been lost")
	}
}

func TestGetInstallOperation(t *testing.T) {
	fakeIoInteractor := fakeIoInteractor{
		rn:  '1',
		err: nil,
	}
	interactor := NewInteractor(fakeIoInteractor)

	operation, err := interactor.GetOperation()
	if err != nil {
		t.Error("Operation returns an error")
	}

	if operation != installers.Install {
		t.Error("Wrong operation")
	}
}

func TestGetUninstallOperation(t *testing.T) {
	fakeIoInteractor := fakeIoInteractor{
		rn:  '2',
		err: nil,
	}
	interactor := NewInteractor(fakeIoInteractor)

	operation, err := interactor.GetOperation()
	if err != nil {
		t.Error("Operation returns an error")
	}

	if operation != installers.Uninstall {
		t.Error("Wrong operation")
	}
}

func TestGetWrongOperation(t *testing.T) {
	errMsg := "Dummy error"
	fakeIoInteractorExecutionNumber = 0
	fakeIoInteractor := fakeIoInteractor{
		rn:                '1',
		err:               errors.New(errMsg),
		getRuneSecondTime: func() rune { return '1' },
	}
	interactor := NewInteractor(fakeIoInteractor)

	_, err := interactor.GetOperation()
	if err == nil || err.Error() != errMsg {
		t.Error("Wrong error message")
	}
}

func TestGetWrongRune(t *testing.T) {
	fakeIoInteractorExecutionNumber = 0
	fakeIoInteractor := fakeIoInteractor{
		rn:                '9',
		err:               nil,
		getRuneSecondTime: func() rune { return '1' },
	}
	interactor := NewInteractor(fakeIoInteractor)

	operation, err := interactor.GetOperation()
	if err != nil {
		t.Error("Operation returns an error")
	}
	if operation != installers.Install {
		t.Error("Wrong operation")
	}
}
