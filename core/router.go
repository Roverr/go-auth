package routing

import (
	"go-auth/core/auth"

	"github.com/julienschmidt/httprouter"
)

// Init router
func Init() *httprouter.Router {
	router := httprouter.New()
	router.POST("/auth/login", auth.LoginHandler)
	router.POST("/auth/register", auth.RegisterHandler)
	return router
}
