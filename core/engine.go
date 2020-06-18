package core

import (
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"sync"
)

type ProxyEngine struct {
	pool sync.Pool
}

func New() *ProxyEngine {
	engine := &ProxyEngine{}
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (e *ProxyEngine) allocateContext() *Context {
	return &Context{engine: e}
}

func (e *ProxyEngine) Run(addr string) (err error) {
	err = http.ListenAndServe(addr, e)
	return
}

func (e *ProxyEngine) RunTLS(addr, certFile, keyFile string) (err error) {
	err = http.ListenAndServeTLS(addr, certFile, keyFile, e)
	return
}

func (e *ProxyEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := e.pool.Get().(*Context)
	c.Writer = w
	c.Req = req

	e.handleHTTPRequest(c)

	e.pool.Put(c)
}

func (e *ProxyEngine) handleHTTPRequest(c *Context) {
	//查询原始请求路径，如：/arithmetic/calculate/10/5
	reqPath := c.Req.URL.Path
	if reqPath == "" {
		return
	}
	//按照分隔符'/'对路径进行分解，获取服务名称serviceName
	pathArray := strings.Split(reqPath, "/")
	serviceName := pathArray[1]
	destPathMap := viper.GetStringMapString("proxy")
	destPath := strings.Join(pathArray[2:], "/")
	c.Req.URL.Scheme = "http"
	c.Req.URL.Host = destPathMap[serviceName]
	c.Req.URL.Path = "/" + destPath

	ReverseProxy().ServeHTTP(c.Writer, c.Req)
}
