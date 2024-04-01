package usecase

import (
	"authwithtoken/domain/auth/model"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
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

	//insert to db
	req.Id = uuid.New().String()

	id, err = u.authRepo.InsertUser(ctx, req)

	if err != nil {
		return "", err
	}

	return
}
