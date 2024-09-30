package test

import (
	"context"
	"testing"

	"github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
)

type WithdrawSuite struct {
	suite.Suite
}

type WithdrawGetSuite struct {
	WithdrawSuite
}

func WithdrawGetSuccessRepositoryMock(repository *mocks.WithdrawRepository) {
	repository.
		On("Get", context.Background(), mock.Anything, mock.Anything).
		Return([]domain.Withdraw{NewWithdrawBuilder().Build()}, nil)
}

func (s *WithdrawGetSuite) TestGet_Success(t provider.T) {
	t.Parallel()
	t.Title("Get success")
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawGetSuccessRepositoryMock(withdrawRepository)
	_, err := withdrawService.Get(context.Background(), 100, 0)
	t.Assert().Nil(err)
}

func WithdrawGetFailureRepositoryMock(repository *mocks.WithdrawRepository) {
	repository.
		On("Get", context.Background(), mock.Anything, mock.Anything).
		Return(nil, domain.ErrNotExist)
}

func (s *WithdrawGetSuite) TestGet_Failure(t provider.T) {
	t.Parallel()
	t.Title("Get failure")
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawGetFailureRepositoryMock(withdrawRepository)
	_, err := withdrawService.Get(context.Background(), 100, 0)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestWithdrawGetSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetWithdraw", new(WithdrawGetSuite))
}

type WithdrawGetByIDSuite struct {
	WithdrawSuite
}

func WithdrawGetByIDSuccessRepositoryMock(repository *mocks.WithdrawRepository, withdrawID domain.ID) {
	repository.
		On("GetByID", context.Background(), withdrawID).
		Return(NewWithdrawBuilder().WithID(withdrawID).Build(), nil)
}

func (s *WithdrawGetByIDSuite) TestGetByID_Success(t provider.T) {
	t.Parallel()
	t.Title("Get by ID success")
	withdrawID := domain.NewID()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawGetByIDSuccessRepositoryMock(withdrawRepository, withdrawID)
	withdraw, err := withdrawService.GetByID(context.Background(), withdrawID)
	t.Assert().Nil(err)
	t.Assert().Equal(withdrawID, withdraw.ID)
}

func WithdrawGetByIDFailureRepositoryMock(repository *mocks.WithdrawRepository, withdrawID domain.ID) {
	repository.
		On("GetByID", context.Background(), withdrawID).
		Return(domain.Withdraw{}, domain.ErrNotExist)
}

func (s *WithdrawGetByIDSuite) TestGetByID_Failure(t provider.T) {
	t.Parallel()
	t.Title("Get by ID failure")
	withdrawID := domain.NewID()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawGetByIDFailureRepositoryMock(withdrawRepository, withdrawID)
	_, err := withdrawService.GetByID(context.Background(), withdrawID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestWithdrawGetByIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetByID", new(WithdrawGetByIDSuite))
}

type WithdrawCreateSuite struct {
	WithdrawSuite
}

func WithdrawCreateSuccessRepositoryMock(repository *mocks.WithdrawRepository, shopID domain.ID) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewWithdrawBuilder().WithShopID(shopID).Build(), nil)
}

func (s *WithdrawCreateSuite) TestCreate_Success(t provider.T) {
	t.Parallel()
	t.Title("Create success")
	param := NewCreateWithdrawParamBuilder().Build()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawCreateSuccessRepositoryMock(withdrawRepository, param.ShopID)
	withdraw, err := withdrawService.Create(context.Background(), param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.ShopID, withdraw.ShopID)
}

func WithdrawCreateFailureRepositoryMock(repository *mocks.WithdrawRepository) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.Withdraw{}, domain.ErrDuplicate)
}

func (s *WithdrawCreateSuite) TestCreate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Create failure")
	param := NewCreateWithdrawParamBuilder().Build()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawCreateFailureRepositoryMock(withdrawRepository)
	_, err := withdrawService.Create(context.Background(), param)
	t.Assert().ErrorIs(err, domain.ErrDuplicate)
}

func TestWithdrawCreateSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "CreateWithdraw", new(WithdrawCreateSuite))
}

type WithdrawUpdateSuite struct {
	WithdrawSuite
}

func WithdrawUpdateSuccessRepositoryMock(repository *mocks.WithdrawRepository, withdrawID domain.ID) {
	repository.
		On("GetByID", context.Background(), withdrawID).
		Return(NewWithdrawBuilder().WithID(withdrawID).Build(), nil)
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewWithdrawBuilder().WithID(withdrawID).Build(), nil)
}

func (s *WithdrawUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Parallel()
	t.Title("Update success")
	withdrawID := domain.NewID()
	param := NewUpdateWithdrawParamBuilder().WithComment("Updated comment").Build()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawUpdateSuccessRepositoryMock(withdrawRepository, withdrawID)
	_, err := withdrawService.Update(context.Background(), withdrawID, param)
	t.Assert().Nil(err)
}

func WithdrawUpdateFailureRepositoryMock(repository *mocks.WithdrawRepository, withdrawID domain.ID) {
	repository.
		On("GetByID", context.Background(), withdrawID).
		Return(domain.Withdraw{}, domain.ErrNotExist)
}

func (s *WithdrawUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Parallel()
	t.Title("Update failure")
	withdrawID := domain.NewID()
	param := NewUpdateWithdrawParamBuilder().WithComment("Updated comment").Build()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawUpdateFailureRepositoryMock(withdrawRepository, withdrawID)
	_, err := withdrawService.Update(context.Background(), withdrawID, param)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestWithdrawUpdateSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "UpdateWithdraw", new(WithdrawUpdateSuite))
}

// GetByShopID Suite
type WithdrawGetByShopIDSuite struct {
	WithdrawSuite
}

func WithdrawGetByShopIDSuccessRepositoryMock(repository *mocks.WithdrawRepository, shopID domain.ID) {
	repository.
		On("GetByShopID", context.Background(), shopID).
		Return([]domain.Withdraw{NewWithdrawBuilder().WithShopID(shopID).Build()}, nil)
}

func (s *WithdrawGetByShopIDSuite) TestGetByShopID_Success(t provider.T) {
	t.Title("Get by shop ID success")
	shopID := domain.NewID()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawGetByShopIDSuccessRepositoryMock(withdrawRepository, shopID)
	withdraws, err := withdrawService.GetByShopID(context.Background(), shopID)
	t.Assert().Nil(err)
	t.Assert().Equal(shopID, withdraws[0].ShopID)
}

func WithdrawGetByShopIDFailureRepositoryMock(repository *mocks.WithdrawRepository, shopID domain.ID) {
	repository.
		On("GetByShopID", context.Background(), shopID).
		Return(nil, domain.ErrNotExist)
}

func (s *WithdrawGetByShopIDSuite) TestGetByShopID_Failure(t provider.T) {
	t.Title("Get by shop ID failure")
	shopID := domain.NewID()
	withdrawRepository := mocks.NewWithdrawRepository(t)
	withdrawService := service.NewWithdrawService(withdrawRepository)
	WithdrawGetByShopIDFailureRepositoryMock(withdrawRepository, shopID)
	_, err := withdrawService.GetByShopID(context.Background(), shopID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestWithdrawGetByShopIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetByShopID", new(WithdrawGetByShopIDSuite))
}
