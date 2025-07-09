package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var ipInfoCmd = &cobra.Command{
	Use:   "ipinfo [host]",
	Short: "Show IP addresses for host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ips, err := net.LookupHost(args[0])
		if err != nil {
			fmt.Println("Failed to lookup host:", err)
			return
		}
		for _, ip := range ips {
			fmt.Println(ip)
		}
	},
}
