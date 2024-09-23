package domain

type ShopItem struct {
	ID        ID
	ShopID    ID
	ProductID ID
	Quantity  int64
}

type Shop struct {
	ID       ID
	SellerID ID
	Name     string
	Description string
	Requisites  string
	Email    string
	Items    []ShopItem
}
