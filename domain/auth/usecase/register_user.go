package usecase

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"context"
	"errors"
	"strings"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

func (u *AuthUsecase) RegisterUser(ctx context.Context, req model.Users) (id string, err error) {

	invalidMessages, isValid := isDataValid(model.Users{
		FullName:     req.FullName,
		UserPassword: req.UserPassword,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		return "", errors.New(errorMsg)
	}

	//cek if email exist
	checkMail, err := u.authRepo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		if err != pg.ErrNoRows {
			return "", err
		}
	}

	if checkMail.Email != "" {
		return "", constant.ErrConflict
	}

	//insert to db
	req.Id = uuid.New().String()

	pHash, err := utils.HashPassword(req.UserPassword)
	if err != nil {
		return
	}

	req.UserPassword = pHash

	id, err = u.authRepo.InsertUser(ctx, req)

	if err != nil {
		return
	}

	return
}
