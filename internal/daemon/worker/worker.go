package worker

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/daemon"
	"github.com/davidsbond/orca/internal/daemon/controller/client"
	"github.com/davidsbond/orca/internal/daemon/worker/api"
	"github.com/davidsbond/orca/internal/daemon/worker/registry"
	"github.com/davidsbond/orca/internal/daemon/worker/scheduler"
	"github.com/davidsbond/orca/internal/task"
	"github.com/davidsbond/orca/internal/workflow"
)

type (
	Config struct {
		ID                string
		ControllerAddress string
		AdvertiseAddress  string
		Port              int
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

	workerSvc := api.NewWorkerService(reg, sched)
	workerAPI := api.New(workerSvc)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return sched.Run(ctx)
	})

	group.Go(func() error {
		return daemon.Run(ctx, daemon.Config{
			GRPCPort: cfg.Port,
			GRPCControllers: []daemon.GRPCController{
				workerAPI,
			},
		})
	})

	group.Go(func() error {
		return controller.RegisterWorker(ctx, client.WorkerInfo{
			ID:               cfg.ID,
			AdvertiseAddress: cfg.AdvertiseAddress,
			Workflows:        reg.Workflows(),
			Tasks:            reg.Tasks(),
		})
	})

	group.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		return errors.Join(
			controller.DeregisterWorker(ctx, cfg.ID),
			controller.Close(),
		)
	})

	return group.Wait()
}
