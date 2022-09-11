// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "SmartLink_Project/domain"

	mock "github.com/stretchr/testify/mock"
)

// LayananUseCase is an autogenerated mock type for the LayananUseCase type
type LayananUseCase struct {
	mock.Mock
}

// AddLayanan provides a mock function with given fields: newLayanan
func (_m *LayananUseCase) AddLayanan(newLayanan domain.Layanan) (int, domain.Layanan) {
	ret := _m.Called(newLayanan)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Layanan) int); ok {
		r0 = rf(newLayanan)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 domain.Layanan
	if rf, ok := ret.Get(1).(func(domain.Layanan) domain.Layanan); ok {
		r1 = rf(newLayanan)
	} else {
		r1 = ret.Get(1).(domain.Layanan)
	}

	return r0, r1
}

type mockConstructorTestingTNewLayananUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewLayananUseCase creates a new instance of LayananUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLayananUseCase(t mockConstructorTestingTNewLayananUseCase) *LayananUseCase {
	mock := &LayananUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}