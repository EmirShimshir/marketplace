package dto

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type SignUpDTO struct {
	Name     string          `json:"name" binding:"required"`
	Surname  string          `json:"surname" binding:"required"`
	Phone    *string         `json:"phone" binding:"omitempty"`
	Role     domain.UserRole `json:"role" binding:"omitempty"`
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required"`
}

type SignInDTO struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Fingerprint string `json:"fingerprint" binding:"required"`
}

type RefreshDTO struct {
	Fingerprint string `json:"fingerprint" binding:"required"`
}
