package orca

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/davidsbond/orca/internal/task"
	"github.com/davidsbond/orca/internal/workflow"
)

func ExecuteTask[Input, Output any](ctx context.Context, t *Task[Input, Output], input Input) (Output, error) {
	var output Output

	inp, err := json.Marshal(input)
	if err != nil {
		return output, fmt.Errorf("failed to marshal input: %w", err)
	}

	run, ok := workflow.RunFromContext(ctx)
	if !ok {
		return output, errors.New("workflow not present in context")
	}

	client, ok := workflow.ClientFromContext(ctx)
	if !ok {
		return output, errors.New("client not present in context")
	}

	var key string
	if t.KeyFunc != nil {
		key = t.KeyFunc(input)
	}

	runID, err := client.ScheduleTask(ctx, workflow.ScheduleTaskParams{
		WorkflowRunID: run.ID,
		TaskName:      t.TaskName,
		IdempotentKey: key,
		Input:         inp,
	})
	if err != nil {
		return output, fmt.Errorf("could not schedule task: %w", err)
	}

	ticker := time.NewTicker(time.Second / 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return output, ctx.Err()
		case <-ticker.C:
			taskRun, err := client.GetTaskRun(ctx, runID)
			if err != nil {
				return output, err
			}

			if taskRun.Status == task.StatusTimeout {
				return output, task.Error{
					Message:  "task timed out",
					TaskName: taskRun.TaskName,
					RunID:    runID,
				}
			}

			if taskRun.Status == task.StatusComplete || taskRun.Status == task.StatusSkipped {
				if err = json.Unmarshal(taskRun.Output, &output); err != nil {
					return output, fmt.Errorf("failed to unmarshal output: %w", err)
				}

				return output, nil
			}

			if taskRun.Status == task.StatusFailed {
				var e task.Error
				if err = json.Unmarshal(taskRun.Output, &e); err != nil {
					return output, fmt.Errorf("failed to unmarshal output: %w", err)
				}

				return output, e
			}
		}
	}
}
