package docker

import (
	"github.com/AnVeliz/docker-installer/installers"
	"github.com/AnVeliz/docker-installer/utils"
)

// dockerInstaller is an installer of Docker application
type dockerInstaller struct {
	uninstallCommand        utils.Command
	updateRepositoryCommand []utils.Command
	installCommand          utils.Command
	checkCommand            utils.Command

	supportedOss []utils.OsInfo

	commandRunner utils.ICommandRunner
}

// CreateInstaller creates Docker installer
func CreateInstaller(commandRunner utils.ICommandRunner) installers.IAppInstaller {
	return &dockerInstaller{
		uninstallCommand: utils.Command{
			ID: "UNINSTALL",

			WelcomeMessage: "Trying to uninstall old version...",
			GoodbyMessage:  "Trying to uninstall old version... Done.",
			Command:        "apt-get",
			Arguments:      []string{"--yes", "--force-yes", "remove", "docker", "docker-engine", "docker.io", "containerd", "runc", "docker-ce", "docker-ce-cli", "containerd.io"},
		},
		updateRepositoryCommand: []utils.Command{
			utils.Command{
				ID: "UPDATE_REPOSITORY",

				WelcomeMessage: "Updating repository...",
				GoodbyMessage:  "Updating repository... Done.",
				Command:        "apt-get",
				Arguments:      []string{"update"},
			},
			utils.Command{
				ID: "INSTALL_DEPENDENCIES",

				WelcomeMessage: "Installing dependencies...",
				GoodbyMessage:  "Installing dependencies... Done.",
				Command:        "apt-get",
				Arguments:      []string{"--yes", "--force-yes", "install", "apt-transport-https", "ca-certificates", "curl", "gnupg-agent", "software-properties-common"},
			},
			utils.Command{
				ID: "ADD_KEYS_REPOSITORY",

				WelcomeMessage: "Adding keys repository...",
				GoodbyMessage:  "Adding keys repository... Done.",
				Command:        "bash",
				Arguments:      []string{"-c", "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -"},
			},
			utils.Command{
				ID: "ADD_KEYS_TO_REPOSITORY",

				WelcomeMessage: "Adding GPG keys...",
				GoodbyMessage:  "Adding GPG keys... Done.",
				Command:        "bash",
				Arguments:      []string{"-c", "-c", "add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\""},
			},
			utils.Command{
				ID: "UPDATE_REPOSITORY",

				WelcomeMessage: "Updating repository after all...",
				GoodbyMessage:  "Updating repository after all... Done.",
				Command:        "apt-get",
				Arguments:      []string{"update"},
			},
		},
		installCommand: utils.Command{
			ID: "INSTALL_DOCKER",

			WelcomeMessage: "Installing Docker...",
			GoodbyMessage:  "Installing Docker... Done.",
			Command:        "apt-get",
			Arguments:      []string{"--yes", "--force-yes", "install", "docker-ce", "docker-ce-cli", "containerd.io"},
		},
		checkCommand: utils.Command{
			ID: "CHECK_DOCKER",

			WelcomeMessage: "Checking Docker...",
			GoodbyMessage:  "Checking Docker... Done.",
			Command:        "docker",
			Arguments:      []string{"--version"},
		},

		supportedOss: []utils.OsInfo{
			utils.OsInfo{
				OsClass:   utils.Linux,
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
func (installer *dockerInstaller) SupportedOs() []utils.OsInfo {
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
