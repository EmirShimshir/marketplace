package dto

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

type WithdrawDTO struct {
	ID      string
	ShopID  string
	Comment string
	Sum     int64
	Status  string
}

func NewWithdrawDTO(withdraw domain.Withdraw) *WithdrawDTO {
	var status string
	switch withdraw.Status {
	case domain.WithdrawStatusStart:
		status = "В обработке"
	case domain.WithdrawStatusReady:
		status = "Принят"
	case domain.WithdrawStatusDone:
		status = "Готов"
	}

	return &WithdrawDTO{
		ID:      string(withdraw.ID),
		ShopID:  string(withdraw.ShopID),
		Comment: withdraw.Comment,
		Sum:     withdraw.Sum,
		Status:  status,
	}
}

func (w *WithdrawDTO) Print() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("ID заявки: %s\n", w.ID)
	fmt.Printf("ID магазина: %s\n", w.ShopID)
	fmt.Printf("Комментарий: %s\n", w.Comment)
	fmt.Printf("Сумма: %d\n", w.Sum)
	fmt.Printf("Cтатус: %s\n", w.Status)
	fmt.Println("------------------------------------------------")
}
