package data

import "SmartLink_Project/domain"

type Layanan struct {
	ID        int     `gorm:"type:int(11);autoIncrement"`
	LayananID string  `gorm:"type:varchar(15);column:layanan_id"`
	Nama      string  `gorm:"type:varchar(50)" json:"nama" form:"nama" validate:"required"`
	Unit      string  `gorm:"type:varchar(3)" json:"unit" form:"unit" validate:"required"`
	Harga     float64 `gorm:"type:double(10,2)" json:"harga" form:"harga" validate:"required"`
	UserID    int     `gorm:"type:int(11)"`
	UserIDS   string  `gorm:"type:varchar(15)"`
	User      User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	ID     int
	UserID string
}

func FromModel(data domain.Layanan) Layanan {
	var res Layanan
	res.ID = int(data.ID)
	res.Nama = data.Nama
	res.Unit = data.Unit
	res.Harga = data.Harga
	res.UserID = data.UserID
	res.UserIDS = data.UserIDS

	return res
}

func (u *Layanan) ToModel() domain.Layanan {
	return domain.Layanan{
		ID:        int(u.ID),
		LayananID: u.LayananID,
		Nama:      u.Nama,
		Unit:      u.Unit,
		Harga:     u.Harga,
		UserID:    u.UserID,
		UserIDS:   u.UserIDS,
	}
}
