package keyvault

import (
	keyvaultmgmt "github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/keyvault/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2021-10-01/keyvault" // nolint: staticcheck

	"github.com/smartpcr/azs-2-tf/internal/azure/common"
)

type Client struct {
	ManagedHsmClient *keyvault.ManagedHsmsClient
	ManagementClient *keyvaultmgmt.BaseClient
	VaultsClient     *keyvault.VaultsClient
	options          *common.ClientOptions
}

func NewClient(o *common.ClientOptions) *Client {
	managedHsmClient := keyvault.NewManagedHsmsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&managedHsmClient.Client, o.ResourceManagerAuthorizer)

	managementClient := keyvaultmgmt.New()
	o.ConfigureClient(&managementClient.Client, o.KeyVaultAuthorizer)

	vaultsClient := keyvault.NewVaultsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vaultsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ManagedHsmClient: &managedHsmClient,
		ManagementClient: &managementClient,
		VaultsClient:     &vaultsClient,
		options:          o,
	}
}

func (client Client) KeyVaultClientForSubscription(subscriptionId string) *keyvault.VaultsClient {
	vaultsClient := keyvault.NewVaultsClientWithBaseURI(client.options.ResourceManagerEndpoint, subscriptionId)
	client.options.ConfigureClient(&vaultsClient.Client, client.options.ResourceManagerAuthorizer)
	return &vaultsClient
}
