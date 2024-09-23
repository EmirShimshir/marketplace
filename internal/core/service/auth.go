package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
	"strings"
)

type AuthService struct {
	authProvider port.IAuthProvider
	userService  port.IUserService
}

func NewAuthService(authProvider port.IAuthProvider, userService port.IUserService) *AuthService {
	return &AuthService{
		authProvider: authProvider,
		userService:  userService,
	}
}

func (a *AuthService) SignIn(ctx context.Context, param port.SignInParam) (domain.AuthDetails, error) {
	user, err := a.userService.GetByEmail(ctx, strings.ToLower(param.Email))
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignIn",
		}).Error(err.Error())
		return domain.AuthDetails{}, domain.ErrEmail
	}

	if user.Password != a.authProvider.GenPasswordHash(param.Password) {
		log.WithFields(log.Fields{
			"from": "SignIn",
		}).Error(domain.ErrPassword.Error())
		return domain.AuthDetails{}, domain.ErrPassword
	}
	ad, err := a.authProvider.CreateJWTSession(domain.AuthPayload{UserID: user.ID, Role: user.Role}, param.Fingerprint)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignIn",
		}).Error(err.Error())
		return domain.AuthDetails{}, err
	}

	log.WithFields(log.Fields{
		"userID": user.ID,
		"Email":  user.Email,
	}).Info("SignIn OK")
	return ad, nil
}

func (a *AuthService) SignUp(ctx context.Context, param port.SignUpParam) error {
	user, err := a.userService.Create(ctx, port.CreateUserParam{
		Name:     param.Name,
		Surname:  param.Surname,
		Phone:    param.Phone,
		Email:    strings.ToLower(param.Email),
		Password: a.authProvider.GenPasswordHash(param.Password),
		Role:     param.Role,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "SignUp",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"userID": user.ID,
		"Email":  user.Email,
	}).Info("SignUp OK")
	return nil
}

func (a *AuthService) LogOut(ctx context.Context, refreshToken domain.Token) error {
	err := a.authProvider.DeleteJWTSession(refreshToken)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "LogOut",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"from":         "LogOut",
		"refreshToken": refreshToken.String(),
	}).Info("LogOut OK")
	return nil
}

func (a *AuthService) Refresh(ctx context.Context, refreshToken domain.Token,
	fingerprint string) (domain.AuthDetails, error) {
	ad, err := a.authProvider.RefreshJWTSession(refreshToken, fingerprint)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "Refresh",
		}).Error(err.Error())
		return domain.AuthDetails{}, err
	}

	log.WithFields(log.Fields{
		"from":         "LogOut",
		"refreshToken": refreshToken.String(),
	}).Info("Refresh OK")
	return ad, nil
}

func (a *AuthService) Verify(ctx context.Context, accessToken domain.Token) error {
	_, err := a.authProvider.VerifyJWTToken(accessToken)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "Verify",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"from":         "LogOut",
		"refreshToken": accessToken.String(),
	}).Info("Verify OK")
	return nil
}

func (a *AuthService) Payload(ctx context.Context, accessToken domain.Token) (domain.AuthPayload, error) {
	ap, err := a.authProvider.VerifyJWTToken(accessToken)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "Payload",
		}).Error(err.Error())
		return domain.AuthPayload{}, err
	}

	log.WithFields(log.Fields{
		"from":         "LogOut",
		"refreshToken": accessToken.String(),
	}).Info("Payload OK")
	return ap, nil
}
