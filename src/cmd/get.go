package cmd

import (
	"github.com/spf13/cobra"

	"github.com/smartpcr/azs-2-tf/src/config"
)

var (
	getCmd = &cobra.Command{
		Use:     "get",
		Short:   "get the state of the Azure Stack environment",
		Long:    `get the state of the Azure Stack environment`,
		Version: config.Version,
	}
	getCommandList = []*cobra.Command{
		getResourceGroupCommand,
		getStorageAccountCommand,
		getVirtualNetworkCommand,
		getNetworkSecurityGroupCommand,
		getSubnetCommand,
		getPublicIPCommand,
		getNetworkInterfaceCommand,
		getVirtualMachineCommand,
		getVirtualMachineExtensionCommand,
		getVirtualMachineScaleSetCommand,
		getVirtualMachineScaleSetExtensionCommand,
		getVirtualMachineScaleSetVMCommand,
		getVirtualMachineScaleSetVMExtensionCommand,
		getAvailabilitySetCommand,
	}
)

func init() {
	RootCmd.AddCommand(getCmd)

	for _, subCmd := range getCommandList {
		getCmd.AddCommand(subCmd)
	}
}
