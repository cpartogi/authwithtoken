package test_auth

import (
	"authwithtoken/domain/auth/mocks"
	"authwithtoken/domain/auth/model"
	"authwithtoken/domain/auth/usecase"
	"authwithtoken/lib/helper"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestUpdateUser(t *testing.T) {
	mockRepo := new(mocks.AuthRepoInterface)

	t.Run("Error forbidden", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		_, err := u.UpdateUser(context.Background(), "abc", model.Users{})

		assert.Error(t, err, "forbidden")
	})

	t.Run("Error repo get user by id", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, errors.New("failed")).Once()

		_, err := u.UpdateUser(context.Background(), token.Token, model.Users{
			Id:           "a",
			FullName:     "b",
			PhoneNumber:  "c",
			UserPassword: "eASd@123",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error data not found", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		_, err := u.UpdateUser(context.Background(), token.Token, model.Users{
			Id:           "a",
			FullName:     "b",
			PhoneNumber:  "c",
			UserPassword: "d",
		})

		assert.Error(t, err, "data not found")
	})

	t.Run("Error invalid data for update", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		_, err := u.UpdateUser(context.Background(), token.Token, model.Users{
			Id:           "a",
			FullName:     "b",
			PhoneNumber:  "c",
			UserPassword: "d",
		})

		assert.Error(t, err, "fullName must be at minimum 3 characters and maximum 60 characters , phoneNumber must be at minimum 10 characters and maximum 13 characters , phoneNumber must start with the Indonesia country code +62 , password must be minimum 6 characters and maximum 64 characters , password containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) character")
	})

	t.Run("Error repo update user", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		mockRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		_, err := u.UpdateUser(context.Background(), token.Token, model.Users{
			Id:           "a",
			FullName:     "Full Name ini",
			PhoneNumber:  "+628781122333",
			UserPassword: "eASd@123",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		mockRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(nil).Once()

		_, err := u.UpdateUser(context.Background(), token.Token, model.Users{
			Id:           "a",
			FullName:     "Full Name ini",
			PhoneNumber:  "+628781122333",
			UserPassword: "eASd@123",
		})

		assert.NilError(t, err)
	})
}
