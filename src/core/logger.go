package core

import (
	"io"
	"time"

	"github.com/lestrrat-go/file-rotatelogs "
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

type reqLog struct {
	Status     int           `json:"status"`
	Method     string        `json:"method"`
	Path       string        `json:"path"`
	Query      string        `json:"query"`
	Target     string        `json:"target"`
	IP         string        `json:"ip"`
	RemoteAddr string        `json:"remote_addr"`
	UserAgent  string        `json:"user_agent"`
	StartTime  time.Time     `json:"start_time"`
	EndTime    time.Time     `json:"end_time"`
	Latency    time.Duration `json:"latency"`
}

func init() {
	proEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format(time.RFC3339))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.InfoLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl > zapcore.DebugLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	debugHook := getWriter("./log/" + "gateway_debug")
	infoHook := getWriter("./log/" + "gateway_request")
	warnHook := getWriter("./log/" + "gateway_error")
	var core zapcore.Core

	core = zapcore.NewTee(
		zapcore.NewCore(proEncoder, zapcore.AddSync(debugHook), debugLevel),
		zapcore.NewCore(proEncoder, zapcore.AddSync(infoHook), infoLevel),
		zapcore.NewCore(proEncoder, zapcore.AddSync(warnHook), warnLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

type Level int8

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
