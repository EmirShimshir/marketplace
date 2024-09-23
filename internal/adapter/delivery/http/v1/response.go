package v1

import (
	"context"
	"errors"
	"fmt"
	errs "github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const (
	ErrBadRequest          = "bad request"
	ErrNotFound            = "not found"
	ErrUnauthorized        = "unauthorized"
	ErrForbidden           = "forbidden"
	ErrInternalServerError = "internal server error"
	ErrRequestTimeout      = "request timeout"
)

var (
	BadRequestError          = errors.New("bad request")
	UnauthorizedError        = errors.New("unauthorized")
	ForbiddenError           = errors.New("forbidden")
	PathIdParamIsEmptyError  = errors.New("empty id query parameter")
	PathIdParamIsInvalidUUID = errors.New("id query parameter is not uuid")
)

var errorStatusMap = map[error]int{
	errs.ErrFilenameEmpty:     http.StatusBadRequest,
	errs.ErrFilepathEmpty:     http.StatusBadRequest,
	errs.ErrFileReaderEmpty:   http.StatusBadRequest,
	errs.ErrSaveFileError:     http.StatusBadRequest,
	errs.ErrEmptyCart:         http.StatusBadRequest,
	errs.ErrAddress:           http.StatusBadRequest,
	errs.ErrName:              http.StatusBadRequest,
	errs.ErrSurname:           http.StatusBadRequest,
	errs.ErrQuantityItems:     http.StatusBadRequest,
	errs.ErrDescription:       http.StatusBadRequest,
	errs.ErrRequisites:        http.StatusBadRequest,
	errs.ErrPrice:             http.StatusBadRequest,
	errs.ErrEmail:             http.StatusBadRequest,
	errs.ErrPassword:          http.StatusBadRequest,
	errs.ErrOrderAlreadyPayed: http.StatusBadRequest,
	errs.ErrInvalidPaymentSum: http.StatusBadRequest,

	errs.ErrDuplicate:          http.StatusBadRequest,
	errs.ErrNotExist:           http.StatusNotFound,
	errs.ErrUpdateFailed:       http.StatusInternalServerError,
	errs.ErrDeleteFailed:       http.StatusInternalServerError,
	errs.ErrPersistenceFailed:  http.StatusInternalServerError,
	errs.ErrTransactionError:   http.StatusInternalServerError,
	errs.ErrInvalidTokenClaims: http.StatusUnauthorized,
	errs.ErrInvalidFingerprint: http.StatusUnauthorized,

	PathIdParamIsEmptyError:  http.StatusBadRequest,
	PathIdParamIsInvalidUUID: http.StatusBadRequest,
}

type RestErr interface {
	Status() int
	Error() string
}

type RestError struct {
	ErrStatus  int       `json:"status,omitempty"`
	ErrMessage string    `json:"error,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d, error: %s", e.ErrStatus, e.ErrMessage)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func NewRestError(status int, err string) RestErr {
	return RestError{
		ErrStatus:  status,
		ErrMessage: err,
		Timestamp:  time.Now().UTC(),
	}
}

func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("field '%s' must be not empty", strings.ToLower(err.Field()))
	case "email":
		return fmt.Sprintf("invalid email")
	case "url":
		return fmt.Sprintf("field '%s' must be URL", strings.ToLower(err.Field()))
	case "oneof":
		return fmt.Sprintf("field '%s' must be enum type", err.Field())
	default:
		return "json validation error"
	}
}

func ParseError(err error) RestErr {
	var validationErrors validator.ValidationErrors

	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, ErrRequestTimeout)
	case errors.Is(err, UnauthorizedError):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized)
	case errors.Is(err, BadRequestError):
		return NewRestError(http.StatusBadRequest, ErrBadRequest)
	case errors.Is(err, ForbiddenError):
		return NewRestError(http.StatusForbidden, ErrForbidden)
	case errors.Is(err, errs.ErrInvalidToken):
		return NewRestError(http.StatusUnauthorized, err.Error())
	case errors.As(err, &validationErrors):
		return NewRestError(http.StatusBadRequest, getValidationMessage(validationErrors[0]))
	default:
		if code, ok := checkError(errorStatusMap, err); ok {
			return NewRestError(code, err.Error())
		}
		if restErr, ok := err.(*RestError); ok {
			return restErr
		}
		return NewRestError(http.StatusInternalServerError, ErrInternalServerError)
	}
}

func checkError(errorStatusMap map[error]int, err error) (int, bool) {
	for key, value := range errorStatusMap {
		if errors.Is(err, key) {
			return value, true
		}
	}

	return 0, false
}

func (h *Handler) errorResponse(context *gin.Context, err error) {
	log.Error(err.Error())
	restErr := ParseError(err)
	context.AbortWithStatusJSON(restErr.Status(), restErr)
}

func (h *Handler) successResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, data)
}

func (h *Handler) createdResponse(context *gin.Context, data interface{}) {
	context.JSON(http.StatusCreated, data)
}
