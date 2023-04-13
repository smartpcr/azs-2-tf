package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Azure/go-autorest/autorest/adal"

	"github.com/smartpcr/azs-2-tf/environment"
	"github.com/smartpcr/azs-2-tf/log"
	"github.com/smartpcr/azs-2-tf/utils"
)

var logger = log.Log

type AppConfig struct {
	SubscriptionId   string                `json:"subscription_id"`
	TenantId         string                `json:"tenant_id"`
	ClientId         string                `json:"client_id"`
	EnvironmentType  utils.EnvironmentType `json:"environment_type"`
	EnvironmentName  string                `json:"environment_name"`
	MetadataEndpoint string                `json:"metadata_endpoint"`

	// Auxiliary tenant IDs used for multi tenant auth
	SupportsAuxiliaryTenants bool     `json:"supports_auxiliary_tenants"`
	AuxiliaryTenantIds       []string `json:"auxiliary_tenant_ids"`

	// The custom Resource Manager Endpoint which should be used
	// only applicable for Azure Stack at this time.
	CustomResourceManagerEndpoint string `json:"custom_resource_manager_endpoint"`

	// Azure CLI Tokens Auth
	SupportsAzureCliToken bool `json:"supports_azure_cli_token"`

	// Managed Service Identity Auth
	SupportsManagedServiceIdentity bool   `json:"supports_managed_service_identity"`
	MsiEndpoint                    string `json:"msi_endpoint"`

	// Service Principal (Client Cert) Auth
	SupportsClientCertAuth bool   `json:"supports_client_cert_auth"`
	ClientCertPath         string `json:"client_cert_path"`
	ClientCertPassword     string `json:"client_cert_password"`

	// Service Principal (Client Secret) Auth
	SupportsClientSecretAuth bool   `json:"supports_client_secret_auth"`
	ClientSecret             string `json:"client_secret"`
	ClientSecretDocsLink     string `json:"client_secret_docs_link"`

	// OIDC Auth
	SupportsOIDCAuth    bool   `json:"supports_oidc_auth"`
	IDTokenRequestURL   string `json:"id_token_request_url"`
	IDTokenRequestToken string `json:"id_token_request_token"`

	// Beta opt-in for Microsoft Graph
	UseMicrosoftGraph bool `json:"use_microsoft_graph"`

	SkipProviderRegistration    bool   `json:"skip_provider_registration"`
	TerraformVersion            string `json:"terraform_version"`
	DisableCorrelationRequestID bool   `json:"disable_correlation_request_id"`
}

type OAuthConfig struct {
	OAuth            *adal.OAuthConfig
	MultiTenantOauth *adal.MultiTenantOAuthConfig
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
