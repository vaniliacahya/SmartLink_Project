package data

import (
	"SmartLink_Project/domain"
	"fmt"
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
		log.Println("error create user")
		return domain.User{}
	}

	if rescreate.RowsAffected < 1 {
		log.Println("no rows effected")
		return domain.User{}
	}

	usercode := "USR00"
	user.UserID = usercode + fmt.Sprint(user.ID)

	resupdate := ud.db.Model(&user).Where("id = ?", user.ID).Update("user_id", user.UserID)

	if resupdate.Error != nil {
		log.Println("error update user")
		return domain.User{}
	}

	if resupdate.RowsAffected < 1 {
		log.Println("no rows effected")
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

func (ud *userData) LoginUserData(logindata domain.User) domain.User {
	var user = FromModel(logindata)
	resfirst := ud.db.Where("username = ?", logindata.Username).First(&user)

	if resfirst.Error != nil {
		log.Println("error get data user")
		return domain.User{}
	}

	return user.ToModel()
}

func (ud *userData) GetPasswordData(username string) string {
	var user User
	resfirst := ud.db.Where("username = ?", username).First(&user)

	if resfirst.Error != nil {
		log.Println("error get data user")
		return ""
	}

	return user.Password
}
