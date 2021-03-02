package system

import (
	"fmt"
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

// GetOsInfo returns operating system info
func GetOsInfo() (OsInfo, UserInfo) {
	switch runtime.GOOS {
	case "linux":
		osName, osVersion := getLinuxVersion()
		userName, userID := getProcessOwner()
		return OsInfo{
				OsClass:   Linux,
				OsName:    osName,
				OsVersion: osVersion,
			}, UserInfo{
				Name:   userName,
				UserID: userID,
			}
	default:
		return OsInfo{
			OsClass: NotSupported,
		}, UserInfo{}
	}
}

func getLinuxVersion() (string, string) {
	nameOut, nameError := exec.Command("lsb_release", "-si").Output()
	versionOut, versionError := exec.Command("lsb_release", "-sr").Output()
	if nameError != nil || versionError != nil {
		fmt.Println("Error during execution lsb_release")
	}

	return strings.Trim(string(nameOut), "\t \n"), strings.Trim(string(versionOut), "\t \n")
}

func getProcessOwner() (string, int) {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	return string(stdout), os.Getuid()
}
