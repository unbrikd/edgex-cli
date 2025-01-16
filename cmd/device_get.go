package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/unbrikd/edgex-cli/internal/edgex"
)

var deviceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get device",
	Run: func(cmd *cobra.Command, args []string) {
		client := edgex.NewClient()
		device, err := client.CoreMetadataService.GetDevice(deviceName)
		if err != nil {
			fmt.Printf("error getting device: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Device Name:    %s\n", device.Name)
		fmt.Printf("Device ID:      %s\n", device.Id)
		fmt.Printf("Device Service: %s\n", device.ServiceName)
		fmt.Printf("Device Profile: %s\n", device.ProfileName)
	},
}

func init() {
	deviceGetCmd.Flags().StringVarP(&deviceName, "name", "n", "", "Device name")
}
