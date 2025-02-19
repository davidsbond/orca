package workflow

import (
	"context"
	"errors"
	"time"
)

type (
	Context struct {
		parent context.Context
		cancel chan struct{}
		err    error
		run    Run
		client Client
	}
)

var (
	ErrTimeout   = errors.New("workflow timeout")
	ErrCancelled = errors.New("workflow cancelled")
)

func NewContext(parent context.Context, run Run, client Client) (*Context, context.CancelFunc) {
	ctx := &Context{
		parent: parent,
		cancel: make(chan struct{}, 1),
		run:    run,
		client: client,
	}

	fn := func() {
		if ctx.err != nil {
			return
		}

		ctx.err = ErrCancelled
		ctx.cancel <- struct{}{}
	}

	return ctx, fn
}

func (c *Context) Run() Run {
	return c.run
}

func (c *Context) SetTimeout(timeout time.Duration) {
	go func() {
		timer := time.NewTimer(timeout)
		defer timer.Stop()

		<-timer.C
		c.err = ErrTimeout
		c.cancel <- struct{}{}
		close(c.cancel)
	}()
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	out := make(chan struct{})

	go func() {
		defer close(out)

		for {
			if c.err != nil || c.parent.Err() != nil {
				out <- struct{}{}
				return
			}
		}
	}()

	return out
}

func (c *Context) Err() error {
	if c.err != nil {
		return c.err
	}

	err := c.parent.Err()
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return ErrTimeout
	case err != nil:
		return err
	}

	return nil
}

func (c *Context) Value(key any) any {
	return c.parent.Value(key)
}

func RunFromContext(ctx context.Context) (Run, bool) {
	if wCtx, ok := ctx.(*Context); ok {
		return wCtx.run, true
	}

	return Run{}, false
}

func ClientFromContext(ctx context.Context) (Client, bool) {
	if wCtx, ok := ctx.(*Context); ok {
		return wCtx.client, true
	}

	return nil, false
}
