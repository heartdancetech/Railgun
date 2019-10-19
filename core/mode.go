package core

import "os"

const EnvMode = "GATEWAY_MODE"

const (
	// DebugMode indicates  mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates  mode is release.
	ReleaseMode = "release"
)
const (
	debugCode = iota
	releaseCode
)

var gatewayMode = debugCode
var modeName = DebugMode

func init() {
	mode := os.Getenv(EnvMode)
	SetMode(mode)
}

func SetMode(value string) {
	switch value {
	case DebugMode, "":
		gatewayMode = debugCode
	case ReleaseMode:
		gatewayMode = releaseCode
	default:
		panic("mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}

func Mode() string {
	return modeName
}
