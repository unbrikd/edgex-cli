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
		exists, err := client.CoreMetadataService.DeviceExists(deviceName)
		if err != nil {
			fmt.Printf("error checking if device exists: %s\n", err)
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
}
