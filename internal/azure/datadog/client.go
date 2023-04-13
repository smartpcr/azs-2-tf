package datadog

import (
	"github.com/Azure/go-autorest/autorest"
	datadog_v2021_03_01 "github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2021-03-01"

	"github.com/smartpcr/azs-2-tf/internal/azure/common"
)

func NewClient(o *common.ClientOptions) *datadog_v2021_03_01.Client {
	client := datadog_v2021_03_01.NewClientWithBaseURI(o.ResourceManagerEndpoint, func(c *autorest.Client) {
		c.Authorizer = o.ResourceManagerAuthorizer
	})
	return &client
}
