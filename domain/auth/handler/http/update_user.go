package http

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"authwithtoken/schema/request"
	"authwithtoken/schema/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	cekHeader := c.Request().Header

	if cekHeader.Get("X-Authorization") == "" {
		return utils.ErrorForbidden(c, constant.ErrForbidden, "")
	}

	if !strings.Contains(cekHeader.Get("X-Authorization"), "JWT") {
		return utils.ErrorForbidden(c, constant.ErrForbidden, "")
	}

	tokenString := strings.Replace(cekHeader.Get("X-Authorization"), "JWT ", "", -1)

	var req request.UpdateUser
	c.Bind(&req)

	updateUser := model.Users{
		FullName:     req.FullName,
		PhoneNumber:  req.PhoneNumber,
		UserPassword: req.Password,
	}

	userId, err := h.authUsecase.UpdateUser(ctx, tokenString, updateUser)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := response.UserUpdate{
		Id: userId,
	}

	return utils.SuccessResponse(c, constant.SuccessUpdateData, data)
}
