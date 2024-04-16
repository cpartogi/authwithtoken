package http

import (
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"authwithtoken/schema/response"
	"strings"

	"github.com/labstack/echo/v4"
)

// GetUser godoc
// @Summary get user data
// @Description get user data
// @Tags Register
// @Accept  json
// @Produce  json
// @Param  X-Authorization header string true "X-Authorization" default(JWT token)
// @Success 200 {object} response.GetUserDataSwagger
// @Failure 403 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /auth/user [get]
func (h *AuthHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	cekHeader := c.Request().Header

	tokenString := strings.Replace(cekHeader.Get("X-Authorization"), "JWT ", "", -1)

	userData, err := h.authUsecase.GetUser(ctx, tokenString)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := response.UserData{
		Id:          userData.Id,
		FullName:    userData.FullName,
		Email:       userData.Email,
		PhoneNumber: userData.PhoneNumber,
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, data)
}
