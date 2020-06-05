package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "LastOrder",
		Short: "",
		Long:  "",
	}
)

func init() {
	rootCmd.AddCommand(runCmd, versionCmd)
}
