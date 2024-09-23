package v1

import (
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) initOrdersRoutes(api *gin.RouterGroup) {
	ordersGroup := api.Group("/order", h.verifyToken)
	{
		ordersCustomerGroup := ordersGroup.Group("/customer", h.verifyRoleCustomer)
		{
			ordersCustomerGroup.GET("", h.getOrderCustomersByCustomerID)
			ordersCustomerGroup.POST("", h.createOrderCustomer)
		}
		ordersSellerGroup := ordersGroup.Group("/seller", h.verifyRoleSeller)
		{
			ordersSellerGroup.GET("/:shop_id", h.getOrderShopsByShopID)
			ordersSellerGroup.PUT("/:order_shop_id", h.updateOrderShopStatusByShopID)
		}
	}
}

// @Summary createOrderCustomer
// @Security ApiKeyAuth
// @Tags order
// @Description create order customer
// @Param input body dto.CreateOrderCustomerDTO true "order customer info"
// @Router /api/v1/order/customer [post]
func (h *Handler) createOrderCustomer(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var orderCustomerDTO dto.CreateOrderCustomerDTO
	err = context.ShouldBindJSON(&orderCustomerDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	order, err := h.orderService.CreateOrderCustomer(context.Request.Context(), port.CreateOrderCustomerParam{
		CustomerID: userID,
		Address:    orderCustomerDTO.Address,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	user, err := h.userService.GetByID(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.cartService.ClearCart(context.Request.Context(), user.CartID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	url, err := h.paymentService.GetOrderPaymentUrl(context.Request.Context(), order.ID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, url.String())
}

// @Summary getOrderCustomersByCustomerID
// @Security ApiKeyAuth
// @Tags order
// @Description get order customer
// @Success 200 {object} []dto.OrderCustomerDTO
// @Router /api/v1/order/customer [get]
func (h *Handler) getOrderCustomersByCustomerID(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	orders, err := h.orderService.GetOrderCustomerByCustomerID(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := make([]dto.OrderCustomerDTO, 0, len(orders))
	for _, order := range orders {
		OrderShopDTOs := make([]dto.OrderShopDTO, 0)
		for _, orderShop := range order.OrderShops {
			OrderShopItemDTOs := make([]dto.OrderShopItemDTO, 0)
			for _, item := range orderShop.OrderShopItems {
				product, err := h.productService.GetByID(context.Request.Context(), item.ProductID)
				if err != nil {
					h.errorResponse(context, err)
					return
				}
				OrderShopItemDTOs = append(OrderShopItemDTOs, *dto.NewOrderShopItemDTO(item, product))
			}
			OrderShopDTOs = append(OrderShopDTOs, *dto.NewOrderShopDTO(orderShop, OrderShopItemDTOs))
		}
		o := dto.NewOrderCustomerDTO(order, OrderShopDTOs)
		res = append(res, *o)
	}

	h.successResponse(context, res)
}

// @Summary getOrderShopsByShopID
// @Security ApiKeyAuth
// @Tags order
// @Description get order shops
// @Param shop_id path string true "shopID"
// @Success 200 {object} []dto.OrderShopDTO
// @Router /api/v1/order/seller/{shop_id} [get]
func (h *Handler) getOrderShopsByShopID(context *gin.Context) {
	shopID, err := getIdFromPath(context, "shop_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.verifyUserIsShopOwner(context, shopID)

	orders, err := h.orderService.GetOrderShopByShopID(context.Request.Context(), shopID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := make([]dto.OrderShopDTO, 0, len(orders))
	for _, orderShop := range orders {
		OrderShopItemDTOs := make([]dto.OrderShopItemDTO, 0)
		for _, item := range orderShop.OrderShopItems {
			product, err := h.productService.GetByID(context.Request.Context(), item.ProductID)
			if err != nil {
				log.WithFields(log.Fields{
					"from": "GetOrderShopsByShopID",
				}).Error(err.Error())
				h.errorResponse(context, UnauthorizedError)
				return
			}
			OrderShopItemDTOs = append(OrderShopItemDTOs, *dto.NewOrderShopItemDTO(item, product))
		}
		o := dto.NewOrderShopDTO(orderShop, OrderShopItemDTOs)
		res = append(res, *o)
	}

	h.successResponse(context, res)
}

// @Summary updateOrderShopStatusByShopID
// @Security ApiKeyAuth
// @Tags order
// @Description update order shop status
// @Param order_shop_id path string true "orderShopID"
// @Param status query string true "status"
// @Success 200 {object} []dto.OrderShopDTO
// @Router /api/v1/order/seller/{order_shop_id} [put]
func (h *Handler) updateOrderShopStatusByShopID(context *gin.Context) {
	orderShopID, err := getIdFromPath(context, "order_shop_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	orderShop, err := h.orderService.GetOrderShopByID(context.Request.Context(), orderShopID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.verifyUserIsShopOwner(context, orderShop.ShopID)

	var status domain.OrderShopStatus
	res, err := getQueryParamInt64(context, "status")
	status = domain.OrderShopStatus(res)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	order, err := h.orderService.UpdateOrderShop(context.Request.Context(), orderShopID, port.UpdateOrderShopParam{
		Status: &status,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	OrderShopItemDTOs := make([]dto.OrderShopItemDTO, 0)
	for _, item := range order.OrderShopItems {
		product, err := h.productService.GetByID(context.Request.Context(), item.ProductID)
		if err != nil {
			h.errorResponse(context, err)
			return
		}
		OrderShopItemDTOs = append(OrderShopItemDTOs, *dto.NewOrderShopItemDTO(item, product))
	}
	o := dto.NewOrderShopDTO(order, OrderShopItemDTOs)

	h.successResponse(context, o)
}
