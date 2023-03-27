package models

import (
	"net/http"

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

	data, err := GetData(days)

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
