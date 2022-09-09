package delivery

import "SmartLink_Project/domain"

type DataUser struct {
	UserID   string `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromModelLogin(data domain.User, token string) DataUser {
	return DataUser{
		UserID:   data.UserID,
		Nama:     data.Nama,
		Username: data.Username,
		Token:    token,
	}
}
