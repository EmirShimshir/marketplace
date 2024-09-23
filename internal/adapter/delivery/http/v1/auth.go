package v1

import (
	"errors"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"net/http"
	"strings"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/sign-in", h.userSignIn)
		authGroup.POST("/logout", h.userLogout)
		authGroup.POST("/sign-up", h.userSignUp)
		authGroup.POST("/refresh", h.userRefresh)
	}
}

// @Summary SignIn
// @Tags auth
// @Description login
// @Param input body dto.SignInDTO true "credentials"
// @Router /api/v1/auth/sign-in [post]
func (h *Handler) userSignIn(context *gin.Context) {
	var signInDTO dto.SignInDTO
	err := context.ShouldBindJSON(&signInDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}
	authDetails, err := h.authService.SignIn(context, port.SignInParam{
		Email:       signInDTO.Email,
		Password:    signInDTO.Password,
		Fingerprint: signInDTO.Fingerprint,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken.String(),
		86400, "/", h.config.Host, false, false)

	h.successResponse(context, authDetails.AccessToken.String())
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @Param input body dto.SignUpDTO true "account info"
// @Router /api/v1/auth/sign-up [post]
func (h *Handler) userSignUp(context *gin.Context) {
	var signUpDTO dto.SignUpDTO
	err := context.ShouldBindJSON(&signUpDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.authService.SignUp(context, port.SignUpParam{
		Name:     signUpDTO.Name,
		Surname:  signUpDTO.Surname,
		Email:    signUpDTO.Email,
		Password: signUpDTO.Password,
		Phone:    null.StringFromPtr(signUpDTO.Phone),
		Role:     signUpDTO.Role,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.createdResponse(context, "successfully signed up")
}

// @Summary Refresh
// @Security ApiKeyAuth
// @Tags auth
// @Description refresh
// @Param input body dto.RefreshDTO true "fingerprint"
// @Router /api/v1/auth/refresh [post]
func (h *Handler) userRefresh(context *gin.Context) {
	h.refreshToken(context)
}

// @Summary Logout
// @Security ApiKeyAuth
// @Tags auth
// @Description logout
// @Router /api/v1/auth/logout [post]
func (h *Handler) userLogout(context *gin.Context) {
	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	err = h.authService.LogOut(context, domain.Token(refreshCookie))
	if err != nil {
		h.errorResponse(context, err)
	}

	context.SetCookie("refreshToken", "", -1, "/", h.config.Host, false, false)

	h.successResponse(context, "successfully logged out")
}

func (h *Handler) refreshToken(context *gin.Context) {
	var refreshDTO dto.RefreshDTO
	err := context.ShouldBindJSON(&refreshDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	authDetails, err := h.authService.Refresh(context, domain.Token(refreshCookie),
		refreshDTO.Fingerprint)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken.String(),
		86400, "/", h.config.Host, false, false)

	h.successResponse(context, authDetails.AccessToken.String())
}

func (h *Handler) verifyToken(context *gin.Context) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	payload, err := h.authService.Payload(context, domain.Token(tokenString))
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	context.Set("userID", payload.UserID.String())
	context.Set("role", int(payload.Role))
}

func (h *Handler) verifyRoleSeller(context *gin.Context) {
	role, err := getRoleFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	if role != domain.UserSeller {
		h.errorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) verifyRoleCustomer(context *gin.Context) {
	role, err := getRoleFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	if role != domain.UserCustomer {
		h.errorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) verifyRoleModerator(context *gin.Context) {
	role, err := getRoleFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	if role != domain.UserModerator {
		h.errorResponse(context, ForbiddenError)
		return
	}
}

func (h *Handler) verifyUserIsShopOwner(context *gin.Context, shopID domain.ID) {

	shop, err := h.shopService.GetShopByID(context.Request.Context(), shopID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	id, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	if shop.SellerID != id {
		h.errorResponse(context, ForbiddenError)
		return
	}
}

func extractAuthToken(context *gin.Context) (string, error) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}

func (h *Handler) extractIdFromAuthHeader(context *gin.Context) (domain.ID, error) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		return "", err
	}

	payload, err := h.authService.Payload(context, domain.Token(tokenString))
	if err != nil {
		h.errorResponse(context, err)
		return "", err
	}

	return payload.UserID, nil
}
