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

func (u *AuthUsecase) Login(ctx context.Context, req model.Users) (res model.UserToken, err error) {

	invalidMessages, isValid := isLoginValid(model.Users{
		UserPassword: req.UserPassword,
		Email:        req.Email,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		return res, errors.New(errorMsg)
	}

	loginData, err := u.authRepo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return
	}

	if loginData.Id == "" {
		return res, constant.ErrNotFound
	}

	err = utils.CheckPasswordHash(req.UserPassword, loginData.UserPassword)
	if err != nil {
		err = u.authRepo.InsertUserLog(ctx, model.UserLogs{
			UserId:       loginData.Id,
			IsSuccess:    false,
			LoginMessage: constant.PasswordWrong,
		})

		if err != nil {
			return
		}

		return res, errors.New(constant.PasswordWrong)
	}

	token, err := helper.GenerateTokenAndRefreshToken(loginData)
	if err != nil {
		return
	}

	err = u.authRepo.InsertUserLog(ctx, model.UserLogs{
		UserId:       loginData.Id,
		IsSuccess:    true,
		LoginMessage: "success",
	})

	if err != nil {
		return
	}

	err = u.authRepo.UpsertUserToken(ctx, model.UserToken{
		Id:                    loginData.Id,
		Token:                 token.Token,
		TokenExpiredAt:        token.TokenExpiredAt,
		RefreshToken:          token.RefreshToken,
		RefreshTokenExpiredAt: token.RefreshTokenExpiredAt,
	})

	if err != nil {
		return
	}

	res = model.UserToken{
		Id:                    loginData.Id,
		Token:                 token.Token,
		TokenExpiredAt:        token.TokenExpiredAt,
		RefreshToken:          token.RefreshToken,
		RefreshTokenExpiredAt: token.RefreshTokenExpiredAt,
	}

	return
}
