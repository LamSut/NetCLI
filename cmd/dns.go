package cmd

import (
	"fmt"
	"net"
	"github.com/spf13/cobra"
)

var dnsLookupCmd = &cobra.Command{
	Use:   "dns [domain]",
	Short: "DNS lookup",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ips, err := net.LookupIP(args[0])
		if err != nil {
			fmt.Println("DNS Lookup failed:", err)
			return
		}
		for _, ip := range ips {
			fmt.Println(ip.String())
		}
	},
}
