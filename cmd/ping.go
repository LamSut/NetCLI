package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping [host]",
	Short: "Ping a host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		conn, err := net.DialTimeout("ip4:icmp", host, 2*time.Second)
		if err != nil {
			fmt.Println("Host unreachable:", err)
			return
		}
		fmt.Println("Host is reachable")
		conn.Close()
	},
}
