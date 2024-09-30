package test

import (
	"context"
	payment "github.com/EmirShimshir/marketplace/internal/adapter/payment/mocks"
	"github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"net/url"
	"testing"

	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
)

// PaymentSuite is a base suite for payment-related tests
type PaymentSuite struct {
	suite.Suite
}

type PaymentGetOrderPaymentUrlSuite struct {
	PaymentSuite
}

// PaymentGetOrderPaymentUrlSuccessMock mocks the repository and gateway for a successful GetOrderPaymentUrl call
func PaymentGetOrderPaymentUrlSuccessMock(orderRepo *mocks.OrderRepository, gateway *payment.PaymentGateway, orderID domain.ID) {
	orderRepo.
		On("GetOrderCustomerByID", context.Background(), orderID).
		Return(domain.OrderCustomer{
			ID:         orderID,
			Payed:      false,
			TotalPrice: 1000,
		}, nil)

	gateway.
		On("GetPaymentUrl", context.Background(), mock.AnythingOfType("domain.PaymentPayload")).
		Return(url.URL{
			Scheme: "https",
			Host:   "payment.com",
			Path:   "/pay",
		}, nil)
}

// TestGetOrderPaymentUrl_Success tests a successful GetOrderPaymentUrl call
func (s *PaymentGetOrderPaymentUrlSuite) TestGetOrderPaymentUrl_Success(t provider.T) {
	t.Parallel()
	t.Title("GetOrderPaymentUrl success")

	orderID := domain.NewID()
	orderRepo := mocks.NewOrderRepository(t)
	gateway := payment.NewPaymentGateway(t)

	paymentService := service.NewPaymentService(gateway, orderRepo)

	PaymentGetOrderPaymentUrlSuccessMock(orderRepo, gateway, orderID)

	u, err := paymentService.GetOrderPaymentUrl(context.Background(), orderID)
	t.Assert().Nil(err)
	t.Assert().Equal("https://payment.com/pay", u.String())
}

func PaymentGetOrderPaymentUrlAlreadyPayedMock(orderRepo *mocks.OrderRepository, orderID domain.ID) {
	orderRepo.
		On("GetOrderCustomerByID", context.Background(), orderID).
		Return(domain.OrderCustomer{
			ID:    orderID,
			Payed: true,
		}, nil)
}

// TestGetOrderPaymentUrl_AlreadyPayed tests a scenario where the order is already paid
func (s *PaymentGetOrderPaymentUrlSuite) TestGetOrderPaymentUrl_AlreadyPayed(t provider.T) {
	t.Parallel()
	t.Title("GetOrderPaymentUrl already paid")

	orderID := domain.NewID()
	orderRepo := mocks.NewOrderRepository(t)
	gateway := payment.NewPaymentGateway(t)

	paymentService := service.NewPaymentService(gateway, orderRepo)

	PaymentGetOrderPaymentUrlAlreadyPayedMock(orderRepo, orderID)

	_, err := paymentService.GetOrderPaymentUrl(context.Background(), orderID)
	t.Assert().ErrorIs(err, domain.ErrOrderAlreadyPayed)
}

func TestPaymentGetOrderPaymentUrlSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetOrderPaymentUrl", new(PaymentGetOrderPaymentUrlSuite))
}

type PaymentProcessOrderPaymentSuite struct {
	PaymentSuite
}

// PaymentProcessOrderPaymentSuccessMock mocks the repository and gateway for a successful ProcessOrderPayment call
func PaymentProcessOrderPaymentSuccessMock(orderRepo *mocks.OrderRepository, gateway *payment.PaymentGateway, orderID domain.ID) {
	gateway.
		On("ProcessPayment", context.Background(), mock.AnythingOfType("string")).
		Return(domain.PaymentPayload{
			OrderID: orderID,
			PaySum:  1000,
		}, nil)

	orderRepo.
		On("GetOrderCustomerByID", context.Background(), orderID).
		Return(domain.OrderCustomer{
			ID:         orderID,
			TotalPrice: 1000,
		}, nil)

	orderRepo.
		On("UpdatePaymentStatus", context.Background(), orderID).
		Return(nil)
}

// TestProcessOrderPayment_Success tests a successful ProcessOrderPayment call
func (s *PaymentProcessOrderPaymentSuite) TestProcessOrderPayment_Success(t provider.T) {
	t.Parallel()
	t.Title("ProcessOrderPayment success")

	paymentKey := "valid-key"
	orderID := domain.NewID()
	orderRepo := mocks.NewOrderRepository(t)
	gateway := payment.NewPaymentGateway(t)

	paymentService := service.NewPaymentService(gateway, orderRepo)

	PaymentProcessOrderPaymentSuccessMock(orderRepo, gateway, orderID)

	err := paymentService.ProcessOrderPayment(context.Background(), paymentKey)
	t.Assert().Nil(err)
}

func PaymentProcessOrderPaymentInvalidSumMock(orderRepo *mocks.OrderRepository, gateway *payment.PaymentGateway, orderID domain.ID) {
	gateway.
		On("ProcessPayment", context.Background(), mock.AnythingOfType("string")).
		Return(domain.PaymentPayload{
			OrderID: orderID,
			PaySum:  2000,
		}, nil)

	orderRepo.
		On("GetOrderCustomerByID", context.Background(), orderID).
		Return(domain.OrderCustomer{
			ID:         orderID,
			TotalPrice: 1000,
		}, nil)
}

// TestProcessOrderPayment_InvalidSum tests a scenario where the payment sum is invalid
func (s *PaymentProcessOrderPaymentSuite) TestProcessOrderPayment_InvalidSum(t provider.T) {
	t.Parallel()
	t.Title("ProcessOrderPayment invalid sum")

	paymentKey := "valid-key"
	orderID := domain.NewID()
	orderRepo := mocks.NewOrderRepository(t)
	gateway := payment.NewPaymentGateway(t)

	paymentService := service.NewPaymentService(gateway, orderRepo)

	PaymentProcessOrderPaymentInvalidSumMock(orderRepo, gateway, orderID)

	err := paymentService.ProcessOrderPayment(context.Background(), paymentKey)
	t.Assert().ErrorIs(err, domain.ErrInvalidPaymentSum)
}

func TestPaymentProcessOrderPaymentSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "ProcessOrderPayment", new(PaymentProcessOrderPaymentSuite))
}
