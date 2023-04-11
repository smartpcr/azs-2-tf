package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/smartpcr/azs-2-tf/config"
)

var (
	appFolder = filepath.Join(os.Getenv("ProgramData"), config.AppFolderName)
	logFolder = filepath.Join(appFolder, "logs")
	cfgFile   string
	RootCmd   = &cobra.Command{
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
			logrus.Errorln("No config file found: %s", cfgFile)
		} else {
			// Config file was found but another error was produced
			logrus.Errorln("Error reading config file")
		}
		os.Exit(1)
	}

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		logrus.Fatalf("unable to decode into struct, %v", err)
	}
}

func initLogger() {
	var fileLogger = &lumberjack.Logger{
		Filename:   filepath.Join(logFolder, "azs-2-tf.log"),
		MaxSize:    50, // megabytes
		MaxAge:     28, //days
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	}

	var consoleLogger = logrus.New()
	consoleLogger.SetOutput(os.Stdout)
	consoleLogger.SetLevel(logrus.InfoLevel)
	consoleLogger.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation:    true,
		ForceColors:               true,
		PadLevelText:              true,
		DisableTimestamp:          true,
		EnvironmentOverrideColors: true,
	})

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Out))
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
