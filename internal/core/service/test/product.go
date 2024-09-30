package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
	"io"
	"strings"
)

// CreateProductParamBuilder helps construct CreateProductParam with default values.
type CreateProductParamBuilder struct {
	param port.CreateProductParam
}

// NewCreateProductParamBuilder returns a new instance of CreateProductParamBuilder.
func NewCreateProductParamBuilder() *CreateProductParamBuilder {
	return &CreateProductParamBuilder{
		param: port.CreateProductParam{
			Name:        "Default Product",
			Description: "Default description",
			Price:       1000,
			Category:    domain.ElectronicCategory,
			PhotoReader: strings.NewReader("fake image data"),
		},
	}
}

// WithName sets the Name field.
func (b *CreateProductParamBuilder) WithName(name string) *CreateProductParamBuilder {
	b.param.Name = name
	return b
}

// WithDescription sets the Description field.
func (b *CreateProductParamBuilder) WithDescription(description string) *CreateProductParamBuilder {
	b.param.Description = description
	return b
}

// WithPrice sets the Price field.
func (b *CreateProductParamBuilder) WithPrice(price int64) *CreateProductParamBuilder {
	b.param.Price = price
	return b
}

// WithCategory sets the Category field.
func (b *CreateProductParamBuilder) WithCategory(category domain.ProductCategory) *CreateProductParamBuilder {
	b.param.Category = category
	return b
}

// WithPhotoReader sets the PhotoReader field.
func (b *CreateProductParamBuilder) WithPhotoReader(reader io.Reader) *CreateProductParamBuilder {
	b.param.PhotoReader = reader
	return b
}

// Build returns the built CreateProductParam.
func (b *CreateProductParamBuilder) Build() port.CreateProductParam {
	return b.param
}

// UpdateProductParamBuilder helps construct UpdateProductParam with default values.
type UpdateProductParamBuilder struct {
	param port.UpdateProductParam
}

// NewUpdateProductParamBuilder returns a new instance of UpdateProductParamBuilder.
func NewUpdateProductParamBuilder() *UpdateProductParamBuilder {
	return &UpdateProductParamBuilder{
		param: port.UpdateProductParam{
			Name:        null.String{},
			Description: null.String{},
			Price:       null.Int{},
			Category:    nil,
			PhotoReader: nil,
		},
	}
}

// WithName sets the Name field.
func (b *UpdateProductParamBuilder) WithName(name string) *UpdateProductParamBuilder {
	b.param.Name = null.StringFrom(name)
	return b
}

// WithDescription sets the Description field.
func (b *UpdateProductParamBuilder) WithDescription(description string) *UpdateProductParamBuilder {
	b.param.Description = null.StringFrom(description)
	return b
}

// WithPrice sets the Price field.
func (b *UpdateProductParamBuilder) WithPrice(price int64) *UpdateProductParamBuilder {
	b.param.Price = null.IntFrom(price)
	return b
}

// WithCategory sets the Category field.
func (b *UpdateProductParamBuilder) WithCategory(category domain.ProductCategory) *UpdateProductParamBuilder {
	b.param.Category = &category
	return b
}

// WithPhotoReader sets the PhotoReader field.
func (b *UpdateProductParamBuilder) WithPhotoReader(reader io.Reader) *UpdateProductParamBuilder {
	b.param.PhotoReader = &reader
	return b
}

// Build returns the built UpdateProductParam.
func (b *UpdateProductParamBuilder) Build() port.UpdateProductParam {
	return b.param
}

// ProductBuilder helps construct Product with default values.
type ProductBuilder struct {
	product domain.Product
}

// NewProductBuilder returns a new instance of ProductBuilder.
func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{
		product: domain.Product{
			ID:          domain.NewID(),
			Name:        "Default Product",
			Description: "Default description",
			Price:       1000,
			Category:    domain.ElectronicCategory,
			PhotoUrl:    "https://example.com/image.jpg",
		},
	}
}

// WithID sets the ID field.
func (b *ProductBuilder) WithID(id domain.ID) *ProductBuilder {
	b.product.ID = id
	return b
}

// WithName sets the Name field.
func (b *ProductBuilder) WithName(name string) *ProductBuilder {
	b.product.Name = name
	return b
}

// WithDescription sets the Description field.
func (b *ProductBuilder) WithDescription(description string) *ProductBuilder {
	b.product.Description = description
	return b
}

// WithPrice sets the Price field.
func (b *ProductBuilder) WithPrice(price int64) *ProductBuilder {
	b.product.Price = price
	return b
}

// WithCategory sets the Category field.
func (b *ProductBuilder) WithCategory(category domain.ProductCategory) *ProductBuilder {
	b.product.Category = category
	return b
}

// WithPhotoUrl sets the PhotoUrl field.
func (b *ProductBuilder) WithPhotoUrl(photoUrl string) *ProductBuilder {
	b.product.PhotoUrl = photoUrl
	return b
}

// Build returns the built Product.
func (b *ProductBuilder) Build() domain.Product {
	return b.product
}
