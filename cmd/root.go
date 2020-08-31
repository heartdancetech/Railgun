package cmd

import (
	"github.com/urfave/cli/v2"
)

var etcdUrlArry = &cli.StringSlice{}
var enableManage bool
var App = cli.NewApp()

func init() {
	App.Usage = "railgun"
	App.HideVersion = true
	App.Commands = []*cli.Command{
		runCmd,
		runTLSCmd,
		versionCmd,
	}
}
