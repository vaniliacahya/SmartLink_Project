package delivery

import (
	"SmartLink_Project/domain"
	"SmartLink_Project/features/common"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type layananHandler struct {
	layananUseCase domain.LayananUseCase
	layananData    domain.LayananData
}

func New(luc domain.LayananUseCase, ld domain.LayananData) domain.LayananHandler {
	return &layananHandler{
		layananUseCase: luc,
		layananData:    ld,
	}
}

func (lh *layananHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newlayanan LayananFormatString
		bind := c.Bind(&newlayanan)

		if bind != nil {
			log.Println("can't bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"status":  "internal server error",
				"message": "gagal tambah layanan",
			})
		}

		// ubah data harga dari string ke float
		val := strings.Replace(newlayanan.Harga, ".", "", -1)
		val2 := strings.Replace(val, ",", ".", -1)
		Harga, err := strconv.ParseFloat(val2, 64)
		if err != nil {
			log.Println("can't convert")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"status":  "internal server error",
				"message": "gagal tambah layanan",
			})
		}

		// buat data baru dengan harga yang sudah di convert ke desimal
		var layanan LayananFormatFloat
		layanan.Nama = newlayanan.Nama
		layanan.Unit = newlayanan.Unit
		layanan.Harga = Harga
		layanan.UserID, layanan.UserIDS = common.ExtractData(c)

		code, datalayanan := lh.layananUseCase.AddLayanan(layanan.ToModelLayanan())
		data := FromModelLayanan(datalayanan)

		if code == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":   code,
				"status": "bad request",
				"data":   data,
			})
		}

		if code == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":   code,
				"status": "internal server error",
				"data":   data,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":   code,
			"status": "success",
			"data":   data,
		})
	}
}
