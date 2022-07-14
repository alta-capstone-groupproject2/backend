package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RemoveTrailingSlash() echo.MiddlewareFunc{
	return middleware.RemoveTrailingSlash()
}