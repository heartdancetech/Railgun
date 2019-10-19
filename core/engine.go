package core

import (
	"fmt"
	"github.com/MisakaSystem/LastOrder/common"
	"github.com/MisakaSystem/LastOrder/discovery"
	"github.com/MisakaSystem/LastOrder/logger"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Engine struct {
	//pool             sync.Pool
}

func New() *Engine {
	engine := &Engine{}
	//engine.pool.New = func() interface{} {
	//	return engine.allocateContext()
	//}
	return engine
}

func (e *Engine) allocateContext() *Context {
	return &Context{}
}

func (e *Engine) Run() {
	if err := common.FlagInit(); err != nil {
		return
	}
	//初始化配置
	etcdList := viper.GetStringSlice("etcd.url")
	SetMode(viper.GetString("runMode"))
	cli, err := discovery.NewClientDis(etcdList)
	if err != nil {
		logger.SelfLogger().WithField("error", err).Error()
		return
	}

	ctx, err := cli.InitServices("/service")
	if err != nil {
		logger.SelfLogger().Error()
		return
	}

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
		errChannel <- http.ListenAndServe(":9091", http01)

	}()

	<-errChannel
}
