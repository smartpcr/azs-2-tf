package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
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

func NewClientBuilder(credential azcore.TokenCredential, opt arm.ClientOptions, subscriptionId string, apiVersion string) *ClientBuilder {
	return &ClientBuilder{
		Credential:     credential,
		Opt:            opt,
		SubscriptionId: subscriptionId,
		ApiVersion:     apiVersion,
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
