package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
	"time"
)

type EmailService struct {
	emailProvider port.IEmailProvider
	orderRepo     port.IOrderRepository
	shopRepo      port.IShopRepository
}

func NewEmailService(emailProvider port.IEmailProvider, orderRepo port.IOrderRepository, shopRepo port.IShopRepository) *EmailService {
	return &EmailService{
		emailProvider: emailProvider,
		orderRepo:     orderRepo,
		shopRepo:      shopRepo,
	}
}

func (e *EmailService) Start(ctx context.Context, stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
			orderShops, err := e.orderRepo.GetNoNotifiedOrderShops(ctx)
			if err != nil {
				log.WithFields(log.Fields{
					"from":    "EmailServiceStart",
					"problem": "GetNoNotifiedOrderShops",
				}).Error(err.Error())
				time.Sleep(time.Minute)
				continue
			}
			if len(orderShops) == 0 {
				time.Sleep(time.Minute)
				continue
			}
			for _, orderShop := range orderShops {
				shop, err := e.shopRepo.GetShopByID(ctx, orderShop.ShopID)
				if err != nil {
					log.WithFields(log.Fields{
						"from":    "EmailServiceStart",
						"problem": "GetShopByID",
					}).Error(err.Error())
					continue
				}
				err = e.emailProvider.SendEmail(ctx, port.CartEmailProviderParam{
					Email:   shop.Email,
					Subject: "New order",
					Body:    "You have new order, check your shop order for details",
				})
				if err != nil {
					log.WithFields(log.Fields{
						"from":    "EmailServiceStart",
						"problem": "SendEmail",
					}).Error(err.Error())
					continue
				}
				orderShop.Notified = true
				_, err = e.orderRepo.UpdateOrderShop(ctx, orderShop)
				if err != nil {
					log.WithFields(log.Fields{
						"from":    "EmailServiceStart",
						"problem": "UpdateOrderShop",
					}).Error(err.Error())
					continue
				}
				log.WithFields(log.Fields{
					"from":  "EmailServiceStart",
					"email": shop.Email,
				}).Info("SendEmail OK")
			}
		}
		time.Sleep(time.Minute)
	}
}
