// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/EmirShimshir/marketplace/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// OrderRepository is an autogenerated mock type for the IOrderRepository type
type OrderRepository struct {
	mock.Mock
}

// CreateOrderCustomer provides a mock function with given fields: ctx, orderCustomer
func (_m *OrderRepository) CreateOrderCustomer(ctx context.Context, orderCustomer domain.OrderCustomer) (domain.OrderCustomer, error) {
	ret := _m.Called(ctx, orderCustomer)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrderCustomer")
	}

	var r0 domain.OrderCustomer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.OrderCustomer) (domain.OrderCustomer, error)); ok {
		return rf(ctx, orderCustomer)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.OrderCustomer) domain.OrderCustomer); ok {
		r0 = rf(ctx, orderCustomer)
	} else {
		r0 = ret.Get(0).(domain.OrderCustomer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.OrderCustomer) error); ok {
		r1 = rf(ctx, orderCustomer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNoNotifiedOrderShops provides a mock function with given fields: ctx
func (_m *OrderRepository) GetNoNotifiedOrderShops(ctx context.Context) ([]domain.OrderShop, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetNoNotifiedOrderShops")
	}

	var r0 []domain.OrderShop
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.OrderShop, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.OrderShop); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.OrderShop)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCustomerByCustomerID provides a mock function with given fields: ctx, customerID
func (_m *OrderRepository) GetOrderCustomerByCustomerID(ctx context.Context, customerID domain.ID) ([]domain.OrderCustomer, error) {
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

// GetOrderCustomerByID provides a mock function with given fields: ctx, orderCustomerID
func (_m *OrderRepository) GetOrderCustomerByID(ctx context.Context, orderCustomerID domain.ID) (domain.OrderCustomer, error) {
	ret := _m.Called(ctx, orderCustomerID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderCustomerByID")
	}

	var r0 domain.OrderCustomer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) (domain.OrderCustomer, error)); ok {
		return rf(ctx, orderCustomerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) domain.OrderCustomer); ok {
		r0 = rf(ctx, orderCustomerID)
	} else {
		r0 = ret.Get(0).(domain.OrderCustomer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, orderCustomerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderShopByID provides a mock function with given fields: ctx, orderShopID
func (_m *OrderRepository) GetOrderShopByID(ctx context.Context, orderShopID domain.ID) (domain.OrderShop, error) {
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
func (_m *OrderRepository) GetOrderShopByShopID(ctx context.Context, shopID domain.ID) ([]domain.OrderShop, error) {
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

// UpdateOrderShop provides a mock function with given fields: ctx, orderShop
func (_m *OrderRepository) UpdateOrderShop(ctx context.Context, orderShop domain.OrderShop) (domain.OrderShop, error) {
	ret := _m.Called(ctx, orderShop)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrderShop")
	}

	var r0 domain.OrderShop
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.OrderShop) (domain.OrderShop, error)); ok {
		return rf(ctx, orderShop)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.OrderShop) domain.OrderShop); ok {
		r0 = rf(ctx, orderShop)
	} else {
		r0 = ret.Get(0).(domain.OrderShop)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.OrderShop) error); ok {
		r1 = rf(ctx, orderShop)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePaymentStatus provides a mock function with given fields: ctx, orderCustomerID
func (_m *OrderRepository) UpdatePaymentStatus(ctx context.Context, orderCustomerID domain.ID) error {
	ret := _m.Called(ctx, orderCustomerID)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePaymentStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) error); ok {
		r0 = rf(ctx, orderCustomerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
