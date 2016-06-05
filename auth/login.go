package auth

import (
	"encoding/json"
	"errors"
	"go-auth/auth/types"
	"go-auth/config"
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/jwt"
	"go-auth/utilities/response"
	"go-auth/utilities/security"
	"go-auth/utilities/validate"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
)

// LoginHandler is the handler function of the login endpoint
func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Can use RegisterRequest here, RealName will be an empty string
	// but this won't matter here
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
	var user dbModels.User
	err = db.Db.Where("user_name = ?", userInformation.UserName).First(&user).Error
	if err != nil {
		res.FinalizeError(w, err, http.StatusInternalServerError)
		return
	}
	spew.Dump(user)
	if user.UserName == "" {
		err = errors.New("Incorrect userName or password.")
		res.FinalizeError(w, err, http.StatusUnauthorized)
		return
	}
	isPasswordValid := security.ValidatePassword(
		user.PasswordHash,
		user.Salt,
		userInformation.Password,
	)
	if !isPasswordValid {
		err = errors.New("Incorrect userName or password.")
		res.FinalizeError(w, err, http.StatusUnauthorized)
		return
	}
	tokenString, jwtErr := jwtUtils.CreateToken(
		user.ID,
		user.UserName,
	)
	if jwtErr != nil {
		res.FinalizeError(w, jwtErr, http.StatusInternalServerError)
		return
	}
	w.Header().Set(configuration.Conf.JwtHeader, tokenString)
	res.Finalize(w, nil)
}
