package core

import (
	"encoding/json"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// NewReverseProxy Creat Reverse Proxy
func ReverseProxy() *httputil.ReverseProxy {
	type reqLog struct {
		Status     int
		Method     string
		Path       string
		Query      string
		Target     string
		IP         string
		RemoteAddr string
		UserAgent  string
		StartTime  time.Time
		EndTime    time.Time
		Latency    time.Duration
	}
	var log reqLog
	log.StartTime = time.Now()
	director := func(req *http.Request) {
		log.Path = req.URL.Path
		log.Method = req.Method
		log.Target = req.URL.Hostname()
		log.Query = req.URL.RawQuery
		log.IP = clientIP(req)
		log.RemoteAddr = clientIP(req)
		log.UserAgent = req.Header.Get("user-agent")
		return
	}

	modifyResponse := func(res *http.Response) error {
		log.EndTime = time.Now()
		log.Latency = log.EndTime.Sub(log.StartTime)
		log.Status = res.StatusCode
		logger.Info("",
			zap.Int("status", log.Status),
			zap.String("method", log.Method),
			zap.String("path", log.Path),
			zap.String("query", log.Query),
			zap.String("ip", log.RemoteAddr),
			zap.String("target", log.Target),
			zap.String("user_agent", log.UserAgent),
			zap.String("start_time", log.StartTime.UTC().Format("2006-01-02T15:04:05.000000-07:00")),
			zap.String("end_time", log.EndTime.UTC().Format("2006-01-02T15:04:05.000000-07:00")),
			zap.Duration("latency", log.Latency))
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
		log.EndTime = time.Now()
		log.Latency = log.EndTime.Sub(log.StartTime)
		logger.Error("", zap.Error(err),
			zap.Int("status", log.Status),
			zap.String("method", log.Method),
			zap.String("path", log.Path),
			zap.String("query", log.Query),
			zap.String("ip", log.IP),
			zap.String("target", log.Target),
			zap.String("user_agent", log.UserAgent),
			zap.String("start_time", log.StartTime.UTC().Format("2006-01-02T15:04:05.000000-07:00")),
			zap.String("end_time", log.EndTime.UTC().Format("2006-01-02T15:04:05.000000-07:00")),
			zap.Duration("latency", log.Latency))

		resB, _ := json.Marshal(resJson)
		if err != nil {
			_, _ = res.Write(resB)
		}
	}

	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse, ErrorHandler: errHandler}
}

func clientIP(req *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
