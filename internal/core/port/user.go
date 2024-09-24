package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
)

type CreateUserParam struct {
	Name     string
	Surname  string
	Phone    null.String
	Email    string
	Password string
	Role     domain.UserRole
}

type UpdateUserParam struct {
	Name    null.String
	Surname null.String
	Phone   null.String
}

type IUserService interface {
	Get(ctx context.Context, limit, offset int64) ([]domain.User, error)
	GetByID(ctx context.Context, userID domain.ID) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, param CreateUserParam) (domain.User, error)
	Update(ctx context.Context, userID domain.ID, param UpdateUserParam) (domain.User, error)
}
