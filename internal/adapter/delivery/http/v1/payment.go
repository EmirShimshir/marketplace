package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initPaymentsRoutes(api *gin.RouterGroup) {
	paymentsGroup := api.Group("/payment", h.verifyToken)
	{
		withdrawsCustomerGroup := paymentsGroup.Group("/customer", h.verifyRoleCustomer)
		{
			withdrawsCustomerGroup.POST("", h.payOrder)
		}
	}
}

// @Summary payOrder
// @Security ApiKeyAuth
// @Tags payment
// @Description pay order
// @Param key query string true "key"
// @Router /api/v1/payment/customer [POST]
func (h *Handler) payOrder(context *gin.Context) {
	key, err := getQueryParamString(context, "key")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.paymentService.ProcessOrderPayment(context.Request.Context(), key)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "paid")
}
