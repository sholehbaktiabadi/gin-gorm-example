package main

import (
	"net/http"
	"v1/api"
	"v1/config"
	"v1/database"
)

func main() {
	db, err := database.DatabaseConnection()
	if err != nil {
		panic("error database coonection")
	}
	var (
		api    = api.New(*db)
		server = http.Server{
			Addr:    ":" + config.Env("SERVER_PORT"),
			Handler: api,
		}
	)
	server.ListenAndServe()
}
