package utils

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

const (
	appName        = "azs-2-tf"
	appFolderName  = ".azs-2-tf"
	configFileName = "config.json"
	logFileName    = "azs-2-tf.log"
)

type Settings interface {
	GetAppName() string
	GetAppFolderPath() string
	GetConfigFolderPath() string
	GetConfigFileName() string
	GetLogFolderPath() string
	GetLogFileName() string
}

var _ Settings = &AppSettings{}

type AppSettings struct {
}

func (a *AppSettings) GetAppName() string {
	return appName
}

func (a *AppSettings) GetAppFolderPath() string {
	appFolderPath := filepath.Join(os.Getenv("ProgramData"), appFolderName)
	return appFolderPath
}

func (a *AppSettings) GetConfigFolderPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(home, appFolderName)
	return configDir
}

func (a *AppSettings) GetConfigFileName() string {
	return configFileName
}

func (a *AppSettings) GetLogFolderPath() string {
	appFolderPath := a.GetAppFolderPath()
	logFolder := filepath.Join(appFolderPath, "logs")
	return logFolder
}

func (a *AppSettings) GetLogFileName() string {
	return logFileName
}
