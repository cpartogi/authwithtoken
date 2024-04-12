package middleware

import (
	"authwithtoken/lib/constant"
	"authwithtoken/lib/pkg/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cekHeader := c.Request().Header

		if cekHeader.Get("X-Authorization") == "" {
			return utils.ErrorForbidden(c, constant.ErrForbidden, "")
		}

		if !strings.Contains(cekHeader.Get("X-Authorization"), "JWT") {
			return utils.ErrorForbidden(c, constant.ErrForbidden, "")
		}

		return next(c)
	}
}

func AuthRefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cekHeader := c.Request().Header

		if cekHeader.Get("Authorization") == "" {
			return utils.ErrorForbidden(c, constant.ErrForbidden, "")
		}

		if !strings.Contains(cekHeader.Get("Authorization"), "Bearer") {
			return utils.ErrorForbidden(c, constant.ErrForbidden, "")
		}

		return next(c)
	}
}
