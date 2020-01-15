package cmd

import (
	"github.com/MisakaSystem/LastOrder/core"
	"github.com/MisakaSystem/LastOrder/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "LastOrder",
		Short: "",
		Long:  "",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			logger.Init("debug", viper.GetString("name"))
			g := core.New()
			_ = g.Run()
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringSlice("etcdUrl", []string{"localhost:2379"}, "use etcd url")
	rootCmd.PersistentFlags().String("runMode", "localhost:2379", "use etcd url")
	rootCmd.PersistentFlags().String("addr", ":8000", "listen addr")
	rootCmd.PersistentFlags().String("name", "last_order", "server name")
	_ = viper.BindPFlag("etcdUrl", rootCmd.PersistentFlags().Lookup("etcdUrl"))
	_ = viper.BindPFlag("runMode", rootCmd.PersistentFlags().Lookup("runMode"))
	_ = viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
	_ = viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
}

func initConfig() {
	viper.SetEnvPrefix("last_order")
	viper.AutomaticEnv()
}
