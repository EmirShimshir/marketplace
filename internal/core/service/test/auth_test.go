package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
)

type AuthDetailsBuilder struct {
	accessToken  domain.Token
	refreshToken domain.Token
}

func NewAuthDetailsBuilder() *AuthDetailsBuilder {
	return &AuthDetailsBuilder{}
}

func (b *AuthDetailsBuilder) WithAccessToken(token domain.Token) *AuthDetailsBuilder {
	b.accessToken = token
	return b
}

func (b *AuthDetailsBuilder) WithRefreshToken(token domain.Token) *AuthDetailsBuilder {
	b.refreshToken = token
	return b
}

func (b *AuthDetailsBuilder) Build() domain.AuthDetails {
	return domain.AuthDetails{
		AccessToken:  b.accessToken,
		RefreshToken: b.refreshToken,
	}
}

type AuthPayloadBuilder struct {
	userID domain.ID
	role   domain.UserRole
}

func NewAuthPayloadBuilder() *AuthPayloadBuilder {
	return &AuthPayloadBuilder{}
}

func (b *AuthPayloadBuilder) WithUserID(id domain.ID) *AuthPayloadBuilder {
	b.userID = id
	return b
}

func (b *AuthPayloadBuilder) WithRole(role domain.UserRole) *AuthPayloadBuilder {
	b.role = role
	return b
}

func (b *AuthPayloadBuilder) Build() domain.AuthPayload {
	return domain.AuthPayload{
		UserID: b.userID,
		Role:   b.role,
	}
}

type SignInParamBuilder struct {
	email       string
	password    string
	fingerprint string
}

func NewSignInParamBuilder() *SignInParamBuilder {
	return &SignInParamBuilder{}
}

func (b *SignInParamBuilder) WithEmail(email string) *SignInParamBuilder {
	b.email = email
	return b
}

func (b *SignInParamBuilder) WithPassword(password string) *SignInParamBuilder {
	b.password = password
	return b
}

func (b *SignInParamBuilder) WithFingerprint(fingerprint string) *SignInParamBuilder {
	b.fingerprint = fingerprint
	return b
}

func (b *SignInParamBuilder) Build() port.SignInParam {
	return port.SignInParam{
		Email:       b.email,
		Password:    b.password,
		Fingerprint: b.fingerprint,
	}
}

type SignUpParamBuilder struct {
	name     string
	surname  string
	email    string
	password string
	phone    null.String
	role     domain.UserRole
}

func NewSignUpParamBuilder() *SignUpParamBuilder {
	return &SignUpParamBuilder{}
}

func (b *SignUpParamBuilder) WithName(name string) *SignUpParamBuilder {
	b.name = name
	return b
}

func (b *SignUpParamBuilder) WithSurname(surname string) *SignUpParamBuilder {
	b.surname = surname
	return b
}

func (b *SignUpParamBuilder) WithEmail(email string) *SignUpParamBuilder {
	b.email = email
	return b
}

func (b *SignUpParamBuilder) WithPassword(password string) *SignUpParamBuilder {
	b.password = password
	return b
}

func (b *SignUpParamBuilder) WithPhone(phone null.String) *SignUpParamBuilder {
	b.phone = phone
	return b
}

func (b *SignUpParamBuilder) WithRole(role domain.UserRole) *SignUpParamBuilder {
	b.role = role
	return b
}

func (b *SignUpParamBuilder) Build() port.SignUpParam {
	return port.SignUpParam{
		Name:     b.name,
		Surname:  b.surname,
		Email:    b.email,
		Password: b.password,
		Phone:    b.phone,
		Role:     b.role,
	}
}
