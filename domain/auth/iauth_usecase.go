package auth

import (
	"authwithtoken/domain/auth/model"
	"context"
)

type AuthUsecaseInteface interface {
	RegisterUser(ctx context.Context, req model.Users) (id string, err error)
	Login(ctx context.Context, req model.Users) (res model.UserToken, err error)
}
