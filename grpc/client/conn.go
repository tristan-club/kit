package client

import (
	"context"
	"fmt"
	"github.com/tristan-club/kit/config"
	"github.com/tristan-club/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"math/rand"
	"time"
)

func GetDefault(clientAddr string) (*grpc.ClientConn, error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             3 * time.Second,  // wait 3 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	return grpc.Dial(clientAddr,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithUnaryInterceptor(interceptor))
}

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// calculate method time spend
	start := time.Now()
	var traceId string

	if ctx.Value("trace_id") == nil {
		traceId = fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(10000))
		ctx = context.WithValue(ctx, "trace_id", traceId)
	} else {
		traceId = ctx.Value("trace_id").(string)
	}

	if !config.IgnoreTraceId() {
		log.Info().Msgf("[%s] grpc start [%s]", traceId, method)
	}
	err := invoker(ctx, method, req, reply, cc, opts...)

	if !config.IgnoreTraceId() {
		log.Info().Msgf("[%s] grpc end [%s] time spend [%d]ms", traceId, method, time.Since(start).Milliseconds())
	}

	return err
}
