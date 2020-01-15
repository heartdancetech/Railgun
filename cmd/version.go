package cmd

import (
	"fmt"
	"github.com/MisakaSystem/LastOrder/common"
	"github.com/spf13/cobra"
)

var (
	gitTag       = common.GetVersion().GitTag
	gitCommit    = common.GetVersion().GitCommit
	gitTreeState = common.GetVersion().GitTreeState
	buildDate    = common.GetVersion().BuildDate
	goVersion    = common.GetVersion().GoVersion
	compiler     = common.GetVersion().Compiler
	platform     = common.GetVersion().Platform
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: `Show version`,
	Long:  `Show version of LastOrder`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: ", gitTag)
		fmt.Println("commit: ", gitCommit)
		fmt.Println("tree state: ", gitTreeState)
		fmt.Println("build date: ", buildDate)
		fmt.Println("go version: ", goVersion)
		fmt.Println("go compiler: ", compiler)
		fmt.Println("platform: ", platform)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
