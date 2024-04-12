package auth

import (
	"authwithtoken/domain/auth/model"
	"context"
)

type AuthUsecaseInteface interface {
	RegisterUser(ctx context.Context, req model.Users) (id string, err error)
	Login(ctx context.Context, req model.Users) (res model.UserToken, err error)
	GetToken(ctx context.Context, refreshToken string) (res model.UserToken, err error)
	GetUser(ctx context.Context, token string) (res model.Users, err error)
	UpdateUser(ctx context.Context, token string, req model.Users) (id string, err error)
}
