// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "qbills/models/domain"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// MembershipCardService is an autogenerated mock type for the MembershipCardService type
type MembershipCardService struct {
	mock.Mock
}

// PrintMembershipCard provides a mock function with given fields: ctx, id
func (_m *MembershipCardService) PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Membership
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, int) (*domain.Membership, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, int) *domain.Membership); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Membership)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadBarcodeToFirebase provides a mock function with given fields: ctx, membership
func (_m *MembershipCardService) UploadBarcodeToFirebase(ctx echo.Context, membership domain.Membership) (string, error) {
	ret := _m.Called(ctx, membership)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, domain.Membership) (string, error)); ok {
		return rf(ctx, membership)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, domain.Membership) string); ok {
		r0 = rf(ctx, membership)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, domain.Membership) error); ok {
		r1 = rf(ctx, membership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMembershipCardService creates a new instance of MembershipCardService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMembershipCardService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MembershipCardService {
	mock := &MembershipCardService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}