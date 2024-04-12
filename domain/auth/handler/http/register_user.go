package http

import (
	"authwithtoken/lib/pkg/utils"

	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/schema/request"
	"authwithtoken/schema/response"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.RegisterUser

	c.Bind(&req)

	registerUser := model.Users{
		FullName:     req.FullName,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		UserPassword: req.Password,
	}

	register, err := h.authUsecase.RegisterUser(ctx, registerUser)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := response.UserToken{
		Id: register,
	}

	return utils.SuccessResponse(c, constant.SuccessAddData, data)
}
