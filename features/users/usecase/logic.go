package usecase

import (
	"SmartLink_Project/domain"
	"SmartLink_Project/features/common"
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

func (uuc *userUseCase) RegisterUser(newuser domain.User) (int, string) {
	var user = data.FromModel(newuser)
	validError := uuc.validate.Struct(user)
	if validError != nil {
		log.Println("error validasi : ", validError)
		return 400, "error validasi"
	}

	if len(user.Nama) > 50 {
		log.Println("nama maksimum 50 karakter")
		return 400, "nama maksimum 50 karakter"
	}

	usernameformat := regexp.MustCompile("^[A-Za-z0-9]*$")
	cekusername := usernameformat.MatchString(user.Username)
	if len(user.Username) > 15 {
		log.Println("username maksimum 15 karakter")
		return 400, "username maksimum 15 karakter"
	} else if !cekusername {
		log.Println("username hanya boleh huruf dan angka")
		return 400, "username hanya boleh huruf dan angka"
	}

	teleponformat := regexp.MustCompile("^[0-9]*$")
	cektelepon := teleponformat.MatchString(user.Telepon)
	if len(user.Telepon) > 15 {
		log.Println("telepon maksimum 15 karakter")
		return 400, "telepon maksimum 15 karakter"
	} else if !cektelepon {
		log.Println("telepon hanya boleh angka")
		return 400, "telepon hanya boleh angka"
	}

	duplicate := uuc.userData.CheckDuplicate(user.ToModel())
	if duplicate {
		log.Println("username sudah terpakai, silahkan pilih username yang lain")
		return 400, "username sudah terpakai, silahkan pilih username yang lain"
	}

	hashed, errorhash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorhash != nil {
		log.Println("gagal hash password: ", errorhash)
		return 500, "gagal hash password"
	}

	user.Password = string(hashed)

	regis := uuc.userData.RegisterUserData(user.ToModel())
	if regis.ID == 0 {
		log.Println("gagal insert data")
		return 500, "gagal insert data"
	}

	return 200, "berhasil terdaftar"
}

func (uuc *userUseCase) LoginUser(datalogin domain.User) (domain.User, int, string) {
	userdata := uuc.userData.LoginUserData(datalogin)
	if userdata.ID == 0 {
		log.Println("gagal menemukan data")
		return domain.User{}, 400, ""
	}

	hashpw := uuc.userData.GetPasswordData(userdata.Username)
	if hashpw == "" {
		log.Println("gagal ambil data")
		return domain.User{}, 500, ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(datalogin.Password))
	if err != nil {
		log.Println("password salah")
		return domain.User{}, 400, ""
	}

	token := common.GenerateToken(userdata.UserID)

	return userdata, 200, token
}
