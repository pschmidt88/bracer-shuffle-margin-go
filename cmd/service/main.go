package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := echo.New()
	api.GET("/:region/:realm", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "")
	})

	api.Logger.Info(api.Start(":8080"))
}
