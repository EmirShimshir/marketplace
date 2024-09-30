package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
)

// CartItemBuilder для создания экземпляров CartItem
type CartItemBuilder struct {
	cartItem domain.CartItem
}

// NewCartItemBuilder создает новый CartItemBuilder
func NewCartItemBuilder() *CartItemBuilder {
	return &CartItemBuilder{
		cartItem: domain.CartItem{
			ID:        domain.NewID(),
			CartID:    domain.NewID(),
			ProductID: domain.NewID(),
			Quantity:  1, // Устанавливаем значение по умолчанию
		},
	}
}

// WithID устанавливает ID
func (b *CartItemBuilder) WithID(id domain.ID) *CartItemBuilder {
	b.cartItem.ID = id
	return b
}

// WithCartID устанавливает CartID
func (b *CartItemBuilder) WithCartID(cartID domain.ID) *CartItemBuilder {
	b.cartItem.CartID = cartID
	return b
}

// WithProductID устанавливает ProductID
func (b *CartItemBuilder) WithProductID(productID domain.ID) *CartItemBuilder {
	b.cartItem.ProductID = productID
	return b
}

// WithQuantity устанавливает Quantity
func (b *CartItemBuilder) WithQuantity(quantity int64) *CartItemBuilder {
	b.cartItem.Quantity = quantity
	return b
}

// Build возвращает созданный экземпляр CartItem
func (b *CartItemBuilder) Build() domain.CartItem {
	return b.cartItem
}

// CartBuilder для создания экземпляров Cart
type CartBuilder struct {
	cart domain.Cart
}

// NewCartBuilder создает новый CartBuilder
func NewCartBuilder() *CartBuilder {
	return &CartBuilder{
		cart: domain.Cart{
			ID:    domain.NewID(),
			Price: 0,
			Items: []domain.CartItem{},
		},
	}
}

// WithID устанавливает ID
func (b *CartBuilder) WithID(id domain.ID) *CartBuilder {
	b.cart.ID = id
	return b
}

// WithPrice устанавливает Price
func (b *CartBuilder) WithPrice(price int64) *CartBuilder {
	b.cart.Price = price
	return b
}

// WithItems добавляет элементы в корзину
func (b *CartBuilder) WithItems(items []domain.CartItem) *CartBuilder {
	b.cart.Items = items
	return b
}

// Build возвращает созданный экземпляр Cart
func (b *CartBuilder) Build() domain.Cart {
	return b.cart
}

// CreateCartItemParamBuilder для создания экземпляров CreateCartItemParam
type CreateCartItemParamBuilder struct {
	param port.CreateCartItemParam
}

// NewCreateCartItemParamBuilder создает новый CreateCartItemParamBuilder
func NewCreateCartItemParamBuilder() *CreateCartItemParamBuilder {
	return &CreateCartItemParamBuilder{
		param: port.CreateCartItemParam{
			CartID:    domain.NewID(),
			ProductID: domain.NewID(),
			Quantity:  1, // Устанавливаем значение по умолчанию
		},
	}
}

// WithCartID устанавливает CartID
func (b *CreateCartItemParamBuilder) WithCartID(cartID domain.ID) *CreateCartItemParamBuilder {
	b.param.CartID = cartID
	return b
}

// WithProductID устанавливает ProductID
func (b *CreateCartItemParamBuilder) WithProductID(productID domain.ID) *CreateCartItemParamBuilder {
	b.param.ProductID = productID
	return b
}

// WithQuantity устанавливает Quantity
func (b *CreateCartItemParamBuilder) WithQuantity(quantity int64) *CreateCartItemParamBuilder {
	b.param.Quantity = quantity
	return b
}

// Build возвращает созданный экземпляр CreateCartItemParam
func (b *CreateCartItemParamBuilder) Build() port.CreateCartItemParam {
	return b.param
}

// UpdateCartItemParamBuilder для создания экземпляров UpdateCartItemParam
type UpdateCartItemParamBuilder struct {
	param port.UpdateCartItemParam
}

// NewUpdateCartItemParamBuilder создает новый UpdateCartItemParamBuilder
func NewUpdateCartItemParamBuilder() *UpdateCartItemParamBuilder {
	return &UpdateCartItemParamBuilder{
		param: port.UpdateCartItemParam{
			Quantity: null.IntFrom(1), // Устанавливаем значение по умолчанию
		},
	}
}

// WithQuantity устанавливает Quantity
func (b *UpdateCartItemParamBuilder) WithQuantity(quantity int64) *UpdateCartItemParamBuilder {
	b.param.Quantity = null.IntFrom(quantity)
	return b
}

// Build возвращает созданный экземпляр UpdateCartItemParam
func (b *UpdateCartItemParamBuilder) Build() port.UpdateCartItemParam {
	return b.param
}
