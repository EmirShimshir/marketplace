package domain

import "time"

type OrderShopStatus int

const (
	OrderShopStatusStart OrderShopStatus = iota
	OrderShopStatusReady
	OrderShopStatusDone
)

type OrderShopItem struct {
	ID          ID
	OrderShopID ID
	ProductID   ID
	Quantity    int64
}

type OrderShop struct {
	ID              ID
	ShopID          ID
	OrderCustomerID ID
	Status          OrderShopStatus
	OrderShopItems  []OrderShopItem
	Notified        bool
}

type OrderCustomer struct {
	ID         ID
	CustomerID ID
	Address    string
	CreatedAt  time.Time
	TotalPrice int64
	Payed      bool
	OrderShops []OrderShop
}
