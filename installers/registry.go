package installers

import "github.com/AnVeliz/docker-installer/utils/system"

// OperationType is an alias for operation types
type OperationType int

const (
	// NotSpecified means that it hasn't been set
	NotSpecified OperationType = iota
	// Install for installation
	Install
	// Uninstall for deinstallation
	Uninstall
)

// IRegistry is an interface of registry
type IRegistry interface {
	Register(installer IAppInstaller)
	FindInstaller(osInfo system.OsInfo) IAppInstaller
}

// Registry contains all installers
type Registry struct {
	IRegistry

	installers []IAppInstaller
}

// Register an installer in the installers repository
func (registry *Registry) Register(installer IAppInstaller) {
	registry.installers = append(registry.installers, installer)
}

// FindInstaller looks up for an appropriate installer
func (registry *Registry) FindInstaller(osInfo system.OsInfo) IAppInstaller {
	for _, currentInstaller := range registry.installers {
		currentInstallerSupportedOs := currentInstaller.SupportedOs()
		for _, currentSupportedOs := range currentInstallerSupportedOs {
			if currentSupportedOs == osInfo {
				return currentInstaller
			}
		}
	}

	return nil
}

// CreateRegistry function is a constructor
func CreateRegistry() *Registry {
	return &Registry{
		installers: make([]IAppInstaller, 0),
	}
}
