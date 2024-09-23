package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type IUserRepository interface {
	Get(Ctx context.Context, limit, offset int64) ([]domain.User, error)
	GetByID(ctx context.Context, userID domain.ID) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	CreateWithoutCart(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, userID domain.ID) error
}

type IProductRepository interface {
	Get(ctx context.Context, limit, offset int64) ([]domain.Product, error)
	GetByID(ctx context.Context, productID domain.ID) (domain.Product, error)
	GetByName(ctx context.Context, name string) ([]domain.Product, error)
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, productID domain.ID) error
}

type ICartRepository interface {
	GetCartByID(ctx context.Context, cartID domain.ID) (domain.Cart, error)
	UpdateCart(ctx context.Context, cart domain.Cart) (domain.Cart, error)
	ClearCart(ctx context.Context, cartID domain.ID) error

	GetCartItemByID(ctx context.Context, cartItemID domain.ID) (domain.CartItem, error)
	CreateCartItem(ctx context.Context, cartItem domain.CartItem) (domain.CartItem, error)
	UpdateCartItem(ctx context.Context, cartItem domain.CartItem) (domain.CartItem, error)
	DeleteCartItem(ctx context.Context, cartItemID domain.ID) error
}

type IShopRepository interface {
	GetShops(ctx context.Context, limit, offset int64) ([]domain.Shop, error)
	GetShopByID(ctx context.Context, shopID domain.ID) (domain.Shop, error)
	GetShopBySellerID(ctx context.Context, sellerID domain.ID) ([]domain.Shop, error)
	CreateShop(ctx context.Context, shop domain.Shop) (domain.Shop, error)
	UpdateShop(ctx context.Context, shop domain.Shop) (domain.Shop, error)
	DeleteShop(ctx context.Context, shopID domain.ID) error

	GetShopItems(ctx context.Context, limit, offset int64) ([]domain.ShopItem, error)
	GetShopItemByID(ctx context.Context, shopItemID domain.ID) (domain.ShopItem, error)
	GetShopItemByProductID(ctx context.Context, productID domain.ID) (domain.ShopItem, error)
	CreateShopItem(ctx context.Context, shopItem domain.ShopItem, product domain.Product) (domain.ShopItem, error)
	UpdateShopItem(ctx context.Context, shopItem domain.ShopItem) (domain.ShopItem, error)
	DeleteShopItem(ctx context.Context, shopItemID domain.ID) error
}

type IOrderRepository interface {
	GetOrderCustomerByCustomerID(ctx context.Context, customerID domain.ID) ([]domain.OrderCustomer, error)
	GetOrderCustomerByID(ctx context.Context, orderCustomerID domain.ID) (domain.OrderCustomer, error)
	GetOrderShopByID(ctx context.Context, orderShopID domain.ID) (domain.OrderShop, error)
	GetNoNotifiedOrderShops(ctx context.Context) ([]domain.OrderShop, error)
	CreateOrderCustomer(ctx context.Context, orderCustomer domain.OrderCustomer) (domain.OrderCustomer, error)
	GetOrderShopByShopID(ctx context.Context, shopID domain.ID) ([]domain.OrderShop, error)
	UpdateOrderShop(ctx context.Context, orderShop domain.OrderShop) (domain.OrderShop, error)
	UpdatePaymentStatus(ctx context.Context, orderCustomerID domain.ID) error
}

type IWithdrawRepository interface {
	Get(ctx context.Context, limit, offset int64) ([]domain.Withdraw, error)
	GetByID(ctx context.Context, WithdrawID domain.ID) (domain.Withdraw, error)
	GetByShopID(ctx context.Context, shopID domain.ID) ([]domain.Withdraw, error)
	Create(ctx context.Context, withdraw domain.Withdraw) (domain.Withdraw, error)
	Update(ctx context.Context, withdraw domain.Withdraw) (domain.Withdraw, error)
	Delete(ctx context.Context, withdrawID domain.ID) error
}
