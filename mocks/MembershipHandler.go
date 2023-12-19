// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// MembershipHandler is an autogenerated mock type for the MembershipHandler type
type MembershipHandler struct {
	mock.Mock
}

// DeleteMembershipHandler provides a mock function with given fields: ctx
func (_m *MembershipHandler) DeleteMembershipHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMembershipByNameHandler provides a mock function with given fields: ctx
func (_m *MembershipHandler) GetMembershipByNameHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMembershipByPhoneNumber provides a mock function with given fields: ctx
func (_m *MembershipHandler) GetMembershipByPhoneNumber(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMembershipHandler provides a mock function with given fields: ctx
func (_m *MembershipHandler) GetMembershipHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMembershipsHandler provides a mock function with given fields: ctx
func (_m *MembershipHandler) GetMembershipsHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTopMembershipsHandler provides a mock function with given fields: ctx
func (_m *MembershipHandler) GetTopMembershipsHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterMembershipHandler provides a mock function with given fields: ctx
func (_m *MembershipHandler) RegisterMembershipHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMembershipHandler provides a mock function with given fields: ctx
func (_m *MembershipHandler) UpdateMembershipHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMembershipHandler creates a new instance of MembershipHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMembershipHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MembershipHandler {
	mock := &MembershipHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}