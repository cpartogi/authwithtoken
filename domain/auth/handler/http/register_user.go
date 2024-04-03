package http

import (
	"authwithtoken/lib/pkg/utils"

	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/schema/request"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.RegisterUser

	c.Bind(&req)

	registerUser := model.Users{
		Id:           req.Id,
		FullName:     req.FullName,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		UserPassword: req.Password,
	}

	id, err := h.authUsecase.RegisterUser(ctx, registerUser)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessAddData, id)
}
