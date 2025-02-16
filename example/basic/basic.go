package basic

import (
	"context"

	"github.com/davidsbond/orca/pkg/task"
	"github.com/davidsbond/orca/pkg/workflow"
)

type (
	GreetWorkflowInput struct {
		Name string
	}

	GreetWorkflowOutput struct {
		Greeting string
	}
)

func Workflow() *workflow.Implementation[GreetWorkflowInput, GreetWorkflowOutput] {
	return &workflow.Implementation[GreetWorkflowInput, GreetWorkflowOutput]{
		WorkflowName: "GreetWorkflow",
		Action:       greetWorkflow,
	}
}

func greetWorkflow(ctx context.Context, input GreetWorkflowInput) (GreetWorkflowOutput, error) {
	output, err := task.Execute(ctx, Task(), GreetTaskInput{Name: input.Name})
	if err != nil {
		return GreetWorkflowOutput{}, err
	}

	return GreetWorkflowOutput{Greeting: output.Greeting}, nil
}

type (
	GreetTaskInput struct {
		Name string
	}

	GreetTaskOutput struct {
		Greeting string
	}
)

func Task() *task.Implementation[GreetTaskInput, GreetTaskOutput] {
	return &task.Implementation[GreetTaskInput, GreetTaskOutput]{
		TaskName: "GreetTask",
		Action:   greetTask,
	}
}

func greetTask(_ context.Context, input GreetTaskInput) (GreetTaskOutput, error) {
	return GreetTaskOutput{Greeting: "Hello " + input.Name + "!"}, nil
}
