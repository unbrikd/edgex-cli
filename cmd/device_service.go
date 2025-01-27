package cmd

import (
	"github.com/spf13/cobra"
)

var deviceServiceCmd = &cobra.Command{
	Use:   "device-service",
	Short: "Device service commands",
}

func init() {
	deviceServiceCmd.AddCommand(deviceServiceGetCmd)
}
