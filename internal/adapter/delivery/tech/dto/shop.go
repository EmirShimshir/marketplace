package dto

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type ShopItemDTO struct {
	ShopID      string
	ProductID   string
	Name        string
	Description string
	Price       int64
	Category    string
	PhotoUrl    string
	Quantity    int64
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

func (si *ShopItemDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("ID магазина: %s\n", si.ShopID)
	fmt.Printf("ID товара: %s\n", si.ProductID)
	fmt.Printf("Назавние: %s\n", si.Name)
	fmt.Printf("Описание: %s\n", si.Description)
	fmt.Printf("Цена: %d\n", si.Price)
	fmt.Printf("Категория: %s\n", si.Category)
	fmt.Printf("Фото: %s\n", si.PhotoUrl)
	fmt.Printf("Количество: %d\n", si.Quantity)
	fmt.Println("------------------------------------------------")
}

type CreateShopItemDTO struct {
	Name        string
	Description string
	Price       int64
	Category    domain.ProductCategory
	PhotoUrl    string
	Quantity    int64
}

func InputCreateShopItemDTO(d *CreateShopItemDTO) error {
	fmt.Print("Введите название: ")
	fmt.Scanln(&d.Name)

	fmt.Print("Введите описание: ")
	fmt.Scanln(&d.Description)

	fmt.Print("Введите цену: ")
	fmt.Scanln(&d.Price)

	err := InputProductCategory(&d.Category)
	if err != nil {
		return err
	}

	fmt.Print("Введите путь к фото: ")
	fmt.Scanln(&d.PhotoUrl)

	fmt.Print("Введите количество: ")
	fmt.Scanln(&d.Quantity)

	return nil
}

type ShopDTO struct {
	ID          string
	Name        string
	Description string
	Requisites  string
	Email       string
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

func (s *ShopDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("ID магазина: %s\n", s.ID)
	fmt.Printf("Назавние: %s\n", s.Name)
	fmt.Printf("Описание: %s\n", s.Description)
	fmt.Printf("Реквизиты: %s\n", s.Requisites)
	fmt.Printf("Почта: %s\n", s.Email)
	fmt.Println("------------------------------------------------")
}

type CreateShopDTO struct {
	Name        string
	Description string
	Requisites  string
	Email       string
}

func InputCreateShopDTO(d *CreateShopDTO) error {
	fmt.Print("Введите название: ")
	fmt.Scanln(&d.Name)

	fmt.Print("Введите описание: ")
	fmt.Scanln(&d.Description)

	fmt.Print("Введите реквизиты: ")
	fmt.Scanln(&d.Requisites)

	err := InputEmail(&d.Email)
	if err != nil {
		return err
	}
	return nil
}
