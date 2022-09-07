package data

import "SmartLink_Project/domain"

type User struct {
	ID       int    `gorm:"autoIncrement"`
	Nama     string `gorm:"type:varchar(50)" json:"nama" form:"nama" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Telepon  string `json:"telepon" form:"telepon" validate:"required"`
}

func FromModel(data domain.User) User {
	var res User
	res.ID = int(data.ID)
	res.Nama = data.Nama
	res.Username = data.Username
	res.Password = data.Password
	res.Telepon = data.Telepon

	return res
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:       int(u.ID),
		Nama:     u.Nama,
		Username: u.Username,
		Password: u.Password,
		Telepon:  u.Telepon,
	}
}
