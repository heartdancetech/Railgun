package core

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

var logger *zap.Logger

//var logLevel = zap.NewAtomicLevel()

func init() {
	proEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
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

//func Debug(msg string, fields ...zap.Field) {
//	logger.Debug(msg, fields...)
//}
//
//func Info(msg string, fields ...zap.Field) {
//	logger.Info(msg, fields...)
//}
//
//func Warn(msg string, fields ...zap.Field) {
//	logger.Warn(msg, fields...)
//}
//func Error(msg string, fields ...zap.Field) {
//	logger.Error(msg, fields...)
//}
//
//func Panic(msg string, fields ...zap.Field) {
//	logger.Panic(msg, fields...)
//}
//func DPanic(msg string, fields ...zap.Field) {
//	logger.DPanic(msg, fields...)
//}
//
//func Fatal(msg string, fields ...zap.Field) {
//	logger.Fatal(msg, fields...)
//}
