package cmd

import (
	"github.com/gsxhnd/owl"
	"github.com/railgun-project/railgun/api"
	"github.com/railgun-project/railgun/core"
	"github.com/urfave/cli/v2"
)

var runCmd = &cli.Command{
	Name:   "run",
	Usage:  "run",
	Before: nil,
	After:  nil,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{Name: "etcds", Aliases: []string{"e"}, Destination: etcdUrlArry},
		&cli.BoolFlag{Name: "dashboard", Value: false, Destination: &enableManage},
	},
	Action: func(ctx *cli.Context) error {
		_ = owl.SetRemoteAddr(etcdUrlArry.Value())
		confKey := ctx.Args().Get(0)
		confStr, err := owl.GetRemote(confKey)
		if err != nil {
		}
		err = owl.ReadInConf([]byte(confStr))
		if err != nil {
			panic(err)
		}
		c := make(chan string)
		go owl.Watcher(confKey, c)
		go func() {
			for i := range c {
				_ = owl.ReadInConf([]byte(i))
			}
		}()

		if enableManage {
			go api.Run()
		}

		core.SetMode(owl.GetString("run_mode"))
		g := core.New()
		_ = g.Run(owl.GetString("addr"))
		return nil
	},
}
