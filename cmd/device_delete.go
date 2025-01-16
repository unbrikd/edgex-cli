package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/unbrikd/edgex-cli/internal/edgex"
)

var deviceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a device",
	Run: func(cmd *cobra.Command, args []string) {
		client := edgex.NewClient()

		if by == "name" {
			err := client.CoreMetadataService.DeleteDeviceFromName(deviceName)
			if err != nil {
				fmt.Printf("error deleting device: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("Deleted: %s\n", deviceName)
		} else if by == "id" {
			err := client.CoreMetadataService.DeleteDeviceFromId(deviceId)
			if err != nil {
				fmt.Printf("error deleting device: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("Deleted: %s\n", deviceId)
		} else {
			fmt.Println("error: flag --by must be 'name' or 'id'")
			os.Exit(1)
		}
	},
}

func init() {
	deviceDeleteCmd.Flags().StringVarP(&deviceName, "name", "n", "", "Device name")
	deviceDeleteCmd.Flags().StringVarP(&deviceId, "id", "i", "", "Device id")
	deviceDeleteCmd.Flags().StringVar(&by, "by", "name", "Delete by name or id (name or id)")
}
