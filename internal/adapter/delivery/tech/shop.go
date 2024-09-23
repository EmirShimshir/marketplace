package tech

import (
	"bytes"
	"context"
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/tech/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
	log "github.com/sirupsen/logrus"
	"os"
)

func (h *Handler) GetShopItem(c *Console) {
	var productID domain.ID
	err := dto.InputID(&productID, "товара")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	shopItem, err := h.shop.GetShopItemByProductID(context.Background(), productID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	product, err := h.product.GetByID(context.Background(), productID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	si := dto.NewShopItemDTO(shopItem, product)
	si.Print()
}

func (h *Handler) GetShopItemsAll(c *Console) {
	shopItems, err := h.shop.GetShopItems(context.Background(), 100, 0)
	if err != nil {
		ErrorResponse(err)
		return
	}
	fmt.Println("Товары:")
	for _, shopItem := range shopItems {
		product, err := h.product.GetByID(context.Background(), shopItem.ProductID)
		if err != nil {
			ErrorResponse(err)
			return
		}
		si := dto.NewShopItemDTO(shopItem, product)
		si.Print()
	}
}

func (h *Handler) GetShopItemsByShopID(c *Console) {
	var shopID domain.ID
	err := dto.InputID(&shopID, "магазина")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopItemsByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	shop, err := h.shop.GetShopByID(context.Background(), shopID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("Магазин:")
	fmt.Printf("Название: %s\n", shop.Name)
	fmt.Printf("Описание: %s\n", shop.Description)

	if len(shop.Items) == 0 {
		fmt.Println("В магазине нет товаров")
		return
	}

	fmt.Println("Товары:")
	for _, shopItem := range shop.Items {
		product, err := h.product.GetByID(context.Background(), shopItem.ProductID)
		if err != nil {
			ErrorResponse(err)
			return
		}
		si := dto.NewShopItemDTO(shopItem, product)
		si.Print()
	}
}

func (h *Handler) GetShopsBySellerID(c *Console) {
	err := h.verifyAuthRole(c, domain.UserSeller)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopsBySellerID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	shops, err := h.shop.GetShopBySellerID(context.Background(), *c.UserID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(shops) == 0 {
		fmt.Println("У вас нет магазинов")
		return
	}

	fmt.Println("Ваши магазины:")
	for _, shop := range shops {
		s := dto.NewShopDTO(shop)
		s.Print()
	}
}

func (h *Handler) CreateShop(console *Console) {
	err := h.verifyAuthRole(console, domain.UserSeller)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShop",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var createShopDTO dto.CreateShopDTO
	err = dto.InputCreateShopDTO(&createShopDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShop",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	shop, err := h.shop.CreateShop(context.Background(), *console.UserID, port.CreateShopParam{
		Name:        createShopDTO.Name,
		Description: createShopDTO.Description,
		Requisites:  createShopDTO.Requisites,
		Email:       createShopDTO.Email,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	s := dto.NewShopDTO(shop)
	s.Print()
}

func (h *Handler) CreateShopItem(console *Console) {
	var shopID domain.ID
	err := dto.InputID(&shopID, "магазина")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	err = h.verifyUserIsShopOwner(console, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var createShopItemDTO dto.CreateShopItemDTO
	err = dto.InputCreateShopItemDTO(&createShopItemDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	data, err := os.ReadFile(createShopItemDTO.PhotoUrl)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	productParam := port.CreateProductParam{
		Name:        createShopItemDTO.Name,
		Description: createShopItemDTO.Description,
		Price:       createShopItemDTO.Price,
		Category:    createShopItemDTO.Category,
		PhotoReader: bytes.NewReader(data),
	}

	param := port.CreateShopItemParam{
		ShopID:       shopID,
		ProductParam: productParam,
		Quantity:     createShopItemDTO.Quantity,
	}

	shopItem, err := h.shop.CreateShopItem(context.Background(), param)
	if err != nil {
		ErrorResponse(err)
		return
	}

	product, err := h.product.GetByID(context.Background(), shopItem.ProductID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	s := dto.NewShopItemDTO(shopItem, product)
	s.Print()
}

func (h *Handler) UpdateShopItem(console *Console) {
	var productID domain.ID
	err := dto.InputID(&productID, "товара")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	shopItem, err := h.shop.GetShopItemByProductID(context.Background(), productID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	err = h.verifyUserIsShopOwner(console, shopItem.ShopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var count int
	fmt.Print("Введите количество: ")
	fmt.Scanf("%d", &count)

	param := port.UpdateShopItemParam{
		Quantity: null.IntFrom(int64(count)),
	}

	shopItem, err = h.shop.UpdateShopItem(context.Background(), shopItem.ID, param)
	if err != nil {
		ErrorResponse(err)
		return
	}

	product, err := h.product.GetByID(context.Background(), shopItem.ProductID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	s := dto.NewShopItemDTO(shopItem, product)
	s.Print()
}

func (h *Handler) DeleteShopItem(console *Console) {
	err := h.verifyAuthRole(console, domain.UserModerator)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "DeleteShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var productID domain.ID
	err = dto.InputID(&productID, "товара")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "DeleteShopItem",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	shopItem, err := h.shop.GetShopItemByProductID(context.Background(), productID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	err = h.shop.DeleteShopItem(context.Background(), shopItem.ID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("Товар успешно удален из магазина")
}

func (h *Handler) verifyUserIsShopOwner(console *Console, shopID domain.ID) error {
	err := h.verifyAuthRole(console, domain.UserSeller)
	if err != nil {
		return err
	}

	shop, err := h.shop.GetShopByID(context.Background(), shopID)
	if err != nil {
		return err
	}

	if shop.SellerID != *console.UserID {
		return ForbiddenError
	}

	return nil
}
