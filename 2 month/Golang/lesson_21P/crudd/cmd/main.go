package main

import (
	"fmt"
	"log"
	"net/http"

	"app/api"
	"app/config"
	"app/storage/postgres"
)

func main() {
	cfg := config.Load()

	dbconn, err := postgres.NewConnectionPostgres(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	api.NewApi(cfg, dbconn)

	fmt.Println("Listen: ", cfg.ServerHost+cfg.Port)
	if err := http.ListenAndServe(cfg.ServerHost+cfg.Port, nil); err != nil {
		log.Fatal("Server does not run: " + err.Error())
	}
}
