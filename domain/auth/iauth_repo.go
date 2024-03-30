package auth

import (
	"authwithtoken/domain/auth/model"
	"context"
)

type AuthRepoInterface interface {
	InsertUser(ctx context.Context, req model.Users) (userId string, err error)
}
