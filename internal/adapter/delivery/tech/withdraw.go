package tech

import (
	"context"
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/tech/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetWithdrawsAll(console *Console) {
	err := h.verifyAuthRole(console, domain.UserModerator)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetWithdrawsAll",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	withdraws, err := h.withdraw.Get(context.Background(), 100, 0)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(withdraws) == 0 {
		fmt.Println("Нет заявок для вывода")
		return
	}

	for _, withdraw := range withdraws {
		w := dto.NewWithdrawDTO(withdraw)
		w.Print()
	}
}

func (h *Handler) GetWithdrawByShopID(console *Console) {
	var shopID domain.ID
	err := dto.InputID(&shopID, "магазина")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetWithdrawByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	err = h.verifyUserIsShopOwner(console, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetWithdrawByShopID",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	withdraws, err := h.withdraw.GetByShopID(context.Background(), shopID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	if len(withdraws) == 0 {
		fmt.Println("Нет заявок для вывода")
		return
	}

	for _, withdraw := range withdraws {
		w := dto.NewWithdrawDTO(withdraw)
		w.Print()
	}
}

func (h *Handler) CreateWithdraw(console *Console) {
	var shopID domain.ID
	err := dto.InputID(&shopID, "магазина")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateWithdraw",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	err = h.verifyUserIsShopOwner(console, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateWithdraw",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var sum int64
	fmt.Print("Введите сумму: ")
	fmt.Scanln(&sum)

	var comment string
	fmt.Print("Введите комментарий: ")
	fmt.Scanln(&comment)

	withdraw, err := h.withdraw.Create(context.Background(), port.CreateWithdrawParam{
		ShopID:  shopID,
		Comment: comment,
		Sum:     sum,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	w := dto.NewWithdrawDTO(withdraw)
	w.Print()
}

func (h *Handler) UpdateWithdraw(console *Console) {
	err := h.verifyAuthRole(console, domain.UserModerator)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateWithdraw",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var withdrawID domain.ID
	err = dto.InputID(&withdrawID, "заявки")
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateWithdraw",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	var status domain.WithdrawStatus
	err = dto.InputWithdrawStatus(&status)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateWithdraw",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	withdraw, err := h.withdraw.Update(context.Background(), withdrawID, port.UpdateWithdrawParam{
		Status: &status,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	w := dto.NewWithdrawDTO(withdraw)
	w.Print()
}
