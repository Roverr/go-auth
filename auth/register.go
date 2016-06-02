package auth

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandler is the handler function of the register endpoint
func RegisterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("Register recieved\n")
}
