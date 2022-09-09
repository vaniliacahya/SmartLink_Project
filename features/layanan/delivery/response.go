package delivery

import "SmartLink_Project/domain"

type DataLayanan struct {
	LayananID string  `json:"id"`
	Nama      string  `json:"nama"`
	Unit      string  `json:"unit"`
	Harga     float64 `json:"harga"`
	UserID    string  `json:"user_id"`
}

func FromModelLayanan(data domain.Layanan) DataLayanan {
	return DataLayanan{
		LayananID: data.LayananID,
		Nama:      data.Nama,
		Unit:      data.Unit,
		Harga:     data.Harga,
		UserID:    data.UserID,
	}
}
