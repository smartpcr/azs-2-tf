package environment

type AzureEnvironment struct {
	Name        string      `json:"name"`
	ArmEndpoint string      `json:"armEndpoint"`
	Environment Environment `json:"environment"`
}

func NewAzureEnvironment(name string, endpoint string) *AzureEnvironment {
	return &AzureEnvironment{
		Name:        name,
		ArmEndpoint: endpoint,
		Environment: Environment{},
	}
}

func (a *AzureEnvironment) GetName() string {
	return a.Name
}

func (a *AzureEnvironment) GetArmEndpoint() string {
	return a.ArmEndpoint
}

func (a *AzureEnvironment) LoadEnvironment() (*Environment, error) {
	uri := getMetadataUri(a.GetArmEndpoint())
	env, err := getSupportedEnvironments(a.Name, uri)
	if err != nil {
		return nil, err
	}
	a.Environment = *env

	return env, nil
}
