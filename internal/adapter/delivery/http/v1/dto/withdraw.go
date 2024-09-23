package dto

import "github.com/EmirShimshir/marketplace/internal/core/domain"

type WithdrawDTO struct {
	ID      string `json:"id"`
	ShopID  string `json:"shop_id"`
	Comment string `json:"comment"`
	Sum     int64  `json:"sum"`
	Status  string `json:"status"`
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

type CreateWithdrawDTO struct {
	Sum     int64  `json:"sum" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}
