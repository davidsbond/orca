package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/davidsbond/orca/example/basic"
	"github.com/davidsbond/orca/pkg/orca"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	defer cancel()

	err := orca.Run(ctx,
		orca.WithWorkflows(
			basic.Workflow(),
		),
		orca.WithTasks(
			basic.Task(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
