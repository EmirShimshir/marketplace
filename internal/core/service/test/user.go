package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
)

// CreateUserParamBuilder helps construct CreateUserParam with default values.
type CreateUserParamBuilder struct {
	param port.CreateUserParam
}

// NewCreateUserParamBuilder returns a new instance of CreateUserParamBuilder.
func NewCreateUserParamBuilder() *CreateUserParamBuilder {
	return &CreateUserParamBuilder{
		param: port.CreateUserParam{
			Name:     "name",
			Surname:  "surname",
			Phone:    null.StringFrom("+1234567890"),
			Email:    "email@gmail.com",
			Password: "password123",
			Role:     domain.UserCustomer,
		},
	}
}

// WithName sets the name field.
func (b *CreateUserParamBuilder) WithName(name string) *CreateUserParamBuilder {
	b.param.Name = name
	return b
}

// WithSurname sets the surname field.
func (b *CreateUserParamBuilder) WithSurname(surname string) *CreateUserParamBuilder {
	b.param.Surname = surname
	return b
}

// WithPhone sets the phone field.
func (b *CreateUserParamBuilder) WithPhone(phone string) *CreateUserParamBuilder {
	b.param.Phone = null.StringFrom(phone)
	return b
}

// WithEmail sets the email field.
func (b *CreateUserParamBuilder) WithEmail(email string) *CreateUserParamBuilder {
	b.param.Email = email
	return b
}

// WithPassword sets the password field.
func (b *CreateUserParamBuilder) WithPassword(password string) *CreateUserParamBuilder {
	b.param.Password = password
	return b
}

// WithPassword sets the password field.
func (b *CreateUserParamBuilder) WithRole(role domain.UserRole) *CreateUserParamBuilder {
	b.param.Role = role
	return b
}

// Build returns the built CreateUserParam.
func (b *CreateUserParamBuilder) Build() port.CreateUserParam {
	return b.param
}

// UpdateUserParamBuilder helps construct UpdateUserParam with default values.
type UpdateUserParamBuilder struct {
	param port.UpdateUserParam
}

// NewUpdateUserParamBuilder returns a new instance of UpdateUserParamBuilder.
func NewUpdateUserParamBuilder() *UpdateUserParamBuilder {
	return &UpdateUserParamBuilder{
		param: port.UpdateUserParam{
			Name:    null.String{},
			Surname: null.String{},
			Phone:   null.String{},
		},
	}
}

// WithName sets the name field.
func (b *UpdateUserParamBuilder) WithName(name string) *UpdateUserParamBuilder {
	b.param.Name = null.StringFrom(name)
	return b
}

// WithSurname sets the surname field.
func (b *UpdateUserParamBuilder) WithSurname(surname string) *UpdateUserParamBuilder {
	b.param.Surname = null.StringFrom(surname)
	return b
}

// WithPhone sets the phone field.
func (b *UpdateUserParamBuilder) WithPhone(phone string) *UpdateUserParamBuilder {
	b.param.Phone = null.StringFrom(phone)
	return b
}

// Build returns the built UpdateUserParam.
func (b *UpdateUserParamBuilder) Build() port.UpdateUserParam {
	return b.param
}

type UserBuilder struct {
	user domain.User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: domain.User{
			ID:       domain.NewID(),
			CartID:   domain.NewID(),
			Name:     "John",
			Surname:  "Doe",
			Phone:    null.StringFrom("+123456789"),
			Email:    "john.doe@example.com",
			Password: "password",
			Role:     domain.UserCustomer,
		},
	}
}

func (b *UserBuilder) WithID(id domain.ID) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithCartID(cartID domain.ID) *UserBuilder {
	b.user.CartID = cartID
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithSurname(surname string) *UserBuilder {
	b.user.Surname = surname
	return b
}

func (b *UserBuilder) WithPhone(phone string) *UserBuilder {
	b.user.Phone = null.StringFrom(phone)
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.user.Password = password
	return b
}

func (b *UserBuilder) WithRole(role domain.UserRole) *UserBuilder {
	b.user.Role = role
	return b
}

func (b *UserBuilder) Build() domain.User {
	return b.user
}
