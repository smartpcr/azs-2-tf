package client

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"github.com/smartpcr/azs-2-tf/config"
	"github.com/smartpcr/azs-2-tf/environment"
	"github.com/smartpcr/azs-2-tf/log"
	"github.com/smartpcr/azs-2-tf/utils"
)

var (
	_ Client = &ClientBuilder{}
)

type Client interface {
	NewKeyVaultSecretsClient() (*armkeyvault.SecretsClient, error)
	NewResourceGroupsClient() (*armresources.ResourceGroupsClient, error)
	NewResourcesClient() (*armresources.Client, error)
}

type ClientBuilder struct {
	Credential     azcore.TokenCredential
	Opt            arm.ClientOptions
	SubscriptionId string
	ApiVersion     string
}

func NewClientBuilder(appConfig *config.AppConfig, env *environment.Environment, appSettings utils.Settings) *ClientBuilder {
	var cloudCfg cloud.Configuration
	if env.EnvironmentType == utils.EnvironmentTypeAzureStack {
		cloudCfg = cloud.Configuration{
			ActiveDirectoryAuthorityHost: env.Authentication.LoginEndpoint,
			Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
				cloud.ResourceManager: {
					Endpoint: env.ResourceManager,
					Audience: env.Authentication.Audiences[0],
				},
			},
		}
	} else {
		cloudCfg = cloud.AzurePublic
	}

	clientOpt := arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			//APIVersion: env.ApiVersion, // cannot force api version for ARM client
			Cloud: cloudCfg,
			Telemetry: policy.TelemetryOptions{
				ApplicationID: appSettings.GetAppName(),
				Disabled:      false,
			},
			Logging: policy.LogOptions{
				IncludeBody: true,
			},
		},
	}

	// default credential read from environment variables
	_ = os.Setenv("AZURE_TENANT_ID", appConfig.TenantId)
	_ = os.Setenv("AZURE_CLIENT_ID", appConfig.ClientId)
	_ = os.Setenv("AZURE_CLIENT_SECRET", appConfig.ClientSecret)
	_ = os.Setenv("AZURE_SUBSCRIPTION_ID", appConfig.SubscriptionId)
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		ClientOptions: clientOpt.ClientOptions,
		TenantID:      appConfig.TenantId,
	})
	if err != nil {
		log.Log.Errorf("Failed to obtain credential for %s: %s", env.Name, err)
		os.Exit(1)
	}

	return &ClientBuilder{
		Credential:     cred,
		Opt:            clientOpt,
		SubscriptionId: appConfig.SubscriptionId,
		ApiVersion:     env.ApiVersion,
	}
}

func (c *ClientBuilder) NewKeyVaultSecretsClient() (*armkeyvault.SecretsClient, error) {
	return armkeyvault.NewSecretsClient(
		c.SubscriptionId,
		c.Credential,
		&c.Opt,
	)
}

func (c *ClientBuilder) NewResourceGroupsClient() (*armresources.ResourceGroupsClient, error) {
	client, err := armresources.NewResourceGroupsClient(
		c.SubscriptionId,
		c.Credential,
		&c.Opt,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientBuilder) NewResourcesClient() (*armresources.Client, error) {
	client, err := armresources.NewClient(
		c.SubscriptionId,
		c.Credential,
		&c.Opt,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
