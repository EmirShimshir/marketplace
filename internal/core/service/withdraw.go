package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

type WithdrawService struct {
	repo port.IWithdrawRepository
}

func NewWithdrawService(repo port.IWithdrawRepository) *WithdrawService {
	return &WithdrawService{
		repo: repo,
	}
}

func (w *WithdrawService) Get(ctx context.Context, limit, offset int64) ([]domain.Withdraw, error) {
	withdraws, err := w.repo.Get(ctx, limit, offset)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "WithdrawGet",
		}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{
		"count": len(withdraws),
	}).Info("WithdrawServiceGet OK")
	return withdraws, nil
}

func (w *WithdrawService) GetByID(ctx context.Context, withdrawID domain.ID) (domain.Withdraw, error) {
	withdraw, err := w.repo.GetByID(ctx, withdrawID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "WithdrawGetByID",
		}).Error(err.Error())
		return domain.Withdraw{}, err
	}

	log.WithFields(log.Fields{
		"id": withdraw.ID,
	}).Info("WithdrawServiceGetByID OK")
	return withdraw, nil
}

func (w *WithdrawService) GetByShopID(ctx context.Context, shopID domain.ID) ([]domain.Withdraw, error) {
	withdraws, err := w.repo.GetByShopID(ctx, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "WithdrawGetByShopID",
		}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{
		"shopID": shopID,
		"count":  len(withdraws),
	}).Info("WithdrawServiceGetByShopID OK")
	return withdraws, nil
}

func (w *WithdrawService) Create(ctx context.Context, param port.CreateWithdrawParam) (domain.Withdraw, error) {
	if param.Sum < 1 {
		log.WithFields(log.Fields{
			"from": "WithdrawCreate",
		}).Error(domain.ErrPrice.Error())
		return domain.Withdraw{}, domain.ErrPrice
	}

	withdraw, err := w.repo.Create(ctx, domain.Withdraw{
		ID:      domain.NewID(),
		ShopID:  param.ShopID,
		Comment: param.Comment,
		Sum:     param.Sum,
		Status:  domain.WithdrawStatusStart,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "WithdrawCreate",
		}).Error(err.Error())
		return domain.Withdraw{}, err
	}

	log.WithFields(log.Fields{
		"shopID": withdraw.ShopID,
		"sum":    withdraw.Sum,
	}).Info("WithdrawServiceCreate OK")
	return withdraw, nil
}

func (w *WithdrawService) Update(ctx context.Context, withdrawID domain.ID, param port.UpdateWithdrawParam) (domain.Withdraw, error) {
	wr, err := w.GetByID(ctx, withdrawID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "WithdrawUpdate",
		}).Error(err.Error())
		return domain.Withdraw{}, err
	}

	if param.Status != nil {
		wr.Status = *param.Status
	}
	if param.Sum.Valid {
		wr.Sum = param.Sum.Int64
	}
	if param.Comment.Valid {
		wr.Comment = param.Comment.String
	}

	if wr.Sum < 1 {
		log.WithFields(log.Fields{
			"from": "WithdrawUpdate",
		}).Error(domain.ErrPrice.Error())
		return domain.Withdraw{}, domain.ErrPrice
	}

	withdraw, err := w.repo.Update(ctx, wr)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "Update",
		}).Error(err.Error())
		return domain.Withdraw{}, err
	}

	log.WithFields(log.Fields{
		"shopID": withdraw.ShopID,
		"sum":    withdraw.Sum,
		"status": withdraw.Status,
	}).Info("WithdrawServiceUpdate OK")
	return withdraw, nil
}

func (w *WithdrawService) Delete(ctx context.Context, withdrawID domain.ID) error {
	err := w.repo.Delete(ctx, withdrawID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "WithdrawDelete",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"id": withdrawID,
	}).Info("WithdrawServiceDelete OK")
	return nil
}
