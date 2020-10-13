package core

import (
	"github.com/gsxhnd/owl"
	"github.com/patrickmn/go-cache"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type ProxyEngine struct {
	Cache *cache.Cache
	pool  sync.Pool
}

func New() *ProxyEngine {
	engine := &ProxyEngine{}
	engine.Cache = cache.New(cache.NoExpiration, cache.NoExpiration)
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
	reqPath := c.Req.URL.Path
	if reqPath == "" {
		return
	}

	//查询原始请求路径，如：/arithmetic/calculate/10/5
	//按照分隔符'/'对路径进行分解，获取服务名称serviceName
	pathArray := strings.Split(reqPath, "/")
	serviceName := pathArray[1]
	destPathMap := owl.GetStringMapString("proxy")
	destPath := strings.Join(pathArray[2:], "/")
	c.Req.URL.Scheme = "http"
	c.Req.URL.Host = destPathMap[serviceName]
	c.Req.URL.Path = "/" + destPath

	c.reqLog.StartTime = time.Now().UTC()
	c.reqLog.Path = c.Req.URL.Path
	c.reqLog.Method = c.Req.Method
	c.reqLog.Target = c.Req.URL.Hostname()
	c.reqLog.Query = c.Req.URL.RawQuery
	c.reqLog.IP = clientIP(c.Req)
	c.reqLog.RemoteAddr = c.Req.URL.Host
	c.reqLog.UserAgent = c.Req.Header.Get("user-agent")

	ReverseProxy(c).ServeHTTP(c.Writer, c.Req)
}

func clientIP(req *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
