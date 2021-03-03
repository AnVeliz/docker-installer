package docker

import (
	"reflect"
	"strings"
	"testing"

	"github.com/AnVeliz/docker-installer/utils/system"
)

type fakeCommandRunner struct {
	fullOutput []string
}

func (commandRunner *fakeCommandRunner) Run(command system.Command) {
	commandString := command.Command + " " + strings.Join(command.Arguments, " ")
	commandRunner.fullOutput = append(commandRunner.fullOutput, commandString)
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSupportedOss(t *testing.T) {
	dockerInstaller := CreateInstaller(&fakeCommandRunner{})
	operatingSystems := dockerInstaller.SupportedOs()

	expected := []system.OsInfo{
		{
			OsClass:   system.Linux,
			OsName:    "Ubuntu",
			OsVersion: "20.10",
		},
	}
	if !reflect.DeepEqual(operatingSystems, expected) {
		t.Error("Suppported operating systems list is wrong")
	}
}

func TestInstall(t *testing.T) {
	commandRunner := &fakeCommandRunner{}

	dockerInstaller := CreateInstaller(commandRunner)
	dockerInstaller.Install()

	expected := []string{
		"apt-get --yes --force-yes remove docker docker-engine docker.io containerd runc docker-ce docker-ce-cli containerd.io",
		"apt-get update",
		"apt-get --yes --force-yes install apt-transport-https ca-certificates curl gnupg-agent software-properties-common",
		"bash -c curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -",
		"bash -c -c add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
		"apt-get update",
		"apt-get --yes --force-yes install docker-ce docker-ce-cli containerd.io",
		"docker --version",
	}
	if !reflect.DeepEqual(commandRunner.fullOutput, expected) {
		t.Error(commandRunner.fullOutput)
	}
}

func TestUninstall(t *testing.T) {
	commandRunner := &fakeCommandRunner{}

	dockerInstaller := CreateInstaller(commandRunner)
	dockerInstaller.Uninstall()

	expected := []string{
		"apt-get --yes --force-yes remove docker docker-engine docker.io containerd runc docker-ce docker-ce-cli containerd.io",
	}
	if !reflect.DeepEqual(commandRunner.fullOutput, expected) {
		t.Error(commandRunner.fullOutput)
	}
}
