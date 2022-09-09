package delivery

import (
	"SmartLink_Project/domain"
	// _middleware "SmartLink_Project/features/common"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, user domain.UserHandler) {
	e.POST("/register", user.Register())
	e.POST("/login", user.Login())
}
