package cmd

import (
	"bytes"
	"github.com/gsxhnd/owl"
	"github.com/railgun-project/railgun/api"
	"github.com/railgun-project/railgun/core"
	"github.com/spf13/viper"
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
		owl.SetAddr(etcdUrlArry.Value())
		confKey := ctx.Args().Get(0)
		confStr, err := owl.GetByKey(confKey)
		if err != nil {
		}
		viper.SetConfigType("yaml")
		err = viper.ReadConfig(bytes.NewBuffer([]byte(confStr)))
		if err != nil {
			panic(err)
		}
		c := make(chan string)
		go owl.Watcher(confKey, c)
		go func() {
			for i := range c {
				_ = viper.ReadConfig(bytes.NewBuffer([]byte(i)))
			}
		}()

		if enableManage {
			go api.Run()
		}

		core.SetMode(viper.GetString("run_mode"))
		g := core.New()
		_ = g.Run(viper.GetString("addr"))
		return nil
	},
}
