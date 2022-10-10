// Code generated by mockery v2.12.3. DO NOT EDIT.

package notification

import (
	context "context"

	model "github.com/goharbor/harbor/src/pkg/notification/policy/model"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Manager) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *Manager) Create(ctx context.Context, _a1 *model.Policy) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.Policy) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Policy) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, policyID
func (_m *Manager) Delete(ctx context.Context, policyID int64) error {
	ret := _m.Called(ctx, policyID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, policyID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Manager) Get(ctx context.Context, id int64) (*model.Policy, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Policy
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.Policy); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Policy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNameAndProjectID provides a mock function with given fields: ctx, name, projectID
func (_m *Manager) GetByNameAndProjectID(ctx context.Context, name string, projectID int64) (*model.Policy, error) {
	ret := _m.Called(ctx, name, projectID)

	var r0 *model.Policy
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *model.Policy); ok {
		r0 = rf(ctx, name, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Policy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, name, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRelatedPolices provides a mock function with given fields: ctx, projectID, eventType
func (_m *Manager) GetRelatedPolices(ctx context.Context, projectID int64, eventType string) ([]*model.Policy, error) {
	ret := _m.Called(ctx, projectID, eventType)

	var r0 []*model.Policy
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) []*model.Policy); ok {
		r0 = rf(ctx, projectID, eventType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Policy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, projectID, eventType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *Manager) List(ctx context.Context, query *q.Query) ([]*model.Policy, error) {
	ret := _m.Called(ctx, query)

	var r0 []*model.Policy
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.Policy); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Policy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Test provides a mock function with given fields: _a0
func (_m *Manager) Test(_a0 *model.Policy) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Policy) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Manager) Update(ctx context.Context, _a1 *model.Policy) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Policy) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewManagerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t NewManagerT) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}