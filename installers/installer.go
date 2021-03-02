package installers

import (
	"github.com/AnVeliz/docker-installer/utils/system"
)

// IAppInstaller should be able install and uninstall software
type IAppInstaller interface {
	Install()
	Uninstall()
	SupportedOs() []system.OsInfo
}
