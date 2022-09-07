package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	udata "SmartLink_Project/features/users/data"
	udeliv "SmartLink_Project/features/users/delivery"
	ucase "SmartLink_Project/features/users/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := udata.New(db)
	userCase := ucase.New(userData, validator)
	userHandler := udeliv.New(userCase, userData)
	udeliv.RouteUser(e, userHandler)
}
