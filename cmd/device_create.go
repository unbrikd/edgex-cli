package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unbrikd/edgex-cli/internal/edgex"
)

var deviceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a device",
	Run: func(cmd *cobra.Command, args []string) {
		client := edgex.NewClient()

		// parse each event in the deviceAutoEvents slice into a slice of edgex.AutoEvent
		autoEvents := make([]edgex.AutoEvent, 0)
		for _, e := range deviceAutoEvents {
			tmp := &edgex.AutoEvent{}
			err := json.Unmarshal([]byte(e), tmp)
			if err != nil {
				fmt.Printf("failed to parse auto event: %v", err)
			}

			if tmp.Interval == "" {
				fmt.Printf("interval is required for auto event")
			}

			if tmp.SourceName == "" {
				fmt.Printf("sourceName is required for auto event")
			}

			autoEvents = append(autoEvents, *tmp)
		}

		protocols := make(map[string]interface{})
		err := json.Unmarshal([]byte(deviceProtocols), &protocols)
		if err != nil {
			fmt.Printf("failed to parse protocols: %v", err)
		}

		device := edgex.Device{
			ApiVersion:     client.CoreMetadataService.ApiVersion,
			Name:           deviceName,
			ServiceName:    deviceServiceName,
			ProfileName:    deviceProfile,
			Labels:         deviceLabels,
			Description:    deviceDescription,
			AdminState:     deviceAdminState,
			OperatingState: deviceOperatingState,
			AutoEvents:     autoEvents,
			Protocols:      protocols,
		}

		if deviceId != "" {
			device.Id = deviceId
		}

		id, err := client.CoreMetadataService.CreateDevice(device)
		if err != nil {
			fmt.Printf("failed to create device: %v", err)
		} else {
			fmt.Printf("%s\n", id)
		}
	},
}

func init() {
	deviceCreateCmd.Flags().StringVarP(&deviceName, "name", "n", "", "Device name")
	deviceCreateCmd.Flags().StringVar(&deviceId, "id", "", "Device id")
	deviceCreateCmd.Flags().StringVarP(&deviceServiceName, "service-name", "s", "", "Service name")
	deviceCreateCmd.Flags().StringVarP(&deviceProfile, "profile", "p", "", "Device profile")
	deviceCreateCmd.Flags().StringSliceVarP(&deviceLabels, "label", "l", []string{}, "Device labels")
	deviceCreateCmd.Flags().StringVarP(&deviceDescription, "description", "d", "", "Device description")
	deviceCreateCmd.Flags().StringVar(&deviceAdminState, "admin-state", "UNLOCKED", "Device admin state")
	deviceCreateCmd.Flags().StringVar(&deviceOperatingState, "operating-state", "UP", "Device operating state")
	deviceCreateCmd.Flags().StringSliceVar(&deviceAutoEvents, "auto-event", make([]string, 0), "Device auto events")
	deviceCreateCmd.Flags().StringVar(&deviceProtocols, "protocols", "{}", "Device protocols")
}
