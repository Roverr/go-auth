package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/core"
	"github.com/roverr/go-auth/database"
	"github.com/roverr/go-auth/utilities/logger"
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
