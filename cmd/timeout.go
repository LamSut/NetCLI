package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var timeoutTestCmd = &cobra.Command{
	Use:   "timeout [host:port]",
	Short: "Test connection timeout",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		start := time.Now()
		_, err := net.DialTimeout("tcp", addr, 3*time.Second)
		duration := time.Since(start)
		if err != nil {
			fmt.Println("Timeout or error:", err)
		} else {
			fmt.Println("Connection successful")
		}
		fmt.Println("Elapsed:", duration)
	},
}
