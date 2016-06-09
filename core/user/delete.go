package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/roverr/go-auth/database"
	"github.com/roverr/go-auth/database/user"
	"github.com/roverr/go-auth/utilities/logger"
	"github.com/roverr/go-auth/utilities/response"
)

func logDelete(user dbModels.User) {
	data := logger.APIPrivateLog{
		UserName: user.UserName,
		ID:       user.ID,
		Status:   http.StatusOK,
		Endpoint: "/user/delete",
		Method:   "DELETE",
	}
	logger.PrivateAPIMessage(data)
}

// Delete is an endpoint where the users can delete themselves from the system
func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user dbModels.User) {
	err := db.Db.Delete(&user).Error
	if err != nil {
		logger.Standard.Error("Error happened during DB delete in /user/delete")
		res.FinalizeError(w, err, http.StatusInternalServerError)
		return
	}
	logDelete(user)
	res.Finalize(w, nil)
}
