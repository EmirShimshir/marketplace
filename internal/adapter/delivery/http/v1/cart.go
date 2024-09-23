package v1

import (
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1/dto"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) initCartsRoutes(api *gin.RouterGroup) {
	cartsGroup := api.Group("/cart", h.verifyToken)
	{
		cartsCustomerGroup := cartsGroup.Group("/customer", h.verifyRoleCustomer)
		{
			cartsCustomerGroup.GET("", h.getCart)
			cartsCustomerGroup.POST("", h.createCartItem)
			cartsCustomerGroup.PUT("", h.updateCartItem)
			cartsCustomerGroup.DELETE("/:cart_product_id", h.deleteCartItem)
		}
	}
}

// @Summary getCart
// @Security ApiKeyAuth
// @Tags cart
// @Description get cart
// @Success 200 {object} dto.CartDTO
// @Router /api/v1/cart/customer  [get]
func (h *Handler) getCart(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	user, err := h.userService.GetByID(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	cart, err := h.cartService.GetCartByID(context.Request.Context(), user.CartID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	products := make([]dto.CartItemDTO, 0, len(cart.Items))
	for _, item := range cart.Items {
		product, err := h.productService.GetByID(context.Request.Context(), item.ProductID)
		if err != nil {
			h.errorResponse(context, err)
			return
		}
		products = append(products, *dto.NewCartItemDTO(item, product))
	}

	h.successResponse(context, dto.CartDTO{TotalPrice: cart.Price, Products: products})
}

// @Summary createCartItem
// @Security ApiKeyAuth
// @Tags cart
// @Description create cart product
// @Param input body dto.CreateCartItemDTO true "cart product info"
// @Success 200 {object} dto.CartItemDTO
// @Router /api/v1/cart/customer [post]
func (h *Handler) createCartItem(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	user, err := h.userService.GetByID(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var createCartItemDTO dto.CreateCartItemDTO
	err = context.ShouldBindJSON(&createCartItemDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	cartItem, err := h.cartService.CreateCartItem(context.Request.Context(), port.CreateCartItemParam{
		CartID:    user.CartID,
		ProductID: createCartItemDTO.ProductID,
		Quantity:  createCartItemDTO.Quantity,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	product, err := h.productService.GetByID(context.Request.Context(), cartItem.ProductID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateCartItem",
		}).Error(err.Error())
		h.errorResponse(context, err)
		return
	}

	res := dto.NewCartItemDTO(cartItem, product)
	h.successResponse(context, res)
}

// @Summary updateCartItem
// @Security ApiKeyAuth
// @Tags cart
// @Description update cart product
// @Param input body dto.UpdateCartItemDTO true "cart product info"
// @Success 200 {object} dto.CartItemDTO
// @Router /api/v1/cart/customer [put]
func (h *Handler) updateCartItem(context *gin.Context) {
	var updateCartItemDTO dto.UpdateCartItemDTO
	err := context.ShouldBindJSON(&updateCartItemDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	cartItem, err := h.cartService.UpdateCartItem(context.Request.Context(), updateCartItemDTO.CartItemID, port.UpdateCartItemParam{
		Quantity: null.IntFrom(updateCartItemDTO.Quantity),
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	product, err := h.productService.GetByID(context.Request.Context(), cartItem.ProductID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := dto.NewCartItemDTO(cartItem, product)
	h.successResponse(context, res)
}

// @Summary deleteCartItem
// @Security ApiKeyAuth
// @Tags cart
// @Description delete cart product
// @Param cart_product_id path string true "cartItemID"
// @Router /api/v1/cart/customer/{cart_product_id} [delete]
func (h *Handler) deleteCartItem(context *gin.Context) {
	cartItemID, err := getIdFromPath(context, "cart_product_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.cartService.DeleteCartItem(context.Request.Context(), cartItemID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "deleted")
}
