package cmd

import (
	"bytes"
	"errors"
	"github.com/gsxhnd/owl"
	"github.com/railgun-project/railgun/api"
	"github.com/railgun-project/railgun/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	runCmd.PersistentFlags().StringArrayVarP(&etcdUrlArry, "etcds", "e", []string{"127.0.0.1:2379"}, "")
	runCmd.PersistentFlags().BoolVar(&enableManage, "enable-manage", false, "")
}

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
		if enableManage {
			go api.Run()
		}

		core.SetMode(viper.GetString("run_mode"))
		g := core.New()
		_ = g.Run(viper.GetString("addr"))
	},
}
