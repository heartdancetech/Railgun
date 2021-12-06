package core

import "os"

const EnvMode = "GATEWAY_MODE"

const (
	// DebugMode indicates  mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates  mode is release.
	ReleaseMode = "release"
)

var gatewayMode = DebugMode

func init() {
	mode := os.Getenv(EnvMode)
	SetMode(mode)
}

func SetMode(value string) {
	switch value {
	case DebugMode, "":
		gatewayMode = DebugMode
	case ReleaseMode:
		gatewayMode = ReleaseMode
	default:
		gatewayMode = DebugMode
	}
}

func Mode() string {
	return gatewayMode
}
