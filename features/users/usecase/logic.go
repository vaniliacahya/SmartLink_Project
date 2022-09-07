package usecase

import (
	"SmartLink_Project/domain"
	"SmartLink_Project/features/users/data"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(ud domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
		validate: v,
	}
}

func (uuc *userUseCase) RegisterUser(newuser domain.User) int {
	var user = data.FromModel(newuser)
	validError := uuc.validate.Struct(user)
	if validError != nil {
		log.Println("validation error : ", validError)
		return 400
	}

	if len(newuser.Nama) > 50 {
		log.Println("nama max 50 karakter")
		return 400
	} else if len(newuser.Username) > 15 {
		log.Println("username max 50 karakter")
		return 400
	} else if len(newuser.Telepon) > 15 {
		log.Println("telepon max 50 karakter")
		return 400
	}

	usernameformat := regexp.MustCompile("^[A-Za-z0-9]*$")
	cekusername := usernameformat.MatchString(newuser.Username)
	if !cekusername {
		log.Println("username only contain letter and number")
		return 400
	}

	teleponformat := regexp.MustCompile("^[0-9]*$")
	cektelepon := teleponformat.MatchString(newuser.Telepon)
	if !cektelepon {
		log.Println("telepon only contain number")
		return 400
	}

	duplicate := uuc.userData.CheckDuplicate(user.ToModel())
	if duplicate {
		log.Println("duplicated data")
		return 400
	}

	hashed, errorhash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorhash != nil {
		log.Println("can't hash: ", errorhash)
		return 500
	}

	user.Password = string(hashed)

	regis := uuc.userData.RegisterUserData(user.ToModel())
	if regis.ID == 0 {
		log.Println("failed insert data")
		return 500
	}

	return 200
}
