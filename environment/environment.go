package environment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type IEnvironment interface {
	GetName() string
	GetEndpoint() string
	LoadEnvironment() (*Environment, error)
}

var (
	_ IEnvironment = &AzureEnvironment{}
	_ IEnvironment = &AzureStackEnvironment{}
)

type IdentityProvider string

const (
	IdentityProviderAAD  IdentityProvider = "aad"
	IdentityProviderADFS IdentityProvider = "adfs"
)

type Authentication struct {
	LoginEndpoint    string           `json:"loginEndpoint"`
	Audiences        []string         `json:"audiences"`
	Tenant           string           `json:"tenant"`
	IdentityProvider IdentityProvider `json:"identityProvider"`
}

type Suffixes struct {
	AzureDataLakeStoreFileSystem        string `json:"azureDataLakeStoreFileSystem"`
	AcrLoginServer                      string `json:"acrLoginServer"`
	SqlServerHostname                   string `json:"sqlServerHostname"`
	AzureDataLakeAnalyticsCatalogAndJob string `json:"azureDataLakeAnalyticsCatalogAndJob"`
	KeyVaultDns                         string `json:"keyVaultDns"`
	Storage                             string `json:"storage"`
	AzureFrontDoorEndpointSuffix        string `json:"azureFrontDoorEndpointSuffix"`
}

type Environment struct {
	Name                    string         `json:"name"`
	Authentication          Authentication `json:"authentication"`
	Profile                 string         `json:"profile"`
	Suffixes                Suffixes       `json:"suffixes"`
	Portal                  string         `json:"portal"`
	Media                   string         `json:"media"`
	GraphAudience           string         `json:"graphAudience"`
	Graph                   string         `json:"graph"`
	Batch                   string         `json:"batch"`
	ResourceManager         string         `json:"resourceManager"`
	VmImageAliasDoc         string         `json:"vmImageAliasDoc"`
	ActiveDirectoryDataLake string         `json:"activeDirectoryDataLake"`
	SqlManagement           string         `json:"sqlManagement"`
	Gallery                 string         `json:"gallery"`
}

func getSupportedEnvironments(name string, uri string) (*Environment, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	req, err := http.NewRequestWithContext(context.Background(), "GET", uri, nil)
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
	for _, e := range environments {
		if e.Name == name {
			env = e
		}
	}

	return &env, nil
}
