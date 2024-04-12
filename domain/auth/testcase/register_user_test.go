package test_auth

import (
	"authwithtoken/domain/auth/mocks"
	"authwithtoken/domain/auth/model"
	"context"
	"errors"
	"testing"

	"authwithtoken/domain/auth/usecase"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestRegisterUser(t *testing.T) {
	mockRepo := new(mocks.AuthRepoInterface)

	t.Run("Error invalid data", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		_, err := u.RegisterUser(context.Background(), model.Users{
			FullName:     "b",
			Email:        "c",
			PhoneNumber:  "d",
			UserPassword: "e",
		})

		assert.Error(t, err, "fullName must be at minimum 3 characters and maximum 60 characters , phoneNumber must be at minimum 10 characters and maximum 13 characters , phoneNumber must start with the Indonesia country code +62 , password must be minimum 6 characters and maximum 64 characters , invalid email address , password containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) character")
	})

	t.Run("Error email required", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		_, err := u.RegisterUser(context.Background(), model.Users{
			FullName:     "Full name and last name",
			PhoneNumber:  "+628781122333",
			UserPassword: "AB781#kec",
		})

		assert.Error(t, err, "email is required , invalid email address")
	})

	t.Run("Error get email by user repo", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "c",
			PhoneNumber:  "d",
			UserPassword: "e",
		}, errors.New("failed")).Once()

		_, err := u.RegisterUser(context.Background(), model.Users{
			FullName:     "Full name and last name",
			PhoneNumber:  "+628781122333",
			UserPassword: "AB781#kec",
			Email:        "abc@def.com",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error email already exist", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "c",
			PhoneNumber:  "d",
			UserPassword: "e",
		}, nil).Once()

		_, err := u.RegisterUser(context.Background(), model.Users{
			FullName:     "Full name and last name",
			PhoneNumber:  "+628781122333",
			UserPassword: "AB781#kec",
			Email:        "abc@def.com",
		})

		assert.Error(t, err, "conflict, data already exist")
	})

	t.Run("Error repo insert user", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{}, nil).Once()

		mockRepo.On("InsertUser", mock.Anything, mock.Anything).Return("", errors.New("failed")).Once()

		_, err := u.RegisterUser(context.Background(), model.Users{
			FullName:     "Full name and last name",
			PhoneNumber:  "+628781122333",
			UserPassword: "AB781#kec",
			Email:        "abc@def.com",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := usecase.NewAuthUsecase(mockRepo)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.Users{}, nil).Once()

		mockRepo.On("InsertUser", mock.Anything, mock.Anything).Return("a", nil).Once()

		_, err := u.RegisterUser(context.Background(), model.Users{
			FullName:     "Full name and last name",
			PhoneNumber:  "+628781122333",
			UserPassword: "AB781#kec",
			Email:        "abc@def.com",
		})

		assert.NilError(t, err)
	})
}
