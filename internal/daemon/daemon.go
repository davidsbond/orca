package daemon

import (
	"context"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/davidsbond/orca/internal/log"
)

type (
	Config struct {
		GRPCPort        int
		GRPCControllers []GRPCController
	}

	GRPCController interface {
		Register(r grpc.ServiceRegistrar)
	}
)

func Run(ctx context.Context, cfg Config) error {
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return runGRPC(ctx, cfg.GRPCPort, cfg.GRPCControllers...)
	})

	return group.Wait()
}

func runGRPC(ctx context.Context, port int, controllers ...GRPCController) error {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(log.Interceptor(ctx)),
			recovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(log.Interceptor(ctx)),
			recovery.StreamServerInterceptor(),
		),
	)
	for _, ctrl := range controllers {
		ctrl.Register(grpcServer)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return grpcServer.Serve(lis)
	})
	group.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()
		return ctx.Err()
	})

	return group.Wait()
}

func Dial(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(log.Interceptor(ctx)),
		),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(log.Interceptor(ctx)),
		),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
