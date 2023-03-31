package cmd

import (
	"github.com/spf13/cobra"

	"github.com/smartpcr/azs-2-tf/src/config"
)

var getVirtualNetworkCommand = &cobra.Command{
	Use:     "vnet [name]",
	Version: config.Version,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
