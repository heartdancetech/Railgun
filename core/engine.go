package core

import (
	"github.com/spf13/viper"
	"net"
	"net/http"
	"sync"
	"time"
)

type ProxyEngine struct {
	// Director must be a function which modifies
	// the request into a new request to be sent
	// using Transport. Its response is then copied
	// back to the original client unmodified.
	Director func(*http.Request)

	// The transport used to perform proxy requests.
	Transport http.RoundTripper

	// FlushInterval specifies the flush interval
	// to flush to the client while copying the
	// response body.
	// If zero, no periodic flushing is done.
	FlushInterval time.Duration

	// dialer is used when values from the
	// defaultDialer need to be overridden per Proxy
	dialer *net.Dialer

	srvResolver srvResolver

	pool sync.Pool
}

func New() *ProxyEngine {
	engine := &ProxyEngine{}
	//engine.pool.New = func() interface{} {
	//	return engine.allocateContext()
	//}
	return engine
}

//func (e *ProxyEngine) allocateContext() *Context {
//	return &Context{engine: e}
//}

func (e *ProxyEngine) Run() error {
	//初始化配置
	SetMode(viper.GetString("run-mode"))
	return nil

	//cli, err := discovery.NewClientDis(viper.GetStringSlice("etcds"))
	//if err != nil {
	//	logger.Panic("", zap.Error(err))
	//	return err
	//}
	//
	//ctx, err := cli.InitServices("/service")
	//if err != nil {
	//	logger.Panic("", zap.Error(err))
	//	return err
	//}
	//
	////创建反向代理
	//proxy := NewReverseProxy(ctx)
	//
	//errChannel := make(chan error)
	////go func() {
	////	c := make(chan os.Signal)
	////	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	////	errChannel <- fmt.Errorf("%s", <-c)
	////}()
	//
	////开始监听
	//var http01 = http.NewServeMux()
	//http01.Handle("/", proxy)
	//errChannel <- http.ListenAndServe(":9091", proxy)
	//
	//<-errChannel
	//return nil
}
