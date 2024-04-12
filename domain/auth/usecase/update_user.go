package usecase

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/helper"
	"authwithtoken/lib/pkg/utils"
	"context"
	"errors"
	"strings"
)

func (u *AuthUsecase) UpdateUser(ctx context.Context, token string, req model.Users) (id string, err error) {

	dataToken, err := helper.GetDataFromToken(token)

	if err != nil {
		return "", constant.ErrForbidden
	}

	res, err := u.authRepo.GetUserById(ctx, dataToken.Id)

	if err != nil {
		return
	}

	if res.Id == "" {
		return "", constant.ErrNotFound
	}

	invalidMessages, isValid := isDataValidForUpdate(model.Users{
		FullName:     req.FullName,
		UserPassword: req.UserPassword,
		PhoneNumber:  req.PhoneNumber,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		return "", errors.New(errorMsg)
	}

	pHash, err := utils.HashPassword(req.UserPassword)
	if err != nil {
		return
	}
	req.UserPassword = pHash
	req.Id = res.Id

	err = u.authRepo.UpdateUser(ctx, req)

	if err != nil {
		return
	}

	return res.Id, nil
}
