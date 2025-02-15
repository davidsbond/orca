package daemon

import (
	"context"
	"fmt"
	"net"

	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
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
	grpcServer := grpc.NewServer()
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
