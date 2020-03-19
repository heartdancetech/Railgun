package cmd

import (
	"github.com/MisakaSystem/LastOrder/common"
	"github.com/MisakaSystem/LastOrder/core"
	"github.com/MisakaSystem/LastOrder/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   common.AppName,
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.PersistentFlags().StringSlice("etcd-url", []string{"localhost:2379"}, "use etcd url")
	rootCmd.PersistentFlags().String("run-mode", "debug", "run mode")
	rootCmd.PersistentFlags().String("addr", ":8000", "listen addr and port")

	_ = viper.BindPFlag("etcd-url", rootCmd.PersistentFlags().Lookup("etcd-url"))
	_ = viper.BindPFlag("run-mode", rootCmd.PersistentFlags().Lookup("run-mode"))
	_ = viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
}

func initConfig() {
	viper.SetEnvPrefix("last_order")
	viper.AutomaticEnv()
}
