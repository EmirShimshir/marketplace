package dto

import "github.com/EmirShimshir/marketplace/internal/core/domain"

type CartDTO struct {
	Products   []CartItemDTO `json:"products"`
	TotalPrice int64         `json:"total_price"`
}

type CartItemDTO struct {
	CartItemID  string `json:"cart_item_id"`
	ProductID   string `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Category    string `json:"category"`
	PhotoUrl    string `json:"photo_url"`
	Quantity    int64  `json:"quantity"`
}

func NewCartItemDTO(cartItem domain.CartItem, product domain.Product) *CartItemDTO {
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

	return &CartItemDTO{
		CartItemID:  string(cartItem.ID),
		ProductID:   string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    category,
		PhotoUrl:    product.PhotoUrl,
		Quantity:    cartItem.Quantity,
	}
}

type CreateCartItemDTO struct {
	ProductID domain.ID `json:"product_id" binding:"required"`
	Quantity  int64     `json:"quantity" binding:"required"`
}

type UpdateCartItemDTO struct {
	CartItemID domain.ID `json:"cart_product_id" binding:"required"`
	Quantity   int64     `json:"quantity" binding:"required"`
}
