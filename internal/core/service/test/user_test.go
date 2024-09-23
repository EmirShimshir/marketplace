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

var unknownUserID = domain.NewID()
var unknownEmail = "test@yandex.ru"

func TestUserService_GetByID(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(userRepo *repoMocks.UserRepository)
		user         domain.User
		hasError     bool
	}{
		{
			name: "user found, ok",
			user: dataUsers[0],
			initRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), dataUsers[0].ID).
					Return(dataUsers[0], nil)
			},
			hasError: false,
		},
		{
			name: "user not found, error",
			user: domain.User{ID: unknownUserID},
			initRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByID", context.Background(), unknownUserID).
					Return(domain.User{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			userRepo := repoMocks.NewUserRepository(t)
			userService := service.NewUserService(userRepo)

			test.initRepoMock(userRepo)

			user, err := userService.GetByID(context.Background(), test.user.ID)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.user.ID, user.ID)
			}
		})
	}
}

func TestUserService_GetByEmail(t *testing.T) {
	testTable := []struct {
		name         string
		initRepoMock func(userRepo *repoMocks.UserRepository)
		user         domain.User
		hasError     bool
	}{
		{
			name: "user found, ok",
			user: dataUsers[0],
			initRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByEmail", context.Background(), dataUsers[0].Email).
					Return(dataUsers[0], nil)
			},
			hasError: false,
		},
		{
			name: "user not found, error",
			user: domain.User{Email: unknownEmail},
			initRepoMock: func(userRepo *repoMocks.UserRepository) {
				userRepo.On("GetByEmail", context.Background(), unknownEmail).
					Return(domain.User{}, errors.New("error"))
			},
			hasError: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			userRepo := repoMocks.NewUserRepository(t)
			userService := service.NewUserService(userRepo)

			test.initRepoMock(userRepo)

			user, err := userService.GetByEmail(context.Background(), test.user.Email)

			if test.hasError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.user.Email, user.Email)
			}
		})
	}
}
