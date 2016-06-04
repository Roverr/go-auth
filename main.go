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
	router := routing.Init()
	config := configuration.InitConfig()
	fmt.Printf("Server starting and listening on %d\n", config.Port)
	dbConn := db.CreateDbConnection()
	fmt.Println("Database connection initialized")
	db.InitalizeModels(dbConn)
	fmt.Println("Models synchronized into database")
	portString := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(portString, router))
}
