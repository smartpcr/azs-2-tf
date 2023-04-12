package config

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/smartpcr/azs-2-tf/utils"
	"github.com/smartpcr/azs-2-tf/utils/mocks"
)

var (
	settings = mocks.NewMockSettings()
	config   = &AppConfig{
		SubscriptionId:   "95ea6fca-2c18-4634-9ce5-e056eebdd92b",
		TenantId:         "00de9b88-26db-443e-bfae-f884ddfe2e8a",
		ClientId:         "d6f471ca-b164-4567-8387-4a51d782c552",
		ClientSecret:     "",
		EnvironmentType:  utils.EnvironmentTypeAzureStack,
		EnvironmentName:  "Northwest",
		MetadataEndpoint: "https://management.northwest.azs-longhaul-17.selfhost.corp.microsoft.com",
	}
)

func TestCreateAppConfig(t *testing.T) {
	err := config.save(settings)
	if err != nil {
		t.Fatal(err)
	}

	config, err := NewAppConfig(settings)
	if err != nil {
		t.Fatal(err)
	}

	if config == nil {
		t.Fatal("Failed to create AppConfig")
	}

	if config.EnvironmentName != "AzureCloud" && config.MetadataEndpoint == "" {
		t.Fatal("ARM endpoint is required for AzureStack environment")
	}
}

func TestReadConfigUsingViper(t *testing.T) {
	err := config.save(settings)
	if err != nil {
		t.Fatal(err)
	}

	configFile := filepath.Join(settings.GetConfigFolderPath(), settings.GetConfigFileName())
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Error reading config file: %s", err)
	}

	rawJson := make(map[string]interface{})
	for key, value := range viper.AllSettings() {
		rawJson[key] = value
	}
	jsonBytes, err := json.Marshal(rawJson)
	if err != nil {
		t.Fatalf("Error marshalling rawJson: %s", err)
	}

	var config AppConfig
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		t.Fatalf("Error unmarshalling jsonBytes into struct: %s", err)
	}

	assert.Equal(t, "00de9b88-26db-443e-bfae-f884ddfe2e8a", config.TenantId)
}
