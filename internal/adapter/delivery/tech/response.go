package tech

import (
	"errors"
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

var (
	BadRequestError    = errors.New("bad request")
	NotFoundError      = errors.New("not Found")
	AlreadyExistsError = errors.New("already exists")
	UnauthorizedError  = errors.New("unauthorized")
	ForbiddenError     = errors.New("forbidden")
)

type ErrConsole interface {
	Error() string
}

type ErrorConsole struct {
	ErrMessage string
}

func (e ErrorConsole) Error() string {
	return fmt.Sprintf("error: %s", e.ErrMessage)
}

func NewConsoleError(err string) ErrConsole {
	return ErrorConsole{
		ErrMessage: err,
	}
}

func ParseError(err error) ErrConsole {
	if errors.Is(err, domain.ErrNotExist) {
		return NewConsoleError(NotFoundError.Error())
	}
	if errors.Is(err, domain.ErrDuplicate) {
		return NewConsoleError(AlreadyExistsError.Error())
	}
	return NewConsoleError(err.Error())
}

func ErrorResponse(err error) {
	consoleErr := ParseError(err)
	fmt.Printf("%v\n", consoleErr)
}
