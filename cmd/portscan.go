package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var portScanCmd = &cobra.Command{
	Use:   "portscan [host]",
	Short: "Scan top ports",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		ports := []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 3306}
		for _, port := range ports {
			address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err == nil {
				fmt.Printf("Port %d open\n", port)
				conn.Close()
			}
		}
	},
}
