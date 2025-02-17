package workflow

import (
	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
	taskv1 "github.com/davidsbond/orca/internal/proto/orca/task/v1"
	workflowv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/v1"
)

type (
	Description struct {
		Run     workflow.Run
		Actions []Action
	}

	Action struct {
		WorkflowRun *Description
		TaskRun     *task.Run
	}
)

func (d Description) ToProto() *workflowv1.WorkflowRunDescription {
	out := &workflowv1.WorkflowRunDescription{
		Run: d.Run.ToProto(),
	}

	out.Actions = make([]*workflowv1.WorkflowAction, len(d.Actions))
	for i, action := range d.Actions {
		out.Actions[i] = action.ToProto()
	}

	return out
}

func (a Action) ToProto() *workflowv1.WorkflowAction {
	action := &workflowv1.WorkflowAction{}
	if a.WorkflowRun != nil {
		action.Action = &workflowv1.WorkflowAction_WorkflowRun{
			WorkflowRun: a.WorkflowRun.ToProto(),
		}
	}

	if a.TaskRun != nil {
		action.Action = &workflowv1.WorkflowAction_TaskRun{
			TaskRun: &taskv1.TaskRunDescription{
				Run: a.TaskRun.ToProto(),
			},
		}
	}

	return action
}
