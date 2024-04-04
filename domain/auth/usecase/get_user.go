package usecase

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/helper"
	"context"
)

func (u *AuthUsecase) GetUser(ctx context.Context, token string) (res model.Users, err error) {

	dataToken, err := helper.GetDataFromToken(token)

	if err != nil {
		return res, constant.ErrForbidden
	}

	res, err = u.authRepo.GetUserById(ctx, dataToken.Id)

	if err != nil {
		return
	}

	if res.Id == "" {
		return res, constant.ErrNotFound
	}

	return
}
