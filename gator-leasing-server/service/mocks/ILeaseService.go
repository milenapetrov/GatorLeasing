// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "GatorLeasing/gator-leasing-server/entity"

	mock "github.com/stretchr/testify/mock"
)

// ILeaseService is an autogenerated mock type for the ILeaseService type
type ILeaseService struct {
	mock.Mock
}

// CreateLease provides a mock function with given fields: request
func (_m *ILeaseService) CreateLease(request *entity.CreateLeaseRequest) (uint, error) {
	ret := _m.Called(request)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.CreateLeaseRequest) (uint, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*entity.CreateLeaseRequest) uint); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(*entity.CreateLeaseRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLease provides a mock function with given fields: request
func (_m *ILeaseService) DeleteLease(request *entity.DeleteLeaseRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.DeleteLeaseRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditLease provides a mock function with given fields: request
func (_m *ILeaseService) EditLease(request *entity.EditLeaseRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.EditLeaseRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllLeases provides a mock function with given fields:
func (_m *ILeaseService) GetAllLeases() ([]*entity.Lease, error) {
	ret := _m.Called()

	var r0 []*entity.Lease
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entity.Lease, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entity.Lease); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Lease)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewILeaseService interface {
	mock.TestingT
	Cleanup(func())
}

// NewILeaseService creates a new instance of ILeaseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewILeaseService(t mockConstructorTestingTNewILeaseService) *ILeaseService {
	mock := &ILeaseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
