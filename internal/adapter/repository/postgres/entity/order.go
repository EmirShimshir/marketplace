package entity

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/google/uuid"
	"time"
)

type PgOrderCustomer struct {
	ID         uuid.UUID `db:"id"`
	CustomerID uuid.UUID `db:"customer_id"`
	Address    string    `db:"address"`
	CreatedAt  time.Time `db:"created_at"`
	TotalPrice int64     `db:"total_price"`
	Payed      bool      `db:"payed"`
}

func (oc *PgOrderCustomer) ToDomain() domain.OrderCustomer {
	return domain.OrderCustomer{
		ID:         domain.ID(oc.ID.String()),
		CustomerID: domain.ID(oc.CustomerID.String()),
		Address:    oc.Address,
		CreatedAt:  oc.CreatedAt,
		TotalPrice: oc.TotalPrice,
		Payed:      oc.Payed,
	}
}

func NewPgOrderCustomer(orderCustomer domain.OrderCustomer) PgOrderCustomer {
	id, _ := uuid.Parse(orderCustomer.ID.String())
	customerID, _ := uuid.Parse(orderCustomer.CustomerID.String())
	return PgOrderCustomer{
		ID:         id,
		CustomerID: customerID,
		Address:    orderCustomer.Address,
		CreatedAt:  orderCustomer.CreatedAt,
		TotalPrice: orderCustomer.TotalPrice,
		Payed:      orderCustomer.Payed,
	}
}

const (
	PgOrderShopStart = "Start"
	PgOrderShopReady = "Ready"
	PgOrderShopDone  = "Done"
)

type PgOrderShop struct {
	ID              uuid.UUID `db:"id"`
	ShopID          uuid.UUID `db:"shop_id"`
	OrderCustomerID uuid.UUID `db:"order_customer_id"`
	Status          string    `db:"status"`
	Notified        bool      `db:"notified"`
}

func (os *PgOrderShop) ToDomain() domain.OrderShop {
	var orderShopStatus domain.OrderShopStatus
	switch os.Status {
	case PgOrderShopStart:
		orderShopStatus = domain.OrderShopStatusStart
	case PgOrderShopReady:
		orderShopStatus = domain.OrderShopStatusReady
	case PgOrderShopDone:
		orderShopStatus = domain.OrderShopStatusDone
	}

	return domain.OrderShop{
		ID:              domain.ID(os.ID.String()),
		ShopID:          domain.ID(os.ShopID.String()),
		OrderCustomerID: domain.ID(os.OrderCustomerID.String()),
		Status:          orderShopStatus,
		Notified:        os.Notified,
	}
}

func NewPgOrderShop(orderShop domain.OrderShop) PgOrderShop {
	id, _ := uuid.Parse(orderShop.ID.String())
	shopID, _ := uuid.Parse(orderShop.ShopID.String())
	orderCustomerID, _ := uuid.Parse(orderShop.OrderCustomerID.String())
	var orderShopStatus string
	switch orderShop.Status {
	case domain.OrderShopStatusStart:
		orderShopStatus = PgOrderShopStart
	case domain.OrderShopStatusReady:
		orderShopStatus = PgOrderShopReady
	case domain.OrderShopStatusDone:
		orderShopStatus = PgOrderShopDone
	}

	return PgOrderShop{
		ID:              id,
		ShopID:          shopID,
		OrderCustomerID: orderCustomerID,
		Status:          orderShopStatus,
		Notified:        orderShop.Notified,
	}
}

type PgOrderShopItem struct {
	ID          uuid.UUID `db:"id"`
	OrderShopID uuid.UUID `db:"order_shop_id"`
	ProductID   uuid.UUID `db:"product_id"`
	Quantity    int64     `db:"quantity"`
}

func (osi *PgOrderShopItem) ToDomain() domain.OrderShopItem {
	return domain.OrderShopItem{
		ID:          domain.ID(osi.ID.String()),
		OrderShopID: domain.ID(osi.OrderShopID.String()),
		ProductID:   domain.ID(osi.ProductID.String()),
		Quantity:    osi.Quantity,
	}
}

func NewPgOrderShopItem(orderShopItem domain.OrderShopItem) PgOrderShopItem {
	id, _ := uuid.Parse(orderShopItem.ID.String())
	orderShopID, _ := uuid.Parse(orderShopItem.OrderShopID.String())
	productID, _ := uuid.Parse(orderShopItem.ProductID.String())
	return PgOrderShopItem{
		ID:          id,
		OrderShopID: orderShopID,
		ProductID:   productID,
		Quantity:    orderShopItem.Quantity,
	}
}
