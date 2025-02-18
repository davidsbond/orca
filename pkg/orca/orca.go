package orca

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/davidsbond/orca/internal/daemon/worker"
	"github.com/davidsbond/orca/internal/log"
	"github.com/davidsbond/orca/internal/task"
	"github.com/davidsbond/orca/internal/workflow"
)

type (
	options struct {
		id                string
		controllerAddress string
		advertiseAddress  string
		port              int
		workflows         []workflow.Workflow
		tasks             []task.Task
		logger            *slog.Logger
	}

	Option func(opts *options)
)

func Run(ctx context.Context, opts ...Option) error {
	o := &options{
		id:                uuid.NewString(),
		controllerAddress: "localhost:4001",
		advertiseAddress:  "localhost:4002",
		port:              4002,
		logger:            slog.Default(),
	}

	for _, opt := range opts {
		opt(o)
	}

	ctx = log.ToContext(ctx, o.logger)

	return worker.Run(ctx, worker.Config{
		ID:                o.id,
		ControllerAddress: o.controllerAddress,
		AdvertiseAddress:  o.advertiseAddress,
		Port:              o.port,
		Workflows:         o.workflows,
		Tasks:             o.tasks,
	})
}

func WithWorkerID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func WithControllerAddress(addr string) Option {
	return func(o *options) {
		o.controllerAddress = addr
	}
}

func WithPort(port int) Option {
	return func(o *options) {
		o.port = port
	}
}

func WithAdvertiseAddress(addr string) Option {
	return func(o *options) {
		o.advertiseAddress = addr
	}
}

func WithWorkflows(workflows ...workflow.Workflow) Option {
	return func(o *options) {
		o.workflows = append(o.workflows, workflows...)
	}
}

func WithTasks(tasks ...task.Task) Option {
	return func(o *options) {
		o.tasks = append(o.tasks, tasks...)
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

type (
	KeyFunc[T any] func(T) string

	Action[Input any, Output any] func(ctx context.Context, input Input) (Output, error)

	ActionOptions struct {
		RetryCount int
		Timeout    time.Duration
	}
)
