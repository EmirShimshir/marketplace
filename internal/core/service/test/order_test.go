package test

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"testing"

	"github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
)

// OrderGetByCustomerID Suite
type OrderGetByCustomerIDSuite struct {
	suite.Suite
}

func OrderGetByCustomerIDSuccessRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("GetOrderCustomerByCustomerID", context.Background(), mock.Anything).
		Return([]domain.OrderCustomer{NewOrderCustomerBuilder().Build()}, nil)
}

func OrderGetByCustomerIDFailureRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("GetOrderCustomerByCustomerID", context.Background(), mock.Anything).
		Return(nil, domain.ErrNotExist)
}

func (s *OrderGetByCustomerIDSuite) TestGetOrderCustomerByCustomerID_Success(t provider.T) {
	t.Parallel()
	t.Title("GetOrderCustomerByCustomerID success")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)
	OrderGetByCustomerIDSuccessRepositoryMock(orderRepository)
	_, err := orderService.GetOrderCustomerByCustomerID(context.Background(), domain.NewID())
	t.Assert().Nil(err)
}

func (s *OrderGetByCustomerIDSuite) TestGetOrderCustomerByCustomerID_Failure(t provider.T) {
	t.Parallel()
	t.Title("GetOrderCustomerByCustomerID failure")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)
	OrderGetByCustomerIDFailureRepositoryMock(orderRepository)
	_, err := orderService.GetOrderCustomerByCustomerID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestOrderGetByCustomerIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetOrderCustomerByCustomerID", new(OrderGetByCustomerIDSuite))
}

// OrderCreateCustomer Suite
type OrderCreateCustomerSuite struct {
	suite.Suite
}

func OrderCreateCustomerSuccessRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("CreateOrderCustomer", context.Background(), mock.Anything).
		Return(NewOrderCustomerBuilder().Build(), nil)
}

func OrderCreateCustomerFailureMock(repository *mocks.OrderRepository) {
	repository.
		On("CreateOrderCustomer", context.Background(), mock.Anything).
		Return(domain.OrderCustomer{}, domain.ErrEmptyCart)
}

func (s *OrderCreateCustomerSuite) TestCreateOrderCustomer_Success(t provider.T) {
	t.Parallel()
	t.Title("CreateOrderCustomer success")
	orderRepository := mocks.NewOrderRepository(t)
	userRepository := mocks.NewUserRepository(t)
	cartRepository := mocks.NewCartRepository(t)
	orderService := service.NewOrderService(orderRepository, userRepository, cartRepository, nil)

	OrderCreateCustomerSuccessRepositoryMock(orderRepository)

	param := port.CreateOrderCustomerParam{
		CustomerID: domain.NewID(),
		Address:    "123 Street",
	}

	_, err := orderService.CreateOrderCustomer(context.Background(), param)
	t.Assert().Nil(err)
}

func (s *OrderCreateCustomerSuite) TestCreateOrderCustomer_Failure(t provider.T) {
	t.Parallel()
	t.Title("CreateOrderCustomer failure")
	orderRepository := mocks.NewOrderRepository(t)
	userRepository := mocks.NewUserRepository(t)
	cartRepository := mocks.NewCartRepository(t)
	orderService := service.NewOrderService(orderRepository, userRepository, cartRepository, nil)

	OrderCreateCustomerFailureMock(orderRepository)

	param := port.CreateOrderCustomerParam{
		CustomerID: domain.NewID(),
		Address:    "",
	}

	_, err := orderService.CreateOrderCustomer(context.Background(), param)
	t.Assert().ErrorIs(err, domain.ErrAddress)
}

func TestOrderCreateCustomerSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "CreateOrderCustomer", new(OrderCreateCustomerSuite))
}

// OrderGetByID Suite
type OrderGetByIDSuite struct {
	suite.Suite
}

func OrderGetByIDSuccessRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("GetOrderShopByID", context.Background(), mock.Anything).
		Return(NewOrderShopBuilder().Build(), nil)
}

func OrderGetByIDFailureRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("GetOrderShopByID", context.Background(), mock.Anything).
		Return(domain.OrderShop{}, domain.ErrNotExist)
}

func (s *OrderGetByIDSuite) TestGetOrderShopByID_Success(t provider.T) {
	t.Parallel()
	t.Title("GetOrderShopByID success")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)

	OrderGetByIDSuccessRepositoryMock(orderRepository)

	_, err := orderService.GetOrderShopByID(context.Background(), domain.NewID())
	t.Assert().Nil(err)
}

func (s *OrderGetByIDSuite) TestGetOrderShopByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("GetOrderShopByID failure")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)

	OrderGetByIDFailureRepositoryMock(orderRepository)

	_, err := orderService.GetOrderShopByID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestOrderGetByIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetOrderShopByID", new(OrderGetByIDSuite))
}

// OrderGetByShopID Suite
type OrderGetByShopIDSuite struct {
	suite.Suite
}

func OrderGetByShopIDSuccessRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("GetOrderShopByShopID", context.Background(), mock.Anything).
		Return([]domain.OrderShop{NewOrderShopBuilder().Build()}, nil)
}

func OrderGetByShopIDFailureRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("GetOrderShopByShopID", context.Background(), mock.Anything).
		Return(nil, domain.ErrNotExist)
}

func (s *OrderGetByShopIDSuite) TestGetOrderShopByShopID_Success(t provider.T) {
	t.Parallel()
	t.Title("GetOrderShopByShopID success")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)

	OrderGetByShopIDSuccessRepositoryMock(orderRepository)

	_, err := orderService.GetOrderShopByShopID(context.Background(), domain.NewID())
	t.Assert().Nil(err)
}

func (s *OrderGetByShopIDSuite) TestGetOrderShopByShopID_Failure(t provider.T) {
	t.Parallel()
	t.Title("GetOrderShopByShopID failure")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)

	OrderGetByShopIDFailureRepositoryMock(orderRepository)

	_, err := orderService.GetOrderShopByShopID(context.Background(), domain.NewID())
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestOrderGetByShopIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetOrderShopByShopID", new(OrderGetByShopIDSuite))
}

// OrderUpdateShop Suite
type OrderUpdateShopSuite struct {
	suite.Suite
}

func OrderUpdateShopSuccessRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("UpdateOrderShop", context.Background(), mock.Anything).
		Return(NewOrderShopBuilder().Build(), nil)
}

func OrderUpdateShopFailureRepositoryMock(repository *mocks.OrderRepository) {
	repository.
		On("UpdateOrderShop", context.Background(), mock.Anything).
		Return(domain.OrderShop{}, domain.ErrNotExist)
}

func (s *OrderUpdateShopSuite) TestUpdateOrderShop_Success(t provider.T) {
	t.Parallel()
	t.Title("UpdateOrderShop success")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)

	OrderUpdateShopSuccessRepositoryMock(orderRepository)

	status := domain.OrderShopStatusReady
	param := port.UpdateOrderShopParam{
		Status: &status,
	}

	_, err := orderService.UpdateOrderShop(context.Background(), domain.NewID(), param)
	t.Assert().Nil(err)
}

func (s *OrderUpdateShopSuite) TestUpdateOrderShop_Failure(t provider.T) {
	t.Parallel()
	t.Title("UpdateOrderShop failure")
	orderRepository := mocks.NewOrderRepository(t)
	orderService := service.NewOrderService(orderRepository, nil, nil, nil)

	OrderUpdateShopFailureRepositoryMock(orderRepository)

	param := port.UpdateOrderShopParam{
		Status: nil,
	}

	_, err := orderService.UpdateOrderShop(context.Background(), domain.NewID(), param)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestOrderUpdateShopSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "UpdateOrderShop", new(OrderUpdateShopSuite))
}
