package helper

import (
	"time"

	"authwithtoken/domain/auth/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var jwtSigningMethod = jwt.SigningMethodHS256

func GenerateToken(login model.Users) (tokenData model.UserToken, err error) {
	tokenData, err = GenerateJWT(login)
	if err != nil {
		return
	}

	return
}

func GenerateJWT(user model.Users) (tokenData model.UserToken, err error) {

	expiredIn := viper.GetInt(`token.expired_in_minutes`)
	exp := time.Now().UTC().Add(time.Duration(expiredIn) * time.Minute)
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

type JWTToken struct {
	Id string `json:"id"`
	jwt.StandardClaims
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
