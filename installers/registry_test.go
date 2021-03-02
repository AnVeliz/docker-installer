package installers

import (
	"reflect"
	"testing"

	"github.com/AnVeliz/docker-installer/utils/system"
)

type dummyInstaller struct {
	osInfo []system.OsInfo
}

func (installer *dummyInstaller) Install() {}

func (installer *dummyInstaller) Uninstall() {}

func (installer *dummyInstaller) SupportedOs() []system.OsInfo {
	return installer.osInfo
}

func TestRegistryAndFind(t *testing.T) {
	dummyOsVersion := system.OsInfo{
		OsClass:   system.Linux,
		OsName:    "Ubuntu",
		OsVersion: "21.10",
	}

	dummyInstaller := &dummyInstaller{
		osInfo: []system.OsInfo{
			dummyOsVersion,
		},
	}

	registry := CreateRegistry()
	registry.Register(dummyInstaller)

	installer := registry.FindInstaller(dummyOsVersion)

	if !reflect.DeepEqual(installer, dummyInstaller) {
		t.Error("Installer registry is corrupted")
	}
}

func TestRegistryAndFindDifferentVersion(t *testing.T) {
	dummyInstaller := &dummyInstaller{
		osInfo: []system.OsInfo{
			{
				OsClass:   system.Linux,
				OsName:    "Ubuntu",
				OsVersion: "21.10",
			},
		},
	}

	registry := CreateRegistry()
	registry.Register(dummyInstaller)

	dummyOsVersion := system.OsInfo{
		OsClass:   system.Linux,
		OsName:    "Ubuntu",
		OsVersion: "10.00",
	}

	installer := registry.FindInstaller(dummyOsVersion)

	if installer != nil {
		t.Error("Installer registry is corrupted")
	}
}
