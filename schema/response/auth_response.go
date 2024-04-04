package response

import "time"

type UserToken struct {
	Id                    string     `json:"id"`
	Token                 string     `json:"token,omitempty"`
	TokenExpiredAt        time.Time  `json:"tokenExpiredAt,omitempty"`
	RefreshToken          string     `json:"refreshToken,omitempty"`
	RefreshTokenExpiredAt *time.Time `json:"refreshTokenExpiredAt,omitempty"`
}
