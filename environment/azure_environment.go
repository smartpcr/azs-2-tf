package environment

import "github.com/smartpcr/azs-2-tf/utils"

type AzureEnvironment struct {
	Name        string      `json:"name"`
	ArmEndpoint string      `json:"armEndpoint"`
	Environment Environment `json:"environment"`
}

func NewAzureEnvironment(name string, endpoint string) (*AzureEnvironment, error) {
	env, err := loadEnvironment(name, endpoint)
	if err != nil {
		return nil, err
	}

	return &AzureEnvironment{
		Name:        name,
		ArmEndpoint: endpoint,
		Environment: *env,
	}, err
}

func (a *AzureEnvironment) GetName() string {
	return a.Name
}

func (a *AzureEnvironment) GetArmEndpoint() string {
	return a.ArmEndpoint
}

func (a *AzureEnvironment) loadEnvironment(name string, endpoint string) (*Environment, error) {
	uri := getMetadataUri(endpoint)
	env, err := getSupportedEnvironments(name, uri)
	if err != nil {
		return nil, err
	}

	env.EnvironmentType = utils.EnvironmentTypeAzure
	env.ApiVersion = metadataApiVersion
	return env, nil
}
