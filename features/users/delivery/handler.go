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
		cost := 10

		if bind != nil {
			log.Println("can't bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"status":  "internal server error",
				"message": "gagal terdaftar",
			})
		}

		code, message := uh.userUseCase.RegisterUser(newuser.ToModelRegis(), cost)

		if code == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    code,
				"status":  "bad request",
				"message": message,
			})
		}

		if code == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    code,
				"status":  "internal server error",
				"message": message,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    code,
			"status":  "success",
			"message": message,
		})

	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginuser LoginFormat
		bind := c.Bind(&loginuser)

		if bind != nil {
			log.Println("can't bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":   500,
				"status": "internal server error",
				"data":   "{}",
			})
		}

		datauser, code, token := uh.userUseCase.LoginUser(loginuser.ToModelLogin())
		data := FromModelLogin(datauser, token)

		if code == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":   code,
				"status": "bad request",
				"data":   data,
			})
		}

		if code == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":   code,
				"status": "internal server error",
				"data":   data,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":   code,
			"status": "success",
			"data":   data,
		})
	}
}
