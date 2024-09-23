package dto

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/mail"
	"regexp"
	"strings"
)

func InputEmail(email *string) error {
	fmt.Print("Введите почту: ")
	var input string
	fmt.Scanln(&input)
	_, err := mail.ParseAddress(input)
	if err != nil {
		return errors.New("invalid email format")
	}
	*email = input
	return nil
}

func InputID(id *domain.ID, idOwner string) error {
	fmt.Printf("Введите ID %s: ", cases.Title(language.Und, cases.NoLower).String(idOwner))
	var input string
	fmt.Scanln(&input)
	_, err := uuid.Parse(input)
	if err != nil {
		return errors.New("invalid uuid format")
	}
	*id = domain.ID(input)
	return nil
}

func InputUserRole(role *domain.UserRole) error {
	fmt.Print("Выберете роль (0 - покупатель; 1 - продавец; 2 - модератор): ")
	var input int
	fmt.Scanf("%d", &input)
	if !(0 <= input && input <= 2) {
		return errors.New("invalid role format")
	}

	*role = domain.UserRole(input)
	return nil
}

func InputProductCategory(category *domain.ProductCategory) error {
	fmt.Print("Выберете категорию (0 - Электроника; 1 - Мода; 2 - Дом; 3 - Здоровье; 4 - Спорт; 5 - Книги): ")
	var input int
	fmt.Scanf("%d", &input)
	if !(0 <= input && input <= 5) {
		return errors.New("invalid category format")
	}

	*category = domain.ProductCategory(input)
	return nil
}

func InputOrderShopStatus(status *domain.OrderShopStatus) error {
	fmt.Print("Выберете статус заказа (0 - В обработке; 1 - Принят; 2 - Готов): ")
	var input int
	fmt.Scanf("%d", &input)
	if !(0 <= input && input <= 2) {
		return errors.New("invalid order status format")
	}

	*status = domain.OrderShopStatus(input)
	return nil
}

func InputWithdrawStatus(status *domain.WithdrawStatus) error {
	fmt.Print("Выберете статус заявки (0 - В обработке; 1 - Принят; 2 - Готов): ")
	var input int
	fmt.Scanf("%d", &input)
	if !(0 <= input && input <= 2) {
		return errors.New("invalid withdraw status format")
	}

	*status = domain.WithdrawStatus(input)
	return nil
}

func checkPhone(phone string) error {
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	phone = strings.ReplaceAll(phone, " ", "")
	if re.Find([]byte(phone)) == nil {
		return errors.New("invalid phone number format")
	}
	return nil
}
