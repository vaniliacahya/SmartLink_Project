package usecase

import (
	"SmartLink_Project/domain"
	"SmartLink_Project/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddLayanan(t *testing.T) {
	repo := new(mocks.LayananData)
	insertData := domain.Layanan{Nama: "", Unit: "", Harga: 0}
	insertData2 := domain.Layanan{Nama: "Cuci Kering Basah Tidak Usah Disetrika Pake Parfum Premium Bahan Knit Ekstra Hati-Hati", Unit: "kg", Harga: 12000}
	insertData3 := domain.Layanan{Nama: "Cuci Kering", Unit: "pc", Harga: 12000}
	insertData4 := domain.Layanan{Nama: "Cuci Kering", Unit: "pcs", Harga: 12000}
	returnData := domain.Layanan{}
	returnData2 := domain.Layanan{ID: 0, LayananID: "", Nama: "", Unit: "", Harga: 0, UserID: 0, UserIDS: ""}
	returnData3 := domain.Layanan{ID: 1, LayananID: "LYN001", Nama: "Cuci Kering", Unit: "pcs", Harga: 12000, UserID: 1, UserIDS: "USR001"}

	t.Run("Valid Error", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, layanan := srv.AddLayanan(insertData)

		assert.Equal(t, 400, int)
		assert.Equal(t, returnData, layanan)
		repo.AssertExpectations(t)
	})

	t.Run("Max nama karakter", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, layanan := srv.AddLayanan(insertData2)

		assert.Equal(t, 400, int)
		assert.Equal(t, returnData, layanan)
		repo.AssertExpectations(t)
	})

	t.Run("Unit diluar kg, pcs, cm, m2 ", func(t *testing.T) {
		srv := New(repo, validator.New())
		int, layanan := srv.AddLayanan(insertData3)

		assert.Equal(t, 400, int)
		assert.Equal(t, returnData, layanan)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Insert Data", func(t *testing.T) {
		repo.On("AddLayananData", mock.Anything).Return(returnData2).Once()
		srv := New(repo, validator.New())
		int, layanan := srv.AddLayanan(insertData4)

		assert.Equal(t, 500, int)
		assert.Equal(t, returnData, layanan)
		repo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		repo.On("AddLayananData", mock.Anything).Return(returnData3).Once()
		srv := New(repo, validator.New())
		int, layanan := srv.AddLayanan(insertData4)

		assert.Equal(t, 200, int)
		assert.Equal(t, returnData3, layanan)
		repo.AssertExpectations(t)
	})
}
