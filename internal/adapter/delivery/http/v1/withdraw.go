package v1

import (
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initWithdrawsRoutes(api *gin.RouterGroup) {
	withdrawsGroup := api.Group("/withdraw", h.verifyToken)
	{
		withdrawsSellerGroup := withdrawsGroup.Group("/seller", h.verifyRoleSeller)
		{
			withdrawsSellerGroup.GET("/:shop_id", h.getWithdrawByShopID)
			withdrawsSellerGroup.POST("/:shop_id", h.createWithdraw)
		}
		withdrawsModeratorGroup := withdrawsGroup.Group("/moderator", h.verifyRoleModerator)
		{
			withdrawsModeratorGroup.GET("", h.getWithdrawsAll)
			withdrawsModeratorGroup.PUT("/:id", h.updateWithdraw)
		}
	}
}

// @Summary getWithdrawsAll
// @Security ApiKeyAuth
// @Tags withdraw
// @Description get withdraws
// @Param limit query string true "limit"
// @Param offset query string true "offset"
// @Success 200 {object} []dto.WithdrawDTO
// @Router /api/v1/withdraw/moderator [get]
func (h *Handler) getWithdrawsAll(context *gin.Context) {
	limit, err := getQueryParamInt64(context, "limit")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	offset, err := getQueryParamInt64(context, "offset")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	withdraws, err := h.withdrawService.Get(context.Request.Context(), limit, offset)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := make([]dto.WithdrawDTO, 0, len(withdraws))
	for _, withdraw := range withdraws {
		w := dto.NewWithdrawDTO(withdraw)
		res = append(res, *w)
	}

	h.successResponse(context, res)
}

// @Summary getWithdrawByShopID
// @Security ApiKeyAuth
// @Tags withdraw
// @Description get withdraws by shopID
// @Param shop_id path string true "shopID"
// @Success 200 {object} []dto.WithdrawDTO
// @Router /api/v1/withdraw/seller/{shop_id} [get]
func (h *Handler) getWithdrawByShopID(context *gin.Context) {
	shopID, err := getIdFromPath(context, "shop_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.verifyUserIsShopOwner(context, shopID)

	withdraws, err := h.withdrawService.GetByShopID(context.Request.Context(), shopID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := make([]dto.WithdrawDTO, 0, len(withdraws))
	for _, withdraw := range withdraws {
		w := dto.NewWithdrawDTO(withdraw)
		res = append(res, *w)
	}

	h.successResponse(context, res)
}

// @Summary createWithdraw
// @Security ApiKeyAuth
// @Tags withdraw
// @Description create withdraw
// @Param shop_id path string true "shopID"
// @Param input body dto.CreateWithdrawDTO true "withdraw info"
// @Success 200 {object} dto.WithdrawDTO
// @Router /api/v1/withdraw/seller/{shop_id} [post]
func (h *Handler) createWithdraw(context *gin.Context) {
	shopID, err := getIdFromPath(context, "shop_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.verifyUserIsShopOwner(context, shopID)

	var createWithdrawDTO dto.CreateWithdrawDTO
	err = context.ShouldBindJSON(&createWithdrawDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	withdraw, err := h.withdrawService.Create(context.Request.Context(), port.CreateWithdrawParam{
		ShopID:  shopID,
		Comment: createWithdrawDTO.Comment,
		Sum:     createWithdrawDTO.Sum,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	w := dto.NewWithdrawDTO(withdraw)

	h.successResponse(context, w)
}

// @Summary updateWithdraw
// @Security ApiKeyAuth
// @Tags withdraw
// @Description update withdraw
// @Param id path string true "withdrawID"
// @Param status query string true "status"
// @Success 200 {object} dto.WithdrawDTO
// @Router /api/v1/withdraw/moderator/{id} [put]
func (h *Handler) updateWithdraw(context *gin.Context) {
	withdrawID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res, err := getQueryParamInt64(context, "status")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	status := domain.WithdrawStatus(res)

	withdraw, err := h.withdrawService.Update(context.Request.Context(), withdrawID, port.UpdateWithdrawParam{
		Status: &status,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	w := dto.NewWithdrawDTO(withdraw)

	h.successResponse(context, w)
}
