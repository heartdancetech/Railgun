package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
	goVersion    = runtime.Version()
	compiler     = runtime.Compiler
	platform     = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: `Show version`,
	Long:  `Show version of gecko`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gecko version: ", gitTag)
		fmt.Println("gecko commit: ", gitCommit)
		fmt.Println("gecko build date: ", gitTreeState)
		fmt.Println("gecko build date: ", buildDate)
		fmt.Println("go version: ", goVersion)
		fmt.Println("go compiler: ", compiler)
		fmt.Println("platform: ", platform)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
