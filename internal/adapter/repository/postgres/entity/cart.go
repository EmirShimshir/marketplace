package entity

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/google/uuid"
)

type PgCart struct {
	ID    uuid.UUID `db:"id"`
	Price int64     `db:"price"`
}

func (c *PgCart) ToDomain() domain.Cart {
	return domain.Cart{
		ID:    domain.ID(c.ID.String()),
		Price: c.Price,
	}
}

func NewPgCart(cart domain.Cart) PgCart {
	id, _ := uuid.Parse(cart.ID.String())
	return PgCart{
		ID:    id,
		Price: cart.Price,
	}
}

type PgCartItem struct {
	ID        uuid.UUID `db:"id"`
	CartID    uuid.UUID `db:"cart_id"`
	ProductID uuid.UUID `db:"product_id"`
	Quantity  int64     `db:"quantity"`
}

func (ci *PgCartItem) ToDomain() domain.CartItem {
	return domain.CartItem{
		ID:        domain.ID(ci.ID.String()),
		CartID:    domain.ID(ci.CartID.String()),
		ProductID: domain.ID(ci.ProductID.String()),
		Quantity:  ci.Quantity,
	}
}

func NewPgCartItem(cartItem domain.CartItem) PgCartItem {
	id, _ := uuid.Parse(cartItem.ID.String())
	cartID, _ := uuid.Parse(cartItem.CartID.String())
	productID, _ := uuid.Parse(cartItem.ProductID.String())
	return PgCartItem{
		ID:        id,
		CartID:    cartID,
		ProductID: productID,
		Quantity:  cartItem.Quantity,
	}
}
