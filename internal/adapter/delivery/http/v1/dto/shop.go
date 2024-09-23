package dto

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"mime/multipart"
)

type ShopDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Requisites  string `json:"requisites"`
	Email       string `json:"email"`
}

func NewShopDTO(shop domain.Shop) *ShopDTO {
	return &ShopDTO{
		ID:          string(shop.ID),
		Name:        shop.Name,
		Description: shop.Description,
		Requisites:  shop.Requisites,
		Email:       shop.Email,
	}
}

type CreateShopDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Requisites  string `json:"requisites" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
}

type CreateShopItemDTO struct {
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description" binding:"required"`
	Price       int64                  `json:"price" binding:"required"`
	Category    domain.ProductCategory `json:"category" binding:"omitempty"`
	Quantity    int64                  `json:"quantity" binding:"required"`
}

type CreateShopItemFormsDTO struct {
	File *multipart.FileHeader `form:"file"`
	Json CreateShopItemDTO     `form:"json"`
}
