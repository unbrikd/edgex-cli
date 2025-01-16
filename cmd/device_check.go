package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/unbrikd/edgex-cli/internal/edgex"
)

var deviceCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if a device exists",
	Run: func(cmd *cobra.Command, args []string) {

		client := edgex.NewClient()

		var exists bool
		var err error

		if by == "name" {
			exists, err = client.CoreMetadataService.DeviceExistsFromName(deviceName)
			if err != nil {
				fmt.Printf("error checking if device exists: %s\n", err)
				os.Exit(1)
			}
		} else if by == "id" {
			exists, err = client.CoreMetadataService.DeviceExistsFromId(deviceId)
			if err != nil {
				fmt.Printf("error checking if device exists: %s\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("error: flag --by must be 'name' or 'id'")
			os.Exit(1)
		}

		if exists {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	},
}

func init() {
	deviceCheckCmd.Flags().StringVarP(&deviceName, "name", "n", "", "Device name")
	deviceCheckCmd.Flags().StringVarP(&deviceId, "id", "i", "", "Device id")
	deviceCheckCmd.Flags().StringVar(&by, "by", "name", "Check by name or id (name or id)")
}
