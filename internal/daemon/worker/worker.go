package worker

import (
	"context"
	"fmt"

	"github.com/davidsbond/orca/internal/daemon/controller/client"
	"github.com/davidsbond/orca/internal/daemon/worker/registry"
	"github.com/davidsbond/orca/internal/daemon/worker/scheduler"
	"github.com/davidsbond/orca/internal/daemon/worker/task"
	"github.com/davidsbond/orca/internal/daemon/worker/workflow"

	"golang.org/x/sync/errgroup"
)

type (
	Config struct {
		ControllerAddress string
		AdvertiseAddress  string
		HTTPPort          int
		GRPCPort          int
		Workflows         []workflow.Workflow
		Tasks             []task.Task
	}
)

func Run(ctx context.Context, cfg Config) error {
	controller, err := client.Dial(cfg.ControllerAddress)
	if err != nil {
		return fmt.Errorf("failed to dial controller: %w", err)
	}

	reg := registry.New()
	reg.RegisterWorkflows(cfg.Workflows...)
	reg.RegisterTasks(cfg.Tasks...)

	sched := scheduler.New(controller)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return sched.Run(ctx)
	})

	group.Go(func() error {
		<-ctx.Done()
		return controller.Close()
	})

	return group.Wait()
}
