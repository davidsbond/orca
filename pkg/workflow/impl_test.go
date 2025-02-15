package workflow_test

import (
	"context"

	"github.com/davidsbond/orca/pkg/task"
	"github.com/davidsbond/orca/pkg/workflow"
)

type (
	TestWorkflowInput struct {
		Name string
	}

	TestWorkflowOutput struct {
		Greeting string
	}

	TestTaskInput struct {
		Name string
	}

	TestTaskOutput struct {
		Greeting string
	}
)

func testWorkflow() workflow.Workflow {
	return &workflow.Implementation[TestWorkflowInput, TestWorkflowOutput]{
		WorkflowName: "TestWorkflow",
		Action: func(ctx context.Context, input TestWorkflowInput) (TestWorkflowOutput, error) {
			output, err := task.Execute[TestTaskOutput](ctx, testTask(), TestTaskInput{Name: input.Name})
			if err != nil {
				return TestWorkflowOutput{}, err
			}

			return TestWorkflowOutput{Greeting: output.Greeting}, nil
		},
	}
}

func testTask() task.Task {
	return &task.Implementation[TestTaskInput, TestTaskOutput]{
		TaskName: "TestTask",
		Action: func(ctx context.Context, input TestTaskInput) (TestTaskOutput, error) {
			return TestTaskOutput{
				Greeting: "Hello " + input.Name,
			}, nil
		},
	}
}
