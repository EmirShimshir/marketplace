package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"time"
)

type OrderCustomerBuilder struct {
	orderCustomer domain.OrderCustomer
}

func NewOrderCustomerBuilder() *OrderCustomerBuilder {
	return &OrderCustomerBuilder{
		orderCustomer: domain.OrderCustomer{
			ID:         domain.NewID(),
			CustomerID: domain.NewID(),
			Address:    "Default Address",
			CreatedAt:  time.Now(),
			TotalPrice: 1000,
			Payed:      false,
			OrderShops: []domain.OrderShop{},
		},
	}
}

func (b *OrderCustomerBuilder) WithID(id domain.ID) *OrderCustomerBuilder {
	b.orderCustomer.ID = id
	return b
}

func (b *OrderCustomerBuilder) WithCustomerID(customerID domain.ID) *OrderCustomerBuilder {
	b.orderCustomer.CustomerID = customerID
	return b
}

func (b *OrderCustomerBuilder) WithAddress(address string) *OrderCustomerBuilder {
	b.orderCustomer.Address = address
	return b
}

func (b *OrderCustomerBuilder) WithTotalPrice(price int64) *OrderCustomerBuilder {
	b.orderCustomer.TotalPrice = price
	return b
}

func (b *OrderCustomerBuilder) Build() domain.OrderCustomer {
	return b.orderCustomer
}

type OrderShopBuilder struct {
	orderShop domain.OrderShop
}

func NewOrderShopBuilder() *OrderShopBuilder {
	return &OrderShopBuilder{
		orderShop: domain.OrderShop{
			ID:              domain.NewID(),
			ShopID:          domain.NewID(),
			OrderCustomerID: domain.NewID(),
			Status:          domain.OrderShopStatusStart,
			OrderShopItems:  []domain.OrderShopItem{},
			Notified:        false,
		},
	}
}

func (b *OrderShopBuilder) WithID(id domain.ID) *OrderShopBuilder {
	b.orderShop.ID = id
	return b
}

func (b *OrderShopBuilder) WithShopID(shopID domain.ID) *OrderShopBuilder {
	b.orderShop.ShopID = shopID
	return b
}

func (b *OrderShopBuilder) WithStatus(status domain.OrderShopStatus) *OrderShopBuilder {
	b.orderShop.Status = status
	return b
}

func (b *OrderShopBuilder) Build() domain.OrderShop {
	return b.orderShop
}

type OrderShopItemBuilder struct {
	orderShopItem domain.OrderShopItem
}

func NewOrderShopItemBuilder() *OrderShopItemBuilder {
	return &OrderShopItemBuilder{
		orderShopItem: domain.OrderShopItem{
			ID:          domain.NewID(),
			OrderShopID: domain.NewID(),
			ProductID:   domain.NewID(),
			Quantity:    1,
		},
	}
}

func (b *OrderShopItemBuilder) WithID(id domain.ID) *OrderShopItemBuilder {
	b.orderShopItem.ID = id
	return b
}

func (b *OrderShopItemBuilder) WithProductID(productID domain.ID) *OrderShopItemBuilder {
	b.orderShopItem.ProductID = productID
	return b
}

func (b *OrderShopItemBuilder) WithQuantity(quantity int64) *OrderShopItemBuilder {
	b.orderShopItem.Quantity = quantity
	return b
}

func (b *OrderShopItemBuilder) Build() domain.OrderShopItem {
	return b.orderShopItem
}
