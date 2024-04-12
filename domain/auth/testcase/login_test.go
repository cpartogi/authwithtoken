package test_auth

import (
	"authwithtoken/domain/auth/mocks"
	"authwithtoken/domain/auth/model"
	"authwithtoken/domain/auth/usecase"
	"authwithtoken/lib/pkg/utils"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestLogin(t *testing.T) {
	mockRepo := new(mocks.AuthRepoInterface)

	t.Run("Error invalid data", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		_, err := u.Login(context.Background(), model.Users{
			Email:        "c",
			UserPassword: "e",
		})

		assert.Error(t, err, "invalid email address , password containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) character")
	})

	t.Run("Error email required", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		_, err := u.Login(context.Background(), model.Users{
			UserPassword: "eASd@123",
		})

		assert.Error(t, err, "email is required , invalid email address")
	})

	t.Run("Error repo get user by email", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, errors.New("failed")).Once()

		_, err := u.Login(context.Background(), model.Users{
			Email:        "abc@def.com",
			UserPassword: "eASd@123",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error data not found", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, nil).Once()

		_, err := u.Login(context.Background(), model.Users{
			Email:        "abc@def.com",
			UserPassword: "eASd@123",
		})

		assert.Error(t, err, "data not found")
	})

	t.Run("Error repo insert user log", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		_, err := u.Login(context.Background(), model.Users{
			Email:        "abc@def.com",
			UserPassword: "eASd@123",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error wrong password", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(nil).Once()

		_, err := u.Login(context.Background(), model.Users{
			Email:        "abc@def.com",
			UserPassword: "eASd@123",
		})

		assert.Error(t, err, "wrong password")
	})

	t.Run("Error repo insert user log password match", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		userPassword := "eASd@123"
		userPassHash, _ := utils.HashPassword(userPassword)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: userPassHash,
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		_, err := u.Login(context.Background(), model.Users{
			Email:        "abc@def.com",
			UserPassword: userPassword,
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error repo UpsertUserToken", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		userPassword := "eASd@123"
		userPassHash, _ := utils.HashPassword(userPassword)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: userPassHash,
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(nil).Once()

		mockRepo.On("UpsertUserToken", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		_, err := u.Login(context.Background(), model.Users{
			Email:        "abc@def.com",
			UserPassword: userPassword,
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		userPassword := "eASd@123"
		userPassHash, _ := utils.HashPassword(userPassword)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: userPassHash,
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(nil).Once()

		mockRepo.On("UpsertUserToken", mock.Anything, mock.Anything).Return(nil).Once()

		_, err := u.Login(context.Background(), model.Users{
			Email:        "abc@def.com",
			UserPassword: userPassword,
		})

		assert.NilError(t, err)
	})
}
