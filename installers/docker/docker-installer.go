package docker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/utils"
)

// IInstaller is an interface for docker installer
type IInstaller interface {
	installers.IAppInstaller
}

// Installer is an installer of Docker application
type Installer struct {
	IInstaller
}

// Install Docker
func (installer *Installer) Install() {
	installer.Uninstall()
	updateRepository()
	installDocker()
	checkDocker()
}

// Uninstall Docker
func (installer *Installer) Uninstall() {
	uninstall()
}

// SupportedOs Docker
func (installer *Installer) SupportedOs() []utils.OsInfo {
	return []utils.OsInfo{
		{
			OsClass:   utils.Linux,
			OsName:    "Ubuntu",
			OsVersion: "20.10",
		},
	}
}

// CreateInstaller creates Docker installer
func CreateInstaller() IInstaller {
	return &Installer{}
}

func uninstall() {
	fmt.Println("Trying to uninstall previous version...")
	// sudo apt-get remove docker docker-engine docker.io containerd runc
	uninstallOut, uninstallError := exec.Command("apt-get", "--yes", "--force-yes", "remove", "docker", "docker-engine", "docker.io", "containerd", "runc", "docker-ce", "docker-ce-cli", "containerd.io").Output()

	fmt.Println("-----\n", strings.TrimRight(strings.TrimLeft(string(uninstallOut), "\t \n"), "\t \n"), "\n-----")
	fmt.Println("Trying to uninstall previous version... Done")

	if uninstallError != nil {
		fmt.Println("Error when trying to uninstall odl version; ", uninstallError)
	}
}

func updateRepository() {
	reloadRepository()
	installDependencies()
	installGpgkey()
	addDockerRepository()
	reloadRepository()
}

func addDockerRepository() {
	// sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
	fmt.Println("Trying to add Docker repository...")
	addDockerRepositoryOut, addDockerRepositoryError := exec.Command("bash", "-c", "add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"").Output()
	fmt.Println(string(addDockerRepositoryOut))

	if addDockerRepositoryError != nil {
		fmt.Println("Error trying to add Docker repository; ", addDockerRepositoryError)
	}
	fmt.Println("Trying to add Docker repository... Done")
}

func installGpgkey() {
	// curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
	fmt.Println("Trying to get and install gpg key...")
	installGpgKeyOut, installGpgKeyError := exec.Command("bash", "-c", "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -").Output()
	fmt.Println(string(installGpgKeyOut))

	if installGpgKeyError != nil {
		fmt.Println("Error when trying to get and install gpg key; ", installGpgKeyError)
	}
	fmt.Println("Trying to get and install gpg key... Done")
}

func installDependencies() {
	// sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
	fmt.Println("Trying to install dependencies...")
	installDependenciesOut, installDependenciesError := exec.Command("apt-get", "--yes", "--force-yes", "install", "apt-transport-https", "ca-certificates", "curl", "gnupg-agent", "software-properties-common").Output()
	fmt.Println(string(installDependenciesOut))

	if installDependenciesError != nil {
		fmt.Println("Error when trying to install dependencies; ", installDependenciesError)
	}
	fmt.Println("Trying to install dependencies... Done")
}

func reloadRepository() {
	// sudo apt-get update
	fmt.Println("Trying to reload the repository...")
	reloadOut, reloadError := exec.Command("apt-get", "update").Output()
	fmt.Println(string(reloadOut))

	if reloadError != nil {
		fmt.Println("Error when trying to reload odl version; ", reloadError)
	}
	fmt.Println("Trying to reload the repository... Done")
}

func installDocker() {
	// sudo apt-get install docker-ce docker-ce-cli containerd.io
	fmt.Println("Trying to install Docker...")
	installationOut, installationError := exec.Command("apt-get", "--yes", "--force-yes", "install", "docker-ce", "docker-ce-cli", "containerd.io").Output()
	fmt.Println(string(installationOut))

	if installationError != nil {
		fmt.Println("Error trying to install Docker; ", installationError)
	}
	fmt.Println("Trying to install Docker... Done")
}

func checkDocker() {
	// sudo docker version
	fmt.Println("Trying Docker...")
	tryingOut, tryingError := exec.Command("docker", "--version").Output()
	fmt.Println(string(tryingOut))

	if tryingError != nil {
		fmt.Println("Error trying Docker; ", tryingError)
	}
	fmt.Println("Trying tDocker... Done")
}
