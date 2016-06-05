package routing

import (
	"errors"
	"go-auth/config"
	"go-auth/core/auth/types"
	"go-auth/core/types"
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/jwt"
	"go-auth/utilities/response"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// JwtCheck is a middleware used for checking the validation
// of the JWT token before letting the request go through
// the protected endpoint
func JwtCheck(next routingTypes.HandleWithAuth) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var err error
		var parsedToken authTypes.ParsedToken
		tokenString := r.Header.Get(configuration.Conf.JwtHeader)
		if tokenString == "" {
			err = errors.New("Invalid JWT token.")
			res.FinalizeError(w, err, http.StatusForbidden)
			return
		}

		parsedToken, err = jwtUtils.ValidateToken(tokenString)
		if err != nil {
			err = errors.New("Invalid JWT token.")
			res.FinalizeError(w, err, http.StatusForbidden)
			return
		}
		if parsedToken.Exp.Before(time.Now()) {
			err = errors.New("Invalid JWT token.")
			res.FinalizeError(w, err, http.StatusForbidden)
			return
		}
		var user dbModels.User
		err = db.Db.Where("ID = ?", parsedToken.ID).First(&user).Error
		if err != nil {
			err = errors.New("Invalid JWT token.")
			res.FinalizeError(w, err, http.StatusForbidden)
			return
		}
		next(w, r, ps, user)
	}
}
