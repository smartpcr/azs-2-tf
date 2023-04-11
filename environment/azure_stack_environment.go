package environment

import (
	"fmt"
	"strings"
)

type AzureStackEnvironment struct {
	Name        string      `json:"name"`
	Environment Environment `json:"environment"`
}

func (a AzureStackEnvironment) GetName() string {
	return a.Name
}

func (a AzureStackEnvironment) GetEndpoint() string {
	return fmt.Sprintf("%s/metadata/endpoints?api-version=2020-06-01", strings.TrimRight(a.GetEndpoint(), "/"))
}

func (a AzureStackEnvironment) LoadEnvironment() (*Environment, error) {
	env, err := getSupportedEnvironments(a.Name, a.GetEndpoint())
	if err != nil {
		return nil, err
	}
	a.Environment = *env

	return env, nil
}
