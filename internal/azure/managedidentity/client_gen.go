package managedidentity

import (
	"github.com/Azure/go-autorest/autorest"
	managedidentity_v2022_01_31_preview "github.com/hashicorp/go-azure-sdk/resource-manager/managedidentity/2022-01-31-preview"

	"github.com/smartpcr/azs-2-tf/internal/azure/common"
)

func NewClient(o *common.ClientOptions) (*managedidentity_v2022_01_31_preview.Client, error) {
	client := managedidentity_v2022_01_31_preview.NewClientWithBaseURI(o.ResourceManagerEndpoint, func(c *autorest.Client) {
		c.Authorizer = o.ResourceManagerAuthorizer
	})
	return &client, nil
}
