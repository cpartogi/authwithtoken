package http

import (
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"authwithtoken/schema/response"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) GetToken(c echo.Context) error {
	ctx := c.Request().Context()

	cekHeader := c.Request().Header

	tokenString := strings.Replace(cekHeader.Get("Authorization"), "Bearer ", "", -1)

	userToken, err := h.authUsecase.GetToken(ctx, tokenString)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := response.UserToken{
		Id:             userToken.Id,
		Token:          userToken.Token,
		TokenExpiredAt: &userToken.TokenExpiredAt,
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, data)
}
