package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
)

// CreateShopParamBuilder helps construct CreateShopParam with default values.
type CreateShopParamBuilder struct {
	param port.CreateShopParam
}

// NewCreateShopParamBuilder returns a new instance of CreateShopParamBuilder.
func NewCreateShopParamBuilder() *CreateShopParamBuilder {
	return &CreateShopParamBuilder{
		param: port.CreateShopParam{
			Name:        "Default Shop",
			Description: "Default description",
			Requisites:  "Default requisites",
			Email:       "shop@example.com",
		},
	}
}

// WithName sets the Name field.
func (b *CreateShopParamBuilder) WithName(name string) *CreateShopParamBuilder {
	b.param.Name = name
	return b
}

// WithDescription sets the Description field.
func (b *CreateShopParamBuilder) WithDescription(description string) *CreateShopParamBuilder {
	b.param.Description = description
	return b
}

// WithRequisites sets the Requisites field.
func (b *CreateShopParamBuilder) WithRequisites(requisites string) *CreateShopParamBuilder {
	b.param.Requisites = requisites
	return b
}

// WithEmail sets the Email field.
func (b *CreateShopParamBuilder) WithEmail(email string) *CreateShopParamBuilder {
	b.param.Email = email
	return b
}

// Build returns the built CreateShopParam.
func (b *CreateShopParamBuilder) Build() port.CreateShopParam {
	return b.param
}

// UpdateShopParamBuilder helps construct UpdateShopParam with default values.
type UpdateShopParamBuilder struct {
	param port.UpdateShopParam
}

// NewUpdateShopParamBuilder returns a new instance of UpdateShopParamBuilder.
func NewUpdateShopParamBuilder() *UpdateShopParamBuilder {
	return &UpdateShopParamBuilder{
		param: port.UpdateShopParam{
			Name:        null.String{},
			Description: null.String{},
			Requisites:  null.String{},
			Email:       null.String{},
		},
	}
}

// WithName sets the Name field.
func (b *UpdateShopParamBuilder) WithName(name string) *UpdateShopParamBuilder {
	b.param.Name = null.StringFrom(name)
	return b
}

// WithDescription sets the Description field.
func (b *UpdateShopParamBuilder) WithDescription(description string) *UpdateShopParamBuilder {
	b.param.Description = null.StringFrom(description)
	return b
}

// WithRequisites sets the Requisites field.
func (b *UpdateShopParamBuilder) WithRequisites(requisites string) *UpdateShopParamBuilder {
	b.param.Requisites = null.StringFrom(requisites)
	return b
}

// WithEmail sets the Email field.
func (b *UpdateShopParamBuilder) WithEmail(email string) *UpdateShopParamBuilder {
	b.param.Email = null.StringFrom(email)
	return b
}

// Build returns the built UpdateShopParam.
func (b *UpdateShopParamBuilder) Build() port.UpdateShopParam {
	return b.param
}

// CreateShopItemParamBuilder helps construct CreateShopItemParam with default values.
type CreateShopItemParamBuilder struct {
	param port.CreateShopItemParam
}

// NewCreateShopItemParamBuilder returns a new instance of CreateShopItemParamBuilder.
func NewCreateShopItemParamBuilder(shopID domain.ID) *CreateShopItemParamBuilder {
	return &CreateShopItemParamBuilder{
		param: port.CreateShopItemParam{
			ShopID:       shopID,
			ProductParam: NewCreateProductParamBuilder().Build(),
			Quantity:     10,
		},
	}
}

// WithQuantity sets the Quantity field.
func (b *CreateShopItemParamBuilder) WithQuantity(quantity int64) *CreateShopItemParamBuilder {
	b.param.Quantity = quantity
	return b
}

// Build returns the built CreateShopItemParam.
func (b *CreateShopItemParamBuilder) Build() port.CreateShopItemParam {
	return b.param
}

// UpdateShopItemParamBuilder helps construct UpdateShopItemParam with default values.
type UpdateShopItemParamBuilder struct {
	param port.UpdateShopItemParam
}

// NewUpdateShopItemParamBuilder returns a new instance of UpdateShopItemParamBuilder.
func NewUpdateShopItemParamBuilder() *UpdateShopItemParamBuilder {
	return &UpdateShopItemParamBuilder{
		param: port.UpdateShopItemParam{
			Quantity: null.Int{},
		},
	}
}

// WithQuantity sets the Quantity field.
func (b *UpdateShopItemParamBuilder) WithQuantity(quantity int64) *UpdateShopItemParamBuilder {
	b.param.Quantity = null.IntFrom(quantity)
	return b
}

// Build returns the built UpdateShopItemParam.
func (b *UpdateShopItemParamBuilder) Build() port.UpdateShopItemParam {
	return b.param
}

// ShopBuilder - билдер для создания тестовых объектов Shop
type ShopBuilder struct {
	shop domain.Shop
}

func NewShopBuilder() *ShopBuilder {
	return &ShopBuilder{
		shop: domain.Shop{
			ID:          domain.NewID(),
			SellerID:    domain.NewID(),
			Name:        "Test Shop",
			Description: "Test Description",
			Requisites:  "Test Requisites",
			Email:       "test@shop.com",
			Items:       []domain.ShopItem{},
		},
	}
}

func (b *ShopBuilder) WithID(id domain.ID) *ShopBuilder {
	b.shop.ID = id
	return b
}

func (b *ShopBuilder) WithSellerID(sellerID domain.ID) *ShopBuilder {
	b.shop.SellerID = sellerID
	return b
}

func (b *ShopBuilder) WithName(name string) *ShopBuilder {
	b.shop.Name = name
	return b
}

func (b *ShopBuilder) WithDescription(description string) *ShopBuilder {
	b.shop.Description = description
	return b
}

func (b *ShopBuilder) WithRequisites(requisites string) *ShopBuilder {
	b.shop.Requisites = requisites
	return b
}

func (b *ShopBuilder) WithEmail(email string) *ShopBuilder {
	b.shop.Email = email
	return b
}

func (b *ShopBuilder) WithItems(items []domain.ShopItem) *ShopBuilder {
	b.shop.Items = items
	return b
}

func (b *ShopBuilder) Build() domain.Shop {
	return b.shop
}

// ShopItemBuilder - билдер для создания тестовых объектов ShopItem
type ShopItemBuilder struct {
	shopItem domain.ShopItem
}

func NewShopItemBuilder() *ShopItemBuilder {
	return &ShopItemBuilder{
		shopItem: domain.ShopItem{
			ID:        domain.NewID(),
			ShopID:    domain.NewID(),
			ProductID: domain.NewID(),
			Quantity:  10,
		},
	}
}

func (b *ShopItemBuilder) WithID(id domain.ID) *ShopItemBuilder {
	b.shopItem.ID = id
	return b
}

func (b *ShopItemBuilder) WithShopID(shopID domain.ID) *ShopItemBuilder {
	b.shopItem.ShopID = shopID
	return b
}

func (b *ShopItemBuilder) WithProductID(productID domain.ID) *ShopItemBuilder {
	b.shopItem.ProductID = productID
	return b
}

func (b *ShopItemBuilder) WithQuantity(quantity int64) *ShopItemBuilder {
	b.shopItem.Quantity = quantity
	return b
}

func (b *ShopItemBuilder) Build() domain.ShopItem {
	return b.shopItem
}
