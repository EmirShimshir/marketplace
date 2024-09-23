package domain

type WithdrawStatus int

const (
	WithdrawStatusStart WithdrawStatus = iota
	WithdrawStatusReady
	WithdrawStatusDone
)

type Withdraw struct {
	ID      ID
	ShopID  ID
	Comment string
	Sum     int64
	Status  WithdrawStatus
}
