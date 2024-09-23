package tech

import (
	"github.com/EmirShimshir/marketplace/internal/core/port"
)

type Handler struct {
	shop     port.IShopService
	product  port.IProductService
	auth     port.IAuthService
	user     port.IUserService
	cart     port.ICartService
	order    port.IOrderService
	payment  port.IPayment
	withdraw port.IWithdrawService
}

type HandlerParams struct {
	Shop     port.IShopService
	Product  port.IProductService
	Auth     port.IAuthService
	User     port.IUserService
	Cart     port.ICartService
	Order    port.IOrderService
	Payment  port.IPayment
	Withdraw port.IWithdrawService
}

func NewHandler(params HandlerParams) *Handler {
	return &Handler{
		shop:     params.Shop,
		product:  params.Product,
		auth:     params.Auth,
		user:     params.User,
		cart:     params.Cart,
		order:    params.Order,
		payment:  params.Payment,
		withdraw: params.Withdraw,
	}
}
