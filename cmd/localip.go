package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var localIPCmd = &cobra.Command{
	Use:   "localip",
	Short: "Show local IP address",
	Run: func(cmd *cobra.Command, args []string) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println("Error fetching local IP:", err)
			return
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Println(ipnet.IP.String())
				}
			}
		}
	},
}
