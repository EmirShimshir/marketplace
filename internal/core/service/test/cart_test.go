package service_test

import (
	"context"
	repoMocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/stretchr/testify/require"
	"testing"
)

var testCartsOut = []domain.Cart{
	{
		ID:    dataCarts[0].ID,
		Price: 132980,
		Items: []domain.CartItem{
			dataCartItems[0],
			dataCartItems[1],
		},
	},
}

func TestCartService_GetCartByID(t *testing.T) {
	testTable := []struct {
		name                string
		initCartRepoMock    func(cartRepo *repoMocks.CartRepository)
		initShopRepoMock    func(shopRepo *repoMocks.ShopRepository)
		initProductRepoMock func(productRepo *repoMocks.ProductRepository)
		param               domain.ID
		out                 domain.Cart
		hasError            bool
	}{
		{
			name: "cart got, ok",
			initCartRepoMock: func(cartRepo *repoMocks.CartRepository) {
				cartRepo.On("GetCartByID", context.Background(), dataCarts[0].ID).
					Return(dataCarts[0], nil)
				cartRepo.On("UpdateCart", context.Background(), testCartsOut[0]).
					Return(testCartsOut[0], nil)

			},
			initProductRepoMock: func(productRepo *repoMocks.ProductRepository) {
				productRepo.On("GetByID", context.Background(), dataCarts[0].Items[0].ProductID).
					Return(dataProducts[0], nil)
				productRepo.On("GetByID", context.Background(), dataCarts[0].Items[1].ProductID).
					Return(dataProducts[1], nil)
			},
			initShopRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[0].ProductID).
					Return(dataShopItems[0], nil).Once()
				shopRepo.On("GetShopItemByProductID", context.Background(), dataCarts[0].Items[1].ProductID).
					Return(dataShopItems[1], nil).Once()
			},
			param:    dataCarts[0].ID,
			out:      testCartsOut[0],
			hasError: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			cartRepo := repoMocks.NewCartRepository(t)
			shopRepo := repoMocks.NewShopRepository(t)
			productRepo := repoMocks.NewProductRepository(t)
			cartService := service.NewCartService(cartRepo, shopRepo, productRepo)

			test.initCartRepoMock(cartRepo)
			test.initShopRepoMock(shopRepo)
			test.initProductRepoMock(productRepo)

			cart, err := cartService.GetCartByID(context.Background(), test.param)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.out, cart)
			}
		})
	}
}
