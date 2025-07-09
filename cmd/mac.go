package cmd

import (
	"fmt"
	"net"
	"github.com/spf13/cobra"
)

var macLookupCmd = &cobra.Command{
	Use:   "mac",
	Short: "Show MAC addresses of interfaces",
	Run: func(cmd *cobra.Command, args []string) {
		interfaces, err := net.Interfaces()
		if err != nil {
			fmt.Println("Failed to get interfaces:", err)
			return
		}
		for _, i := range interfaces {
			fmt.Printf("%s - %s\n", i.Name, i.HardwareAddr)
		}
	},
}
