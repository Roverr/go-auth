package routingTypes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/roverr/go-auth/database/user"
)

// HandleWithAuth is a handler which is almost the same the httprouter.Handle
// except that is has another parameter where the user is being
// parsed before an endpoint which needs authentication
type HandleWithAuth func(http.ResponseWriter, *http.Request, httprouter.Params, dbModels.User)
