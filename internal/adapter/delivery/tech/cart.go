package tech

import (
	"context"
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/tech/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) CreateCartItem(console *Console) {
	err := h.verifyAuthRole(console, domain.UserCustomer)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateCartItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	user, err := h.user.GetByID(context.Background(), *console.UserID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	var createCartItemDTO dto.CreateCartItemDTO
	err = dto.InputCreateCartItemDTO(&createCartItemDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateCartItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	cartItem, err := h.cart.CreateCartItem(context.Background(), port.CreateCartItemParam{
		CartID:    user.CartID,
		ProductID: createCartItemDTO.ProductID,
		Quantity:  createCartItemDTO.Quantity,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	product, err := h.product.GetByID(context.Background(), cartItem.ProductID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateCartItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	c := dto.NewCartItemDTO(cartItem, product)
	c.Print()
}

func (h *Handler) UpdateCartItem(console *Console) {
	err := h.verifyAuthRole(console, domain.UserCustomer)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateCartItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var updateCartItemDTO dto.UpdateCartItemDTO
	err = dto.InputUpdateCartItemDTO(&updateCartItemDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateCartItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	cartItem, err := h.cart.UpdateCartItem(context.Background(), updateCartItemDTO.CartItemID, port.UpdateCartItemParam{
		Quantity: null.IntFrom(updateCartItemDTO.Quantity),
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	product, err := h.product.GetByID(context.Background(), cartItem.ProductID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	c := dto.NewCartItemDTO(cartItem, product)
	c.Print()
}

func (h *Handler) DeleteCartItem(console *Console) {
	err := h.verifyAuthRole(console, domain.UserCustomer)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "DeleteCartItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var cartItemID domain.ID
	err = dto.InputID(&cartItemID, "товара в корзине")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "DeleteCartItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	err = h.cart.DeleteCartItem(context.Background(), cartItemID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("Товар успешно удален")
}

func (h *Handler) GetCart(console *Console) {
	err := h.verifyAuthRole(console, domain.UserCustomer)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetCart",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	user, err := h.user.GetByID(context.Background(), *console.UserID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	cart, err := h.cart.GetCartByID(context.Background(), user.CartID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(cart.Items) == 0 {
		fmt.Println("Корзина пуста")
		return
	}

	fmt.Println("Товары в корзине:")
	for _, item := range cart.Items {
		product, err := h.product.GetByID(context.Background(), item.ProductID)
		if err != nil {
			ErrorResponse(err)
			return
		}
		ci := dto.NewCartItemDTO(item, product)
		ci.Print()
	}
	fmt.Printf("Общая сумма товаров: %d\n", cart.Price)
}
