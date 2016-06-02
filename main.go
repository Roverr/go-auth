package main

import (
	"fmt"
	"go-auth/core"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Server starting and listening on %d\n", 8080)
	router := routing.Init()
	fmt.Printf("Routes initialized!\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
