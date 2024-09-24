package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
)

type UpdateCartParam struct {
	Price null.Int
}

type CreateCartItemParam struct {
	CartID    domain.ID
	ProductID domain.ID
	Quantity  int64
}

type UpdateCartItemParam struct {
	Quantity null.Int
}

type ICartService interface {
	GetCartByID(ctx context.Context, cartID domain.ID) (domain.Cart, error)
	ClearCart(ctx context.Context, cartID domain.ID) error
	CreateCartItem(ctx context.Context, param CreateCartItemParam) (domain.CartItem, error)
	UpdateCartItem(ctx context.Context, cartItemID domain.ID, param UpdateCartItemParam) (domain.CartItem, error)
	DeleteCartItem(ctx context.Context, CartItemID domain.ID) error
}
