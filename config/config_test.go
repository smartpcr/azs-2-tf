package config

import (
	"github.com/smartpcr/azs-2-tf/utils/mocks"
	"testing"
)

var (
	settings = mocks.NewMockSettings()
	config   = &AppConfig{
		SubscriptionId:        "95ea6fca-2c18-4634-9ce5-e056eebdd92b",
		TenantId:              "00de9b88-26db-443e-bfae-f884ddfe2e8a",
		ClientId:              "d6f471ca-b164-4567-8387-4a51d782c552",
		ClientSecret:          "",
		AzureStackEnvironment: "Northwest",
		AzureStackArmEndpoint: "https://management.northwest.azs-longhaul-17.selfhost.corp.microsoft.com",
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

	if config.AzureStackEnvironment != "AzureCloud" && config.AzureStackArmEndpoint == "" {
		t.Fatal("ARM endpoint is required for AzureStack environment")
	}
}
