package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAzureStackEnvironment(t *testing.T) {
	azEnv, err := NewAzureStackEnvironment(
		"Northwest",
		"management.northwest.azs-longhaul-17.selfhost.corp.microsoft.com")
	if err != nil {
		t.Fatal(err)
	}

	env := azEnv.Environment
	assert.NotNil(t, env)
	assert.NotNil(t, env.Authentication)
	assert.NotNil(t, env.Authentication.Audiences)
	assert.Equal(t, 1, len(env.Authentication.Audiences))
	assert.Equal(t, "https://management.azlr.onmicrosoft.com/8cfa8f6a-6010-4fd7-b962-827f81f87a6f", env.Authentication.Audiences[0])
}
