package core

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"time"
)

// NewReverseProxy Creat Reverse Proxy
func ReverseProxy(ctx *Context) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		return
	}

	modifyResponse := func(res *http.Response) error {
		ctx.reqLog.EndTime = time.Now().UTC()
		ctx.reqLog.Latency = ctx.reqLog.EndTime.Sub(ctx.reqLog.StartTime)
		ctx.reqLog.Status = res.StatusCode
		logger.Info("success_req", zap.Any("request", ctx.reqLog))
		return nil
	}

	errHandler := func(res http.ResponseWriter, request *http.Request, err error) {
		var resJson = struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Code:    10010,
			Message: err.Error(),
			Data:    nil,
		}

		ctx.reqLog.EndTime = time.Now()
		ctx.reqLog.Latency = ctx.reqLog.EndTime.Sub(ctx.reqLog.StartTime)
		logger.Info("fail_req", zap.Any("request", ctx.reqLog))

		resB, _ := json.Marshal(resJson)
		if err != nil {
			_, _ = res.Write(resB)
		}
	}

	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse, ErrorHandler: errHandler}
}
