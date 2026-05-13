package diode

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// defaultClientKeepaliveDialOption configures HTTP/2 pings so idle client connections
// are less likely to be torn down by middleboxes (L4/L7 proxies, NAT). Intervals match
// netbox-assurance-plugin reconciler gRPC client and are compatible with diode-pro
// reconciler-pro server keepalive enforcement (MinTime 10s; client Time 30s).
func defaultClientKeepaliveDialOption() grpc.DialOption {
	return grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                30 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	})
}
