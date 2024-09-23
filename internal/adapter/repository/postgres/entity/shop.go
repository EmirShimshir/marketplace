package entity

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/google/uuid"
)

type PgShop struct {
	ID          uuid.UUID `db:"id"`
	SellerID    uuid.UUID `db:"seller_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Requisites  string    `db:"requisites"`
	Email       string    `db:"email"`
}

func (s *PgShop) ToDomain() domain.Shop {
	return domain.Shop{
		ID:          domain.ID(s.ID.String()),
		SellerID:    domain.ID(s.SellerID.String()),
		Name:        s.Name,
		Description: s.Description,
		Requisites:  s.Requisites,
		Email:       s.Email,
	}
}

func NewPgShop(shop domain.Shop) PgShop {
	id, _ := uuid.Parse(shop.ID.String())
	sellerID, _ := uuid.Parse(shop.SellerID.String())
	return PgShop{
		ID:          id,
		SellerID:    sellerID,
		Name:        shop.Name,
		Description: shop.Description,
		Requisites:  shop.Requisites,
		Email:       shop.Email,
	}
}

type PgShopItem struct {
	ID        uuid.UUID `db:"id"`
	ShopID    uuid.UUID `db:"shop_id"`
	ProductID uuid.UUID `db:"product_id"`
	Quantity  int64     `db:"quantity"`
}

func (si *PgShopItem) ToDomain() domain.ShopItem {
	return domain.ShopItem{
		ID:        domain.ID(si.ID.String()),
		ShopID:    domain.ID(si.ShopID.String()),
		ProductID: domain.ID(si.ProductID.String()),
		Quantity:  si.Quantity,
	}
}

func NewPgShopItem(shopItem domain.ShopItem) PgShopItem {
	id, _ := uuid.Parse(shopItem.ID.String())
	shopID, _ := uuid.Parse(shopItem.ShopID.String())
	productID, _ := uuid.Parse(shopItem.ProductID.String())
	return PgShopItem{
		ID:        id,
		ShopID:    shopID,
		ProductID: productID,
		Quantity:  shopItem.Quantity,
	}
}
