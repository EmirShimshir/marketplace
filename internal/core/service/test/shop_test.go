package test

import (
	"context"
	repomocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	storagemocks "github.com/EmirShimshir/marketplace/internal/adapter/storage/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"testing"
)

// ShopSuite is the base suite for Shop tests.
type ShopSuite struct {
	suite.Suite
}

// GetShopByID Suite
type ShopGetByIDSuite struct {
	ShopSuite
}

func ShopGetByIDSuccessRepositoryMock(repository *repomocks.ShopRepository, shopID domain.ID) {
	repository.
		On("GetShopByID", context.Background(), shopID).
		Return(NewShopBuilder().WithID(shopID).Build(), nil)
}

func (s *ShopGetByIDSuite) TestGetShopByID_Success(t provider.T) {
	t.Title("Get shop by ID success")
	shopID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetByIDSuccessRepositoryMock(shopRepository, shopID)
	shop, err := shopService.GetShopByID(context.Background(), shopID)
	t.Assert().Nil(err)
	t.Assert().Equal(shopID, shop.ID)
}

func ShopGetByIDFailureRepositoryMock(repository *repomocks.ShopRepository, shopID domain.ID) {
	repository.
		On("GetShopByID", context.Background(), shopID).
		Return(domain.Shop{}, domain.ErrNotExist)
}

