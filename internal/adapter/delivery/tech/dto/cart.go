package dto

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type CartItemDTO struct {
	CartItemID  string
	ProductID   string
	Name        string
	Description string
	Price       int64
	Category    string
	PhotoUrl    string
	Quantity    int64
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

func (ci *CartItemDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("ID тоавра в корзине: %s\n", ci.CartItemID)
	fmt.Printf("ID товара: %s\n", ci.ProductID)
	fmt.Printf("Назавние: %s\n", ci.Name)
	fmt.Printf("Описание: %s\n", ci.Description)
	fmt.Printf("Цена: %d\n", ci.Price)
	fmt.Printf("Категория: %s\n", ci.Category)
	fmt.Printf("Фото: %s\n", ci.PhotoUrl)
	fmt.Printf("Количество: %d\n", ci.Quantity)
	fmt.Println("------------------------------------------------")
}

type CreateCartItemDTO struct {
	ProductID domain.ID
	Quantity  int64
}

func InputCreateCartItemDTO(ci *CreateCartItemDTO) error {
	err := InputID(&ci.ProductID, "товара")
	if err != nil {
		return err
	}

	fmt.Print("Введите количество: ")
	fmt.Scanln(&ci.Quantity)

	return nil
}

type UpdateCartItemDTO struct {
	CartItemID domain.ID
	Quantity   int64
}

func InputUpdateCartItemDTO(ci *UpdateCartItemDTO) error {
	err := InputID(&ci.CartItemID, "товара в корзине")
	if err != nil {
		return err
	}

	fmt.Print("Введите количество: ")
	fmt.Scanln(&ci.Quantity)

	return nil
}
