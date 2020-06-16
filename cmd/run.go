package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/MisakaSystem/LastOrder/core"
	"github.com/MisakaSystem/LastOrder/logger"
	"github.com/gsxhnd/owl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	runCmd.PersistentFlags().StringArrayVarP(&etcdUrlArry, "etcds", "e", []string{"127.0.0.1:2379"}, "")
}

var etcdUrlArry []string
var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "run",
	Long:    "run",
	Example: "run",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || len(args) > 1 {
			return errors.New("need key or too manay args")
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		owl.SetAddr(etcdUrlArry)
		confKey := args[0]
		confStr, err := owl.GetByKey(confKey)
		if err != nil {
		}
		fmt.Print("config: ", confStr)
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
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server name: ", viper.GetString("name"))
		logger.Init("debug", viper.GetString("name"))
		g := core.New()
		_ = g.Run()
	},
}
