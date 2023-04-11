package config

import (
	"encoding/json"
	"fmt"
	"github.com/smartpcr/azs-2-tf/log"
	"github.com/smartpcr/azs-2-tf/utils"
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
