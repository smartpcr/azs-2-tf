package environment

import "github.com/Azure/go-autorest/autorest/azure"

type IEnvironment interface {
	LoadEnvironment() (*azure.Environment, error)
}

var _ IEnvironment = &Environment{}

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

func (e Environment) LoadEnvironment() (*azure.Environment, error) {
	//TODO implement me
	panic("implement me")
}
