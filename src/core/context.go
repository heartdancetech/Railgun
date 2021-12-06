package core

import (
	"net/http"
)

type Context struct {
	engine *ProxyEngine
	Writer http.ResponseWriter
	Req    *http.Request
	reqLog
}
