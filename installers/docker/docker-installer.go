package docker

import (
	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/utils/system"
)

// dockerInstaller is an installer of Docker application
type dockerInstaller struct {
	uninstallCommand        system.Command
	updateRepositoryCommand []system.Command
	installCommand          system.Command
	checkCommand            system.Command

	supportedOss []system.OsInfo

	commandRunner system.ICommandRunner
}

// CreateInstaller creates Docker installer
func CreateInstaller(commandRunner system.ICommandRunner) installers.IAppInstaller {
	return &dockerInstaller{
		uninstallCommand: system.Command{
			ID: "UNINSTALL",

			WelcomeMessage: "Trying to uninstall old version...",
			GoodbyeMessage: "Trying to uninstall old version... Done.",
			Command:        "apt-get",
			Arguments:      []string{"--yes", "--force-yes", "remove", "docker", "docker-engine", "docker.io", "containerd", "runc", "docker-ce", "docker-ce-cli", "containerd.io"},
		},
		updateRepositoryCommand: []system.Command{
			{
				ID: "UPDATE_REPOSITORY",

				WelcomeMessage: "Updating repository...",
				GoodbyeMessage: "Updating repository... Done.",
				Command:        "apt-get",
				Arguments:      []string{"update"},
			},
			{
				ID: "INSTALL_DEPENDENCIES",

				WelcomeMessage: "Installing dependencies...",
				GoodbyeMessage: "Installing dependencies... Done.",
				Command:        "apt-get",
				Arguments:      []string{"--yes", "--force-yes", "install", "apt-transport-https", "ca-certificates", "curl", "gnupg-agent", "software-properties-common"},
			},
			{
				ID: "ADD_KEYS_REPOSITORY",

				WelcomeMessage: "Adding keys repository...",
				GoodbyeMessage: "Adding keys repository... Done.",
				Command:        "bash",
				Arguments:      []string{"-c", "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -"},
			},
			{
				ID: "ADD_KEYS_TO_REPOSITORY",

				WelcomeMessage: "Adding GPG keys...",
				GoodbyeMessage: "Adding GPG keys... Done.",
				Command:        "bash",
				Arguments:      []string{"-c", "-c", "add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\""},
			},
			{
				ID: "UPDATE_REPOSITORY",

				WelcomeMessage: "Updating repository after all...",
				GoodbyeMessage: "Updating repository after all... Done.",
				Command:        "apt-get",
				Arguments:      []string{"update"},
			},
		},
		installCommand: system.Command{
			ID: "INSTALL_DOCKER",

			WelcomeMessage: "Installing Docker...",
			GoodbyeMessage: "Installing Docker... Done.",
			Command:        "apt-get",
			Arguments:      []string{"--yes", "--force-yes", "install", "docker-ce", "docker-ce-cli", "containerd.io"},
		},
		checkCommand: system.Command{
			ID: "CHECK_DOCKER",

			WelcomeMessage: "Checking Docker...",
			GoodbyeMessage: "Checking Docker... Done.",
			Command:        "docker",
			Arguments:      []string{"--version"},
		},

		supportedOss: []system.OsInfo{
			{
				OsClass:   system.Linux,
				OsName:    "Ubuntu",
				OsVersion: "20.10",
			},
		},

		commandRunner: commandRunner,
	}
}

// Install Docker
func (installer *dockerInstaller) Install() {
	installer.uninstall()
	installer.updateRepository()
	installer.installDocker()
	installer.checkDocker()
}

// Uninstall Docker
func (installer *dockerInstaller) Uninstall() {
	installer.uninstall()
}

// SupportedOs Docker
func (installer *dockerInstaller) SupportedOs() []system.OsInfo {
	return installer.supportedOss
}

func (installer *dockerInstaller) uninstall() {
	installer.commandRunner.Run(installer.uninstallCommand)
}

func (installer *dockerInstaller) updateRepository() {
	for _, value := range installer.updateRepositoryCommand {
		installer.commandRunner.Run(value)
	}
}

func (installer *dockerInstaller) installDocker() {
	installer.commandRunner.Run(installer.installCommand)
}

func (installer *dockerInstaller) checkDocker() {
	installer.commandRunner.Run(installer.checkCommand)
}
