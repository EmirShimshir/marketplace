package domain

type CartItem struct {
	ID        ID
	CartID    ID
	ProductID ID
	Quantity  int64
}

type Cart struct {
	ID    ID
	Price int64
	Items []CartItem
}
