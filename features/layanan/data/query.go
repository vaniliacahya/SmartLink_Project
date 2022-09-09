package data

import (
	"SmartLink_Project/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type layananData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.LayananData {
	return &layananData{
		db: db,
	}
}

func (ld *layananData) AddLayananData(newlayanan domain.Layanan) domain.Layanan {
	var layanan = FromModel(newlayanan)
	rescreate := ld.db.Create(&layanan)

	if rescreate.Error != nil {
		log.Println("error create layanan")
		return domain.Layanan{}
	}

	if rescreate.RowsAffected < 1 {
		log.Println("no rows effected")
		return domain.Layanan{}
	}

	layanancode := "LYN00"
	layanan.LayananID = layanancode + fmt.Sprint(layanan.ID)

	resupdate := ld.db.Model(&layanan).Where("id = ?", layanan.ID).Update("layanan_id", layanan.LayananID)

	if resupdate.Error != nil {
		log.Println("error update layanan")
		return domain.Layanan{}
	}

	if resupdate.RowsAffected < 1 {
		log.Println("no rows effected")
		return domain.Layanan{}
	}
	return layanan.ToModel()
}
