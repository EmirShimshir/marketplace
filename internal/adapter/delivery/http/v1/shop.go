package v1

import (
	"bytes"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1/dto"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"io"
)

func (h *Handler) initShopsRoutes(api *gin.RouterGroup) {
	shopsGroup := api.Group("/shop", h.verifyToken)
	{
		shopsSellerGroup := shopsGroup.Group("/seller", h.verifyRoleSeller)
		{
			shopsSellerGroup.GET("", h.getShopsBySellerID)
			shopsSellerGroup.POST("", h.createShop)
			shopsSellerGroup.POST("/product/:shop_id", h.createShopItem)
			shopsSellerGroup.PUT("/product/:product_id", h.updateShopItem)
		}
		shopsModeratorGroup := shopsGroup.Group("/moderator", h.verifyRoleModerator)
		{
			shopsModeratorGroup.DELETE("/product/:product_id", h.deleteShopItem)
		}
	}
}

// @Summary getShopsBySellerID
// @Security ApiKeyAuth
// @Tags shop
// @Description get shops by sellerID
// @Success 200 {object} []dto.ShopDTO
// @Router /api/v1/shop/seller  [get]
func (h *Handler) getShopsBySellerID(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	shops, err := h.shopService.GetShopBySellerID(context.Request.Context(), userID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := make([]dto.ShopDTO, 0, len(shops))

	for _, shop := range shops {
		res = append(res, *dto.NewShopDTO(shop))
	}

	h.successResponse(context, res)
}

// @Summary createShop
// @Security ApiKeyAuth
// @Tags shop
// @Description create shop
// @Param input body dto.CreateShopDTO true "shop info"
// @Success 200 {object} dto.ShopDTO
// @Router /api/v1/shop/seller [post]
func (h *Handler) createShop(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	var createShopDTO dto.CreateShopDTO
	err = context.ShouldBindJSON(&createShopDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	shop, err := h.shopService.CreateShop(context.Request.Context(), userID, port.CreateShopParam{
		Name:        createShopDTO.Name,
		Description: createShopDTO.Description,
		Requisites:  createShopDTO.Requisites,
		Email:       createShopDTO.Email,
	})
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := dto.NewShopDTO(shop)

	h.successResponse(context, res)
}

// @Summary createShopProduct
// @Security ApiKeyAuth
// @Tags shop
// @Description create shop product
// @Accept multipart/form-data
// @Param shop_id path string true "shopID"
// @Param json formData string true "product info"
// @Param file formData file true "Upload file"
// @Success 200 {object} dto.ShopItemDTO
// @Router /api/v1/shop/seller/product/{shop_id} [post]
func (h *Handler) createShopItem(context *gin.Context) {
	shopID, err := getIdFromPath(context, "shop_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}
	h.verifyUserIsShopOwner(context, shopID)

	var createShopItemFormsDTO dto.CreateShopItemFormsDTO
	err = context.ShouldBind(&createShopItemFormsDTO)
	if err != nil {
		h.errorResponse(context, err)
		return
	}
	f, err := createShopItemFormsDTO.File.Open()
	data, err := io.ReadAll(f)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	productParam := port.CreateProductParam{
		Name:        createShopItemFormsDTO.Json.Name,
		Description: createShopItemFormsDTO.Json.Description,
		Price:       createShopItemFormsDTO.Json.Price,
		Category:    createShopItemFormsDTO.Json.Category,
		PhotoReader: bytes.NewReader(data),
	}

	param := port.CreateShopItemParam{
		ShopID:       shopID,
		ProductParam: productParam,
		Quantity:     createShopItemFormsDTO.Json.Quantity,
	}

	shopItem, err := h.shopService.CreateShopItem(context.Request.Context(), param)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	product, err := h.productService.GetByID(context.Request.Context(), shopItem.ProductID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := dto.NewShopItemDTO(shopItem, product)
	h.successResponse(context, res)
}

// @Summary updateShopItem
// @Security ApiKeyAuth
// @Tags shop
// @Description update shop product
// @Param product_id path string true "productID"
// @Param count query string true "count product"
// @Success 200 {object} dto.ShopItemDTO
// @Router /api/v1/shop/seller/product/{product_id} [put]
func (h *Handler) updateShopItem(context *gin.Context) {
	productID, err := getIdFromPath(context, "product_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	shopItem, err := h.shopService.GetShopItemByProductID(context.Request.Context(), productID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.verifyUserIsShopOwner(context, shopItem.ShopID)

	count, err := getQueryParamInt64(context, "count")
	if err != nil {
		h.errorResponse(context, err)
		return
	}
	param := port.UpdateShopItemParam{
		Quantity: null.IntFrom(count),
	}

	shopItem, err = h.shopService.UpdateShopItem(context.Request.Context(), shopItem.ID, param)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	product, err := h.productService.GetByID(context.Request.Context(), shopItem.ProductID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	res := dto.NewShopItemDTO(shopItem, product)
	h.successResponse(context, res)
}

// @Summary deleteShopItem
// @Security ApiKeyAuth
// @Tags shop
// @Description delete shop product
// @Param product_id path string true "productID"
// @Router /api/v1/shop/moderator/product/{product_id} [delete]
func (h *Handler) deleteShopItem(context *gin.Context) {
	productID, err := getIdFromPath(context, "product_id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	shopItem, err := h.shopService.GetShopItemByProductID(context.Request.Context(), productID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.shopService.DeleteShopItem(context.Request.Context(), shopItem.ID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "deleted")
}
