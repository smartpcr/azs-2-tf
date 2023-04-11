package environment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAzureStackEnvironment(t *testing.T) {
	azEnv := AzureStackEnvironment{
		Name:        "Northwest",
		Environment: Environment{},
	}
	env, err := azEnv.LoadEnvironment()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, env)
	assert.Equal(t, azEnv.Name, env.Name)
}
