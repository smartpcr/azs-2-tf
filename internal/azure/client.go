package azure

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/validation"
	aadb2c_v2021_04_01_preview "github.com/hashicorp/go-azure-sdk/resource-manager/aadb2c/2021-04-01-preview"
	analysisservices_v2017_08_01 "github.com/hashicorp/go-azure-sdk/resource-manager/analysisservices/2017-08-01"
	azurestackhci_v2022_12_01 "github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2022-12-01"
	datadog_v2021_03_01 "github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2021-03-01"
	dns_v2018_05_01 "github.com/hashicorp/go-azure-sdk/resource-manager/dns/2018-05-01"

	"github.com/smartpcr/azs-2-tf/internal/azure/azurestackhci"
	"github.com/smartpcr/azs-2-tf/internal/azure/databoxedge"
	"github.com/smartpcr/azs-2-tf/internal/azure/datadog"
	"github.com/smartpcr/azs-2-tf/internal/azure/dns"
	"github.com/smartpcr/azs-2-tf/internal/azure/keyvault"
	"github.com/smartpcr/azs-2-tf/internal/azure/network"
	"github.com/smartpcr/azs-2-tf/internal/azure/resource"

	"github.com/smartpcr/azs-2-tf/internal/azure/aadb2c"
	"github.com/smartpcr/azs-2-tf/internal/azure/advisor"
	"github.com/smartpcr/azs-2-tf/internal/azure/analysisServices"
	apiManagement "github.com/smartpcr/azs-2-tf/internal/azure/apimanagement"
	"github.com/smartpcr/azs-2-tf/internal/azure/authorization"
	"github.com/smartpcr/azs-2-tf/internal/azure/common"
	"github.com/smartpcr/azs-2-tf/internal/azure/compute"
	"github.com/smartpcr/azs-2-tf/internal/azure/features"
)

type Client struct {
	autoClient

	// StopContext is used for propagating control from Terraform Core (e.g. Ctrl/Cmd+C)
	StopContext context.Context

	Account  *ResourceManagerAccount
	Features features.UserFeatures

	AadB2c           *aadb2c_v2021_04_01_preview.Client
	Advisor          *advisor.Client
	AnalysisServices *analysisservices_v2017_08_01.Client
	ApiManagement    *apiManagement.Client
	Authorization    *authorization.Client
	AzureStackHCI    *azurestackhci_v2022_12_01.Client
	Compute          *compute.Client
	DataboxEdge      *databoxedge.Client
	Datadog          *datadog_v2021_03_01.Client
	Dns              *dns_v2018_05_01.Client
	KeyVault         *keyvault.Client
	Network          *network.Client
	Resource         *resource.Client
}

// NOTE: it should be possible for this method to become Private once the top level Client's removed

func (client *Client) Build(ctx context.Context, o *common.ClientOptions) error {
	autorest.Count429AsRetry = false
	// Disable the Azure SDK for Go's validation since it's unhelpful for our use-case
	validation.Disabled = true

	if err := buildAutoClients(&client.autoClient, o); err != nil {
		return fmt.Errorf("building auto-clients: %+v", err)
	}

	client.Features = o.Features
	client.StopContext = ctx

	var err error

	if client.AadB2c, err = aadb2c.NewClient(o); err != nil {
		return fmt.Errorf("building clients for AadB2c: %+v", err)
	}
	client.Advisor = advisor.NewClient(o)
	client.AnalysisServices = analysisServices.NewClient(o)
	client.ApiManagement = apiManagement.NewClient(o)
	client.Authorization = authorization.NewClient(o)
	client.AzureStackHCI = azurestackhci.NewClient(o)
	client.Compute = compute.NewClient(o)
	client.DataboxEdge = databoxedge.NewClient(o)
	client.Datadog = datadog.NewClient(o)

	if client.Dns, err = dns.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Dns: %+v", err)
	}
	client.KeyVault = keyvault.NewClient(o)
	client.Network = network.NewClient(o)
	client.Resource = resource.NewClient(o)

	return nil
}
