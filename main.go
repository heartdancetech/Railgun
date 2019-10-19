package main

import (
	"fmt"
	"github.com/MisakaSystem/LastOrder/common"
	"github.com/MisakaSystem/LastOrder/discovery"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := common.FlagInit(); err != nil {
		return
	}
	//初始化配置
	etcdList := viper.GetStringSlice("etcd.url")
	common.SetMode(viper.GetString("runMode"))
	cli, _ := discovery.NewClientDis(etcdList)
	ctx, _ := cli.InitServices("/service")

	//创建反向代理
	proxy := common.NewReverseProxy(ctx)

	errChannel := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		errChannel <- fmt.Errorf("%s", <-c)
	}()

	//开始监听
	go func() {
		var http01 = http.NewServeMux()
		http01.Handle("/", proxy)
		//http01.HandleFunc("/", )
		errChannel <- http.ListenAndServe(":9090", http01)

	}()

	<-errChannel
}
