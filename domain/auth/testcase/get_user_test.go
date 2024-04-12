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

func TestGetUser(t *testing.T) {
	mockRepo := new(mocks.AuthRepoInterface)

	t.Run("Error forbidden", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		_, err := u.GetUser(context.Background(), "abc")

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

		_, err := u.GetUser(context.Background(), token.Token)

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

		_, err := u.GetUser(context.Background(), token.Token)

		assert.Error(t, err, "data not found")
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

		_, err := u.GetUser(context.Background(), token.Token)

		assert.NilError(t, err)
	})
}
