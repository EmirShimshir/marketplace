package test

import (
	"context"
	"testing"

	"github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/guregu/null"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
)

// CartGetByID Suite
type CartGetByIDSuite struct {
	suite.Suite
}

func CartGetByIDSuccessRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("GetCartByID", context.Background(), mock.Anything).
		Return(domain.Cart{
			ID:    domain.NewID(),
			Price: 100,
			Items: []domain.CartItem{NewCartItemBuilder().Build()},
		}, nil)
}

func CartGetByIDFailureRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("GetCartByID", context.Background(), mock.Anything).
		Return(domain.Cart{}, domain.ErrNotExist)
}

func (s *CartGetByIDSuite) TestGetCartByID_Success(t provider.T) {
	t.Parallel()
	t.Title("GetCartByID success")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartGetByIDSuccessRepositoryMock(cartRepository)

	_, err := cartService.GetCartByID(context.Background(), domain.NewID())
	t.Assert().Nil(err)
}

func (s *CartGetByIDSuite) TestGetCartByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("GetCartByID failure")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartGetByIDFailureRepositoryMock(cartRepository)

	_, err := cartService.GetCartByID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestCartGetByIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetCartByID", new(CartGetByIDSuite))
}

// CartClear Suite
type CartClearSuite struct {
	suite.Suite
}

func CartClearSuccessRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("GetCartByID", context.Background(), mock.Anything).
		Return(NewCartBuilder().WithItems([]domain.CartItem{{ID: domain.NewID()}}).Build(), nil)

	repository.
		On("UpdateCart", context.Background(), mock.Anything).
		Return(domain.Cart{}, nil)
}

func CartClearFailureRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("GetCartByID", context.Background(), mock.Anything).
		Return(domain.Cart{}, domain.ErrNotExist)
}

func (s *CartClearSuite) TestClearCart_Success(t provider.T) {
	t.Parallel()
	t.Title("ClearCart success")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartClearSuccessRepositoryMock(cartRepository)

	err := cartService.ClearCart(context.Background(), domain.NewID())
	t.Assert().Nil(err)
}

func (s *CartClearSuite) TestClearCart_Failure(t provider.T) {
	t.Parallel()
	t.Title("ClearCart failure")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartClearFailureRepositoryMock(cartRepository)

	err := cartService.ClearCart(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestCartClearSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "ClearCart", new(CartClearSuite))
}

// CartCreateItem Suite
type CartCreateItemSuite struct {
	suite.Suite
}

func CartCreateItemSuccessRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("CreateCartItem", context.Background(), mock.Anything).
		Return(NewCartItemBuilder().Build(), nil)
}

func CartCreateItemFailureRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("CreateCartItem", context.Background(), mock.Anything).
		Return(domain.CartItem{}, domain.ErrQuantityItems)
}

func (s *CartCreateItemSuite) TestCreateCartItem_Success(t provider.T) {
	t.Parallel()
	t.Title("CreateCartItem success")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartCreateItemSuccessRepositoryMock(cartRepository)

	param := port.CreateCartItemParam{
		CartID:    domain.NewID(),
		ProductID: domain.NewID(),
		Quantity:  1,
	}

	_, err := cartService.CreateCartItem(context.Background(), param)
	t.Assert().Nil(err)
}

func (s *CartCreateItemSuite) TestCreateCartItem_Failure(t provider.T) {
	t.Parallel()
	t.Title("CreateCartItem failure")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartCreateItemFailureRepositoryMock(cartRepository)

	param := port.CreateCartItemParam{
		CartID:    domain.NewID(),
		ProductID: domain.NewID(),
		Quantity:  0, // Неверное количество
	}

	_, err := cartService.CreateCartItem(context.Background(), param)
	t.Assert().ErrorIs(err, domain.ErrQuantityItems)
}

func TestCartCreateItemSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "CreateCartItem", new(CartCreateItemSuite))
}

// CartUpdateItem Suite
type CartUpdateItemSuite struct {
	suite.Suite
}

func CartUpdateItemSuccessRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("GetCartItemByID", context.Background(), mock.Anything).
		Return(NewCartItemBuilder().WithQuantity(1).Build(), nil)

	repository.
		On("UpdateCartItem", context.Background(), mock.Anything).
		Return(NewCartItemBuilder().Build(), nil)
}

func CartUpdateItemFailureRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("GetCartItemByID", context.Background(), mock.Anything).
		Return(domain.CartItem{}, domain.ErrNotExist)
}

func (s *CartUpdateItemSuite) TestUpdateCartItem_Success(t provider.T) {
	t.Parallel()
	t.Title("UpdateCartItem success")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartUpdateItemSuccessRepositoryMock(cartRepository)

	param := port.UpdateCartItemParam{
		Quantity: null.IntFrom(2),
	}

	_, err := cartService.UpdateCartItem(context.Background(), domain.NewID(), param)
	t.Assert().Nil(err)
}

func (s *CartUpdateItemSuite) TestUpdateCartItem_Failure(t provider.T) {
	t.Parallel()
	t.Title("UpdateCartItem failure")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartUpdateItemFailureRepositoryMock(cartRepository)

	param := port.UpdateCartItemParam{
		Quantity: null.IntFrom(0), // Неверное количество
	}

	_, err := cartService.UpdateCartItem(context.Background(), domain.NewID(), param)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestCartUpdateItemSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "UpdateCartItem", new(CartUpdateItemSuite))
}

// CartDeleteItem Suite
type CartDeleteItemSuite struct {
	suite.Suite
}

func CartDeleteItemSuccessRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("DeleteCartItem", context.Background(), mock.Anything).
		Return(nil)
}

func CartDeleteItemFailureRepositoryMock(repository *mocks.CartRepository) {
	repository.
		On("DeleteCartItem", context.Background(), mock.Anything).
		Return(domain.ErrNotExist)
}

func (s *CartDeleteItemSuite) TestDeleteCartItem_Success(t provider.T) {
	t.Parallel()
	t.Title("DeleteCartItem success")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartDeleteItemSuccessRepositoryMock(cartRepository)

	err := cartService.DeleteCartItem(context.Background(), domain.NewID())
	t.Assert().Nil(err)
}

func (s *CartDeleteItemSuite) TestDeleteCartItem_Failure(t provider.T) {
	t.Parallel()
	t.Title("DeleteCartItem failure")
	cartRepository := mocks.NewCartRepository(t)
	shopRepository := mocks.NewShopRepository(t)
	productRepository := mocks.NewProductRepository(t)
	cartService := service.NewCartService(cartRepository, shopRepository, productRepository)

	CartDeleteItemFailureRepositoryMock(cartRepository)

	err := cartService.DeleteCartItem(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestCartDeleteItemSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "DeleteCartItem", new(CartDeleteItemSuite))
}
