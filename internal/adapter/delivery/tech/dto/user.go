package dto

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
)

type UpdateUserDTO struct {
	Name    null.String
	Surname null.String
	Phone   null.String
}

func InputUpdateUserDTO(d *UpdateUserDTO) error {
	var name string
	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)
	if name != "" {
		d.Name = null.StringFrom(name)
	}

	var surname string
	fmt.Print("Введите фамилию: ")
	fmt.Scanln(&surname)
	if surname != "" {
		d.Surname = null.StringFrom(surname)
	}

	var phone string
	fmt.Print("Введите номер телефона: ")
	fmt.Scanln(&phone)
	if phone != "" {
		err := checkPhone(phone)
		if err != nil {
			return err
		}
		d.Phone.String = phone
		d.Phone.Valid = true
	}
	return nil
}

type UserDTO struct {
	Name    string
	Surname string
	Email   string
	Phone   string
	Role    string
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
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Phone:   user.Phone.String,
		Role:    role,
	}
}

func (u *UserDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("Имя: %s\n", u.Name)
	fmt.Printf("Фамилия: %s\n", u.Surname)
	fmt.Printf("Почта: %s\n", u.Email)
	fmt.Printf("Телефон: %s\n", u.Phone)
	fmt.Printf("Роль: %s\n", u.Role)
	fmt.Println("------------------------------------------------")
}
