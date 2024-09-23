package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
)

type SignInParam struct {
	Email       string
	Password    string
	Fingerprint string
}

type SignUpParam struct {
	Name     string
	Surname  string
	Email    string
	Password string
	Phone    null.String
	Role     domain.UserRole
}

type IAuthService interface {
	SignIn(ctx context.Context, param SignInParam) (domain.AuthDetails, error)
	SignUp(ctx context.Context, param SignUpParam) error
	LogOut(ctx context.Context, refreshToken domain.Token) error
	Refresh(ctx context.Context, refreshToken domain.Token, fingerprint string) (domain.AuthDetails, error)
	Verify(ctx context.Context, accessToken domain.Token) error
	Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error)
}
