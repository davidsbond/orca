package main_test

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/davidsbond/orca/internal/log"

	"github.com/davidsbond/orca/internal/daemon/worker"
	task2 "github.com/davidsbond/orca/internal/task"
	workflow2 "github.com/davidsbond/orca/internal/workflow"
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

	TestChildWorkflowInput struct {
		Name string
	}

	TestChildWorkflowOutput struct {
		Extra string
	}
)

func testWorkflow() *workflow.Implementation[TestWorkflowInput, TestWorkflowOutput] {
	return &workflow.Implementation[TestWorkflowInput, TestWorkflowOutput]{
		WorkflowName: "TestWorkflow",
		Action: func(ctx context.Context, input TestWorkflowInput) (TestWorkflowOutput, error) {
			taskOutput, err := task.Execute(ctx, testTask(), TestTaskInput{Name: input.Name})
			if err != nil {
				return TestWorkflowOutput{}, err
			}

			workflowOutput, err := workflow.Execute(ctx, testChildWorkflow(), TestChildWorkflowInput{
				Name: input.Name,
			})
			if err != nil {
				return TestWorkflowOutput{}, err
			}

			return TestWorkflowOutput{Greeting: taskOutput.Greeting + " " + workflowOutput.Extra}, nil
		},
	}
}

func testTask() *task.Implementation[TestTaskInput, TestTaskOutput] {
	return &task.Implementation[TestTaskInput, TestTaskOutput]{
		TaskName: "TestTask",
		Action: func(ctx context.Context, input TestTaskInput) (TestTaskOutput, error) {
			return TestTaskOutput{
				Greeting: "Hello " + input.Name,
			}, nil
		},
	}
}

func testChildWorkflow() *workflow.Implementation[TestChildWorkflowInput, TestChildWorkflowOutput] {
	return &workflow.Implementation[TestChildWorkflowInput, TestChildWorkflowOutput]{
		WorkflowName: "TestChildWorkflow",
		Action: func(ctx context.Context, input TestChildWorkflowInput) (TestChildWorkflowOutput, error) {
			taskOutput, err := task.Execute(ctx, testTask(), TestTaskInput{Name: input.Name})
			if err != nil {
				return TestChildWorkflowOutput{}, err
			}

			var builder strings.Builder
			for i := len(taskOutput.Greeting) - 1; i >= 0; i-- {
				builder.WriteRune(rune(taskOutput.Greeting[i]))
			}

			panic("bingus bongus")

			return TestChildWorkflowOutput{
				Extra: builder.String(),
			}, nil
		},
	}
}

func TestWorker(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), time.Hour)
	defer cancel()

	ctx = log.ToContext(ctx, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	err := worker.Run(ctx, worker.Config{
		ID:                "65e1dee9-d017-46c3-84c0-e39edd548075",
		ControllerAddress: "localhost:8081",
		AdvertiseAddress:  "localhost:8082",
		Port:              8082,
		Workflows: []workflow2.Workflow{
			testWorkflow(),
			testChildWorkflow(),
		},
		Tasks: []task2.Task{
			testTask(),
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
