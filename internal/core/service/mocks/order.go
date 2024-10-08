// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/EmirShimshir/marketplace/internal/core/domain"
	mock "github.com/stretchr/testify/mock"

	port "github.com/EmirShimshir/marketplace/internal/core/port"
)

// OrderService is an autogenerated mock type for the IOrderService type
type OrderService struct {
	mock.Mock
}

// CreateOrderCustomer provides a mock function with given fields: ctx, param
func (_m *OrderService) CreateOrderCustomer(ctx context.Context, param port.CreateOrderCustomerParam) (domain.OrderCustomer, error) {
	ret := _m.Called(ctx, param)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrderCustomer")
	}

	var r0 domain.OrderCustomer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, port.CreateOrderCustomerParam) (domain.OrderCustomer, error)); ok {
		return rf(ctx, param)
	}
	if rf, ok := ret.Get(0).(func(context.Context, port.CreateOrderCustomerParam) domain.OrderCustomer); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(domain.OrderCustomer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, port.CreateOrderCustomerParam) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCustomerByCustomerID provides a mock function with given fields: ctx, customerID
func (_m *OrderService) GetOrderCustomerByCustomerID(ctx context.Context, customerID domain.ID) ([]domain.OrderCustomer, error) {
	ret := _m.Called(ctx, customerID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderCustomerByCustomerID")
	}

	var r0 []domain.OrderCustomer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) ([]domain.OrderCustomer, error)); ok {
		return rf(ctx, customerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) []domain.OrderCustomer); ok {
		r0 = rf(ctx, customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.OrderCustomer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderShopByID provides a mock function with given fields: ctx, orderShopID
func (_m *OrderService) GetOrderShopByID(ctx context.Context, orderShopID domain.ID) (domain.OrderShop, error) {
	ret := _m.Called(ctx, orderShopID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderShopByID")
	}

	var r0 domain.OrderShop
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) (domain.OrderShop, error)); ok {
		return rf(ctx, orderShopID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) domain.OrderShop); ok {
		r0 = rf(ctx, orderShopID)
	} else {
		r0 = ret.Get(0).(domain.OrderShop)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, orderShopID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderShopByShopID provides a mock function with given fields: ctx, shopID
func (_m *OrderService) GetOrderShopByShopID(ctx context.Context, shopID domain.ID) ([]domain.OrderShop, error) {
	ret := _m.Called(ctx, shopID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderShopByShopID")
	}

	var r0 []domain.OrderShop
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) ([]domain.OrderShop, error)); ok {
		return rf(ctx, shopID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) []domain.OrderShop); ok {
		r0 = rf(ctx, shopID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.OrderShop)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, shopID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrderShop provides a mock function with given fields: ctx, orderShopID, param
func (_m *OrderService) UpdateOrderShop(ctx context.Context, orderShopID domain.ID, param port.UpdateOrderShopParam) (domain.OrderShop, error) {
	ret := _m.Called(ctx, orderShopID, param)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrderShop")
	}

	var r0 domain.OrderShop
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID, port.UpdateOrderShopParam) (domain.OrderShop, error)); ok {
		return rf(ctx, orderShopID, param)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID, port.UpdateOrderShopParam) domain.OrderShop); ok {
		r0 = rf(ctx, orderShopID, param)
	} else {
		r0 = ret.Get(0).(domain.OrderShop)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID, port.UpdateOrderShopParam) error); ok {
		r1 = rf(ctx, orderShopID, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderService creates a new instance of OrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderService(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderService {
	mock := &OrderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
