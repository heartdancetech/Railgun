package core

import (
	"context"
	"net"
)

type srvResolver interface {
	LookupSRV(context.Context, string, string, string) (string, []*net.SRV, error)
}
