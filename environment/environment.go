package environment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/smartpcr/azs-2-tf/utils"
)

type IEnvironment interface {
	GetName() string
	GetArmEndpoint() string
}

var (
	_                  IEnvironment = &AzureEnvironment{}
	_                  IEnvironment = &AzureStackEnvironment{}
	metadataPath                    = "metadata/endpoints"
	metadataApiVersion              = "2020-06-01"
)

type Authentication struct {
	LoginEndpoint    string                 `json:"loginEndpoint"`
	Audiences        []string               `json:"audiences"`
	Tenant           string                 `json:"tenant"`
	IdentityProvider utils.IdentityProvider `json:"identityProvider"`
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
	Name                    string                `json:"name"`
	EnvironmentType         utils.EnvironmentType `json:"environmentType"`
	ApiVersion              string                `json:"apiVersion"`
	Authentication          Authentication        `json:"authentication"`
	Profile                 string                `json:"profile"`
	Suffixes                Suffixes              `json:"suffixes"`
	Portal                  string                `json:"portal"`
	Media                   string                `json:"media"`
	GraphAudience           string                `json:"graphAudience"`
	Graph                   string                `json:"graph"`
	Batch                   string                `json:"batch"`
	ResourceManager         string                `json:"resourceManager"`
	VmImageAliasDoc         string                `json:"vmImageAliasDoc"`
	ActiveDirectoryDataLake string                `json:"activeDirectoryDataLake"`
	SqlManagement           string                `json:"sqlManagement"`
	Gallery                 string                `json:"gallery"`
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

	return &env, nil
}

func getMetadataUri(armEndpoint string) string {
	endpoint := strings.TrimRight(strings.TrimLeft(armEndpoint, "https://"), "/")
	uri := fmt.Sprintf("https://%s/%s?api-version=%s", endpoint, metadataPath, metadataApiVersion)
	return uri
}
