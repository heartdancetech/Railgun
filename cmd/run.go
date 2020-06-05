package cmd

import (
	"bytes"
	"fmt"
	"github.com/MisakaSystem/LastOrder/core"
	"github.com/MisakaSystem/LastOrder/logger"
	"github.com/coreos/etcd/clientv3"
	"github.com/gsxhnd/owl/backend"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var (
	etcdUrlArry []string
	confKey     string
)
var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "run",
	Long:    "run",
	Example: "run",
	Run: func(cmd *cobra.Command, args []string) {
		getConfig()
		fmt.Printf("server name: %v", viper.GetString("name"))
		logger.Init("debug", viper.GetString("name"))
		g := core.New()
		_ = g.Run()
	},
}

func init() {
	runCmd.PersistentFlags().StringArrayVar(&etcdUrlArry, "etcds", []string{"127.0.0.1:2379"}, "")
	runCmd.PersistentFlags().StringVar(&confKey, "conf_key", "", "")
}

func getConfig() {
	conf := clientv3.Config{
		Endpoints:        etcdUrlArry,
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second,
	}
	e, _ := backend.NewEtcdConn(conf)
	confStr, _ := e.Get(confKey)
	viper.SetConfigType("yaml")
	_ = viper.ReadConfig(bytes.NewBuffer([]byte(confStr)))
}
