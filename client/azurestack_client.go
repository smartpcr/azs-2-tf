package client

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-azure-helpers/authentication"

	"github.com/smartpcr/azs-2-tf/config"
	"github.com/smartpcr/azs-2-tf/internal/azurestack"
	"github.com/smartpcr/azs-2-tf/log"
)

var logger = log.Log

func NewAzureStackClient(ctx context.Context, config *config.AppConfig) (*azurestack.Client, error) {
	builder := &authentication.Builder{
		SubscriptionID:     config.SubscriptionId,
		ClientID:           config.ClientId,
		ClientSecret:       config.ClientSecret,
		TenantID:           config.TenantId,
		Environment:        config.EnvironmentName,
		MetadataHost:       config.MetadataEndpoint,
		AuxiliaryTenantIDs: config.AuxiliaryTenantIds,
		MsiEndpoint:        config.MsiEndpoint,
		ClientCertPassword: config.ClientCertPassword,
		ClientCertPath:     config.ClientCertPath,

		// Feature Toggles
		SupportsClientCertAuth:   true,
		SupportsClientSecretAuth: true,
		// SupportsManagedServiceIdentity: d.Get("use_msi").(bool), todo supported in stack?
		SupportsAzureCliToken:    true,
		SupportsAuxiliaryTenants: len(config.AuxiliaryTenantIds) > 0,

		// Doc Links
		ClientSecretDocsLink: "https://registry.terraform.io/providers/hashicorp/azurestack/latest/docs/guides/service_principal_client_secret",
	}

	logger.Info("building auth config...")
	authConfig, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("building Azurestack Client: %s", err)
	}

	terraformVersion := config.TerraformVersion
	if terraformVersion == "" {
		// Terraform 0.12 introduced this field to the protocol
		// We can therefore assume that if it's missing it's 0.10 or 0.11
		terraformVersion = "0.11+compatible"
	}

	logger.Info("building azurestack client...")
	clientBuilder := azurestack.ClientBuilder{
		AuthConfig:                  authConfig,
		SkipProviderRegistration:    config.SkipProviderRegistration,
		TerraformVersion:            terraformVersion,
		DisableCorrelationRequestID: config.DisableCorrelationRequestID,

		// this field is intentionally not exposed in the provider block, since it's only used for
		// platform level tracing
		CustomCorrelationRequestID: os.Getenv("ARM_CORRELATION_REQUEST_ID"),
	}

	client, err := azurestack.Build(ctx, clientBuilder)
	if err != nil {
		return nil, fmt.Errorf("failed building Azurestack Client: %s", err)
	}

	client.StopContext = ctx

	return client, nil
}
