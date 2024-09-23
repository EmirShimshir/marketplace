package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
	log "github.com/sirupsen/logrus"
)

type CartService struct {
	cartRepo    port.ICartRepository
	shopRepo    port.IShopRepository
	productRepo port.IProductRepository
}

func NewCartService(cartRepo port.ICartRepository, shopRepo port.IShopRepository, productRepo port.IProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		shopRepo:    shopRepo,
		productRepo: productRepo,
	}
}

func (c *CartService) GetCartByID(ctx context.Context, cartID domain.ID) (domain.Cart, error) {
	cart, err := c.cartRepo.GetCartByID(ctx, cartID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetCartByID",
		}).Error(err.Error())
		return domain.Cart{}, err
	}

	var totalPrice int64 = 0
	for _, cartItem := range cart.Items {
		shopItem, err := c.shopRepo.GetShopItemByProductID(ctx, cartItem.ProductID)
		if err != nil {
			log.WithFields(log.Fields{
				"from": "GetCartByID",
			}).Error(err.Error())
			return domain.Cart{}, err
		}
		if cartItem.Quantity > shopItem.Quantity {
			cartItem, err = c.UpdateCartItem(ctx, cartItem.ID, port.UpdateCartItemParam{
				Quantity: null.IntFrom(shopItem.Quantity),
			})
			if err != nil {
				log.WithFields(log.Fields{
					"from": "GetCartByID",
				}).Error(err.Error())
				return domain.Cart{}, err
			}
		}
		product, err := c.productRepo.GetByID(ctx, cartItem.ProductID)
		if err != nil {
			log.WithFields(log.Fields{
				"from": "GetCartByID",
			}).Error(err.Error())
			return domain.Cart{}, err
		}
		totalPrice += product.Price * cartItem.Quantity
	}

	cart.Price = totalPrice
	cu, err := c.cartRepo.UpdateCart(ctx, cart)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetCartByID",
		}).Error(err.Error())
		return domain.Cart{}, err
	}

	log.WithFields(log.Fields{
		"CartID": cart.ID,
	}).Info("GetCartByID OK")
	return cu, nil
}

func (c *CartService) ClearCart(ctx context.Context, cartID domain.ID) error {
	cart, err := c.cartRepo.GetCartByID(ctx, cartID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ClearCart",
		}).Error(err.Error())
		return err
	}

	for _, item := range cart.Items {
		err = c.DeleteCartItem(ctx, item.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"from": "ClearCart",
			}).Error(err.Error())
			return err
		}
	}

	cart.Price = 0
	_, err = c.cartRepo.UpdateCart(ctx, cart)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ClearCart",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"CartID": cart.ID,
	}).Info("ClearCart OK")
	return nil
}

func (c *CartService) GetCartItemByID(ctx context.Context, cartItemID domain.ID) (domain.CartItem, error) {
	cartItem, err := c.cartRepo.GetCartItemByID(ctx, cartItemID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetCartItemByID",
		}).Error(err.Error())
		return domain.CartItem{}, err
	}

	log.WithFields(log.Fields{
		"CartID": cartItem.CartID,
	}).Info("GetCartItemByID OK")
	return cartItem, nil
}

func (c *CartService) CreateCartItem(ctx context.Context, param port.CreateCartItemParam) (domain.CartItem, error) {
	shopItem, err := c.shopRepo.GetShopItemByProductID(ctx, param.ProductID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateCartItem",
		}).Error(err.Error())
		return domain.CartItem{}, err
	}

	if param.Quantity < 1 || param.Quantity > shopItem.Quantity {
		log.WithFields(log.Fields{
			"from": "CreateCartItem",
		}).Error(domain.ErrQuantityItems.Error())
		return domain.CartItem{}, domain.ErrQuantityItems
	}

	ci, err := c.cartRepo.CreateCartItem(ctx, domain.CartItem{
		ID:        domain.NewID(),
		CartID:    param.CartID,
		ProductID: param.ProductID,
		Quantity:  param.Quantity,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateCartItem",
		}).Error(err.Error())
		return domain.CartItem{}, err
	}

	log.WithFields(log.Fields{
		"CartID":    ci.CartID,
		"ProductID": ci.ProductID,
		"Quantity":  ci.Quantity,
	}).Info("CreateCartItem OK")
	return ci, nil
}

func (c *CartService) UpdateCartItem(ctx context.Context, cartItemID domain.ID, param port.UpdateCartItemParam) (domain.CartItem, error) {
	cartItem, err := c.cartRepo.GetCartItemByID(ctx, cartItemID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateCartItem",
		}).Error(err.Error())
		return domain.CartItem{}, err
	}

	if param.Quantity.Valid {
		cartItem.Quantity = param.Quantity.Int64
	}

	if cartItem.Quantity < 1 {
		log.WithFields(log.Fields{
			"from": "UpdateCartItem",
		}).Error(err.Error())
		return domain.CartItem{}, c.DeleteCartItem(ctx, cartItemID)
	}

	shopItem, err := c.shopRepo.GetShopItemByProductID(ctx, cartItem.ProductID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateCartItem",
		}).Error(err.Error())
		return domain.CartItem{}, err
	}

	if cartItem.Quantity > shopItem.Quantity {
		return domain.CartItem{}, domain.ErrQuantityItems
	}

	ci, err := c.cartRepo.UpdateCartItem(ctx, cartItem)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateCartItem",
		}).Error(err.Error())
		return domain.CartItem{}, err
	}

	log.WithFields(log.Fields{
		"CartID":    ci.CartID,
		"ProductID": ci.ProductID,
		"Quantity":  ci.Quantity,
	}).Info("UpdateCartItem OK")
	return ci, nil
}

func (c *CartService) DeleteCartItem(ctx context.Context, cartItemID domain.ID) error {
	err := c.cartRepo.DeleteCartItem(ctx, cartItemID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "DeleteCartItem",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"cartItemID": cartItemID,
	}).Info("DeleteCartItem OK")
	return nil
}
