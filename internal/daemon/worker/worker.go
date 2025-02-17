package worker

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/daemon"
	"github.com/davidsbond/orca/internal/daemon/controller/api/controller"
	"github.com/davidsbond/orca/internal/daemon/worker/api/worker"
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
	client, err := controller.Dial(ctx, cfg.ControllerAddress)
	if err != nil {
		return fmt.Errorf("failed to dial controller: %w", err)
	}

	reg := registry.New()
	reg.RegisterWorkflows(cfg.Workflows...)
	reg.RegisterTasks(cfg.Tasks...)

	sched := scheduler.New(client)

	workerSvc := worker.NewHandler(reg, sched)
	workerAPI := worker.NewAPI(workerSvc)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return sched.Run(ctx)
	})

	group.Go(func() error {
		return daemon.Run(ctx, daemon.Config{
			GRPCPort: cfg.Port,
			Controllers: []daemon.Controller{
				workerAPI,
			},
		})
	})

	group.Go(func() error {
		return client.RegisterWorker(ctx, controller.WorkerInfo{
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
			client.DeregisterWorker(ctx, cfg.ID),
			client.Close(),
		)
	})

	return group.Wait()
}
