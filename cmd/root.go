package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "LastOrder",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(runCmd, versionCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
