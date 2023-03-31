package cmd

import (
	"github.com/spf13/cobra"

	"github.com/smartpcr/azs-2-tf/src/config"
)

var getResourceGroupCommand = &cobra.Command{
	Use:     "group [name]",
	Short:   "Query the state of a resource group",
	Long:    `Query the state of a resource group`,
	Version: config.Version,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
