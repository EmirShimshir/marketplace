package tech

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/tech/dto"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetUser(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetUser",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}
	userID := *c.UserID

	user, err := h.user.GetByID(context.Background(), userID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	userDTO := dto.NewUserDTO(user)
	userDTO.Print()
}

func (h *Handler) UpdateUser(c *Console) {
	err := h.verifyAuth(c)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateUser",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}
	userID := *c.UserID

	var updateUserDTO dto.UpdateUserDTO
	err = dto.InputUpdateUserDTO(&updateUserDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateUser",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	user, err := h.user.Update(context.Background(), userID, port.UpdateUserParam{
		Name:    updateUserDTO.Name,
		Surname: updateUserDTO.Surname,
		Phone:   updateUserDTO.Phone,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	userDTO := dto.NewUserDTO(user)
	userDTO.Print()
}
