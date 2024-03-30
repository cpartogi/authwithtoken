package usecase

import (
	"authwithtoken/domain/auth/model"
	"context"
	"errors"
	"strings"
)

func (u *AuthUsecase) RegisterUser(ctx context.Context, req model.Users) (id string, err error) {

	invalidMessages, isValid := isDataValid(model.Users{
		FullName:    req.FullName,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		return "", errors.New(errorMsg)
	}

	return
}
