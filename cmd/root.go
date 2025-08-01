// Package cmd це пакет для cli/shell
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "wallet",
		Short: "створити гаманець",
}

func init() {
	rootCmd.AddCommand(SendCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
