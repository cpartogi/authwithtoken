package repo

import (
	"authwithtoken/domain/auth"
	"authwithtoken/domain/auth/model"
	"context"
	"fmt"
	"time"

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

	query := `INSERT INTO users (id, full_name, email, phone_number, user_password, created_by, created_at) values ('%s', '%s', '%s', '%s','%s', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.FullName, req.Email, req.PhoneNumber, req.UserPassword, req.Id)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
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

	return
}

func (r *AuthRepo) GetUserById(ctx context.Context, id string) (res model.Users, err error) {

	err = r.gopg.ModelContext(ctx, &res).Where("id=?", id).First()

	if err != nil {
		if err != pg.ErrNoRows {
			return
		}
	}

	return
}

func (r *AuthRepo) InsertUserLog(ctx context.Context, req model.UserLogs) (err error) {

	req.Id = uuid.New().String()

	query := `INSERT INTO user_logs (id, user_id, is_success, login_message, created_at) values ('%s', '%s', '%t', '%s', now())`
	query = fmt.Sprintf(query, req.Id, req.UserId, req.IsSuccess, req.LoginMessage)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		return
	}

	return

}

func (r *AuthRepo) UpsertUserToken(ctx context.Context, req model.UserToken) (err error) {

	// check data exist
	err = r.gopg.ModelContext(ctx, &model.UserToken{}).Where("user_id=?", req.Id).First()
	var query string
	id := uuid.New().String()

	tokenExpiredAt := req.TokenExpiredAt.UTC().Format(time.RFC3339)
	refreshTokenExpiredAt := req.RefreshTokenExpiredAt.UTC().Format(time.RFC3339)

	if err != nil {
		if err == pg.ErrNoRows {
			query = `INSERT INTO user_tokens (id, user_id, token, token_expired_at, refresh_token, refresh_token_expired_at, created_at) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', now()) `
			query = fmt.Sprintf(query, id, req.Id, req.Token, tokenExpiredAt, req.RefreshToken, refreshTokenExpiredAt)
		} else {
			return
		}
	} else {

		if req.RefreshToken != "" {
			query = `UPDATE user_tokens SET  token = '%s', token_expired_at = '%s', refresh_token = '%s',  refresh_token_expired_at = '%s', updated_at = now()  WHERE user_id = '%s' `
			query = fmt.Sprintf(query, req.Token, tokenExpiredAt, req.RefreshToken, refreshTokenExpiredAt, req.Id)
		} else {
			query = `UPDATE user_tokens SET  token = '%s', token_expired_at = '%s', updated_at = now()  WHERE user_id = '%s' `
			query = fmt.Sprintf(query, req.Token, tokenExpiredAt, req.Id)
		}
	}

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		return
	}

	return
}

func (r *AuthRepo) UpdateUser(ctx context.Context, req model.Users) (err error) {

	query := `UPDATE users SET full_name  = '%s', phone_number = '%s', user_password  = '%s', updated_at = now(), updated_by = '%s' WHERE id = '%s' `
	query = fmt.Sprintf(query, req.FullName, req.PhoneNumber, req.UserPassword, req.Id, req.Id)

	_, err = r.gopg.ExecContext(ctx, query)

	if err != nil {
		return
	}

	return
}
