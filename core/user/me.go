package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/core/user/types"
	"github.com/roverr/go-auth/database/user"
	"github.com/roverr/go-auth/utilities/logger"
	"github.com/roverr/go-auth/utilities/response"
)

func logMe(user dbModels.User) {
	data := logger.APIPrivateLog{
		UserName: user.UserName,
		ID:       user.ID,
		Status:   http.StatusOK,
		Endpoint: "/user/me",
		Method:   "GET",
	}
	logger.PrivateAPIMessage(data)
}

// Me is an endpoint where the client can request data
// about the current user. This is necessery for stateless
// functionality
func Me(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user dbModels.User) {
	clientObject := userTypes.ClientObject{
		UserName: user.UserName,
		RealName: user.RealName,
		ID:       user.ID,
	}
	token := r.Header.Get(configuration.Conf.JwtHeader)
	w.Header().Set(configuration.Conf.JwtHeader, token)
	logMe(user)
	res.Finalize(w, clientObject)
}
