package api

import (
	"Software-Market-Go-API/pkg/customresponse"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetTokenFromHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		authHeader := ctx.Request().Header.Get("Authorization")

		if authHeader == "" {

			resp := new(customresponse.Response)
			resp.Success = false
			resp.Message = "no authorization header"
			resp.Result = nil

			return ctx.JSON(http.StatusBadRequest, resp)
		}

		auth := strings.SplitN(authHeader, " ", 2)
		if len(auth) != 2 || auth[0] != "Bearer" {

			resp := new(customresponse.Response)
			resp.Success = false
			resp.Message = "malformed authorization header"
			resp.Result = nil

			return ctx.JSON(http.StatusBadRequest, resp)
		}

		ctx.Set("token-string", auth[1])

		return next(ctx)
	}
}
