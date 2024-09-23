package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
)

type CreateWithdrawParam struct {
	ShopID  domain.ID
	Comment string
	Sum     int64
}

type UpdateWithdrawParam struct {
	Comment null.String
	Sum     null.Int
	Status  *domain.WithdrawStatus
}

type IWithdrawService interface {
	Get(ctx context.Context, limit, offset int64) ([]domain.Withdraw, error)
	GetByID(ctx context.Context, WithdrawID domain.ID) (domain.Withdraw, error)
	GetByShopID(ctx context.Context, shopID domain.ID) ([]domain.Withdraw, error)
	Create(ctx context.Context, param CreateWithdrawParam) (domain.Withdraw, error)
	Update(ctx context.Context, WithdrawID domain.ID, param UpdateWithdrawParam) (domain.Withdraw, error)
	Delete(ctx context.Context, WithdrawID domain.ID) error
}
