package paymentMock

import (
	"context"
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"net/url"
	"strconv"
	"strings"
)

type MockGateway struct {
}

func NewtMockGateway() *MockGateway {
	return &MockGateway{}
}

func (g *MockGateway) GetPaymentUrl(ctx context.Context, payload domain.PaymentPayload) (url.URL, error) {
	return url.URL{
		Scheme: "https",
		Host:   "MockGateway.ru",
		Path:   fmt.Sprintf("%s&%s", payload.OrderID, strconv.Itoa(int(payload.PaySum))),
	}, nil
}

func (g *MockGateway) ProcessPayment(ctx context.Context, key string) (domain.PaymentPayload, error) {
	res := strings.Split(key, "&")
	if len(res) != 2 {
		return domain.PaymentPayload{}, domain.ErrInvalidPaymentSum
	}
	sum, err := strconv.ParseInt(res[1], 10, 64)
	if err != nil {
		return domain.PaymentPayload{}, err
	}
	return domain.PaymentPayload{
		OrderID: domain.ID(res[0]),
		PaySum:  sum,
	}, nil
}
