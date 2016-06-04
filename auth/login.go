package auth

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LoginHandler is the handler function of the login endpoint
func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("Login recieved\n")
}
