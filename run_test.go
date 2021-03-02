package main

import (
	"errors"
	"testing"

	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/utils/system"
	"github.com/AnVeliz/docker-installer/utils/system/interactivity"
)

type fakeOsDetector struct {
	fakeOsInfo   system.OsInfo
	fakeUserInfo system.UserInfo
	err          error
}

func (osDetector fakeOsDetector) GetOsInfo() (system.OsInfo, system.UserInfo, error) {
	return osDetector.fakeOsInfo, osDetector.fakeUserInfo, osDetector.err
}

type fakeUserInteractor struct {
	operationType installers.OperationType
	err           error
	io            fakeIoInteractor
}

func (interactor fakeUserInteractor) GetOperation() (installers.OperationType, error) {
	return interactor.operationType, interactor.err
}

func (interactor fakeUserInteractor) IO() interactivity.IO {
	return interactor.io
}

type fakeIoInteractor struct {
	rn  rune
	err error
}

func (fakeInteractor fakeIoInteractor) PrintMessage(message ...string) {
}

func (fakeInteractor fakeIoInteractor) GetRune() (rune, error) {
	return fakeInteractor.rn, fakeInteractor.err
}

type fakeRegistry struct {
	installer installers.IAppInstaller
}

func (registry fakeRegistry) Register(installer installers.IAppInstaller) {
}

func (registry fakeRegistry) FindInstaller(osInfo system.OsInfo) installers.IAppInstaller {
	return registry.installer
}

type dummyInstaller struct {
	osInfo []system.OsInfo
}

func (installer *dummyInstaller) Install() {}

func (installer *dummyInstaller) Uninstall() {}

func (installer *dummyInstaller) SupportedOs() []system.OsInfo {
	return installer.osInfo
}

var rightSingleOsRegistry installers.IRegistry = fakeRegistry{
	installer: &dummyInstaller{
		osInfo: []system.OsInfo{
			{
				OsClass:   system.Linux,
				OsVersion: "20.10",
				OsName:    "Ubuntu",
			},
		},
	},
}

var notFoundOsRegistry installers.IRegistry = fakeRegistry{}

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

	err := Run(rightSingleOsRegistry, detector, interactor)
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

	err := Run(rightSingleOsRegistry, detector, interactor)
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

	err := Run(notFoundOsRegistry, detector, interactor)
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

	err := Run(rightSingleOsRegistry, detector, interactor)
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

	err := Run(rightSingleOsRegistry, detector, interactor)
	if err == nil || err.Error() != "Unexpected operation requested" {
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
		err: nil,
	}

	errMsg := "Dummy error"
	interactor := fakeUserInteractor{
		operationType: installers.Install,
		err:           errors.New(errMsg),
	}

	err := Run(rightSingleOsRegistry, detector, interactor)
	if err == nil || err.Error() != errMsg {
		t.Error("Run function result is wrong:", err.Error())
	}
}

func TestUserOsDetectorError(t *testing.T) {
	errMsg := "Dummy error"
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
		err: errors.New(errMsg),
	}

	interactor := fakeUserInteractor{
		operationType: installers.Install,
		err:           nil,
	}

	err := Run(rightSingleOsRegistry, detector, interactor)
	if err == nil || err.Error() != errMsg {
		t.Error("Run function result is wrong:", err.Error())
	}
}

func TestSetupRegistry(t *testing.T) {
	registry := SetupRegistry()
	installer := registry.FindInstaller(system.OsInfo{
		OsClass:   system.Linux,
		OsVersion: "20.10",
		OsName:    "Ubuntu",
	})

	if installer == nil {
		t.Error("Registry Setup is not correct")
	}
}
