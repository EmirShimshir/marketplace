package dto

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
)

type SignUpDTO struct {
	Name     string
	Surname  string
	Phone    null.String
	Role     domain.UserRole
	Email    string
	Password string
}

func InputSignUpDTO(d *SignUpDTO) error {
	fmt.Print("Введите имя: ")
	fmt.Scanln(&d.Name)

	fmt.Print("Введите фамилию: ")
	fmt.Scanln(&d.Surname)

	fmt.Print("Введите номер телефона: ")
	var input string
	fmt.Scanln(&input)
	if input != "" {
		err := checkPhone(input)
		if err != nil {
			return err
		}
		d.Phone.String = input
		d.Phone.Valid = true
	}

	err := InputUserRole(&d.Role)
	if err != nil {
		return err
	}

	err = InputEmail(&d.Email)
	if err != nil {
		return err
	}

	fmt.Print("Введите пароль: ")
	fmt.Scanln(&d.Password)
	d.Phone.Valid = true
	return nil
}

type SignInDTO struct {
	Email    string
	Password string
}

func InputSignInDTO(d *SignInDTO) error {
	err := InputEmail(&d.Email)
	if err != nil {
		return err
	}

	fmt.Print("Введите пароль: ")
	fmt.Scanln(&d.Password)
	return nil
}
