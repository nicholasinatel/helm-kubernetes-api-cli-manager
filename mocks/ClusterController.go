// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ClusterController is an autogenerated mock type for the ClusterController type
type ClusterController struct {
	mock.Mock
}

// GetNodesAndCandidate provides a mock function with given fields:
func (_m *ClusterController) GetNodesAndCandidate() (string, int, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPodsAndCandidate provides a mock function with given fields:
func (_m *ClusterController) GetPodsAndCandidate() (string, int, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewClusterController interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterController creates a new instance of ClusterController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterController(t mockConstructorTestingTNewClusterController) *ClusterController {
	mock := &ClusterController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
