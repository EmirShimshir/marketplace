package service_test

import (
	"context"
	"errors"
	repoMocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/EmirShimshir/marketplace/internal/core/service/mocks"
	"testing"
	"time"
)

var dataUpdatedOrderShops = []domain.OrderShop{
	dataOrderShops[0],
	dataOrderShops[1],
}

func TestEmailService_Start(t *testing.T) {
	testTable := []struct {
		name                  string
		initEmailProviderMock func(emailProvider *mocks.EmailProvider)
		initOrderRepoMock     func(orderRepo *repoMocks.OrderRepository)
		initShopRepoMock      func(shopRepo *repoMocks.ShopRepository)
	}{
		{
			name: "started, ok",
			initEmailProviderMock: func(emailProvider *mocks.EmailProvider) {
				emailProvider.On("SendEmail", context.Background(), port.CartEmailProviderParam{
					Email:   dataShops[0].Email,
					Subject: "New order",
					Body:    "You have new order, check your shop order for details",
				}).Return(nil).Once()
				emailProvider.On("SendEmail", context.Background(), port.CartEmailProviderParam{
					Email:   dataShops[1].Email,
					Subject: "New order",
					Body:    "You have new order, check your shop order for details",
				}).Return(nil).Once()
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
				orderRepo.On("GetNoNotifiedOrderShops", context.Background()).
					Return(dataOrderShops, nil)
				dataUpdatedOrderShops[0].Notified = true
				orderRepo.On("UpdateOrderShop", context.Background(), dataUpdatedOrderShops[0]).
					Return(dataUpdatedOrderShops[0], nil).Once()
				dataUpdatedOrderShops[1].Notified = true
				orderRepo.On("UpdateOrderShop", context.Background(), dataUpdatedOrderShops[1]).
					Return(dataUpdatedOrderShops[1], nil).Once()
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopByID", context.Background(), dataOrderShops[0].ShopID).
					Return(dataShops[0], nil).Once()
				shopRepo.On("GetShopByID", context.Background(), dataOrderShops[1].ShopID).
					Return(dataShops[1], nil).Once()
			},
		},
		{
			name: "started one email error, ok",
			initEmailProviderMock: func(emailProvider *mocks.EmailProvider) {
				emailProvider.On("SendEmail", context.Background(), port.CartEmailProviderParam{
					Email:   dataShops[0].Email,
					Subject: "New order",
					Body:    "You have new order, check your shop order for details",
				}).Return(nil).Once()
				emailProvider.On("SendEmail", context.Background(), port.CartEmailProviderParam{
					Email:   dataShops[1].Email,
					Subject: "New order",
					Body:    "You have new order, check your shop order for details",
				}).Return(errors.New("error")).Once()
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
				orderRepo.On("GetNoNotifiedOrderShops", context.Background()).
					Return(dataOrderShops, nil)
				dataUpdatedOrderShops[0].Notified = true
				orderRepo.On("UpdateOrderShop", context.Background(), dataUpdatedOrderShops[0]).
					Return(dataUpdatedOrderShops[0], nil).Once()
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopByID", context.Background(), dataOrderShops[0].ShopID).
					Return(dataShops[0], nil).Once()
				shopRepo.On("GetShopByID", context.Background(), dataOrderShops[1].ShopID).
					Return(dataShops[1], nil).Once()
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			emailProvider := mocks.NewEmailProvider(t)
			orderRepo := repoMocks.NewOrderRepository(t)
			shopRepo := repoMocks.NewShopRepository(t)
			emailService := service.NewEmailService(emailProvider, orderRepo, shopRepo)

			test.initEmailProviderMock(emailProvider)
			test.initOrderRepoMock(orderRepo)
			test.initShopRepoMock(shopRepo)

			stopCh := make(chan struct{})
			go emailService.Start(context.Background(), stopCh)
			time.Sleep(1 * time.Second)
			stopCh <- struct{}{}
		})
	}
}
