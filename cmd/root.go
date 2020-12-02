package cmd

import (
	"github.com/railgun-project/railgun/utils"
	"github.com/urfave/cli/v2"
)

var etcdUrlArry = &cli.StringSlice{}
var enableManage bool
var App = cli.NewApp()

func init() {
	cli.VersionPrinter = utils.PrintVersion
	App.Usage = "railgun"
	App.Version = utils.GetVersionTag()
	App.Commands = []*cli.Command{
		runCmd,
		runTLSCmd,
	}
}
