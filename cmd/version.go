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

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: `Show version`,
	Long:  `Show version of railgun`,
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
