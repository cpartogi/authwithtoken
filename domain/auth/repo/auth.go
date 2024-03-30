package repo

import (
	"authwithtoken/domain/auth"
	"authwithtoken/domain/auth/model"
	"context"
	"fmt"

	"github.com/go-pg/pg"
)

type AuthRepo struct {
	gopg *pg.DB
}

func NewAuthRepo(gopg *pg.DB) auth.AuthRepoInterface {
	return &AuthRepo{
		gopg: gopg,
	}
}

func (r *AuthRepo) InsertUser(ctx context.Context, req model.Users) (userId string, err error) {

	tx, err := r.gopg.Begin()

	if err != nil {
		return
	}

	query := `INSERT INTO users (id, full_name, email, phone_number, user_password, created_by, created_at) values ('%s', '%s', '%s', '%s','%s', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.FullName, req.Email, req.PhoneNumber, req.Password, req.Id)

	_, err = tx.ExecContext(ctx, query)

	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}

	userId = req.Id

	return
}
