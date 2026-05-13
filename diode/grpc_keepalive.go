package diode

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// defaultClientKeepaliveDialOption configures HTTP/2 pings for the Diode **ingester**
// client (`NewClient`). Do not attach this to OTLP exporter dials—generic OTLP collectors
// often reject idle pings (`GOAWAY` / too_many_pings). Intervals match the assurance
// reconciler gRPC client and are compatible with diode-pro reconciler-pro server policy
// (MinTime 10s; client Time 30s).
func defaultClientKeepaliveDialOption() grpc.DialOption {
	return grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                30 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	})
}
