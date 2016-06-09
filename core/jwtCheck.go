package routing

import (
	"errors"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/core/auth/types"
	"github.com/roverr/go-auth/core/types"
	"github.com/roverr/go-auth/database"
	"github.com/roverr/go-auth/database/user"
	"github.com/roverr/go-auth/utilities/jwt"
	"github.com/roverr/go-auth/utilities/response"
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
