package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type SalesResponse struct {
	Total         float32            `json:"total"`
	ComprasPorTDC map[string]float32 `json:"compras_por_tdc"`
	NoCompraron   int                `json:"no_compraron"`
	CompraMasAlta float32            `json:"compra_mas_alta"`
}

func NewSalesResponse() *SalesResponse {
	return &SalesResponse{
		ComprasPorTDC: make(map[string]float32),
	}
}

func (s *SalesResponse) GenerateResponse(days string) *SalesResponse {

	dateFormat := "2019-12-01"
	intDays, err := strconv.Atoi(days)

	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	listDates := []string{}

	if intDays == 0 {
		listDates = append(listDates, "2019-12-01")
	} else {
		date, err := time.Parse("2006-01-02", dateFormat)
		if err != nil {
			echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		for i := 0; i < intDays; i++ {
			if i == 0 {
				date = date.AddDate(0, 0, i)
			} else {
				date = date.AddDate(0, 0, 1)
			}
			listDates = append(listDates, date.Format("2006-01-02"))
		}
	}

	data, err := GetData(listDates)

	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, d := range data {
		if d.Monto > 0 {
			s.Total += d.Monto
			_, ok := s.ComprasPorTDC[d.Tdc]
			if ok {
				s.ComprasPorTDC[d.Tdc] += d.Monto
			} else {
				s.ComprasPorTDC[d.Tdc] = d.Monto
			}
			if d.Monto > s.CompraMasAlta {
				s.CompraMasAlta = d.Monto
			}
		} else {
			s.NoCompraron++
		}
	}

	return s
}
