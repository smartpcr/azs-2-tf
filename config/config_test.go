package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAppConfig(t *testing.T) {
	config, err := NewAppConfig()

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

func TestCreateAppConfigFromEnvironment(t *testing.T) {
	config, err := NewAppConfig()
	if err != nil {
		t.Fatal(err)
	}

	env, err := config.LoadEnvironment()
	if err != nil {
		t.Fatal(err)
	}

	if env == nil {
		t.Fatal("Failed to load Environment")
	}

	assert.Equal(t, env.Name, config.AzureStackEnvironment)
}
