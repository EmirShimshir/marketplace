package entity

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/google/uuid"
)

const (
	PgWithdrawStart = "Start"
	PgWithdrawReady = "Ready"
	PgWithdrawDone  = "Done"
)

type PgWithdraw struct {
	ID      uuid.UUID `db:"id"`
	ShopID  uuid.UUID `db:"shop_id"`
	Comment string    `db:"comment"`
	Sum     int64     `db:"sum"`
	Status  string    `db:"status"`
}

func (w *PgWithdraw) ToDomain() domain.Withdraw {
	var withdrawStatus domain.WithdrawStatus
	switch w.Status {
	case PgWithdrawStart:
		withdrawStatus = domain.WithdrawStatusStart
	case PgWithdrawReady:
		withdrawStatus = domain.WithdrawStatusReady
	case PgWithdrawDone:
		withdrawStatus = domain.WithdrawStatusDone
	}

	return domain.Withdraw{
		ID:      domain.ID(w.ID.String()),
		ShopID:  domain.ID(w.ShopID.String()),
		Comment: w.Comment,
		Sum:     w.Sum,
		Status:  withdrawStatus,
	}
}

func NewPgWithdraw(withdraw domain.Withdraw) PgWithdraw {
	id, _ := uuid.Parse(withdraw.ID.String())
	shopID, _ := uuid.Parse(withdraw.ShopID.String())
	var withdrawStatus string
	switch withdraw.Status {
	case domain.WithdrawStatusStart:
		withdrawStatus = PgWithdrawStart
	case domain.WithdrawStatusReady:
		withdrawStatus = PgWithdrawReady
	case domain.WithdrawStatusDone:
		withdrawStatus = PgWithdrawDone
	}

	return PgWithdraw{
		ID:      id,
		ShopID:  shopID,
		Comment: withdraw.Comment,
		Sum:     withdraw.Sum,
		Status:  withdrawStatus,
	}
}
