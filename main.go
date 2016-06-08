package main

import (
	"fmt"
	"go-auth/config"
	"go-auth/core"
	"go-auth/database"
	"go-auth/utilities/logger"
	"log"
	"net/http"
)

func main() {
	config := configuration.InitConfig()
	logger.InitLogger()
	db.CreateDbConnection()
	dbErr := db.InitalizeModels()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	router := routing.Init()
	port := fmt.Sprintf("%d", config.Port)
	logger.Standard.Info("Server starting and listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
