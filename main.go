package main

import (
	"fmt"
	"go-auth/config"
	"go-auth/core"
	"go-auth/database"
	"log"
	"net/http"
)

func main() {
	config := configuration.InitConfig()
	dbConn := db.CreateDbConnection()
	dbErr := db.InitalizeModels(dbConn)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	router := routing.Init()
	portString := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Server starting and listening on %d\n", config.Port)
	log.Fatal(http.ListenAndServe(portString, router))
}
