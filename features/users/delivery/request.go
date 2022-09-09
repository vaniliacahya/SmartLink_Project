package delivery

import (
	"SmartLink_Project/domain"
)

type UserFormat struct {
	Nama     string `json:"nama" form:"nama" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Telepon  string `json:"telepon" form:"telepon" validate:"required"`
}

func (uf *UserFormat) ToModelRegis() domain.User {
	return domain.User{
		Nama:     uf.Nama,
		Username: uf.Username,
		Password: uf.Password,
		Telepon:  uf.Telepon,
	}
}

type LoginFormat struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (lf *LoginFormat) ToModelLogin() domain.User {
	return domain.User{
		Username: lf.Username,
		Password: lf.Password,
	}
}
