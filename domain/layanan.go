package domain

import "github.com/labstack/echo/v4"

type Layanan struct {
	ID        int
	LayananID string
	Nama      string
	Unit      string
	Harga     float64
	User      UserLayanan
}

type UserLayanan struct {
	ID     int
	UserID string
}

type LayananHandler interface {
	Add() echo.HandlerFunc
}

type LayananUseCase interface {
	AddLayanan(newLayanan Layanan) (int, Layanan)
}

type LayananData interface {
	AddLayananData(newLayanan Layanan) Layanan
}
