package routing

import (
	"github.com/julienschmidt/httprouter"
	"github.com/roverr/go-auth/core/auth"
	"github.com/roverr/go-auth/core/user"
)

// Init router
func Init() *httprouter.Router {
	router := httprouter.New()
	router.POST("/auth/login", auth.LoginHandler)
	router.POST("/auth/register", auth.RegisterHandler)
	router.GET("/user/me", JwtCheck(user.Me))
	router.DELETE("/user/delete", JwtCheck(user.Delete))
	return router
}
