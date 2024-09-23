package tech

import (
	"context"
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/tech/dto"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) Logout(c *Console) {
	c.UserID = nil
	fmt.Println("Вы успешно вышли из системы")
}

func (h *Handler) SignUp(c *Console) {
	var signUpDTO dto.SignUpDTO
	err := dto.InputSignUpDTO(&signUpDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignUp",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	err = h.auth.SignUp(context.Background(), port.SignUpParam{
		Name:     signUpDTO.Name,
		Surname:  signUpDTO.Surname,
		Email:    signUpDTO.Email,
		Password: signUpDTO.Password,
		Role:     signUpDTO.Role,
		Phone:    signUpDTO.Phone,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("Вы успешно зарегистрированы")
}

func (h *Handler) SignIn(c *Console) {
	if c.UserID != nil {
		log.WithFields(log.Fields{
			"SignIn": "not auth",
		}).Error(BadRequestError)
		ErrorResponse(BadRequestError)
		return
	}

	var signInDTO dto.SignInDTO
	err := dto.InputSignInDTO(&signInDTO)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignIn",
		}).Error(err.Error())
		ErrorResponse(err)
		return
	}

	authDetails, err := h.auth.SignIn(context.Background(), port.SignInParam{
		Email:    signInDTO.Email,
		Password: signInDTO.Password,
	})
	if err != nil {
		ErrorResponse(err)
		return
	}

	payload, err := h.auth.Payload(context.Background(), authDetails.AccessToken)
	if err != nil {
		ErrorResponse(err)
		return
	}

	user, err := h.user.GetByID(context.Background(), payload.UserID)
	if err != nil {
		ErrorResponse(err)
		return
	}

	c.UserID = &user.ID
	fmt.Println("Вы успешно вошли в систему")
}

func (h *Handler) verifyAuthRole(c *Console, role domain.UserRole) error {
	if c.UserID == nil {
		log.WithFields(log.Fields{
			"from": "verifyAuthRole",
		}).Error(UnauthorizedError)
		return UnauthorizedError
	}

	u, err := h.user.GetByID(context.Background(), *c.UserID)
	if err != nil {
		ErrorResponse(err)
		return err
	}

	if u.Role != role {
		return ForbiddenError
	}

	return nil
}

func (h *Handler) verifyAuth(c *Console) error {
	if c.UserID == nil {
		log.WithFields(log.Fields{
			"from": "verifyAuth",
		}).Error(UnauthorizedError)
		return UnauthorizedError
	}

	return nil
}
