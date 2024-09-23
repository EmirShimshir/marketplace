package dto

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type OrderShopItemDTO struct {
	ProductID   string
	Name        string
	Description string
	Price       int64
	Category    string
	PhotoUrl    string
	Quantity    int64
}

func NewOrderShopItemDTO(orderShopItem domain.OrderShopItem, product domain.Product) *OrderShopItemDTO {
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

	return &OrderShopItemDTO{
		ProductID:   string(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    category,
		PhotoUrl:    product.PhotoUrl,
		Quantity:    orderShopItem.Quantity,
	}
}

func (si *OrderShopItemDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("ID товара: %s\n", si.ProductID)
	fmt.Printf("Назавние: %s\n", si.Name)
	fmt.Printf("Описание: %s\n", si.Description)
	fmt.Printf("Цена: %d\n", si.Price)
	fmt.Printf("Категория: %s\n", si.Category)
	fmt.Printf("Фото: %s\n", si.PhotoUrl)
	fmt.Printf("Количество: %d\n", si.Quantity)
	fmt.Println("------------------------------------------------")
}

type OrderShopDTO struct {
	ID             string
	ShopID         string
	Status         string
	OrderShopItems []OrderShopItemDTO
}

func NewOrderShopDTO(orderShop domain.OrderShop, orderShopItemDTOs []OrderShopItemDTO) *OrderShopDTO {
	var status string
	switch orderShop.Status {
	case domain.OrderShopStatusStart:
		status = "В обработке"
	case domain.OrderShopStatusReady:
		status = "Принят"
	case domain.OrderShopStatusDone:
		status = "Готов"
	}

	return &OrderShopDTO{
		ID:             string(orderShop.ID),
		ShopID:         string(orderShop.ShopID),
		Status:         status,
		OrderShopItems: orderShopItemDTOs,
	}
}

func (si *OrderShopDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("ID заказа магазина: %s\n", si.ID)
	fmt.Printf("ID магазина: %s\n", si.ShopID)
	fmt.Printf("Статус: %s\n", si.Status)
	fmt.Println("------------------------------------------------")
	fmt.Println("Товары в заказе:")
	for _, oi := range si.OrderShopItems {
		oi.Print()
	}
}

type OrderCustomerDTO struct {
	ID            string
	Address       string
	CreatedAt     string
	TotalPrice    int64
	Payed         string
	OrderShopDTOs []OrderShopDTO
}

func NewOrderCustomerDTO(orderCustomer domain.OrderCustomer, OrderShopDTOs []OrderShopDTO) *OrderCustomerDTO {
	payed := "Ожидает оплаты"
	if orderCustomer.Payed {
		payed = "Оплачен"
	}
	return &OrderCustomerDTO{
		ID:            string(orderCustomer.ID),
		Address:       orderCustomer.Address,
		CreatedAt:     orderCustomer.CreatedAt.String(),
		TotalPrice:    orderCustomer.TotalPrice,
		Payed:         payed,
		OrderShopDTOs: OrderShopDTOs,
	}
}

func (si *OrderCustomerDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("ID заказа покупателя: %s\n", si.ID)
	fmt.Printf("Адрес: %s\n", si.Address)
	fmt.Printf("Дата создания: %s\n", si.CreatedAt)
	fmt.Printf("Итоговая сумма: %d\n", si.TotalPrice)
	fmt.Printf("Статус: %s\n", si.Payed)
	fmt.Println("------------------------------------------------")
	fmt.Println("Заказы в в магазинах:")
	for _, oi := range si.OrderShopDTOs {
		oi.Print()
	}
}
