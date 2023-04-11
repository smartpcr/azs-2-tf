package config

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
)

type AppConfig struct {
	SubscriptionId        string `json:"subscription_id"`
	TenantId              string `json:"tenant_id"`
	ClientId              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	AzureStackEnvironment string `json:"azure_stack_environment"`
	AzureStackArmEndpoint string `json:"azure_stack_arm_endpoint"`
}

func NewAppConfig() (*AppConfig, error) {
	// specify config.json file
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	configDir := filepath.Join(home, AppFolderName)
	cfgFile := filepath.Join(configDir, ConfigFileName)

	// read config.json file as json
	// Open the JSON file
	file, err := os.Open(cfgFile)
	if err != nil {
		logrus.Errorln("Unable to open config file: %s", cfgFile)
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		logrus.Errorln("Error reading file:", err)
		return nil, err
	}

	// Unmarshal JSON content into a Person struct
	appConfig := &AppConfig{}
	err = json.Unmarshal(content, &appConfig)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return appConfig, nil
}

func (appConfig *AppConfig) LoadEnvironment() (*azure.Environment, error) {
	// load azure stack environment
	env, err := azure.EnvironmentFromName(appConfig.AzureStackEnvironment)
	if err != nil {
		logrus.Errorln("Error loading azure stack environment: %s", appConfig.AzureStackEnvironment)
		return nil, err
	}

	// set azure stack arm endpoint
	env.ResourceManagerEndpoint = appConfig.AzureStackArmEndpoint

	return &env, nil
}
