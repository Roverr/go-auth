package auth

import (
	"encoding/json"
	"errors"
	"go-auth/utilities/response"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type requestProperties struct {
	RealName string `json:"realName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func validateRequest(request requestProperties) error {
	if request.UserName == "" || request.Password == "" {
		return errors.New("Request body did not contain userName or password.")
	}
	return nil
}

// RegisterHandler is the handler function of the register endpoint
func RegisterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userInformation requestProperties
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.FinalizeError(w, err, http.StatusInternalServerError)
		return
	}
	//var user dbModels.User
	err = json.Unmarshal(body, &userInformation)
	if err != nil {
		res.FinalizeError(w, err, http.StatusInternalServerError)
		return
	}
	err = validateRequest(userInformation)
	if err != nil {
		res.FinalizeError(w, err, http.StatusBadRequest)
		return
	}
	res.Finalize(w, userInformation)
}
