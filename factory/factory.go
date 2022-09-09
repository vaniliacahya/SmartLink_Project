package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	udata "SmartLink_Project/features/users/data"
	udeliv "SmartLink_Project/features/users/delivery"
	ucase "SmartLink_Project/features/users/usecase"

	ldata "SmartLink_Project/features/layanan/data"
	ldeliv "SmartLink_Project/features/layanan/delivery"
	lcase "SmartLink_Project/features/layanan/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := udata.New(db)
	userCase := ucase.New(userData, validator)
	userHandler := udeliv.New(userCase, userData)
	udeliv.RouteUser(e, userHandler)

	layananData := ldata.New(db)
	layananCase := lcase.New(layananData, validator)
	layananHandler := ldeliv.New(layananCase, layananData)
	ldeliv.RouteLayanan(e, layananHandler)
}
