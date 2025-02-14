package registry

import (
	"sync"

	"github.com/davidsbond/orca/pkg/task"
	"github.com/davidsbond/orca/pkg/workflow"
)

type (
	Registry struct {
		mux       sync.RWMutex
		workflows map[string]workflow.Workflow
		tasks     map[string]task.Task
	}
)

func New() *Registry {
	return &Registry{
		workflows: make(map[string]workflow.Workflow),
		tasks:     make(map[string]task.Task),
	}
}

func (r *Registry) RegisterWorkflows(workflows ...workflow.Workflow) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, w := range workflows {
		r.workflows[w.Name()] = w
	}
}

func (r *Registry) RegisterTasks(tasks ...task.Task) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, t := range tasks {
		r.tasks[t.Name()] = t
	}
}

func (r *Registry) GetWorkflow(name string) (workflow.Workflow, bool) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	w, ok := r.workflows[name]
	return w, ok
}

func (r *Registry) GetTask(name string) (task.Task, bool) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	w, ok := r.tasks[name]
	return w, ok
}
