package environment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/smartpcr/azs-2-tf/utils"
)

type AzureStackEnvironment struct {
	Name        string      `json:"name"`
	ArmEndpoint string      `json:"armEndpoint"`
	Environment Environment `json:"environment"`
}

func NewAzureStackEnvironment(name string, endpoint string) (*AzureStackEnvironment, error) {
	env, err := loadEnvironment(name, endpoint)
	if err != nil {
		return nil, err
	}

	return &AzureStackEnvironment{
		Name:        name,
		ArmEndpoint: endpoint,
		Environment: *env,
	}, nil
}

func (a AzureStackEnvironment) GetName() string {
	return a.Name
}

func (a AzureStackEnvironment) GetArmEndpoint() string {
	return a.ArmEndpoint
}

func GetAzureStackEnvironment(ctx context.Context, name string, endpoint string) (*azure.Environment, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	uri := getMetadataUri(endpoint)
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieving environments from Azure MetaData service: %+v", err)
	}

	var environments []Environment
	if err := json.NewDecoder(resp.Body).Decode(&environments); err != nil {
		return nil, err
	}

	var env Environment
	if len(environments) == 0 {
		return nil, fmt.Errorf("no environments were returned from Azure MetaData service")
	} else if len(environments) == 1 {
		env = environments[0]
	} else {
		for _, e := range environments {
			if e.Name == name {
				env = e
			}
		}
	}
	env.ResourceManager = endpoint

	azEnv := &azure.Environment{
		Name:                      env.Name,
		ManagementPortalURL:       env.Portal,
		PublishSettingsURL:        "",
		ServiceManagementEndpoint: "https://" + env.ResourceManager,
		ResourceManagerEndpoint:   env.ResourceManager,
		ActiveDirectoryEndpoint:   env.Authentication.LoginEndpoint,
		GalleryEndpoint:           env.Gallery,
		KeyVaultEndpoint:          env.Suffixes.KeyVaultDns,
		GraphEndpoint:             env.Graph,
		StorageEndpointSuffix:     env.Suffixes.Storage,
		TokenAudience:             env.Authentication.Audiences[0],
	}

	return azEnv, nil
}

func loadEnvironment(name string, endpoint string) (*Environment, error) {
	uri := getMetadataUri(endpoint)
	env, err := getSupportedEnvironments(name, uri)
	if err != nil {
		return nil, err
	}

	if env.ResourceManager == "" {
		env.ResourceManager = endpoint
	}
	env.EnvironmentType = utils.EnvironmentTypeAzureStack
	env.ApiVersion = metadataApiVersion

	return env, nil
}
