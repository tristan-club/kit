package server

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/tristan-club/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"
	"time"
)

var (
	customFunc grpc_recovery.RecoveryHandlerFuncContext
)

var GrpcServer *grpc.Server

func Init() *grpc.Server {

	// Define customfunc to handle panic
	//Todo panic error
	customFunc = func(ctx context.Context, p interface{}) error {
		log.Error().Msgf("[PANIC] %s\n\n%s", p, string(debug.Stack()))
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(customFunc),
	}

	keepaliveConfig := keepalive.EnforcementPolicy{
		MinTime:             10 * time.Second,
		PermitWithoutStream: true,
	}

	// Create a server. Recovery handlers should typically be last in the chain_info so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic

	grpcServer := grpc.NewServer(

		grpc.KeepaliveEnforcementPolicy(keepaliveConfig),

		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
			//otgrpc.OpenTracingServerInterceptor(thisTracer),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
			//grpc_opentracing.StreamServerInterceptor(topts...),
		),
	)

	return grpcServer
}

func Start(s *grpc.Server, port string) error {
	grpcAddr := fmt.Sprintf("0.0.0.0:%s", port)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}

	reflection.Register(s)
	log.Info().Msgf("start grpc server at %s", grpcAddr)

	return s.Serve(lis)
}
