package usecase

import "authwithtoken/domain/auth"

type AuthUsecase struct {
	authRepo auth.AuthRepoInterface
}

func NewAuthUsecase(authRepo auth.AuthRepoInterface) auth.AuthUsecaseInteface {
	return &AuthUsecase{
		authRepo: authRepo,
	}
}
