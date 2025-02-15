package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/davidsbond/orca/internal/workflow"

	"github.com/davidsbond/orca/internal/task"
)

type (
	Task interface {
		Name() string
	}
)

func Execute[Output any](ctx context.Context, t Task, input any) (Output, error) {
	var output Output

	params, err := json.Marshal(input)
	if err != nil {
		return output, fmt.Errorf("failed to marshal input: %w", err)
	}

	run, ok := workflow.RunFromContext(ctx)
	if !ok {
		return output, errors.New("workflow not present in context")
	}

	client, ok := task.ClientFromContext(ctx)
	if !ok {
		return output, errors.New("client not present in context")
	}

	runID, err := client.ScheduleTask(ctx, run.ID, t.Name(), params)
	if err != nil {
		return output, fmt.Errorf("could not schedule task: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return output, ctx.Err()
		default:
			taskRun, err := client.GetTaskRun(ctx, runID)
			if err != nil {
				return output, err
			}

			switch taskRun.Status {
			case task.StatusRunning, task.StatusPending, task.StatusUnspecified:
				continue
			case task.StatusComplete:
				if err = json.Unmarshal(taskRun.Output, &output); err != nil {
					return output, fmt.Errorf("failed to unmarshal output: %w", err)
				}

				return output, nil
			case task.StatusFailed:
				return output, errors.New(string(taskRun.Output))
			default:
				continue
			}
		}
	}
}
