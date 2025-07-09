package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var reverseDNSCmd = &cobra.Command{
	Use:   "reverse [ip]",
	Short: "Reverse DNS lookup",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := net.LookupAddr(args[0])
		if err != nil {
			fmt.Println("Reverse DNS lookup failed:", err)
			return
		}
		for _, host := range hosts {
			fmt.Println(host)
		}
	},
}
