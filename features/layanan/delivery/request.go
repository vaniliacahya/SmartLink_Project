package delivery

import (
	"SmartLink_Project/domain"
)

type LayananFormatString struct {
	Nama  string `json:"nama" form:"nama" validate:"required"`
	Unit  string `json:"unit" form:"unit" validate:"required"`
	Harga string `json:"harga" form:"harga" validate:"required"`
}

type LayananFormatFloat struct {
	Nama   string  `json:"nama" form:"nama" validate:"required"`
	Unit   string  `json:"unit" form:"unit" validate:"required"`
	Harga  float64 `json:"harga" form:"harga" validate:"required"`
	UserID string
}

func (lf *LayananFormatFloat) ToModelLayanan() domain.Layanan {
	return domain.Layanan{
		Nama:   lf.Nama,
		Unit:   lf.Unit,
		Harga:  lf.Harga,
		UserID: lf.UserID,
	}
}
