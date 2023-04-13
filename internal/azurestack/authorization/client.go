package authorization

import (
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"

	"github.com/smartpcr/azs-2-tf/internal/azurestack/common"
)

type Client struct {
	ServicePrincipalsClient *graphrbac.ServicePrincipalsClient
}

func NewClient(o *common.ClientOptions) *Client {
	servicePrincipalsClient := graphrbac.NewServicePrincipalsClientWithBaseURI(o.GraphEndpoint, o.TenantID)
	o.ConfigureClient(&servicePrincipalsClient.Client, o.GraphAuthorizer)

	return &Client{
		ServicePrincipalsClient: &servicePrincipalsClient,
	}
}
