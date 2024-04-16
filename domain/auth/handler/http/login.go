package http

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"authwithtoken/schema/request"
	"authwithtoken/schema/response"

	"github.com/labstack/echo/v4"
)

// UserLogin godoc
// @Summary user login
// @Description user login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body request.Login true "Request Body"
// @Success 200 {object} response.LoginSuccessSwagger
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /auth/login [post]
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
		TokenExpiredAt:        &login.TokenExpiredAt,
		RefreshToken:          login.RefreshToken,
		RefreshTokenExpiredAt: &login.RefreshTokenExpiredAt,
	}

	return utils.SuccessResponse(c, constant.SuccessLogin, data)
}
