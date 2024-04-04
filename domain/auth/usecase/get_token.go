package usecase

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/helper"
	"context"
)

func (u *AuthUsecase) GetToken(ctx context.Context, refreshToken string) (res model.UserToken, err error) {

	dataToken, err := helper.GetDataFromToken(refreshToken)

	if err != nil {
		return res, constant.ErrForbidden
	}

	loginData, err := u.authRepo.GetUserById(ctx, dataToken.Id)

	if err != nil {
		return
	}

	if loginData.Id == "" {
		return res, constant.ErrNotFound
	}

	token, err := helper.GenerateToken(loginData)
	if err != nil {
		return
	}

	res = model.UserToken{
		Id:             loginData.Id,
		Token:          token.Token,
		TokenExpiredAt: token.TokenExpiredAt,
		RefreshToken:   refreshToken,
	}

	return
}
