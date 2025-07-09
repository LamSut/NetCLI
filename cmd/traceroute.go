package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tracerouteCmd = &cobra.Command{
	Use:   "traceroute [host]",
	Short: "Trace route to host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Traceroute is not fully implemented due to raw socket limitations in Go.")
	},
}
