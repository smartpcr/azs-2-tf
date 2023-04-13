package aadb2c

import (
	aadb2c_v2021_04_01_preview "github.com/hashicorp/go-azure-sdk/resource-manager/aadb2c/2021-04-01-preview"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"

	"github.com/smartpcr/azs-2-tf/internal/azure/common"
)

func NewClient(o *common.ClientOptions) (*aadb2c_v2021_04_01_preview.Client, error) {
	client, err := aadb2c_v2021_04_01_preview.NewClientWithBaseURI(o.Environment.ResourceManager, func(c *resourcemanager.Client) {
		o.Configure(c, o.Authorizers.ResourceManager)
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
