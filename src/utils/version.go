package utils

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"runtime"
)

var (
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
)

func GetVersionTag() string {
	return gitTag
}

func PrintVersion(c *cli.Context) {
	fmt.Println("version: ", gitTag)
	fmt.Println("commit: ", gitCommit)
	fmt.Println("tree state: ", gitTreeState)
	fmt.Println("build date: ", buildDate)
	fmt.Println("go version: ", runtime.Version())
	fmt.Println("go compiler: ", runtime.Compiler)
	fmt.Println("platform: ", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}
