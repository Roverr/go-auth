package auth

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutHandler is the handler function of the logout endpoint
func LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("Logout recieved\n")
}
