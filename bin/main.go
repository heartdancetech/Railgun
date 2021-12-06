package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gsxhnd/owl"
	"github.com/heart-dance-x/railgun/src/api"
	"github.com/heart-dance-x/railgun/src/core"
	"github.com/urfave/cli/v2"
)

var App = cli.NewApp()
var etcdUrlArry = &cli.StringSlice{}

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println("version: ", gitTag)
		fmt.Println("commit: ", gitCommit)
		fmt.Println("tree state: ", gitTreeState)
		fmt.Println("build date: ", buildDate)
		fmt.Println("go version: ", runtime.Version())
		fmt.Println("go compiler: ", runtime.Compiler)
		fmt.Println("platform: ", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
	}
	App.Name = "vpp-miniprogram-api"
	App.Version = gitTag
	App.Flags = []cli.Flag{
		&cli.StringSliceFlag{Name: "etcds", Aliases: []string{"e"}, Destination: etcdUrlArry},
	}
	App.Action = func(ctx *cli.Context) error {
		_ = owl.SetRemoteAddr(etcdUrlArry.Value())
		confKey := ctx.Args().Get(0)

		confStr, err := owl.GetRemote(confKey)
		if err != nil {
			return err
		}
		err = owl.ReadInConf([]byte(confStr))
		if err != nil {
			return err
		}

		c := make(chan string)
		go owl.Watcher(confKey, c)
		go func() {
			for i := range c {
				_ = owl.ReadInConf([]byte(i))
			}
		}()

		var certFile = owl.GetString("certFile")
		var keyFile = owl.GetString("keyFile")

		if owl.GetBool("enable_manage") {
			go api.RunTLS(certFile, keyFile)
		}

		core.SetMode(owl.GetString("run_mode"))
		g := core.New()
		_ = g.RunTLS(owl.GetString("addr"), certFile, keyFile)
		return nil
	}
}

func main() {
	err := App.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
