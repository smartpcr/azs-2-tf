package authorization

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2020-04-01-preview/authorization" // nolint: staticcheck // nolint: staticcheck

	"github.com/smartpcr/azs-2-tf/internal/azure/common"
)

type Client struct {
	RoleAssignmentsClient *authorization.RoleAssignmentsClient
	RoleDefinitionsClient *authorization.RoleDefinitionsClient
}

func NewClient(o *common.ClientOptions) *Client {
	roleAssignmentsClient := authorization.NewRoleAssignmentsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&roleAssignmentsClient.Client, o.ResourceManagerAuthorizer)

	roleDefinitionsClient := authorization.NewRoleDefinitionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&roleDefinitionsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		RoleAssignmentsClient: &roleAssignmentsClient,
		RoleDefinitionsClient: &roleDefinitionsClient,
	}
}
