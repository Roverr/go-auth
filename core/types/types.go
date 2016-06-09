package routingTypes

import (
	"net/http"

	"github.com/Roverr/go-auth/database/user"
	"github.com/julienschmidt/httprouter"
)

// HandleWithAuth is a handler which is almost the same the httprouter.Handle
// except that is has another parameter where the user is being
// parsed before an endpoint which needs authentication
type HandleWithAuth func(http.ResponseWriter, *http.Request, httprouter.Params, dbModels.User)
