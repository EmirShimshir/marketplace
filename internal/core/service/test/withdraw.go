package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/guregu/null"
)

type CreateWithdrawParamBuilder struct {
	param port.CreateWithdrawParam
}

func NewCreateWithdrawParamBuilder() *CreateWithdrawParamBuilder {
	return &CreateWithdrawParamBuilder{
		param: port.CreateWithdrawParam{
			ShopID:  domain.NewID(),
			Comment: "Default comment",
			Sum:     1000,
		},
	}
}

func (b *CreateWithdrawParamBuilder) WithShopID(shopID domain.ID) *CreateWithdrawParamBuilder {
	b.param.ShopID = shopID
	return b
}

func (b *CreateWithdrawParamBuilder) WithComment(comment string) *CreateWithdrawParamBuilder {
	b.param.Comment = comment
	return b
}

func (b *CreateWithdrawParamBuilder) WithSum(sum int64) *CreateWithdrawParamBuilder {
	b.param.Sum = sum
	return b
}

func (b *CreateWithdrawParamBuilder) Build() port.CreateWithdrawParam {
	return b.param
}

type UpdateWithdrawParamBuilder struct {
	param port.UpdateWithdrawParam
}

func NewUpdateWithdrawParamBuilder() *UpdateWithdrawParamBuilder {
	return &UpdateWithdrawParamBuilder{
		param: port.UpdateWithdrawParam{
			Comment: null.StringFrom("Updated comment"),
			Sum:     null.IntFrom(2000),
			Status:  nil,
		},
	}
}

func (b *UpdateWithdrawParamBuilder) WithComment(comment string) *UpdateWithdrawParamBuilder {
	b.param.Comment = null.StringFrom(comment)
	return b
}

func (b *UpdateWithdrawParamBuilder) WithSum(sum int64) *UpdateWithdrawParamBuilder {
	b.param.Sum = null.IntFrom(sum)
	return b
}

func (b *UpdateWithdrawParamBuilder) WithStatus(status domain.WithdrawStatus) *UpdateWithdrawParamBuilder {
	b.param.Status = &status
	return b
}

func (b *UpdateWithdrawParamBuilder) Build() port.UpdateWithdrawParam {
	return b.param
}

type WithdrawBuilder struct {
	withdraw domain.Withdraw
}

func NewWithdrawBuilder() *WithdrawBuilder {
	return &WithdrawBuilder{
		withdraw: domain.Withdraw{
			ID:      domain.NewID(),
			ShopID:  domain.NewID(),
			Comment: "Default comment",
			Sum:     1000,
			Status:  domain.WithdrawStatusStart,
		},
	}
}

func (b *WithdrawBuilder) WithID(id domain.ID) *WithdrawBuilder {
	b.withdraw.ID = id
	return b
}

func (b *WithdrawBuilder) WithShopID(shopID domain.ID) *WithdrawBuilder {
	b.withdraw.ShopID = shopID
	return b
}

func (b *WithdrawBuilder) WithComment(comment string) *WithdrawBuilder {
	b.withdraw.Comment = comment
	return b
}

func (b *WithdrawBuilder) WithSum(sum int64) *WithdrawBuilder {
	b.withdraw.Sum = sum
	return b
}

func (b *WithdrawBuilder) WithStatus(status domain.WithdrawStatus) *WithdrawBuilder {
	b.withdraw.Status = status
	return b
}

func (b *WithdrawBuilder) Build() domain.Withdraw {
	return b.withdraw
}
