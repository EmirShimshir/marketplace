package test

import (
	"context"
	"testing"

	repomocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	storagemocks "github.com/EmirShimshir/marketplace/internal/adapter/storage/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

// ProductSuite is the base suite for Product tests.
type ProductSuite struct {
	suite.Suite
}

// GetByID Suite
type ProductGetByIDSuite struct {
	ProductSuite
}

func ProductGetByIDSuccessRepositoryMock(repository *repomocks.ProductRepository, productID domain.ID) {
	repository.
		On("GetByID", context.Background(), productID).
		Return(NewProductBuilder().WithID(productID).Build(), nil)
}

func (s *ProductGetByIDSuite) TestGetByID_Success(t provider.T) {
	t.Title("Get by ID success")
	productID := domain.NewID()
	productRepository := repomocks.NewProductRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	productService := service.NewProductService(productRepository, storage)
	ProductGetByIDSuccessRepositoryMock(productRepository, productID)
	product, err := productService.GetByID(context.Background(), productID)
	t.Assert().Nil(err)
	t.Assert().Equal(productID, product.ID)
}

func ProductGetByIDFailureRepositoryMock(repository *repomocks.ProductRepository, productID domain.ID) {
	repository.
		On("GetByID", context.Background(), productID).
		Return(domain.Product{}, domain.ErrNotExist)
}

func (s *ProductGetByIDSuite) TestGetByID_Failure(t provider.T) {
	t.Title("Get by ID failure")
	productID := domain.NewID()
	productRepository := repomocks.NewProductRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	productService := service.NewProductService(productRepository, storage)
	ProductGetByIDFailureRepositoryMock(productRepository, productID)
	_, err := productService.GetByID(context.Background(), productID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestProductGetByIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetByID", new(ProductGetByIDSuite))
}
