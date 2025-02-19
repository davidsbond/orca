package basic

import (
	"context"
	"time"

	"github.com/davidsbond/orca/pkg/orca"
)

type (
	GreetWorkflowInput struct {
		Name string
	}

	GreetWorkflowOutput struct {
		Greeting string
	}
)

func Workflow() *orca.Workflow[GreetWorkflowInput, GreetWorkflowOutput] {
	return &orca.Workflow[GreetWorkflowInput, GreetWorkflowOutput]{
		WorkflowName: "GreetWorkflow",
		Action:       greetWorkflow,
		Options: orca.ActionOptions{
			RetryCount: 3,
			Timeout:    time.Hour,
		},
	}
}

func greetWorkflow(ctx context.Context, input GreetWorkflowInput) (GreetWorkflowOutput, error) {
	output, err := orca.ExecuteTask(ctx, Task(), GreetTaskInput{Name: input.Name})
	if err != nil {
		return GreetWorkflowOutput{}, err
	}

	for {
		select {
		case <-ctx.Done():
			return GreetWorkflowOutput{Greeting: output.Greeting}, ctx.Err()
		}
	}

	//return GreetWorkflowOutput{Greeting: output.Greeting}, nil
}

type (
	GreetTaskInput struct {
		Name string
	}

	GreetTaskOutput struct {
		Greeting string
	}
)

func Task() *orca.Task[GreetTaskInput, GreetTaskOutput] {
	return &orca.Task[GreetTaskInput, GreetTaskOutput]{
		TaskName: "GreetTask",
		Action:   greetTask,
		KeyFunc: func(input GreetTaskInput) string {
			return input.Name
		},
		Options: orca.ActionOptions{
			RetryCount: 3,
			Timeout:    time.Minute,
		},
	}
}

func greetTask(_ context.Context, input GreetTaskInput) (GreetTaskOutput, error) {
	return GreetTaskOutput{Greeting: "Hello " + input.Name + "!"}, nil
}
