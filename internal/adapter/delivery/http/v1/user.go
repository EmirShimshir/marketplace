package v1

import (
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1/dto"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	usersGroup := api.Group("/user", h.verifyToken)
	{
		usersGroup.GET("", h.getUser)
		usersGroup.PUT("", h.updateUser)
	}
}

// @Summary GetUser
// @Security ApiKeyAuth
// @Tags user
// @Description get user
// @Success 200 {object} dto.UserDTO
// @Router /api/v1/user  [get]
func (h *Handler) getUser(context *gin.Context) {
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

	userDTO := dto.NewUserDTO(user)
	h.successResponse(context, userDTO)
}

// @Summary UpdateUser
// @Security ApiKeyAuth
// @Tags user
// @Description update user
// @Param input body dto.UpdateUserDTO true "user info"
// @Success 200 {object} dto.UserDTO
// @Router /api/v1/user [put]
func (h *Handler) updateUser(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	var updateUserDTO dto.UpdateUserDTO
	err = context.ShouldBindJSON(&updateUserDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	user, err := h.userService.Update(context.Request.Context(), userID, port.UpdateUserParam{
		Name:    null.StringFromPtr(updateUserDTO.Name),
		Surname: null.StringFromPtr(updateUserDTO.Surname),
		Phone:   null.StringFromPtr(updateUserDTO.Phone),
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userDTO := dto.NewUserDTO(user)
	h.successResponse(context, userDTO)
}
