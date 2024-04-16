package response

import "time"

type UserToken struct {
	Id                    string     `json:"id"`
	Token                 string     `json:"token,omitempty"`
	TokenExpiredAt        *time.Time `json:"tokenExpiredAt,omitempty"`
	RefreshToken          string     `json:"refreshToken,omitempty"`
	RefreshTokenExpiredAt *time.Time `json:"refreshTokenExpiredAt,omitempty"`
}

type UserData struct {
	Id          string `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserUpdate struct {
	Id string `json:"id"`
}

type RegisterUserSuccessSwagger struct {
	Base
	Data UserUpdate `json:"data"`
}

type LoginSuccessSwagger struct {
	Base
	Data UserToken `json:"data"`
}

type GetTokenResponse struct {
	Id             string     `json:"id"`
	Token          string     `json:"token,omitempty"`
	TokenExpiredAt *time.Time `json:"tokenExpiredAt,omitempty"`
}

type GetTokenSwagger struct {
	Base
	Data GetTokenResponse `json:"data"`
}

type GetUserDataSwagger struct {
	Base
	Data UserData `json:"data"`
}
