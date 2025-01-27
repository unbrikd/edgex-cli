package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gems-edgexcli",
	Short: "GEMS EdgeX utility CLI",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(deviceCmd)
	rootCmd.AddCommand(deviceServiceCmd)
}
