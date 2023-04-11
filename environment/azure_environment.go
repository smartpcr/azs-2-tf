package environment

import (
	"fmt"
)

type AzureEnvironment struct {
	Name        string      `json:"name"`
	Environment Environment `json:"environment"`
}

func (a *AzureEnvironment) GetName() string {
	return a.Name
}

func (a *AzureEnvironment) GetEndpoint() string {
	return fmt.Sprintf("https://%s/metadata/endpoints?api-version=2020-06-01", a.Environment.ResourceManager)
}

func (a *AzureEnvironment) LoadEnvironment() (*Environment, error) {
	env, err := getSupportedEnvironments(a.Name, a.GetEndpoint())
	if err != nil {
		return nil, err
	}
	a.Environment = *env

	return env, nil
}
