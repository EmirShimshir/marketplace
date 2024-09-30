package test

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/EmirShimshir/marketplace/internal/adapter/repository/mocks"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/service"
)

type UserSuite struct {
	suite.Suite
}

// Get Suite
type UserGetSuite struct {
	UserSuite
}

func UserGetSuccessRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("Get", context.Background(), mock.Anything, mock.Anything).
		Return([]domain.User{NewUserBuilder().Build()}, nil)
}

func (s *UserGetSuite) TestGet_Success(t provider.T) {
	t.Parallel()
	t.Title("Get success")
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserGetSuccessRepositoryMock(userRepository)
	_, err := userService.Get(context.Background(), 100, 0)
	t.Assert().Nil(err)
}

func UserGetFailureRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("Get", context.Background(), mock.Anything, mock.Anything).
		Return(nil, domain.ErrNotExist)
}

func (s *UserGetSuite) TestGet_Failure(t provider.T) {
	t.Parallel()
	t.Title("Get failure")
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserGetFailureRepositoryMock(userRepository)
	_, err := userService.Get(context.Background(), 100, 0)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestUserGetAllSuite(t *testing.T) {
	suite.RunNamedSuite(t, "Get", new(UserGetSuite))
}

// GetByID Suite
type UserGetByIDSuite struct {
	UserSuite
}

func UserGetByIDSuccessRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("GetByID", context.Background(), userID).
		Return(NewUserBuilder().WithID(userID).Build(), nil)
}

func (s *UserGetByIDSuite) TestGetByID_Success(t provider.T) {
	t.Title("Get by id success")
	userID := domain.NewID()
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserGetByIDSuccessRepositoryMock(userRepository, userID)
	user, err := userService.GetByID(context.Background(), userID)
	t.Assert().Nil(err)
	t.Assert().Equal(userID, user.ID)
}

func UserGetByIDFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("GetByID", context.Background(), userID).
		Return(domain.User{}, domain.ErrNotExist)
}

func (s *UserGetByIDSuite) TestGetByID_Failure(t provider.T) {
	t.Title("Get by id failure")
	userID := domain.NewID()
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserGetByIDFailureRepositoryMock(userRepository, userID)
	_, err := userService.GetByID(context.Background(), userID)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestUserGetByIDSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetByID", new(UserGetByIDSuite))
}

// GetByEmail Suite
type UserGetByEmailSuite struct {
	UserSuite
}

func UserGetByEmailSuccessRepositoryMock(repository *mocks.UserRepository, email string) {
	repository.
		On("GetByEmail", context.Background(), email).
		Return(NewUserBuilder().
			WithEmail(email).
			Build(), nil)
}

func (s *UserGetByEmailSuite) TestGetByEmail_Success(t provider.T) {
	t.Title("Get by email success")
	email := "test@example.com"
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserGetByEmailSuccessRepositoryMock(userRepository, email)
	user, err := userService.GetByEmail(context.Background(), email)
	t.Assert().Nil(err)
	t.Assert().Equal(email, user.Email)
}

func UserGetByEmailFailureRepositoryMock(repository *mocks.UserRepository, email string) {
	repository.
		On("GetByEmail", context.Background(), email).
		Return(domain.User{}, domain.ErrEmail)
}

func (s *UserGetByEmailSuite) TestGetByEmail_Failure(t provider.T) {
	t.Title("Get by email failure")
	email := "test@example.com"
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserGetByEmailFailureRepositoryMock(userRepository, email)
	_, err := userService.GetByEmail(context.Background(), email)
	t.Assert().ErrorIs(err, domain.ErrEmail)
}

func TestUserGetByEmailSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "GetByEmail", new(UserGetByEmailSuite))
}

// Create Suite
type UserCreateSuite struct {
	UserSuite
}

func UserCreateSuccessRepositoryMock(repository *mocks.UserRepository, email string) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(NewUserBuilder().WithEmail(email).Build(), nil)
}

func (s *UserCreateSuite) TestCreate_Success(t provider.T) {
	t.Title("Create success")
	param := NewCreateUserParamBuilder().Build()
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserCreateSuccessRepositoryMock(userRepository, param.Email)
	user, err := userService.Create(context.Background(), param)
	t.Assert().Nil(err)
	t.Assert().Equal(param.Email, user.Email)
}

func UserCreateFailureRepositoryMock(repository *mocks.UserRepository) {
	repository.
		On("Create", context.Background(), mock.Anything).
		Return(domain.User{}, domain.ErrDuplicate)
}

func (s *UserCreateSuite) TestCreate_Failure(t provider.T) {
	t.Title("Create failure")
	param := NewCreateUserParamBuilder().Build()
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserCreateFailureRepositoryMock(userRepository)
	_, err := userService.Create(context.Background(), param)
	t.Assert().ErrorIs(err, domain.ErrDuplicate)
}

func TestUserCreateSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "Create", new(UserCreateSuite))
}

// Update Suite
type UserUpdateSuite struct {
	UserSuite
}

func UserUpdateSuccessRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("GetByID", context.Background(), userID).
		Return(NewUserBuilder().WithID(userID).Build(), nil)
	repository.
		On("Update", context.Background(), mock.Anything).
		Return(NewUserBuilder().WithID(userID).Build(), nil)
}

func (s *UserUpdateSuite) TestUpdate_Success(t provider.T) {
	t.Title("Update success")
	userID := domain.NewID()
	param := NewUpdateUserParamBuilder().WithName("name").Build()
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserUpdateSuccessRepositoryMock(userRepository, userID)
	_, err := userService.Update(context.Background(), userID, param)
	t.Assert().Nil(err)
}

func UserUpdateFailureRepositoryMock(repository *mocks.UserRepository, userID domain.ID) {
	repository.
		On("GetByID", context.Background(), userID).
		Return(domain.User{}, domain.ErrNotExist)
}

func (s *UserUpdateSuite) TestUpdate_Failure(t provider.T) {
	t.Title("Update failure")
	userID := domain.NewID()
	param := NewUpdateUserParamBuilder().WithName("name").Build()
	userRepository := mocks.NewUserRepository(t)
	userService := service.NewUserService(userRepository)
	UserUpdateFailureRepositoryMock(userRepository, userID)
	_, err := userService.Update(context.Background(), userID, param)
	t.Assert().ErrorIs(err, domain.ErrNotExist)
}

func TestUserUpdateSuite(t *testing.T) {
	t.Parallel()
	suite.RunNamedSuite(t, "Update", new(UserUpdateSuite))
}
