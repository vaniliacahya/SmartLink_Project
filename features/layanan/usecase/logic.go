package usecase

import (
	"SmartLink_Project/domain"
	"SmartLink_Project/features/layanan/data"
	"log"

	"github.com/go-playground/validator/v10"
)

type layananUseCase struct {
	layananData domain.LayananData
	validate    *validator.Validate
}

func New(ld domain.LayananData, v *validator.Validate) domain.LayananUseCase {
	return &layananUseCase{
		layananData: ld,
		validate:    v,
	}
}

func (luc *layananUseCase) AddLayanan(newlayanan domain.Layanan) (int, domain.Layanan) {
	var layanan = data.FromModel(newlayanan)
	validError := luc.validate.Struct(layanan)
	if validError != nil {
		log.Println("error validasi : ", validError)
		return 400, domain.Layanan{}
	}

	if len(layanan.Nama) > 50 {
		log.Println("nama maksimum 50 karakter")
		return 400, domain.Layanan{}
	}

	unit := []string{"kg", "pcs", "cm", "m2"}
	count := 0
	for _, v := range unit {
		if layanan.Unit == v {
			count = 1
			break
		}
	}

	if count != 1 {
		log.Println("unit hanya boleh kg, pcs, cm , m2")
		return 400, domain.Layanan{}
	}

	layanandata := luc.layananData.AddLayananData(layanan.ToModel())
	if layanandata.ID == 0 {
		log.Println("gagal insert data")
		return 500, domain.Layanan{}
	}

	return 200, layanandata
}
