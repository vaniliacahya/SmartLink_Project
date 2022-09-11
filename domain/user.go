package domain

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int
	UserID   string
	Nama     string
	Username string
	Password string
	Telepon  string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newUser User, cost int) (int, string)
	LoginUser(userData User) (User, int, string)
}

type UserData interface {
	RegisterUserData(newUser User) User
	CheckDuplicate(newUser User) bool
	LoginUserData(userData User) User
	GetPasswordData(username string) string
}
