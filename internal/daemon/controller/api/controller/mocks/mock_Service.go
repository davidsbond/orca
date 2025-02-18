// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	context "context"
	json "encoding/json"

	controller "github.com/davidsbond/orca/internal/daemon/controller/api/controller"

	mock "github.com/stretchr/testify/mock"

	task "github.com/davidsbond/orca/internal/daemon/controller/database/task"

	worker "github.com/davidsbond/orca/internal/daemon/controller/database/worker"

	workflow "github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// DeregisterWorker provides a mock function with given fields: ctx, id
func (_m *Service) DeregisterWorker(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeregisterWorker")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_DeregisterWorker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeregisterWorker'
type Service_DeregisterWorker_Call struct {
	*mock.Call
}

// DeregisterWorker is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Service_Expecter) DeregisterWorker(ctx interface{}, id interface{}) *Service_DeregisterWorker_Call {
	return &Service_DeregisterWorker_Call{Call: _e.mock.On("DeregisterWorker", ctx, id)}
}

func (_c *Service_DeregisterWorker_Call) Run(run func(ctx context.Context, id string)) *Service_DeregisterWorker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Service_DeregisterWorker_Call) Return(_a0 error) *Service_DeregisterWorker_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_DeregisterWorker_Call) RunAndReturn(run func(context.Context, string) error) *Service_DeregisterWorker_Call {
	_c.Call.Return(run)
	return _c
}

// GetTaskRun provides a mock function with given fields: ctx, id
func (_m *Service) GetTaskRun(ctx context.Context, id string) (task.Run, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskRun")
	}

	var r0 task.Run
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (task.Run, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) task.Run); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(task.Run)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_GetTaskRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTaskRun'
type Service_GetTaskRun_Call struct {
	*mock.Call
}

// GetTaskRun is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Service_Expecter) GetTaskRun(ctx interface{}, id interface{}) *Service_GetTaskRun_Call {
	return &Service_GetTaskRun_Call{Call: _e.mock.On("GetTaskRun", ctx, id)}
}

func (_c *Service_GetTaskRun_Call) Run(run func(ctx context.Context, id string)) *Service_GetTaskRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Service_GetTaskRun_Call) Return(_a0 task.Run, _a1 error) *Service_GetTaskRun_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_GetTaskRun_Call) RunAndReturn(run func(context.Context, string) (task.Run, error)) *Service_GetTaskRun_Call {
	_c.Call.Return(run)
	return _c
}

// GetWorkflowRun provides a mock function with given fields: ctx, id
func (_m *Service) GetWorkflowRun(ctx context.Context, id string) (workflow.Run, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetWorkflowRun")
	}

	var r0 workflow.Run
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (workflow.Run, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) workflow.Run); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(workflow.Run)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_GetWorkflowRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWorkflowRun'
type Service_GetWorkflowRun_Call struct {
	*mock.Call
}

// GetWorkflowRun is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Service_Expecter) GetWorkflowRun(ctx interface{}, id interface{}) *Service_GetWorkflowRun_Call {
	return &Service_GetWorkflowRun_Call{Call: _e.mock.On("GetWorkflowRun", ctx, id)}
}

func (_c *Service_GetWorkflowRun_Call) Run(run func(ctx context.Context, id string)) *Service_GetWorkflowRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Service_GetWorkflowRun_Call) Return(_a0 workflow.Run, _a1 error) *Service_GetWorkflowRun_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_GetWorkflowRun_Call) RunAndReturn(run func(context.Context, string) (workflow.Run, error)) *Service_GetWorkflowRun_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterWorker provides a mock function with given fields: ctx, w
func (_m *Service) RegisterWorker(ctx context.Context, w worker.Worker) error {
	ret := _m.Called(ctx, w)

	if len(ret) == 0 {
		panic("no return value specified for RegisterWorker")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, worker.Worker) error); ok {
		r0 = rf(ctx, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_RegisterWorker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterWorker'
type Service_RegisterWorker_Call struct {
	*mock.Call
}

// RegisterWorker is a helper method to define mock.On call
//   - ctx context.Context
//   - w worker.Worker
func (_e *Service_Expecter) RegisterWorker(ctx interface{}, w interface{}) *Service_RegisterWorker_Call {
	return &Service_RegisterWorker_Call{Call: _e.mock.On("RegisterWorker", ctx, w)}
}

func (_c *Service_RegisterWorker_Call) Run(run func(ctx context.Context, w worker.Worker)) *Service_RegisterWorker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(worker.Worker))
	})
	return _c
}

func (_c *Service_RegisterWorker_Call) Return(_a0 error) *Service_RegisterWorker_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_RegisterWorker_Call) RunAndReturn(run func(context.Context, worker.Worker) error) *Service_RegisterWorker_Call {
	_c.Call.Return(run)
	return _c
}

