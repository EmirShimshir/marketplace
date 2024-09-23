package service_test

import (
	"context"
	repoMocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/EmirShimshir/marketplace/internal/core/service/mocks"
	"github.com/guregu/null"
	"github.com/stretchr/testify/require"
	"testing"
)

var testUpdateShopItemParam = []port.UpdateShopItemParam{
	{
		Quantity: null.IntFrom(2),
	},
	{
		Quantity: null.IntFrom(-1),
	},
	{
		Quantity: null.IntFrom(3),
	},
}

var testShopItemsOuts = []domain.ShopItem{
	{
		ID:        dataShopItems[0].ID,
		ShopID:    dataShopItems[0].ShopID,
		ProductID: dataShopItems[0].ProductID,
		Quantity:  testUpdateShopItemParam[0].Quantity.Int64,
	},
	{
		ID:        dataShopItems[0].ID,
		ShopID:    dataShopItems[0].ShopID,
		ProductID: dataShopItems[0].ProductID,
		Quantity:  testUpdateShopItemParam[2].Quantity.Int64,
	},
}

func TestShopService_UpdateShopItem(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(shopRepo *repoMocks.ShopRepository)
		param        port.UpdateShopItemParam
		out          domain.ShopItem
		hasError     bool
	}{
		{
			name:  "shop item updated, ok",
			param: testUpdateShopItemParam[0],
			initRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByID", context.Background(), dataShopItems[0].ID).
					Return(dataShopItems[0], nil)
				shopRepo.On("UpdateShopItem", context.Background(), testShopItemsOuts[0]).
					Return(testShopItemsOuts[0], nil)
			},

			out:      testShopItemsOuts[0],
			hasError: false,
		},
		{
			name:  "shop item not updated, error quantity",
			param: testUpdateShopItemParam[1],
			initRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByID", context.Background(), dataShopItems[0].ID).
					Return(dataShopItems[0], nil)
			},
			out:      testShopItemsOuts[0],
			hasError: true,
		},
		{
			name:  "shop item updated with product, ok",
			param: testUpdateShopItemParam[2],
			initRepoMock: func(shopRepo *repoMocks.ShopRepository) {
				shopRepo.On("GetShopItemByID", context.Background(), dataShopItems[0].ID).
					Return(dataShopItems[0], nil)
				shopRepo.On("UpdateShopItem", context.Background(), testShopItemsOuts[1]).
					Return(testShopItemsOuts[1], nil)
			},
			out:      testShopItemsOuts[1],
			hasError: false,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			shopRepo := repoMocks.NewShopRepository(t)
			storage := mocks.NewObjectStorage(t)
			shopService := service.NewShopService(shopRepo, storage)

			test.initRepoMock(shopRepo)

			shopItem, err := shopService.UpdateShopItem(context.Background(), test.out.ID, test.param)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.out, shopItem)
			}
		})
	}
}