func (s *ShopGetByIDSuite) TestGetShopByID_Failure(t provider.T) {
	t.Title("Get shop by ID failure")
	shopID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetByIDFailureRepositoryMock(shopRepository, shopID)
	_, err := shopService.GetShopByID(context.Background(), shopID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestShopGetByIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetByID", new(ShopGetByIDSuite))
}

// GetShopBySellerID Suite
type ShopGetBySellerIDSuite struct {
	ShopSuite
}

func ShopGetBySellerIDSuccessRepositoryMock(repository *repomocks.ShopRepository, sellerID domain.ID) {
	repository.
		On("GetShopBySellerID", context.Background(), sellerID).
		Return([]domain.Shop{
			NewShopBuilder().WithSellerID(sellerID).Build(),
		}, nil)
}

func (s *ShopGetBySellerIDSuite) TestGetShopBySellerID_Success(t provider.T) {
	t.Title("Get shop by Seller ID success")
	sellerID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetBySellerIDSuccessRepositoryMock(shopRepository, sellerID)
	shops, err := shopService.GetShopBySellerID(context.Background(), sellerID)
	t.Assert().Nil(err)
	t.Assert().Equal(1, len(shops))
	t.Assert().Equal(sellerID, shops[0].SellerID)
}

func ShopGetBySellerIDFailureRepositoryMock(repository *repomocks.ShopRepository, sellerID domain.ID) {
	repository.
		On("GetShopBySellerID", context.Background(), sellerID).
		Return(nil, domain.ErrNotExist)
}

func (s *ShopGetBySellerIDSuite) TestGetShopBySellerID_Failure(t provider.T) {
	t.Title("Get shop by Seller ID failure")
	sellerID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetBySellerIDFailureRepositoryMock(shopRepository, sellerID)
	_, err := shopService.GetShopBySellerID(context.Background(), sellerID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestShopGetBySellerIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetBySellerID", new(ShopGetBySellerIDSuite))
}

// CreateShop Suite
type ShopCreateSuite struct {
	ShopSuite
}

func ShopCreateSuccessRepositoryMock(repository *repomocks.ShopRepository, sellerID domain.ID, param port.CreateShopParam) {
	repository.
		On("CreateShop", context.Background(), mock.Anything).
		Return(NewShopBuilder().WithSellerID(sellerID).WithName(param.Name).Build(), nil)
}

func (s *ShopCreateSuite) TestCreateShop_Success(t provider.T) {
	t.Title("Create shop success")
	sellerID := domain.NewID()
	param := NewCreateShopParamBuilder().Build()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopCreateSuccessRepositoryMock(shopRepository, sellerID, param)
	shop, err := shopService.CreateShop(context.Background(), sellerID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(sellerID, shop.SellerID)
	t.Assert().Equal(param.Name, shop.Name)
}

func ShopCreateFailureRepositoryMock(repository *repomocks.ShopRepository, sellerID domain.ID, param port.CreateShopParam) {
	repository.
		On("CreateShop", context.Background(), mock.Anything).
		Return(domain.Shop{}, domain.ErrName)
}

func (s *ShopCreateSuite) TestCreateShop_Failure(t provider.T) {
	t.Title("Create shop failure due to missing name")
	sellerID := domain.NewID()
	param := NewCreateShopParamBuilder().Build()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopCreateFailureRepositoryMock(shopRepository, sellerID, param)
	_, err := shopService.CreateShop(context.Background(), sellerID, param)
	t.Assert().ErrorIs(err, domain.ErrName)
}

func TestShopCreateSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "CreateShop", new(ShopCreateSuite))
}

// GetShopItems Suite
type ShopGetItemsSuite struct {
	ShopSuite
}

func ShopGetItemsSuccessRepositoryMock(repository *repomocks.ShopRepository, limit, offset int64) {
	repository.
		On("GetShopItems", context.Background(), limit, offset).
		Return([]domain.ShopItem{NewShopItemBuilder().Build()}, nil)
}

func (s *ShopGetItemsSuite) TestGetShopItems_Success(t provider.T) {
	t.Title("Get shop items success")
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetItemsSuccessRepositoryMock(shopRepository, 10, 0)
	items, err := shopService.GetShopItems(context.Background(), 10, 0)
	t.Assert().Nil(err)
	t.Assert().Equal(1, len(items))
}

func ShopGetItemsFailureRepositoryMock(repository *repomocks.ShopRepository, limit, offset int64) {
	repository.
		On("GetShopItems", context.Background(), limit, offset).
		Return(nil, domain.ErrNotExist)
}

func (s *ShopGetItemsSuite) TestGetShopItems_Failure(t provider.T) {
	t.Title("Get shop items failure due to database error")
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetItemsFailureRepositoryMock(shopRepository, 10, 0)
	_, err := shopService.GetShopItems(context.Background(), 10, 0)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestShopGetItemsSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetShopItems", new(ShopGetItemsSuite))
}

// GetShopItemByProductID Suite
type ShopGetItemByProductIDSuite struct {
	ShopSuite
}

func ShopGetItemByProductIDSuccessRepositoryMock(repository *repomocks.ShopRepository, productID domain.ID) {
	repository.
		On("GetShopItemByProductID", context.Background(), productID).
		Return(NewShopItemBuilder().WithProductID(productID).Build(), nil)
}

func (s *ShopGetItemByProductIDSuite) TestGetShopItemByProductID_Success(t provider.T) {
	t.Title("Get shop item by Product ID success")
	productID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetItemByProductIDSuccessRepositoryMock(shopRepository, productID)
	item, err := shopService.GetShopItemByProductID(context.Background(), productID)
	t.Assert().Nil(err)
	t.Assert().Equal(productID, item.ProductID)
}

func ShopGetItemByProductIDFailureRepositoryMock(repository *repomocks.ShopRepository, productID domain.ID) {
	repository.
		On("GetShopItemByProductID", context.Background(), productID).
		Return(domain.ShopItem{}, domain.ErrNotExist)
}

func (s *ShopGetItemByProductIDSuite) TestGetShopItemByProductID_Failure(t provider.T) {
	t.Title("Get shop item by Product ID failure due to non-existing product")
	productID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopGetItemByProductIDFailureRepositoryMock(shopRepository, productID)
	_, err := shopService.GetShopItemByProductID(context.Background(), productID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestShopGetItemByProductIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetShopItemByProductID", new(ShopGetItemByProductIDSuite))
}

// CreateShopItem Suite
type ShopCreateItemSuite struct {
	ShopSuite
}

func ShopCreateItemSuccessRepositoryMock(repository *repomocks.ShopRepository, shopItem domain.ShopItem) {
	repository.
		On("CreateShopItem", context.Background(), mock.Anything, mock.Anything).
		Return(shopItem, nil)
}

func ShopCreateItemSuccessStorageMock(storage *storagemocks.ObjectStorage) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url(""), nil)
}

func (s *ShopCreateItemSuite) TestCreateShopItem_Success(t provider.T) {
	t.Title("Create shop item success")
	param := NewCreateShopItemParamBuilder(domain.NewID()).Build()
	shopItem := NewShopItemBuilder().WithShopID(param.ShopID).Build()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopCreateItemSuccessRepositoryMock(shopRepository, shopItem)
	ShopCreateItemSuccessStorageMock(storage)
	item, err := shopService.CreateShopItem(context.Background(), param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.ShopID, item.ShopID)
}

func ShopCreateItemFailureRepositoryMock(repository *repomocks.ShopRepository) {
	repository.
		On("CreateShopItem", context.Background(), mock.Anything, mock.Anything).
		Return(domain.ShopItem{}, nil)
}

