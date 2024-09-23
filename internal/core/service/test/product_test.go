package service_test

import (
	"context"
	"errors"
	repoMocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	storageMocks "github.com/EmirShimshir/marketplace/internal/adapter/storage/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/stretchr/testify/require"
	"testing"
)

var unknownProductID = domain.NewID()

func TestProductService_GetByID(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(ProductRepo *repoMocks.ProductRepository)
		out          domain.Product
		hasError     bool
	}{
		{
			name: "product found, ok",
			out:  dataProducts[0],
			initRepoMock: func(ProductRepo *repoMocks.ProductRepository) {
				ProductRepo.On("GetByID", context.Background(), dataProducts[0].ID).
					Return(dataProducts[0], nil)
			},
			hasError: false,
		},
		{
			name: "product not found, error",
			out:  domain.Product{ID: unknownProductID},
			initRepoMock: func(ProductRepo *repoMocks.ProductRepository) {
				ProductRepo.On("GetByID", context.Background(), unknownProductID).
					Return(domain.Product{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			productRepo := repoMocks.NewProductRepository(t)
			storage := storageMocks.NewObjectStorage(t)
			productService := service.NewProductService(productRepo, storage)

			test.initRepoMock(productRepo)

			product, err := productService.GetByID(context.Background(), test.out.ID)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.out.ID, product.ID)
			}
		})
	}
}
