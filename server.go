package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"software.racoony/bracershuffle/pkg/battlenet"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := echo.New()
	api.GET("/:region/:realm", func(context echo.Context) error {
		realmName := context.Param("realm")
		region := context.Param("region")

		battleNet := battlenet.NewAPI(region)
		connectedRealm, err := battleNet.FindConnectedRealm(realmName)

		if err != nil {
			log.Fatal(err)
		}

		battleNet.ListAuctions(connectedRealm.ID)

		return context.String(http.StatusOK, ":+1:")
	})

	api.Logger.Info(api.Start(":8080"))
}
