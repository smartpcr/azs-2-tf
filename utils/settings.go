package utils

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type IdentityProvider string

const (
	IdentityProviderAAD  IdentityProvider = "aad"
	IdentityProviderADFS IdentityProvider = "adfs"
)

type EnvironmentType string

const (
	EnvironmentTypeAzureStack EnvironmentType = "AzureStack"
	EnvironmentTypeAzure      EnvironmentType = "Azure"
)

const (
	appName              string = "azs-2-tf"
	appFolderName        string = ".azs-2-tf"
	configFileName       string = "config.json"
	logFileName          string = "azs-2-tf.log"
	Terraform_Env_Prefix string = "ARM"
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
