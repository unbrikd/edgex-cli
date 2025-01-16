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

		if getAll {
			devices, err := client.CoreMetadataService.GetAllDevices(100)
			if err != nil {
				fmt.Printf("error getting all devices: %s\n", err)
				os.Exit(1)
			}
			for _, device := range *devices {
				prettyPrintDevice(&device)
				fmt.Println("---")
			}
		} else if by == "id" {
			device, err := client.CoreMetadataService.GetDeviceFromId(deviceId)
			if err != nil {
				fmt.Printf("error getting device: %s\n", err)
				os.Exit(1)
			}
			prettyPrintDevice(device)
		} else if by == "name" {
			device, err := client.CoreMetadataService.GetDeviceFromName(deviceName)
			if err != nil {
				fmt.Printf("error getting device: %s\n", err)
				os.Exit(1)
			}
			prettyPrintDevice(device)
		} else {
			fmt.Println("error: flag --by must be 'name' or 'id'")
			os.Exit(1)
		}
	},
}

func init() {
	deviceGetCmd.Flags().StringVarP(&deviceName, "name", "n", "", "Device name")
	deviceGetCmd.Flags().StringVarP(&deviceId, "id", "i", "", "Device id")
	deviceGetCmd.Flags().BoolVarP(&getAll, "all", "a", false, "Get all devices")
	deviceGetCmd.Flags().StringVar(&by, "by", "name", "Get by name or id (name or id)")
}

func prettyPrintDevice(device *edgex.Device) {
	fmt.Printf("Device Name:    %s\n", device.Name)
	fmt.Printf("Device ID:      %s\n", device.Id)
	fmt.Printf("Device Service: %s\n", device.ServiceName)
	fmt.Printf("Device Profile: %s\n", device.ProfileName)
}
