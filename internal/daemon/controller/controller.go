package controller

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/daemon/controller/api"
	"github.com/davidsbond/orca/internal/daemon/controller/database"
	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
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

	controllerSvc := api.NewControllerService(workers, tasks, workflows)
	controllerAPI := api.New(controllerSvc)

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		return db.Close(ctx)
	})

	return group.Wait()
}
