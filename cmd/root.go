package cmd

import (
	"github.com/spf13/cobra"
)

var etcdUrlArry []string
var enableManage bool
var rootCmd = &cobra.Command{
	Use:   "railgun",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(runCmd, runTLSCmd, versionCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
