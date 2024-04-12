package auth

import (
	"authwithtoken/domain/auth/model"
	"context"
)

type AuthRepoInterface interface {
	InsertUser(ctx context.Context, req model.Users) (userId string, err error)
	GetUserByEmail(ctx context.Context, email string) (res model.Users, err error)
	GetUserById(ctx context.Context, id string) (res model.Users, err error)
	InsertUserLog(ctx context.Context, req model.UserLogs) (err error)
	UpsertUserToken(ctx context.Context, req model.UserToken) (err error)
	UpdateUser(ctx context.Context, req model.Users) (err error)
}
