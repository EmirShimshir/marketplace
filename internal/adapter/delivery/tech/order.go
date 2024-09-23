package tech

import (
	"context"
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/tech/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) CreateOrderCustomer(console *Console) {
	err := h.verifyAuthRole(console, domain.UserCustomer)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateOrderCustomer",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var address string
	fmt.Println("Введите адрес")
	fmt.Scanln(&address)

	order, err := h.order.CreateOrderCustomer(context.Background(), port.CreateOrderCustomerParam{
		CustomerID: *console.UserID,
		Address:    address,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	user, err := h.user.GetByID(context.Background(), *console.UserID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	err = h.cart.ClearCart(context.Background(), user.CartID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	url, err := h.payment.GetOrderPaymentUrl(context.Background(), order.ID)
	if err != nil {
		ErrorResponse(err)
		return
	}
	fmt.Println("Заказ успешно создан, произведите оплату по ссылке:", url.String())
}

func (h *Handler) GetOrderCustomersByCustomerID(console *Console) {
	err := h.verifyAuthRole(console, domain.UserCustomer)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetOrderCustomersByCustomerID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	orders, err := h.order.GetOrderCustomerByCustomerID(context.Background(), *console.UserID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(orders) == 0 {
		log.WithFields(log.Fields{
			"from": "GetOrderCustomersByCustomerID",
		}).Error(err.Error())
		fmt.Println("У вас нет заказов")
		return
	}

	for _, order := range orders {
		OrderShopDTOs := make([]dto.OrderShopDTO, 0)
		for _, orderShop := range order.OrderShops {
			OrderShopItemDTOs := make([]dto.OrderShopItemDTO, 0)
			for _, item := range orderShop.OrderShopItems {
				product, err := h.product.GetByID(context.Background(), item.ProductID)
				if err != nil {
					ErrorResponse(err)
					return
				}
				OrderShopItemDTOs = append(OrderShopItemDTOs, *dto.NewOrderShopItemDTO(item, product))
			}
			OrderShopDTOs = append(OrderShopDTOs, *dto.NewOrderShopDTO(orderShop, OrderShopItemDTOs))
		}
		o := dto.NewOrderCustomerDTO(order, OrderShopDTOs)
		o.Print()
	}
}

func (h *Handler) GetOrderShopsByShopID(console *Console) {
	var shopID domain.ID
	err := dto.InputID(&shopID, "магазина")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetOrderShopsByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	err = h.verifyUserIsShopOwner(console, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetOrderShopsByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	orders, err := h.order.GetOrderShopByShopID(context.Background(), shopID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	for _, orderShop := range orders {
		OrderShopItemDTOs := make([]dto.OrderShopItemDTO, 0)
		for _, item := range orderShop.OrderShopItems {
			product, err := h.product.GetByID(context.Background(), item.ProductID)
			if err != nil {
				log.WithFields(log.Fields{
					"from": "GetOrderShopsByShopID",
				}).Error(err.Error())
				ErrorResponse(err)
				return
			}
			OrderShopItemDTOs = append(OrderShopItemDTOs, *dto.NewOrderShopItemDTO(item, product))
		}
		o := dto.NewOrderShopDTO(orderShop, OrderShopItemDTOs)
		o.Print()
	}

}

func (h *Handler) UpdateOrderShopStatusByShopID(console *Console) {
	var orderShopID domain.ID
	err := dto.InputID(&orderShopID, "заказа магазина")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateOrderShopStatusByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	orderShop, err := h.order.GetOrderShopByID(context.Background(), orderShopID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	err = h.verifyUserIsShopOwner(console, orderShop.ShopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateOrderShopStatusByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var status domain.OrderShopStatus
	err = dto.InputOrderShopStatus(&status)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateOrderShopStatusByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	order, err := h.order.UpdateOrderShop(context.Background(), orderShopID, port.UpdateOrderShopParam{
		Status: &status,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	OrderShopItemDTOs := make([]dto.OrderShopItemDTO, 0)
	for _, item := range order.OrderShopItems {
		product, err := h.product.GetByID(context.Background(), item.ProductID)
		if err != nil {
			ErrorResponse(err)
			return
		}
		OrderShopItemDTOs = append(OrderShopItemDTOs, *dto.NewOrderShopItemDTO(item, product))
	}
	o := dto.NewOrderShopDTO(order, OrderShopItemDTOs)
	o.Print()
}
