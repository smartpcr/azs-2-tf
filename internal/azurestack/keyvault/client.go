package keyvault

import (
	keyvaultmgmt "github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/keyvault/keyvault"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/keyvault/mgmt/keyvault"

	"github.com/smartpcr/azs-2-tf/internal/azurestack/common"
)

type Client struct {
	ManagementClient *keyvaultmgmt.BaseClient
	VaultsClient     *keyvault.VaultsClient
}

func NewClient(o *common.ClientOptions) *Client {
	managementClient := keyvaultmgmt.New()
	o.ConfigureClient(&managementClient.Client, o.KeyVaultAuthorizer)

	vaultsClient := keyvault.NewVaultsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&vaultsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ManagementClient: &managementClient,
		VaultsClient:     &vaultsClient,
	}
}
