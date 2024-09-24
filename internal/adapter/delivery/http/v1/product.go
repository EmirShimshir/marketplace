package v1

import (
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initProductsRoutes(api *gin.RouterGroup) {
	productGroup := api.Group("/product")
	{
		productGroup.GET("/:id", h.getShopItem)
		productGroup.GET("/all", h.getShopItemsAll)
		productGroup.GET("/shop/:id", h.getShopItemsByShopID)
	}
}

// @Summary GetShopProduct
// @Tags product
// @Description get shop product
// @Param id path string true "productID"
// @Success 200 {object} dto.ShopItemDTO
// @Router /api/v1/product/{id} [get]
func (h *Handler) getShopItem(context *gin.Context) {
	productID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	shopItem, err := h.shopService.GetShopItemByProductID(context.Request.Context(), productID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	product, err := h.productService.GetByID(context.Request.Context(), productID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	ShopItemDTO := dto.NewShopItemDTO(shopItem, product)
	h.successResponse(context, ShopItemDTO)
}

// @Summary GetShopProduct
// @Tags product
// @Description get shop products all
// @Param limit query string true "limit"
// @Param offset query string true "offset"
// @Success 200 {object} []dto.ShopItemDTO
// @Router /api/v1/product/all [get]
func (h *Handler) getShopItemsAll(context *gin.Context) {
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

	shopItems, err := h.shopService.GetShopItems(context.Request.Context(), limit, offset)
	if err != nil {
		h.errorResponse(context, err)
		return
	}
	res := make([]dto.ShopItemDTO, 0, len(shopItems))
	for _, shopItem := range shopItems {
		product, err := h.productService.GetByID(context.Request.Context(), shopItem.ProductID)
		if err != nil {
			h.errorResponse(context, err)
			return
		}
		res = append(res, *dto.NewShopItemDTO(shopItem, product))
	}

	h.successResponse(context, res)
}

// @Summary GetShopItemsByShopID
// @Tags product
// @Description get shop products by shop ID
// @Param id path string true "shopID"
// @Success 200 {object} dto.ShopItemsByShopIdDTO
// @Router /api/v1/product/shop/{id} [get]
func (h *Handler) getShopItemsByShopID(context *gin.Context) {
	shopID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	shop, err := h.shopService.GetShopByID(context.Request.Context(), shopID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	name := shop.Name
	description := shop.Description

	products := make([]dto.ShopItemDTO, 0, len(shop.Items))
	for _, shopItem := range shop.Items {
		product, err := h.productService.GetByID(context.Request.Context(), shopItem.ProductID)
		if err != nil {
			h.errorResponse(context, err)
			return
		}
		products = append(products, *dto.NewShopItemDTO(shopItem, product))
	}

	h.successResponse(context, dto.ShopItemsByShopIdDTO{Name: name, Description: description, Products: products})
}
