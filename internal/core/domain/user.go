package domain

import "github.com/guregu/null"

type UserRole int

const (
	UserCustomer UserRole = iota
	UserSeller
	UserModerator
)

type User struct {
	ID     ID
	CartID ID
	Name   string
	Surname  string
	Phone    null.String
	Email    string
	Password string
	Role UserRole
}
