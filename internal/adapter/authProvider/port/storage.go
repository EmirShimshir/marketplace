package port

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"time"
)

type AuthSession struct {
	RefreshToken string
	RefreshExp   int64
	Fingerprint  string
	Payload      domain.AuthPayload
}

type ISessionStorage interface {
	Get(refreshToken string) (AuthSession, error)
	Put(refreshToken string, session AuthSession, expireTime time.Duration) error
	Delete(refreshToken string) error
}
