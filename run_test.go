package main

import (
	"errors"
	"testing"

	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/utils/system"
)

type fakeOsDetector struct {
	fakeOsInfo   system.OsInfo
	fakeUserInfo system.UserInfo
}

func (osDetector fakeOsDetector) GetOsInfo() (system.OsInfo, system.UserInfo) {
	return osDetector.fakeOsInfo, osDetector.fakeUserInfo
}

type fakeUserInteractor struct {
	operationType installers.OperationType
	err           error
}

func (interactor fakeUserInteractor) GetOperation() (installers.OperationType, error) {
	return interactor.operationType, interactor.err
}

func TestInstallSupportedOss(t *testing.T) {
	detector := fakeOsDetector{
		fakeOsInfo: system.OsInfo{
			OsClass:   system.Linux,
			OsVersion: "20.10",
			OsName:    "Ubuntu",
		},
		fakeUserInfo: system.UserInfo{
			Name:   "root",
			UserID: 0,
		},
	}

	interactor := fakeUserInteractor{
		operationType: installers.Install,
		err:           nil,
	}

	err := Run(detector, interactor)
	if err != nil {
		t.Error("Run function result is wrong:", err.Error())
	}
}

func TestInstallNotSupportedUser(t *testing.T) {
	detector := fakeOsDetector{
		fakeOsInfo: system.OsInfo{
			OsClass:   system.Linux,
			OsVersion: "20.10",
			OsName:    "Ubuntu",
		},
		fakeUserInfo: system.UserInfo{
			Name:   "user",
			UserID: 1,
		},
	}

	interactor := fakeUserInteractor{
		operationType: installers.Install,
		err:           nil,
	}

	err := Run(detector, interactor)
	if err == nil || err.Error() != "Application should run as root" {
		t.Error("Run function result is wrong:", err.Error())
	}
}

func TestInstallNotSupportedOss(t *testing.T) {
	detector := fakeOsDetector{
		fakeOsInfo: system.OsInfo{
			OsClass:   system.Linux,
			OsVersion: "10.10",
			OsName:    "Manjaro",
		},
		fakeUserInfo: system.UserInfo{
			Name:   "root",
			UserID: 0,
		},
	}

	interactor := fakeUserInteractor{
		operationType: installers.Install,
		err:           nil,
	}

	err := Run(detector, interactor)
	if err == nil || err.Error() != "No compatible installer has been found" {
		t.Error("Run function result is wrong:", err.Error())
	}
}

func TestUninstallSupportedOss(t *testing.T) {
	detector := fakeOsDetector{
		fakeOsInfo: system.OsInfo{
			OsClass:   system.Linux,
			OsVersion: "20.10",
			OsName:    "Ubuntu",
		},
		fakeUserInfo: system.UserInfo{
			Name:   "root",
			UserID: 0,
		},
	}

	interactor := fakeUserInteractor{
		operationType: installers.Uninstall,
		err:           nil,
	}

	err := Run(detector, interactor)
	if err != nil {
		t.Error("Run function result is wrong:", err.Error())
	}
}

func TestNotSpecifiedOperationSupportedOss(t *testing.T) {
	detector := fakeOsDetector{
		fakeOsInfo: system.OsInfo{
			OsClass:   system.Linux,
			OsVersion: "20.10",
			OsName:    "Ubuntu",
		},
		fakeUserInfo: system.UserInfo{
			Name:   "root",
			UserID: 0,
		},
	}

	interactor := fakeUserInteractor{
		operationType: installers.NotSpecified,
		err:           nil,
	}

	err := Run(detector, interactor)
	if err != nil {
		t.Error("Run function result is wrong:", err.Error())
	}
}

func TestUserInputError(t *testing.T) {
	detector := fakeOsDetector{
		fakeOsInfo: system.OsInfo{
			OsClass:   system.Linux,
			OsVersion: "10.10",
			OsName:    "Manjaro",
		},
		fakeUserInfo: system.UserInfo{
			Name:   "root",
			UserID: 0,
		},
	}

	errMsg := "Dummy error"
	interactor := fakeUserInteractor{
		operationType: installers.Install,
		err:           errors.New(errMsg),
	}

	err := Run(detector, interactor)
	if err == nil || err.Error() != errMsg {
		t.Error("Run function result is wrong:", err.Error())
	}
}
