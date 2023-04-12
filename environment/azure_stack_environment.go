package environment

type AzureStackEnvironment struct {
	Name        string      `json:"name"`
	ArmEndpoint string      `json:"armEndpoint"`
	Environment Environment `json:"environment"`
}

func NewAzureStackEnvironment(name string, endpoint string) *AzureStackEnvironment {
	return &AzureStackEnvironment{
		Name:        name,
		ArmEndpoint: endpoint,
		Environment: Environment{},
	}
}

func (a AzureStackEnvironment) GetName() string {
	return a.Name
}

func (a AzureStackEnvironment) GetArmEndpoint() string {
	return a.ArmEndpoint
}

func (a AzureStackEnvironment) LoadEnvironment() (*Environment, error) {
	uri := getMetadataUri(a.GetArmEndpoint())
	env, err := getSupportedEnvironments(a.Name, uri)
	if err != nil {
		return nil, err
	}
	a.Environment = *env

	return env, nil
}
