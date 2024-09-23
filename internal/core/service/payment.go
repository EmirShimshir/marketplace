package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type PaymentService struct {
	gateway   port.IPaymentGateway
	orderRepo port.IOrderRepository
}

func NewPaymentService(gateway port.IPaymentGateway, orderRepo port.IOrderRepository) *PaymentService {
	return &PaymentService{
		gateway:   gateway,
		orderRepo: orderRepo,
	}
}

func (p *PaymentService) GetOrderPaymentUrl(ctx context.Context, orderID domain.ID) (url.URL, error) {
	orderCustomer, err := p.orderRepo.GetOrderCustomerByID(ctx, orderID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetOrderPaymentUrl",
		}).Error(err.Error())
		return url.URL{}, err
	}

	if orderCustomer.Payed {
		log.WithFields(log.Fields{
			"from": "GetOrderPaymentUrl",
		}).Error(domain.ErrOrderAlreadyPayed.Error())
		return url.URL{}, domain.ErrOrderAlreadyPayed
	}

	u, err := p.gateway.GetPaymentUrl(ctx, domain.PaymentPayload{
		OrderID: orderCustomer.ID,
		PaySum:  orderCustomer.TotalPrice,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetOrderPaymentUrl",
		}).Error(err.Error())
		return url.URL{}, err
	}

	log.WithFields(log.Fields{
		"url":     u.String(),
		"orderID": orderID,
	}).Info("GetOrderPaymentUrl OK")
	return u, nil
}

func (p *PaymentService) ProcessOrderPayment(ctx context.Context, key string) error {
	payload, err := p.gateway.ProcessPayment(ctx, key)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProcessOrderPayment",
		}).Error(err.Error())
		return err
	}

	orderCustomer, err := p.orderRepo.GetOrderCustomerByID(ctx, payload.OrderID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProcessOrderPayment",
		}).Error(err.Error())
		return err
	}

	if orderCustomer.TotalPrice < payload.PaySum {
		log.WithFields(log.Fields{
			"from": "ProcessOrderPayment",
		}).Error(domain.ErrInvalidPaymentSum.Error())
		return domain.ErrInvalidPaymentSum
	}

	err = p.orderRepo.UpdatePaymentStatus(ctx, orderCustomer.ID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProcessOrderPayment",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"orderID": orderCustomer.ID,
	}).Info("ProcessOrderPayment OK")
	return nil
}
