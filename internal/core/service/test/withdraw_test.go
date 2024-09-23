package service_test

import (
	"context"
	"errors"
	repoMocks "github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/stretchr/testify/require"
	"testing"
)

var unknownWithdrawID = domain.NewID()
var unknownWithdrawShopID = domain.NewID()

func TestWithdrawService_GetByID(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(userRepo *repoMocks.WithdrawRepository)
		withdraw     domain.Withdraw
		hasError     bool
	}{
		{
			name:     "withdraw found, ok",
			withdraw: dataWithdraws[0],
			initRepoMock: func(withdrawRepo *repoMocks.WithdrawRepository) {
				withdrawRepo.On("GetByID", context.Background(), dataWithdraws[0].ID).
					Return(dataWithdraws[0], nil)
			},
			hasError: false,
		},
		{
			name:     "withdraw not found, error",
			withdraw: domain.Withdraw{ID: unknownWithdrawID},
			initRepoMock: func(withdrawRepo *repoMocks.WithdrawRepository) {
				withdrawRepo.On("GetByID", context.Background(), unknownWithdrawID).
					Return(domain.Withdraw{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			withdrawRepo := repoMocks.NewWithdrawRepository(t)
			withdrawService := service.NewWithdrawService(withdrawRepo)

			test.initRepoMock(withdrawRepo)

			withdraw, err := withdrawService.GetByID(context.Background(), test.withdraw.ID)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.withdraw.ID, withdraw.ID)
			}
		})
	}
}

func TestWithdrawService_GetByEmail(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(userRepo *repoMocks.WithdrawRepository)
		withdraw     domain.Withdraw
		hasError     bool
	}{
		{
			name:     "withdraw found, ok",
			withdraw: dataWithdraws[0],
			initRepoMock: func(withdrawRepo *repoMocks.WithdrawRepository) {
				withdrawRepo.On("GetByShopID", context.Background(), dataWithdraws[0].ShopID).
					Return([]domain.Withdraw{dataWithdraws[0]}, nil)
			},
			hasError: false,
		},
		{
			name:     "withdraw not found, error",
			withdraw: domain.Withdraw{ShopID: unknownWithdrawShopID},
			initRepoMock: func(withdrawRepo *repoMocks.WithdrawRepository) {
				withdrawRepo.On("GetByShopID", context.Background(), unknownWithdrawShopID).
					Return([]domain.Withdraw{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			withdrawRepo := repoMocks.NewWithdrawRepository(t)
			withdrawService := service.NewWithdrawService(withdrawRepo)

			test.initRepoMock(withdrawRepo)

			withdraw, err := withdrawService.GetByShopID(context.Background(), test.withdraw.ShopID)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.withdraw.ShopID, withdraw[0].ShopID)
			}
		})
	}
}
