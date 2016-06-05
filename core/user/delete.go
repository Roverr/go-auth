package user

import (
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Delete is an endpoint where the users can delete themselves from the system
func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user dbModels.User) {
	db.Db.Delete(&user)
	res.Finalize(w, nil)
}
