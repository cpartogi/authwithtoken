package repo

import (
	"authwithtoken/domain/auth"
	"authwithtoken/domain/auth/model"
	"context"
	"fmt"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
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
	query = fmt.Sprintf(query, req.Id, req.FullName, req.Email, req.PhoneNumber, req.UserPassword, req.Id)

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

func (r *AuthRepo) GetUserByEmail(ctx context.Context, email string) (res model.Users, err error) {

	err = r.gopg.ModelContext(ctx, &res).Where("email=?", email).First()

	if err != nil {
		if err != pg.ErrNoRows {
			return
		}
	}

	return res, nil
}

func (r *AuthRepo) GetUserById(ctx context.Context, id string) (res model.Users, err error) {

	err = r.gopg.ModelContext(ctx, &res).Where("id=?", id).First()

	if err != nil {
		if err != pg.ErrNoRows {
			return
		}
	}

	return res, nil
}

func (r *AuthRepo) InsertUserLog(ctx context.Context, req model.UserLogs) (err error) {

	tx, err := r.gopg.Begin()

	if err != nil {
		return
	}

	req.Id = uuid.New().String()

	query := `INSERT INTO user_logs (id, user_id, is_success, login_message, created_at) values ('%s', '%s', '%t', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.UserId, req.IsSuccess, req.LoginMessage)

	_, err = tx.ExecContext(ctx, query)

	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return
	}

	return

}
