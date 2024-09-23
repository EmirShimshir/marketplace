// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/EmirShimshir/marketplace/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// CartRepository is an autogenerated mock type for the ICartRepository type
type CartRepository struct {
	mock.Mock
}

// ClearCart provides a mock function with given fields: ctx, cartID
func (_m *CartRepository) ClearCart(ctx context.Context, cartID domain.ID) error {
	ret := _m.Called(ctx, cartID)

	if len(ret) == 0 {
		panic("no return value specified for ClearCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) error); ok {
		r0 = rf(ctx, cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCartItem provides a mock function with given fields: ctx, cartItem
func (_m *CartRepository) CreateCartItem(ctx context.Context, cartItem domain.CartItem) (domain.CartItem, error) {
	ret := _m.Called(ctx, cartItem)

	if len(ret) == 0 {
		panic("no return value specified for CreateCartItem")
	}

	var r0 domain.CartItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CartItem) (domain.CartItem, error)); ok {
		return rf(ctx, cartItem)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CartItem) domain.CartItem); ok {
		r0 = rf(ctx, cartItem)
	} else {
		r0 = ret.Get(0).(domain.CartItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CartItem) error); ok {
		r1 = rf(ctx, cartItem)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCartItem provides a mock function with given fields: ctx, cartItemID
func (_m *CartRepository) DeleteCartItem(ctx context.Context, cartItemID domain.ID) error {
	ret := _m.Called(ctx, cartItemID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCartItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) error); ok {
		r0 = rf(ctx, cartItemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCartByID provides a mock function with given fields: ctx, cartID
func (_m *CartRepository) GetCartByID(ctx context.Context, cartID domain.ID) (domain.Cart, error) {
	ret := _m.Called(ctx, cartID)

	if len(ret) == 0 {
		panic("no return value specified for GetCartByID")
	}

	var r0 domain.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) (domain.Cart, error)); ok {
		return rf(ctx, cartID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) domain.Cart); ok {
		r0 = rf(ctx, cartID)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCartItemByID provides a mock function with given fields: ctx, cartItemID
func (_m *CartRepository) GetCartItemByID(ctx context.Context, cartItemID domain.ID) (domain.CartItem, error) {
	ret := _m.Called(ctx, cartItemID)

	if len(ret) == 0 {
		panic("no return value specified for GetCartItemByID")
	}

	var r0 domain.CartItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) (domain.CartItem, error)); ok {
		return rf(ctx, cartItemID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) domain.CartItem); ok {
		r0 = rf(ctx, cartItemID)
	} else {
		r0 = ret.Get(0).(domain.CartItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, cartItemID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCart provides a mock function with given fields: ctx, cart
func (_m *CartRepository) UpdateCart(ctx context.Context, cart domain.Cart) (domain.Cart, error) {
	ret := _m.Called(ctx, cart)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCart")
	}

	var r0 domain.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Cart) (domain.Cart, error)); ok {
		return rf(ctx, cart)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Cart) domain.Cart); ok {
		r0 = rf(ctx, cart)
	} else {
		r0 = ret.Get(0).(domain.Cart)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Cart) error); ok {
		r1 = rf(ctx, cart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCartItem provides a mock function with given fields: ctx, cartItem
func (_m *CartRepository) UpdateCartItem(ctx context.Context, cartItem domain.CartItem) (domain.CartItem, error) {
	ret := _m.Called(ctx, cartItem)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCartItem")
	}

	var r0 domain.CartItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CartItem) (domain.CartItem, error)); ok {
		return rf(ctx, cartItem)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CartItem) domain.CartItem); ok {
		r0 = rf(ctx, cartItem)
	} else {
		r0 = ret.Get(0).(domain.CartItem)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CartItem) error); ok {
		r1 = rf(ctx, cartItem)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCartRepository creates a new instance of CartRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCartRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CartRepository {
	mock := &CartRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