func ShopCreateItemFailureStorageMock(storage *storagemocks.ObjectStorage) {
	storage.
		On("SaveFile", context.Background(), mock.Anything).
		Return(domain.Url(""), nil)
}

func (s *ShopCreateItemSuite) TestCreateShopItem_Failure(t provider.T) {
	t.Title("Create shop item failure due to invalid input")
	param := NewCreateShopItemParamBuilder(domain.NewID()).Build()
	param.ProductParam.Name = ""
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopCreateItemFailureRepositoryMock(shopRepository)
	ShopCreateItemFailureStorageMock(storage)
	_, err := shopService.CreateShopItem(context.Background(), param)
	t.Assert().ErrorIs(err, domain.ErrName)
}

func TestShopCreateItemSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "CreateShopItem", new(ShopCreateItemSuite))
}

// UpdateShopItem Suite
type ShopUpdateItemSuite struct {
	ShopSuite
}

func ShopUpdateItemSuccessRepositoryMock(repository *repomocks.ShopRepository, shopItem domain.ShopItem) {
	repository.
		On("GetShopItemByID", context.Background(), shopItem.ID).
		Return(shopItem, nil)
	repository.
		On("UpdateShopItem", context.Background(), shopItem).
		Return(shopItem, nil)
}

func (s *ShopUpdateItemSuite) TestUpdateShopItem_Success(t provider.T) {
	t.Title("Update shop item success")
	shopItemID := domain.NewID()
	param := NewUpdateShopItemParamBuilder().WithQuantity(5).Build()
	shopItem := NewShopItemBuilder().WithID(shopItemID).WithQuantity(10).Build()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	shopItem.Quantity = param.Quantity.Int64
	ShopUpdateItemSuccessRepositoryMock(shopRepository, shopItem)
	item, err := shopService.UpdateShopItem(context.Background(), shopItemID, param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Quantity.Int64, item.Quantity)
}

func ShopUpdateItemFailureRepositoryMock(repository *repomocks.ShopRepository, shopItem domain.ShopItem) {
	repository.
		On("GetShopItemByID", context.Background(), shopItem.ID).
		Return(shopItem, nil)
}

func (s *ShopUpdateItemSuite) TestUpdateShopItem_Failure(t provider.T) {
	t.Title("Update shop item failure due to invalid input")
	shopItemID := domain.NewID()
	param := NewUpdateShopItemParamBuilder().WithQuantity(-1).Build()
	shopItem := NewShopItemBuilder().WithID(shopItemID).WithQuantity(10).Build()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	shopItem.Quantity = param.Quantity.Int64
	ShopUpdateItemFailureRepositoryMock(shopRepository, shopItem)
	_, err := shopService.UpdateShopItem(context.Background(), shopItemID, param)
	t.Assert().ErrorIs(err, domain.ErrQuantityItems)
}

func TestShopUpdateItemSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "UpdateShopItem", new(ShopUpdateItemSuite))
}

// DeleteShopItem Suite
type ShopDeleteItemSuite struct {
	ShopSuite
}

func ShopDeleteItemSuccessRepositoryMock(repository *repomocks.ShopRepository, shopItemID domain.ID) {
	repository.
		On("DeleteShopItem", context.Background(), shopItemID).
		Return(nil)
}

func (s *ShopDeleteItemSuite) TestDeleteShopItem_Success(t provider.T) {
	t.Title("Delete shop item success")
	shopItemID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopDeleteItemSuccessRepositoryMock(shopRepository, shopItemID)
	err := shopService.DeleteShopItem(context.Background(), shopItemID)
	t.Assert().Nil(err)
}

func ShopDeleteItemFailureRepositoryMock(repository *repomocks.ShopRepository, shopItemID domain.ID) {
	repository.
		On("DeleteShopItem", context.Background(), shopItemID).
		Return(domain.ErrNotExist)
}

func (s *ShopDeleteItemSuite) TestDeleteShopItem_Failure(t provider.T) {
	t.Title("Delete shop item failure due to non-existing item")
	shopItemID := domain.NewID()
	shopRepository := repomocks.NewShopRepository(t)
	storage := storagemocks.NewObjectStorage(t)
	shopService := service.NewShopService(shopRepository, storage)
	ShopDeleteItemFailureRepositoryMock(shopRepository, shopItemID)
	err := shopService.DeleteShopItem(context.Background(), shopItemID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestShopDeleteItemSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "DeleteShopItem", new(ShopDeleteItemSuite))
}
