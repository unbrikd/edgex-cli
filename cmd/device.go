package cmd

import (
	"github.com/spf13/cobra"
)

var (
	deviceName           string
	deviceId             string
	deviceProfile        string
	deviceLabels         []string
	deviceDescription    string
	deviceAdminState     string
	deviceOperatingState string
	deviceAutoEvents     string
	deviceProtocols      string
	deviceServiceName    string
	getAll               bool
	by                   string
)

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Device commands",
}

func init() {
	deviceCmd.AddCommand(deviceGetCmd)
	deviceCmd.AddCommand(deviceCheckCmd)
	deviceCmd.AddCommand(deviceCreateCmd)
	deviceCmd.AddCommand(deviceDeleteCmd)
}
