package delivery

import (
	"SmartLink_Project/domain"
	_middleware "SmartLink_Project/features/common"

	"github.com/labstack/echo/v4"
)

func RouteLayanan(e *echo.Echo, layanan domain.LayananHandler) {
	e.POST("/layanan", layanan.Add(), _middleware.JWTMiddleware())
}
