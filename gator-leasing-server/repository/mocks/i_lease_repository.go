// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "GatorLeasing/gator-leasing-server/model"

	mock "github.com/stretchr/testify/mock"
)

// ILeaseRepository is an autogenerated mock type for the ILeaseRepository type
type ILeaseRepository struct {
	mock.Mock
}

// CreateLease provides a mock function with given fields: lease
func (_m *ILeaseRepository) CreateLease(lease *model.Lease) (uint, error) {
	ret := _m.Called(lease)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Lease) (uint, error)); ok {
		return rf(lease)
	}
	if rf, ok := ret.Get(0).(func(*model.Lease) uint); ok {
		r0 = rf(lease)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(*model.Lease) error); ok {
		r1 = rf(lease)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditLease provides a mock function with given fields: lease
func (_m *ILeaseRepository) EditLease(lease *model.Lease) error {
	ret := _m.Called(lease)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Lease) error); ok {
		r0 = rf(lease)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllLeases provides a mock function with given fields:
func (_m *ILeaseRepository) GetAllLeases() ([]model.Lease, error) {
	ret := _m.Called()

	var r0 []model.Lease
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Lease, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Lease); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Lease)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewILeaseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewILeaseRepository creates a new instance of ILeaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewILeaseRepository(t mockConstructorTestingTNewILeaseRepository) *ILeaseRepository {
	mock := &ILeaseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
