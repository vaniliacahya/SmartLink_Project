package usecase

import (
	"SmartLink_Project/domain"
	"SmartLink_Project/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{Nama: "", Username: "", Password: "", Telepon: ""}
	insertData2 := domain.User{Nama: "Vanilia Cahya Nugroho Vanilia Cahya Nugroho Vanilia Cahya Nugroho", Username: "vanili23", Password: "123", Telepon: "081249690397"}
	insertData3 := domain.User{Nama: "Vanilia Cahya Nugroho", Username: "vanili2345678910", Password: "123", Telepon: "081249690397"}
	insertData4 := domain.User{Nama: "Vanilia Cahya Nugroho", Username: "vanili234!", Password: "123", Telepon: "081249690397"}
	insertData5 := domain.User{Nama: "Vanilia Cahya Nugroho", Username: "vanili234", Password: "123", Telepon: "081249690397081249690397"}
	insertData6 := domain.User{Nama: "Vanilia Cahya Nugroho", Username: "vanili234", Password: "123", Telepon: "081249690397!"}
	insertData7 := domain.User{Nama: "Vanilia Cahya Nugroho", Username: "vanili234", Password: "123", Telepon: "081249690397"}
	returnData := domain.User{ID: 0, Nama: "", Username: "", Password: "", Telepon: ""}
	returnData2 := domain.User{ID: 1, Nama: "Vanilia Cahya Nugroho", Username: "vanili234", Password: "123", Telepon: "081249690397"}

	t.Run("Valid Error", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData, 10)

		assert.Equal(t, 400, int)
		assert.Equal(t, "error validasi", string)
		repo.AssertExpectations(t)
	})

	t.Run("Max nama karakter", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData2, 10)

		assert.Equal(t, 400, int)
		assert.Equal(t, "nama maksimum 50 karakter", string)
		repo.AssertExpectations(t)
	})

	t.Run("Max username karakter", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData3, 10)

		assert.Equal(t, 400, int)
		assert.Equal(t, "username maksimum 15 karakter", string)
		repo.AssertExpectations(t)
	})

	t.Run("Username hanya huruf dan angka", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData4, 10)

		assert.Equal(t, 400, int)
		assert.Equal(t, "username hanya boleh huruf dan angka", string)
		repo.AssertExpectations(t)
	})

	t.Run("Max telepon karakter", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData5, 10)

		assert.Equal(t, 400, int)
		assert.Equal(t, "telepon maksimum 15 karakter", string)
		repo.AssertExpectations(t)
	})

	t.Run("Telepon hanya boleh angka", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData6, 10)

		assert.Equal(t, 400, int)
		assert.Equal(t, "telepon hanya boleh angka", string)
		repo.AssertExpectations(t)
	})

	t.Run("Username sudah ada", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData7, 10)

		assert.Equal(t, 400, int)
		assert.Equal(t, "username sudah terpakai, silahkan pilih username yang lain", string)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal hash password", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData7, 40)

		assert.Equal(t, 500, int)
		assert.Equal(t, "gagal hash password", string)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal insert", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterUserData", mock.Anything).Return(returnData).Once()
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData7, 10)

		assert.Equal(t, 500, int)
		assert.Equal(t, "gagal insert data", string)
		repo.AssertExpectations(t)
	})

	t.Run("Sukses insert", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterUserData", mock.Anything).Return(returnData2).Once()
		srv := New(repo, validator.New())
		int, string := srv.RegisterUser(insertData7, 10)

		assert.Equal(t, 200, int)
		assert.Equal(t, "berhasil terdaftar", string)
		repo.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := domain.User{Nama: "Vanilia Cahya Nugroho", Username: "vanili123", Password: "12312", Telepon: "081249690397"}
	returnData := domain.User{}
	returnData2 := domain.User{ID: 0, UserID: "", Nama: "", Username: "", Password: "", Telepon: ""}
	returnData3 := domain.User{ID: 1, UserID: "USR001", Nama: "Vanilia Cahya Nugroho", Username: "vanili123", Password: "12312", Telepon: "081249690397"}

	t.Run("Gagal menemukan data", func(t *testing.T) {
		repo.On("LoginUserData", mock.Anything).Return(returnData2).Once()
		srv := New(repo, validator.New())
		user, int, string := srv.LoginUser(insertData)

		assert.Equal(t, returnData, user)
		assert.Equal(t, 400, int)
		assert.Equal(t, "", string)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal ambil data", func(t *testing.T) {
		repo.On("LoginUserData", mock.Anything).Return(returnData3).Once()
		repo.On("GetPasswordData", mock.Anything).Return("").Once()
		srv := New(repo, validator.New())
		user, int, string := srv.LoginUser(insertData)

		assert.Equal(t, returnData, user)
		assert.Equal(t, 500, int)
		assert.Equal(t, "", string)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal compare password", func(t *testing.T) {
		repo.On("LoginUserData", mock.Anything).Return(returnData3).Once()
		repo.On("GetPasswordData", mock.Anything).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlcklEIjoiVVNSMDAxIiwiZXhwIjoxNjYyODg5MzI5fQ").Once()
		srv := New(repo, validator.New())
		user, int, string := srv.LoginUser(insertData)

		assert.Equal(t, returnData, user)
		assert.Equal(t, 400, int)
		assert.Equal(t, "", string)
		repo.AssertExpectations(t)
	})

	t.Run("Succes", func(t *testing.T) {
		repo.On("LoginUserData", mock.Anything).Return(returnData3).Once()
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$DIhf9ePQvSY0s7oQmje6UuU70YMp2w8a0oftKflVGI68JOegXMGQG").Once()
		srv := New(repo, validator.New())
		user, int, string := srv.LoginUser(insertData)

		assert.Equal(t, returnData3, user)
		assert.Equal(t, 200, int)
		assert.NotEmpty(t, string)
		repo.AssertExpectations(t)
	})
}
