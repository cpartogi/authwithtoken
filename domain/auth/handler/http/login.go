package http

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"authwithtoken/schema/request"
	"authwithtoken/schema/response"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.Login

	c.Bind(&req)

	loginUser := model.Users{
		Email:        req.Email,
		UserPassword: req.Password,
	}

	login, err := h.authUsecase.Login(ctx, loginUser)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := response.UserToken{
		Id:                    login.Id,
		Token:                 login.Token,
		TokenExpiredAt:        login.TokenExpiredAt,
		RefreshToken:          login.RefreshToken,
		RefreshTokenExpiredAt: &login.RefreshTokenExpiredAt,
	}

	return utils.SuccessResponse(c, constant.SuccessAddData, data)
}
