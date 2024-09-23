package dto

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type UpdateUserDTO struct {
	Name    *string `json:"name" binding:"omitempty"`
	Surname *string `json:"surname" binding:"omitempty"`
	Phone   *string `json:"phone" binding:"omitempty"`
}

type UserDTO struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Role    string `json:"role"`
}

func NewUserDTO(user domain.User) *UserDTO {
	var role string
	switch user.Role {
	case domain.UserCustomer:
		role = "Покупатель"
	case domain.UserSeller:
		role = "Продавец"
	case domain.UserModerator:
		role = "Модератор"
	}
	return &UserDTO{
		ID:      user.ID.String(),
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Phone:   user.Phone.String,
		Role:    role,
	}
}
