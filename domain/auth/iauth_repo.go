package auth

import (
	"authwithtoken/domain/auth/model"
	"context"
)

type AuthRepoInterface interface {
	InsertUser(ctx context.Context, req model.Users) (userId string, err error)
	GetUserByEmail(ctx context.Context, email string) (res model.Users, err error)
}
