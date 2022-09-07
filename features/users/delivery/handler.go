package delivery

import (
	"SmartLink_Project/domain"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUseCase domain.UserUseCase
	userData    domain.UserData
}

func New(uuc domain.UserUseCase, ud domain.UserData) domain.UserHandler {
	return &userHandler{
		userUseCase: uuc,
		userData:    ud,
	}
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		bind := c.Bind(&newuser)

		if bind != nil {
			log.Println("can't bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"status":  "internal server error",
				"message": "gagal terdaftar",
			})
		}

		code := uh.userUseCase.RegisterUser(newuser.ToModel())

		if code == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    code,
				"status":  "bad request",
				"message": "gagal terdaftar",
			})
		}

		if code == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    code,
				"status":  "internal server error",
				"message": "gagal terdaftar",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    code,
			"status":  "success",
			"message": "berhasil terdaftar",
		})

	}
}
