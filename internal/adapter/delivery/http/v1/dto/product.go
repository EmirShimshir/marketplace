package dto

import "github.com/EmirShimshir/marketplace/internal/core/domain"

type ShopItemDTO struct {
	ShopID      string `json:"shop_id"`
	ProductID   string `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Category    string `json:"category"`
	PhotoUrl    string `json:"photo_url"`
	Quantity    int64  `json:"quantity"`
}

func NewShopItemDTO(shopItem domain.ShopItem, product domain.Product) *ShopItemDTO {
	var category string
	switch product.Category {
	case domain.ElectronicCategory:
		category = "Электроника"
	case domain.FashionCategory:
		category = "Мода"
	case domain.HomeCategory:
		category = "Дом"
	case domain.HealthCategory:
		category = "Здоровье"
	case domain.SportCategory:
		category = "Спорт"
	case domain.BooksCategory:
		category = "Книги"
	}

	return &ShopItemDTO{
		ShopID:      string(shopItem.ShopID),
		ProductID:   string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    category,
		PhotoUrl:    product.PhotoUrl,
		Quantity:    shopItem.Quantity,
	}
}

type ProductDTO struct {
	ProductID   string `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Category    string `json:"category"`
	PhotoUrl    string `json:"photo_url"`
}

func NewProductDTO(product domain.Product) *ProductDTO {
	var category string
	switch product.Category {
	case domain.ElectronicCategory:
		category = "Электроника"
	case domain.FashionCategory:
		category = "Мода"
	case domain.HomeCategory:
		category = "Дом"
	case domain.HealthCategory:
		category = "Здоровье"
	case domain.SportCategory:
		category = "Спорт"
	case domain.BooksCategory:
		category = "Книги"
	}

	return &ProductDTO{
		ProductID:   string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    category,
		PhotoUrl:    product.PhotoUrl,
	}
}

type ShopItemsByShopIdDTO struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Products    []ShopItemDTO `json:"products"`
}
