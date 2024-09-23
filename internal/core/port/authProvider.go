package port

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type IAuthProvider interface {
	CreateJWTSession(payload domain.AuthPayload, fingerprint string) (domain.AuthDetails, error)
	RefreshJWTSession(refreshToken domain.Token, fingerprint string) (domain.AuthDetails, error)
	DeleteJWTSession(refreshToken domain.Token) error
	VerifyJWTToken(accessToken domain.Token) (domain.AuthPayload, error)
	GenPasswordHash(password string) string
}
