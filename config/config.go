package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/smartpcr/azs-2-tf/environment"
	"github.com/smartpcr/azs-2-tf/log"
	"github.com/smartpcr/azs-2-tf/utils"
)

type AppConfig struct {
	SubscriptionId   string                `json:"subscription_id"`
	TenantId         string                `json:"tenant_id"`
	ClientId         string                `json:"client_id"`
	ClientSecret     string                `json:"client_secret"`
	EnvironmentType  utils.EnvironmentType `json:"environment_type"`
	EnvironmentName  string                `json:"environment_name"`
	MetadataEndpoint string                `json:"metadata_endpoint"`
}

func NewAppConfig(settings utils.Settings) (*AppConfig, error) {
	configDir := settings.GetConfigFolderPath()
	cfgFile := filepath.Join(configDir, settings.GetConfigFileName())

	// read config.json file as json
	// Open the JSON file
	file, err := os.Open(cfgFile)
	if err != nil {
		log.Log.Errorf("Unable to open config file: %s", cfgFile)
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		log.Log.Errorf("Error reading file: %v", err)
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

func (config *AppConfig) save(settings utils.Settings) error {
	configDir := settings.GetConfigFolderPath()
	err := utils.EnsureDirectory(configDir)
	if err != nil {
		return err
	}

	cfgFile := filepath.Join(configDir, settings.GetConfigFileName())
	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(cfgFile, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (config *AppConfig) GetEnvironment() (*environment.Environment, error) {
	if config.EnvironmentType == utils.EnvironmentTypeAzure {
		azureEnv, err := environment.NewAzureEnvironment(config.EnvironmentName, config.MetadataEndpoint)
		if err != nil {
			log.Log.Errorf("Failed to load azure environment: %s", err)
			return nil, err
		}
		return &azureEnv.Environment, nil
	} else if config.EnvironmentType == utils.EnvironmentTypeAzureStack {
		azsEnv, err := environment.NewAzureStackEnvironment(config.EnvironmentName, config.MetadataEndpoint)
		if err != nil {
			log.Log.Errorf("Failed to load azure stack environment: %s", err)
			return nil, err
		}
		return &azsEnv.Environment, nil
	} else {
		return nil, fmt.Errorf("Unknown environment type: %s", config.EnvironmentType)
	}
}
