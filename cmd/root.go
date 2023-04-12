package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/smartpcr/azs-2-tf/client"
	"github.com/smartpcr/azs-2-tf/log"
	"github.com/smartpcr/azs-2-tf/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/smartpcr/azs-2-tf/config"
)

var (
	cfgFile     string
	appSettings = utils.AppSettings{}
	RootCmd     = &cobra.Command{
		Use:     appSettings.GetAppName(),
		Version: config.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return bindViperToCobra(cmd)
		},
	}
	appConfig     config.AppConfig
	clientBuilder *client.ClientBuilder
)

func init() {
	RootCmd.Short = "azure stack -> terraform"
	RootCmd.Long = "sync azure stack resources to terraform files"
	cobra.OnInitialize(initLogger, initConfig, initClientBuilder)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azs-2-tf/config.json)")
}

func initLogger() {
	log.Log.Info("initialized logger")
}

func initConfig() {
	if cfgFile == "" {
		configDir := appSettings.GetConfigFolderPath()
		cfgFile = filepath.Join(configDir, appSettings.GetConfigFileName())
	}
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(utils.Terraform_Env_Prefix)
	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("json")
	log.Log.Infof("viper config file used: %s", viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Log.Errorf("No config file found: %s", cfgFile)
		} else {
			// Config file was found but another error was produced
			log.Log.Error("Error reading config file")
		}
		os.Exit(1)
	}

	// don't know why viper.Unmarshal doesn't work, fallback to use json
	rawJson := make(map[string]interface{})
	for key, value := range viper.AllSettings() {
		rawJson[key] = value
	}
	jsonBytes, err := json.Marshal(rawJson)
	if err != nil {
		log.Log.Errorf("Error marshalling rawJson: %s", err)
		os.Exit(1)
	}

	err = json.Unmarshal(jsonBytes, &appConfig)
	if err != nil {
		log.Log.Errorf("Failed to unmarshall json into appConfig: %s", err)
		os.Exit(1)
	}

	log.Log.Debugf("%+v\n", appConfig)
}

func initClientBuilder() {
	env, err := appConfig.GetEnvironment()
	if err != nil {
		log.Log.Errorf("environment is not set: %s", err)
		os.Exit(1)
	}

	var cloudCfg cloud.Configuration
	if env.EnvironmentType == utils.EnvironmentTypeAzureStack {
		cloudCfg = cloud.Configuration{
			ActiveDirectoryAuthorityHost: env.Authentication.LoginEndpoint,
			Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
				cloud.ResourceManager: {
					Endpoint: env.ResourceManager,
					Audience: env.Authentication.Audiences[0],
				},
			},
		}
	} else {
		cloudCfg = cloud.AzurePublic
	}

	clientOpt := arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			//APIVersion: env.ApiVersion,
			Cloud: cloudCfg,
			Telemetry: policy.TelemetryOptions{
				ApplicationID: appSettings.GetAppName(),
				Disabled:      false,
			},
			Logging: policy.LogOptions{
				IncludeBody: true,
			},
		},
	}

	// default credential read from environment variables
	_ = os.Setenv("AZURE_TENANT_ID", appConfig.TenantId)
	_ = os.Setenv("AZURE_CLIENT_ID", appConfig.ClientId)
	_ = os.Setenv("AZURE_CLIENT_SECRET", appConfig.ClientSecret)
	_ = os.Setenv("AZURE_SUBSCRIPTION_ID", appConfig.SubscriptionId)
	cred, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		ClientOptions: clientOpt.ClientOptions,
		TenantID:      appConfig.TenantId,
	})
	if err != nil {
		log.Log.Errorf("Failed to obtain credential for %s: %s", env.Name, err)
		os.Exit(1)
	}

	clientBuilder = client.NewClientBuilder(cred, clientOpt, appConfig.SubscriptionId, env.ApiVersion)
}

func bindViperToCobra(cmd *cobra.Command) error {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Override only if value hasn't changed
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				return
			}
		}
	})

	return nil
}
