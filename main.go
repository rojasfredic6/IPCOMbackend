package main

import (
	"IPCOMBack/models"
	"encoding/csv"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Wellcome to Fredic Rojas' backend test")
	})

	e.GET("/resume/2019-12-01", func(c echo.Context) error {
		days := c.QueryParam("dias")
		newResponse := models.NewSalesResponse()

		if days == "" {
			return c.String(http.StatusBadRequest, "Error in query parameters")
		}
		response := newResponse.GenerateResponse(days)
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/csv", func(c echo.Context) error {
		file, err := c.FormFile("file")

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		src, err := file.Open()

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		arch := csv.NewReader(src)
		data, err := arch.ReadAll()

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())

		}

		resp, err := models.GenerateResJSON(data)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
