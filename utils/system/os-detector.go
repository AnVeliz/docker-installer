package system

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// OsClass type
type OsClass string

// Supported operating systems
const (
	NotSupported OsClass = ""
	Linux        OsClass = "Linux"
)

// OsInfo is a storage of operating system info
type OsInfo struct {
	OsClass   OsClass
	OsName    string
	OsVersion string
}

// UserInfo provides a bit of information about current user
type UserInfo struct {
	Name   string
	UserID int
}

// IOsDetector is an interface for Os detection system
type IOsDetector interface {
	GetOsInfo() (OsInfo, UserInfo, error)
}

type osDetector struct{}

// NewOsDetector is a factory for osDetector
func NewOsDetector() IOsDetector {
	return &osDetector{}
}

// GetOsInfo returns operating system info
func (osDetector osDetector) GetOsInfo() (OsInfo, UserInfo, error) {
	switch runtime.GOOS {
	case "linux":
		osName, osVersion, versionErr := getLinuxVersion()
		if versionErr != nil {
			return OsInfo{}, UserInfo{}, versionErr
		}

		userName, userID, procOwnerErr := getProcessOwner()
		if procOwnerErr != nil {
			return OsInfo{}, UserInfo{}, procOwnerErr
		}

		return OsInfo{
				OsClass:   Linux,
				OsName:    osName,
				OsVersion: osVersion,
			}, UserInfo{
				Name:   userName,
				UserID: userID,
			},
			nil
	default:
		return OsInfo{
			OsClass: NotSupported,
		}, UserInfo{}, nil
	}
}

func getLinuxVersion() (string, string, error) {
	nameOut, nameError := exec.Command("lsb_release", "-si").Output()
	versionOut, versionError := exec.Command("lsb_release", "-sr").Output()
	if nameError != nil || versionError != nil {
		return "", "", errors.New("Error during execution lsb_release")
	}

	return strings.Trim(string(nameOut), "\t \n"), strings.Trim(string(versionOut), "\t \n"), nil
}

func getProcessOwner() (string, int, error) {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		return "", -1, err
	}
	return string(stdout), os.Getuid(), nil
}
