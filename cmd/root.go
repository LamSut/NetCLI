package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netcli",
	Short: "Advanced Network CLI Tool",
	Long:  `A multifunctional network CLI tool for diagnostics and information retrieval.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(dnsLookupCmd)
	rootCmd.AddCommand(reverseDNSCmd)
	rootCmd.AddCommand(tracerouteCmd)
	rootCmd.AddCommand(portScanCmd)
	rootCmd.AddCommand(ipInfoCmd)
	rootCmd.AddCommand(localIPCmd)
	rootCmd.AddCommand(macLookupCmd)
	rootCmd.AddCommand(hostLookupCmd)
	rootCmd.AddCommand(timeoutTestCmd)
}