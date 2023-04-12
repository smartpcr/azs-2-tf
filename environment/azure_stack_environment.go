package environment

import "github.com/smartpcr/azs-2-tf/utils"

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
