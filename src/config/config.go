package config

type AppConfig struct {
	SubscriptionId        string `json:"subscription_id"`
	TenantId              string `json:"tenant_id"`
	ClientId              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	AzureStackEnvironment string `json:"azure_stack_environment"`
	AzureStackArmEndpoint string `json:"azure_stack_arm_endpoint"`
}