// ScheduleTask provides a mock function with given fields: ctx, params
func (_m *Service) ScheduleTask(ctx context.Context, params controller.ScheduleTaskParams) (string, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleTask")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, controller.ScheduleTaskParams) (string, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, controller.ScheduleTaskParams) string); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, controller.ScheduleTaskParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_ScheduleTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScheduleTask'
type Service_ScheduleTask_Call struct {
	*mock.Call
}

// ScheduleTask is a helper method to define mock.On call
//   - ctx context.Context
//   - params controller.ScheduleTaskParams
func (_e *Service_Expecter) ScheduleTask(ctx interface{}, params interface{}) *Service_ScheduleTask_Call {
	return &Service_ScheduleTask_Call{Call: _e.mock.On("ScheduleTask", ctx, params)}
}

func (_c *Service_ScheduleTask_Call) Run(run func(ctx context.Context, params controller.ScheduleTaskParams)) *Service_ScheduleTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(controller.ScheduleTaskParams))
	})
	return _c
}

func (_c *Service_ScheduleTask_Call) Return(_a0 string, _a1 error) *Service_ScheduleTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_ScheduleTask_Call) RunAndReturn(run func(context.Context, controller.ScheduleTaskParams) (string, error)) *Service_ScheduleTask_Call {
	_c.Call.Return(run)
	return _c
}

// ScheduleWorkflow provides a mock function with given fields: ctx, params
func (_m *Service) ScheduleWorkflow(ctx context.Context, params controller.ScheduleWorkflowParams) (string, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleWorkflow")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, controller.ScheduleWorkflowParams) (string, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, controller.ScheduleWorkflowParams) string); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, controller.ScheduleWorkflowParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_ScheduleWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScheduleWorkflow'
type Service_ScheduleWorkflow_Call struct {
	*mock.Call
}

// ScheduleWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - params controller.ScheduleWorkflowParams
func (_e *Service_Expecter) ScheduleWorkflow(ctx interface{}, params interface{}) *Service_ScheduleWorkflow_Call {
	return &Service_ScheduleWorkflow_Call{Call: _e.mock.On("ScheduleWorkflow", ctx, params)}
}

func (_c *Service_ScheduleWorkflow_Call) Run(run func(ctx context.Context, params controller.ScheduleWorkflowParams)) *Service_ScheduleWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(controller.ScheduleWorkflowParams))
	})
	return _c
}

func (_c *Service_ScheduleWorkflow_Call) Return(_a0 string, _a1 error) *Service_ScheduleWorkflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_ScheduleWorkflow_Call) RunAndReturn(run func(context.Context, controller.ScheduleWorkflowParams) (string, error)) *Service_ScheduleWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// SetTaskRunStatus provides a mock function with given fields: ctx, id, status, output
func (_m *Service) SetTaskRunStatus(ctx context.Context, id string, status task.Status, output json.RawMessage) error {
	ret := _m.Called(ctx, id, status, output)

	if len(ret) == 0 {
		panic("no return value specified for SetTaskRunStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, task.Status, json.RawMessage) error); ok {
		r0 = rf(ctx, id, status, output)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_SetTaskRunStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTaskRunStatus'
type Service_SetTaskRunStatus_Call struct {
	*mock.Call
}

// SetTaskRunStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - status task.Status
//   - output json.RawMessage
func (_e *Service_Expecter) SetTaskRunStatus(ctx interface{}, id interface{}, status interface{}, output interface{}) *Service_SetTaskRunStatus_Call {
	return &Service_SetTaskRunStatus_Call{Call: _e.mock.On("SetTaskRunStatus", ctx, id, status, output)}
}

func (_c *Service_SetTaskRunStatus_Call) Run(run func(ctx context.Context, id string, status task.Status, output json.RawMessage)) *Service_SetTaskRunStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(task.Status), args[3].(json.RawMessage))
	})
	return _c
}

func (_c *Service_SetTaskRunStatus_Call) Return(_a0 error) *Service_SetTaskRunStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_SetTaskRunStatus_Call) RunAndReturn(run func(context.Context, string, task.Status, json.RawMessage) error) *Service_SetTaskRunStatus_Call {
	_c.Call.Return(run)
	return _c
}

// SetWorkflowRunStatus provides a mock function with given fields: ctx, id, status, output
func (_m *Service) SetWorkflowRunStatus(ctx context.Context, id string, status workflow.Status, output json.RawMessage) error {
	ret := _m.Called(ctx, id, status, output)

	if len(ret) == 0 {
		panic("no return value specified for SetWorkflowRunStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, workflow.Status, json.RawMessage) error); ok {
		r0 = rf(ctx, id, status, output)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_SetWorkflowRunStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetWorkflowRunStatus'
type Service_SetWorkflowRunStatus_Call struct {
	*mock.Call
}

// SetWorkflowRunStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - status workflow.Status
//   - output json.RawMessage
func (_e *Service_Expecter) SetWorkflowRunStatus(ctx interface{}, id interface{}, status interface{}, output interface{}) *Service_SetWorkflowRunStatus_Call {
	return &Service_SetWorkflowRunStatus_Call{Call: _e.mock.On("SetWorkflowRunStatus", ctx, id, status, output)}
}

func (_c *Service_SetWorkflowRunStatus_Call) Run(run func(ctx context.Context, id string, status workflow.Status, output json.RawMessage)) *Service_SetWorkflowRunStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(workflow.Status), args[3].(json.RawMessage))
	})
	return _c
}

func (_c *Service_SetWorkflowRunStatus_Call) Return(_a0 error) *Service_SetWorkflowRunStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_SetWorkflowRunStatus_Call) RunAndReturn(run func(context.Context, string, workflow.Status, json.RawMessage) error) *Service_SetWorkflowRunStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
