package core

import (
	"net/http"
	"net/http/httputil"
)

// NewReverseProxy 创建反向代理处理方法
func ReverseProxy() *httputil.ReverseProxy {
	//创建Director
	director := func(req *http.Request) {
		return
	}

	modifyResponse := func(res *http.Response) error {
		return nil
	}

	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
}
