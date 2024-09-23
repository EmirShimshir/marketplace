package domain

import "errors"

var (
	ErrFilenameEmpty   = errors.New("validation filename is empty error")
	ErrFilepathEmpty   = errors.New("validation filepath is empty error")
	ErrFileReaderEmpty = errors.New("validation file reader is nil error")
	ErrSaveFileError   = errors.New("failed to save file to object storage")
)

var (
	ErrEmptyCart          = errors.New("cart is empty error")
	ErrAddress            = errors.New("address error")
	ErrName               = errors.New("name error")
	ErrSurname            = errors.New("surname error")
	ErrQuantityItems      = errors.New("quantity items error")
	ErrDescription        = errors.New("description error")
	ErrRequisites         = errors.New("requisites error")
	ErrPrice              = errors.New("price error")
	ErrEmail              = errors.New("email error")
	ErrPassword           = errors.New("password error")
	ErrInvalidFingerprint = errors.New("invalid client fingerprint")
	ErrInvalidTokenClaims = errors.New("token or claims are invalid")
	ErrOrderAlreadyPayed  = errors.New("order already payed")
	ErrInvalidPaymentSum  = errors.New("received invalid payment")
)

var (
	ErrDuplicate         = errors.New("record already exists")
	ErrNotExist          = errors.New("record does not exist")
	ErrUpdateFailed      = errors.New("record update failed")
	ErrDeleteFailed      = errors.New("record delete failed")
	ErrPersistenceFailed = errors.New("persistence internal error")
	ErrTransactionError  = errors.New("transaction error occurred")
	ErrInvalidToken      = errors.New("invalid jwt token")
)
