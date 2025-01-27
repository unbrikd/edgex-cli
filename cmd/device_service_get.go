package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/unbrikd/edgex-cli/internal/edgex"
)

var deviceServiceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get device service",
	Run: func(cmd *cobra.Command, args []string) {
		client := edgex.NewClient()

		ds, err := client.CoreMetadataService.GetDeviceServiceFromName(deviceServiceName)
		if err != nil {
			fmt.Printf("error getting device service: %s\n", err)
			os.Exit(1)
		}

		prettyPrintDeviceService(ds)
	},
}

func init() {
	deviceServiceGetCmd.Flags().StringVarP(&deviceServiceName, "name", "n", "", "Device name")
}

func prettyPrintDeviceService(ds *edgex.DeviceService) {
	fmt.Printf("Device Service Name:    %s\n", ds.Name)
	fmt.Printf("Device Service ID:      %s\n", ds.Id)
}
