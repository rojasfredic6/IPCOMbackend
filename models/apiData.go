package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	ClientId int     `json:"client_id"`
	Phone    string  `json:"phone"`
	Nombre   string  `json:"nombre"`
	Compro   bool    `json:"compro"`
	Tdc      string  `json:"tdc"`
	Monto    float32 `json:"monto"`
	Date     string  `json:"date"`
}

var (
	URL = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/"
)

func GetData(date []string) ([]APIResponse, error) {

	var newResponse []APIResponse

	c := make(chan []APIResponse, len(date))

	go routinRequest(date, c)

	for i := 0; i < len(date); i++ {
		newResponse = append(newResponse, <-c...)
	}

	return newResponse, nil
}

func routinRequest(date []string, c chan []APIResponse) {
	client := &http.Client{}
	for _, r := range date {
		var tempResponse []APIResponse
		req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", URL, r), nil)

		if err != nil {
			echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		unmarshalError := json.Unmarshal(body, &tempResponse)

		if unmarshalError != nil {
			log.Fatal(unmarshalError.Error())
		}

		c <- tempResponse
	}
}
