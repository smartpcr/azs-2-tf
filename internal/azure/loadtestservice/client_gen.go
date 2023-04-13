package loadtestservice

import (
	"github.com/Azure/go-autorest/autorest"
	loadtestservice_v2021_12_01_preview "github.com/hashicorp/go-azure-sdk/resource-manager/loadtestservice/2021-12-01-preview"

	"github.com/smartpcr/azs-2-tf/internal/azure/common"
)

func NewClient(o *common.ClientOptions) (*loadtestservice_v2021_12_01_preview.Client, error) {
	client := loadtestservice_v2021_12_01_preview.NewClientWithBaseURI(o.ResourceManagerEndpoint, func(c *autorest.Client) {
		c.Authorizer = o.ResourceManagerAuthorizer
	})
	return &client, nil
}
