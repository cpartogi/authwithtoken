package http

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"authwithtoken/schema/request"

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

	id, err := h.authUsecase.Login(ctx, loginUser)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessAddData, id)
}
