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
	err := db.Db.Delete(&user).Error
	if err != nil {
		res.FinalizeError(w, err, http.StatusInternalServerError)
	}
	res.Finalize(w, nil)
}
