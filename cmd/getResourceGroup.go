package cmd

import (
	"context"
	"fmt"
	"io"

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
		//group, err := getResourceGroup(groupName)
		group, err := getGroup(context.Background(), groupName)
		if err != nil {

			log.Log.Errorf("failed to get resource group %s: %v", groupName, err)
		}

		log.Log.Infof("resource group %s: %v", groupName, group)
	},
}

func getGroup(ctx context.Context, groupName string) (armresources.ResourceGroup, error) {
	groupClient := azsClient.Resource.GroupsClient
	resp, err := groupClient.Get(ctx, groupName)
	if err != nil {
		return armresources.ResourceGroup{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return armresources.ResourceGroup{}, fmt.Errorf("failed to get resource group %s: %s", groupName, resp.Status)
	}

	var group = armresources.ResourceGroup{
		Name:     resp.Name,
		ID:       resp.ID,
		Location: resp.Location,
	}
	return group, nil
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
