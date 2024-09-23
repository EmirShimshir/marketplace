package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type CreateOrderCustomerParam struct {
	CustomerID domain.ID
	Address    string
}

type CreateOrderShopParam struct {
	ShopID          domain.ID
	OrderCustomerID domain.ID
	Status          domain.OrderShopStatus
	CartItems       []domain.CartItem
}

type UpdateOrderShopParam struct {
	Status *domain.OrderShopStatus
}

type CreateOrderShopItemParam struct {
	OrderShopID domain.ID
	ProductID   domain.ID
	Quantity    int64
}

type IOrderService interface {
	GetOrderCustomerByCustomerID(ctx context.Context, customerID domain.ID) ([]domain.OrderCustomer, error)
	GetOrderCustomerByID(ctx context.Context, ID domain.ID) (domain.OrderCustomer, error)
	CreateOrderCustomer(ctx context.Context, param CreateOrderCustomerParam) (domain.OrderCustomer, error)
	GetOrderShopByID(ctx context.Context, orderShopID domain.ID) (domain.OrderShop, error)
	GetOrderShopByShopID(ctx context.Context, shopID domain.ID) ([]domain.OrderShop, error)
	UpdateOrderShop(ctx context.Context, orderShopID domain.ID, param UpdateOrderShopParam) (domain.OrderShop, error)
}
