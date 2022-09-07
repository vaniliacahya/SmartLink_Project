package domain

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int
	Nama     string
	Username string
	Password string
	Telepon  string
}

type UserHandler interface {
	Register() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newUser User) int
}

type UserData interface {
	RegisterUserData(newUser User) User
	CheckDuplicate(newUser User) bool
}
