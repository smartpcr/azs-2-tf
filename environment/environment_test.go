package environment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAzureStackEnvironment(t *testing.T) {
	azEnv := NewAzureStackEnvironment(
		"Northwest",
		"management.northwest.azs-longhaul-17.selfhost.corp.microsoft.com")
	env, err := azEnv.LoadEnvironment()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, env)
	assert.NotNil(t, env.Authentication)
	assert.NotNil(t, env.Authentication.Audiences)
	assert.Equal(t, 1, len(env.Authentication.Audiences))
	assert.Equal(t, "https://management.azlr.onmicrosoft.com/8cfa8f6a-6010-4fd7-b962-827f81f87a6f", env.Authentication.Audiences[0])
}
