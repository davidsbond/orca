// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	workflow "github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
)

// WorkflowRepository is an autogenerated mock type for the WorkflowRepository type
type WorkflowRepository struct {
	mock.Mock
}

type WorkflowRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *WorkflowRepository) EXPECT() *WorkflowRepository_Expecter {
	return &WorkflowRepository_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *WorkflowRepository) Get(ctx context.Context, id string) (workflow.Run, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
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

// WorkflowRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type WorkflowRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *WorkflowRepository_Expecter) Get(ctx interface{}, id interface{}) *WorkflowRepository_Get_Call {
	return &WorkflowRepository_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *WorkflowRepository_Get_Call) Run(run func(ctx context.Context, id string)) *WorkflowRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *WorkflowRepository_Get_Call) Return(_a0 workflow.Run, _a1 error) *WorkflowRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WorkflowRepository_Get_Call) RunAndReturn(run func(context.Context, string) (workflow.Run, error)) *WorkflowRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, run
func (_m *WorkflowRepository) Save(ctx context.Context, run workflow.Run) error {
	ret := _m.Called(ctx, run)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workflow.Run) error); ok {
		r0 = rf(ctx, run)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WorkflowRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type WorkflowRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - run workflow.Run
func (_e *WorkflowRepository_Expecter) Save(ctx interface{}, run interface{}) *WorkflowRepository_Save_Call {
	return &WorkflowRepository_Save_Call{Call: _e.mock.On("Save", ctx, run)}
}

func (_c *WorkflowRepository_Save_Call) Run(run func(ctx context.Context, run workflow.Run)) *WorkflowRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workflow.Run))
	})
	return _c
}

func (_c *WorkflowRepository_Save_Call) Return(_a0 error) *WorkflowRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WorkflowRepository_Save_Call) RunAndReturn(run func(context.Context, workflow.Run) error) *WorkflowRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewWorkflowRepository creates a new instance of WorkflowRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWorkflowRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WorkflowRepository {
	mock := &WorkflowRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
