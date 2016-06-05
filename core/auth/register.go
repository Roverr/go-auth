package auth

import (
	"encoding/json"
	"errors"
	"go-auth/core/auth/types"
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/response"
	"go-auth/utilities/security"
	"go-auth/utilities/validate"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandler is the handler function of the register endpoint
func RegisterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userInformation authTypes.RegisterRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.FinalizeError(w, err, http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &userInformation)
	if err != nil {
		res.FinalizeError(w, err, http.StatusBadRequest)
		return
	}
	err = requestValidate.UsernamePassword(userInformation)
	if err != nil {
		res.FinalizeError(w, err, http.StatusBadRequest)
		return
	}
	passwordHash, pwErr := security.GeneratePassword(userInformation.Password)
	if pwErr != nil {
		res.FinalizeError(w, pwErr, http.StatusInternalServerError)
		return
	}
	var user = dbModels.User{
		RealName:     userInformation.RealName,
		UserName:     userInformation.UserName,
		Salt:         passwordHash.Salt,
		PasswordHash: passwordHash.Hash,
	}
	dbErr := db.Db.Create(&user).Error
	if dbErr != nil {
		msg := dbErr.Error()
		if strings.Contains(msg, "Duplicate entry") {
			dbErr = errors.New("User already registered in the system.")
		}
		res.FinalizeError(w, dbErr, http.StatusBadRequest)
		return
	}
	res.Finalize(w, nil)
}
