package delivery

import "SmartLink_Project/domain"

type DataUser struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromModelLogin(data domain.User, token string) DataUser {
	return DataUser{
		ID:       data.ID,
		Nama:     data.Nama,
		Username: data.Username,
		Token:    token,
	}
}
