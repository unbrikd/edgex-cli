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
		err := client.CoreMetadataService.DeleteDevice(deviceName)
		if err != nil {
			fmt.Printf("error deleting device: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Device %s deleted\n", deviceName)
	},
}

func init() {
	deviceDeleteCmd.Flags().StringVarP(&deviceName, "name", "n", "", "Device name")
}
