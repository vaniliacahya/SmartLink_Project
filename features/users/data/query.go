package data

import (
	"SmartLink_Project/domain"
	"log"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) RegisterUserData(newuser domain.User) domain.User {
	var user = FromModel(newuser)
	rescreate := ud.db.Create(&user)

	if rescreate.Error != nil {
		log.Println("error create user", rescreate.Error)
		return domain.User{}
	}

	if rescreate.RowsAffected < 1 {
		log.Println("no rows effected", rescreate.Error)
		return domain.User{}
	}

	return user.ToModel()
}

func (ud *userData) CheckDuplicate(newuser domain.User) bool {
	var user User
	resfind := ud.db.Where("username = ?", newuser.Username).Find(&user)

	if resfind.RowsAffected > 0 {
		log.Println("duplicated username")
		return true
	}
	return false
}
