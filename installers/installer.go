package installers

import (
	"github.com/AnVeliz/docker-installer/utils"
)

// IAppInstaller should be able install and uninstall software
type IAppInstaller interface {
	Install()
	Uninstall()
	SupportedOs() []utils.OsInfo
}
