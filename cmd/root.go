package cmd

import (
	"fmt"
	"github.com/smartpcr/azs-2-tf/log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/smartpcr/azs-2-tf/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	RootCmd = &cobra.Command{
		Use:     config.AppName,
		Version: config.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return bindViperToCobra(cmd)
		},
	}
	appConfig config.AppConfig
)

func init() {
	RootCmd.Short = "azure stack -> terraform"
	RootCmd.Long = "sync azure stack resources to terraform files"
	cobra.OnInitialize(initConfig, initLogger)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azs-2-tf/config.json)")
}

func initConfig() {
	if cfgFile == "" {
		// Search config under "~/.azs-2-tf/" (without extension).
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		configDir := filepath.Join(home, config.AppFolderName)
		cfgFile = filepath.Join(configDir, config.ConfigFileName)
	}
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(config.AppName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("json")
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

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		log.Log.Fatalf("unable to decode into struct, %v", err)
	}
}

func initLogger() {
	log.Log.Info("initialized logger")
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
