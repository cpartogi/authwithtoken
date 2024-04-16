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

// UpdateUser godoc
// @Summary update user data
// @Description update user data
// @Tags Register
// @Accept  json
// @Produce  json
// @Param  X-Authorization header string true "X-Authorization" default(JWT token)
// @Param request body request.UpdateUser true "Request Body"
// @Success 200 {object} response.RegisterUserSuccessSwagger
// @Failure 400 {object} response.Base
// @Failure 403 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /auth/user [put]
func (h *AuthHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	cekHeader := c.Request().Header

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
