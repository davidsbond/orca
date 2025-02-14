package workflow

import "context"

type (
	Workflow interface {
		Run(ctx context.Context, input []byte) ([]byte, error)
		Name() string
	}

	Status int

	Run struct {
		ID string
	}

	ctxKey struct{}
)

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusScheduled
	StatusRunning
	StatusComplete
	StatusFailed
)

func RunToContext(ctx context.Context, r Run) context.Context {
	return context.WithValue(ctx, ctxKey{}, r)
}

func RunFromContext(ctx context.Context) (Run, bool) {
	run, ok := ctx.Value(ctxKey{}).(Run)
	return run, ok
}
