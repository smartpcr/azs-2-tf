package cmd

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/spf13/cobra"

	"github.com/smartpcr/azs-2-tf/config"
	"github.com/smartpcr/azs-2-tf/log"
)

var getResourceGroupCommand = &cobra.Command{
	Use:     "group [name]",
	Short:   "Query the state of a resource group",
	Long:    `Query the state of a resource group`,
	Version: config.Version,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		groupName := args[0]
		group, err := getResourceGroup(groupName)
		if err != nil {
			log.Log.Errorf("failed to get resource group %s: %v", groupName, err)
		}

		log.Log.Infof("resource group %s: %v", groupName, group)
	},
}

func getResourceGroup(groupName string) (armresources.ResourceGroup, error) {
	groupClient, err := clientBuilder.NewResourceGroupsClient()
	if err != nil {
		return armresources.ResourceGroup{}, err
	}

	resp, err := groupClient.Get(context.Background(), groupName, nil)
	if err != nil {
		return armresources.ResourceGroup{}, err
	}

	return resp.ResourceGroup, nil
}
