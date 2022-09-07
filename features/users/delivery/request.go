package delivery

import (
	"SmartLink_Project/domain"
)

type UserFormat struct {
	Nama     string `gorm:"type:varchar(50)" json:"nama" form:"nama" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Telepon  string `json:"telepon" form:"telepon" validate:"required"`
}

func (uf *UserFormat) ToModel() domain.User {
	return domain.User{
		Nama:     uf.Nama,
		Username: uf.Username,
		Password: uf.Password,
		Telepon:  uf.Telepon,
	}
}
