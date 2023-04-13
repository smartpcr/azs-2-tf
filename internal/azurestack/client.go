package azurestack

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/validation"

	"github.com/smartpcr/azs-2-tf/internal/azurestack/authorization"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/common"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/compute"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/dns"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/features"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/keyvault"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/loadbalancer"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/network"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/resource"
	"github.com/smartpcr/azs-2-tf/internal/azurestack/storage"
)

type Client struct {
	// StopContext is used for propagating control from Terraform Core (e.g. Ctrl/Cmd+C)
	StopContext context.Context

	Account       *ResourceManagerAccount
	Authorization *authorization.Client
	Compute       *compute.Client
	Dns           *dns.Client
	KeyVault      *keyvault.Client
	LoadBalancer  *loadbalancer.Client
	Network       *network.Client
	Resource      *resource.Client
	Storage       *storage.Client

	Features features.UserFeatures
}

func (client *Client) Build(ctx context.Context, o *common.ClientOptions) error {
	autorest.Count429AsRetry = false
	// Disable the Azure SDK for Go's validation since it's unhelpful for our use-case
	validation.Disabled = true

	client.StopContext = ctx

	client.Authorization = authorization.NewClient(o)
	client.Compute = compute.NewClient(o)
	client.Dns = dns.NewClient(o)
	client.KeyVault = keyvault.NewClient(o)
	client.LoadBalancer = loadbalancer.NewClient(o)
	client.Network = network.NewClient(o)
	client.Resource = resource.NewClient(o)
	client.Storage = storage.NewClient(o)

	return nil
}
