package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

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
	URL = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/2019-12-01"
)

func GetData(days string) ([]APIResponse, error) {
	client := &http.Client{}
	values := url.Values{
		"dias": {days},
	}

	var newResponse []APIResponse
	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	req.URL.RawQuery = values.Encode()

	resp, err := client.Do(req)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	unmarshalError := json.Unmarshal(body, &newResponse)

	if unmarshalError != nil {
		log.Fatal(unmarshalError.Error())
	}

	fmt.Println(resp.Request.URL, len(newResponse))

	return newResponse, nil
}
