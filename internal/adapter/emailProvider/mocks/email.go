// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	port "github.com/EmirShimshir/marketplace/internal/core/port"
	mock "github.com/stretchr/testify/mock"
)

// EmailProvider is an autogenerated mock type for the IEmailProvider type
type EmailProvider struct {
	mock.Mock
}

// SendEmail provides a mock function with given fields: ctx, param
func (_m *EmailProvider) SendEmail(ctx context.Context, param port.CartEmailProviderParam) error {
	ret := _m.Called(ctx, param)

	if len(ret) == 0 {
		panic("no return value specified for SendEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, port.CartEmailProviderParam) error); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEmailProvider creates a new instance of EmailProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmailProvider {
	mock := &EmailProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
