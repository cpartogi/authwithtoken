package helper

import (
	"time"

	"authwithtoken/domain/auth/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var jwtSigningMethod = jwt.SigningMethodHS256

type JWTToken struct {
	Id          string `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	jwt.StandardClaims
}

func GenerateToken(login model.Users) (tokenData model.UserToken, err error) {
	tokenData, err = GenerateJWT(login)
	if err != nil {
		return
	}

	refreshTokenData, err := GenerateRefresh(login)

	if err != nil {
		return
	}

	tokenData.RefreshToken = refreshTokenData.RefreshToken
	tokenData.RefreshTokenExpiredAt = refreshTokenData.RefreshTokenExpiredAt

	return
}

func GenerateJWT(user model.Users) (tokenData model.UserToken, err error) {

	exp := time.Now().UTC().Add(viper.GetDuration("token.expiry"))
	claims := JWTToken{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().UTC().Unix(),
			ExpiresAt: exp.Unix(),
		},
		Id:          user.Id,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString([]byte(viper.GetString(`token.key`)))
	if err != nil {
		return
	}

	tokenData = model.UserToken{
		Id:             user.Id,
		Token:          signedToken,
		TokenExpiredAt: exp,
	}

	return
}

func GenerateRefresh(user model.Users) (refreshTokenData model.UserToken, err error) {

	exp := time.Now().UTC().Add(viper.GetDuration("token.refresh_token_expiry"))
	claims := JWTToken{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().UTC().Unix(),
			ExpiresAt: exp.Unix(),
		},
		Id: user.Id,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString([]byte(viper.GetString("token.key")))
	if err != nil {
		return
	}

	refreshTokenData = model.UserToken{
		RefreshToken:          signedToken,
		RefreshTokenExpiredAt: exp,
	}

	return
}

func ParseToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, jwt.ErrInvalidKeyType
		}

		return []byte(viper.GetString(`token.key`)), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return token, err
}

func GetDataFromToken(token string) (userData model.Users, err error) {

	resToken, err := ParseToken(token)
	if err != nil {
		return
	}

	claims := resToken.Claims.(jwt.MapClaims)

	userData.Id = claims["id"].(string)

	return
}
