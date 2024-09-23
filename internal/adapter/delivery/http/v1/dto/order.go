package dto

import "github.com/EmirShimshir/marketplace/internal/core/domain"

type CreateOrderCustomerDTO struct {
	Address string `json:"address" binding:"required"`
}

type OrderShopItemDTO struct {
	ProductID   string `json:"product_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
	PhotoUrl    string `json:"photo_url" binding:"required"`
	Quantity    int64  `json:"quantity" binding:"required"`
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

type OrderShopDTO struct {
	ID             string             `json:"id" binding:"required"`
	ShopID         string             `json:"shop_id" binding:"required"`
	Status         string             `json:"status" binding:"required"`
	OrderShopItems []OrderShopItemDTO `json:"order_shop_items" binding:"required"`
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

type OrderCustomerDTO struct {
	ID            string         `json:"id" binding:"required"`
	Address       string         `json:"address" binding:"required"`
	CreatedAt     string         `json:"created_at" binding:"required"`
	TotalPrice    int64          `json:"total_price" binding:"required"`
	Payed         string         `json:"payed" binding:"required"`
	OrderShopDTOs []OrderShopDTO `order_shops:"id" binding:"required"`
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
