package delivery

import (
	"SmartLink_Project/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, user domain.UserHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.POST},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", user.Register())
	e.POST("/login", user.Login())
}
