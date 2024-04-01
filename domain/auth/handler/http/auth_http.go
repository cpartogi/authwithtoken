package http

import (
	"authwithtoken/domain/auth"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase auth.AuthUsecaseInteface
}

func NewAuthHander(e *echo.Echo, us auth.AuthUsecaseInteface) {
	// handler := &AuthHandler{
	// 	authUsecase: us,
	// }

}
