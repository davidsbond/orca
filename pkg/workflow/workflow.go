package workflow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/davidsbond/orca/internal/workflow"
)

func Execute[Input, Output any](ctx context.Context, w *Implementation[Input, Output], input Input) (Output, error) {
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
	if w.KeyFunc != nil {
		key = w.KeyFunc(input)
	}

	runID, err := client.ScheduleWorkflow(ctx, workflow.ScheduleWorkflowParams{
		WorkflowRunID: run.ID,
		WorkflowName:  w.WorkflowName,
		IdempotentKey: key,
		Input:         inp,
	})
	if err != nil {
		return output, fmt.Errorf("could not schedule workflow: %w", err)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return output, ctx.Err()
		case <-ticker.C:
			workflowRun, err := client.GetWorkflowRun(ctx, runID)
			if err != nil {
				return output, err
			}

			switch workflowRun.Status {
			case workflow.StatusRunning, workflow.StatusPending, workflow.StatusUnspecified:
				continue
			case workflow.StatusComplete, workflow.StatusSkipped:
				if err = json.Unmarshal(workflowRun.Output, &output); err != nil {
					return output, fmt.Errorf("failed to unmarshal output: %w", err)
				}

				return output, nil
			case workflow.StatusFailed:
				var e workflow.Error
				if err = json.Unmarshal(workflowRun.Output, &e); err != nil {
					return output, fmt.Errorf("failed to unmarshal output: %w", err)
				}

				return output, e
			default:
				continue
			}
		}
	}
}
