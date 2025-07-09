package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hostLookupCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Show system hostname",
	Run: func(cmd *cobra.Command, args []string) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println("Failed to get hostname:", err)
			return
		}
		fmt.Println("Hostname:", hostname)
	},
}
