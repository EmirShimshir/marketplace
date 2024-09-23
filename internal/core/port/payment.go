package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"net/url"
)

type IPayment interface {
	GetOrderPaymentUrl(ctx context.Context, orderID domain.ID) (url.URL, error)
	ProcessOrderPayment(ctx context.Context, key string) error
}
