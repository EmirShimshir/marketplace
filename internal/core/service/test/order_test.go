package service_test

import (
	"context"
	"errors"
	repoMocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var testCreateOrderCustomerParam = []port.CreateOrderCustomerParam{
	{
		CustomerID: dataUsers[0].ID,
		Address:    "Pushkina 1-2-3",
	},
	{
		CustomerID: dataUsers[2].ID,
		Address:    "Pushkina 1-2-4",
	},
	{
		CustomerID: domain.NewID(),
		Address:    "Pushkina 1-2-5",
	},
}

var testCreateOrderCustomersOut = []domain.OrderCustomer{
	dataOrderCustomers[0],
}

func TestOrderService_CreateOrderCustomer(t *testing.T) {
	testTable := []struct {
		name              string
		initOrderRepoMock func(orderRepo *repoMocks.OrderRepository)
		initUserRepoMock  func(userRepo *repoMocks.UserRepository)
		initCartRepoMock  func(cartRepo *repoMocks.CartRepository)
		initShopRepoMock  func(shopRepo *repoMocks.ShopRepository)
		param             port.CreateOrderCustomerParam
		out               domain.OrderCustomer
		hasError          bool
	}{
		{
			name:  "order customer created, ok",
			param: testCreateOrderCustomerParam[0],
			initUserRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), testCreateOrderCustomerParam[0].CustomerID).
					Return(dataUsers[0], nil)
			},
			initCartRepoMock: func(cartRepo *repoMocks.CartRepository) {
				cartRepo.On("GetCartByID", context.Background(), dataUsers[0].CartID).
					Return(dataCarts[0], nil)
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[0].ProductID).
					Return(dataShopItems[0], nil).Once()
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[1].ProductID).
					Return(dataShopItems[1], nil).Once()
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
				orderRepo.On("CreateOrderCustomer", context.Background(), mock.AnythingOfType("domain.OrderCustomer")).
					Return(dataOrderCustomers[0], nil)
			},
			out:      testCreateOrderCustomersOut[0],
			hasError: false,
		},
		{
			name:  "empty cart, error",
			param: testCreateOrderCustomerParam[1],
			initUserRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), testCreateOrderCustomerParam[1].CustomerID).
					Return(dataUsers[2], nil)
			},
			initCartRepoMock: func(cartRepo *repoMocks.CartRepository) {
				cartRepo.On("GetCartByID", context.Background(), dataUsers[2].CartID).
					Return(dataCarts[2], nil)
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
			},
			out:      domain.OrderCustomer{},
			hasError: true,
		},
		{
			name:  "userRepo GetCartByID, error",
			param: testCreateOrderCustomerParam[2],
			initUserRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), testCreateOrderCustomerParam[2].CustomerID).
					Return(domain.User{}, errors.New("error"))
			},
			initCartRepoMock: func(cartRepo *repoMocks.CartRepository) {
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
			},
			out:      domain.OrderCustomer{},
			hasError: true,
		},
		{
			name:  "cartRepo GetCartByID, error",
			param: testCreateOrderCustomerParam[0],
			initUserRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), testCreateOrderCustomerParam[0].CustomerID).
					Return(dataUsers[0], nil)
			},
			initCartRepoMock: func(cartRepo *repoMocks.CartRepository) {
				cartRepo.On("GetCartByID", context.Background(), dataUsers[0].CartID).
					Return(dataCarts[0], nil)
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[0].ProductID).
					Return(dataShopItems[0], nil).Once()
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[1].ProductID).
					Return(domain.ShopItem{}, errors.New("error")).Once()
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
			},
			out:      domain.OrderCustomer{},
			hasError: true,
		},
		{
			name:  "getCartItemsByShopID, error",
			param: testCreateOrderCustomerParam[0],
			initUserRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), testCreateOrderCustomerParam[0].CustomerID).
					Return(dataUsers[0], nil)
			},
			initCartRepoMock: func(cartRepo *repoMocks.CartRepository) {
				cartRepo.On("GetCartByID", context.Background(), dataUsers[0].CartID).
					Return(dataCarts[0], nil)
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[0].ProductID).
					Return(dataShopItems[0], nil).Once()
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[1].ProductID).
					Return(domain.ShopItem{}, errors.New("error")).Once()
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
			},
			out:      domain.OrderCustomer{},
			hasError: true,
		},
		{
			name:  "orderRepo CreateOrderCustomer, error",
			param: testCreateOrderCustomerParam[0],
			initUserRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), testCreateOrderCustomerParam[0].CustomerID).
					Return(dataUsers[0], nil)
			},
			initCartRepoMock: func(cartRepo *repoMocks.CartRepository) {
				cartRepo.On("GetCartByID", context.Background(), dataUsers[0].CartID).
					Return(dataCarts[0], nil)
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[0].ProductID).
					Return(dataShopItems[0], nil).Once()
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[1].ProductID).
					Return(dataShopItems[1], nil).Once()
			},
			initOrderRepoMock: func(orderRepo *repoMocks.OrderRepository) {
				orderRepo.On("CreateOrderCustomer", context.Background(), mock.AnythingOfType("domain.OrderCustomer")).
					Return(domain.OrderCustomer{}, errors.New("error"))
			},
			out:      domain.OrderCustomer{},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			orderRepo := repoMocks.NewOrderRepository(t)
			userRepo := repoMocks.NewUserRepository(t)
			cartRepo := repoMocks.NewCartRepository(t)
			shopRepo := repoMocks.NewShopRepository(t)
			orderService := service.NewOrderService(orderRepo, userRepo, cartRepo, shopRepo)

			test.initOrderRepoMock(orderRepo)
			test.initUserRepoMock(userRepo)
			test.initCartRepoMock(cartRepo)
			test.initShopRepoMock(shopRepo)

			orderCustomer, err := orderService.CreateOrderCustomer(context.Background(), test.param)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.out, orderCustomer)
			}
		})
	}
}
