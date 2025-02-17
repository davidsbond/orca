package controller

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/daemon"
	"github.com/davidsbond/orca/internal/daemon/controller/api/controller"
	workflow_api "github.com/davidsbond/orca/internal/daemon/controller/api/workflow"
	"github.com/davidsbond/orca/internal/daemon/controller/database"
	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
	"github.com/davidsbond/orca/internal/daemon/controller/runner"
)

type (
	Config struct {
		GRPCPort    int
		HTTPPort    int
		DatabaseURL string
	}
)

func Run(ctx context.Context, config Config) error {
	db, err := database.Open(ctx, config.DatabaseURL)
	if err != nil {
		return err
	}

	workers := worker.NewPostgresRepository(db)
	tasks := task.NewPostgresRepository(db)
	workflows := workflow.NewPostgresRepository(db)

	run := runner.New(workers, workflows, tasks)

	controllerHandler := controller.NewHandler(workers, tasks, workflows)
	controllerAPI := controller.NewAPI(controllerHandler)

	workflowHandler := workflow_api.NewHandler(workflows)
	workflowAPI := workflow_api.NewAPI(workflowHandler)

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return daemon.Run(ctx, daemon.Config{
			GRPCPort:  config.GRPCPort,
			HTTPPort:  config.HTTPPort,
			ServeHTTP: true,
			Controllers: []daemon.Controller{
				controllerAPI,
				workflowAPI,
			},
		})
	})

	group.Go(func() error {
		return run.Run(ctx)
	})

	group.Go(func() error {
		<-ctx.Done()
		db.Close()
		return nil
	})

	return group.Wait()
}
