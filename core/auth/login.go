package auth

import (
	"encoding/json"
	"errors"
	"go-auth/config"
	"go-auth/core/auth/types"
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/jwt"
	"go-auth/utilities/logger"
	"go-auth/utilities/response"
	"go-auth/utilities/security"
	"go-auth/utilities/validate"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func logSuccessLogin(user dbModels.User) {
	data := logger.APIPrivateLog{
		Status:   http.StatusOK,
		Endpoint: "/auth/login",
		Method:   "POST",
		UserName: user.UserName,
		ID:       user.ID,
	}
	logger.PrivateAPIMessage(data)
}

func logFailLogin(status int, message string) {
	data := logger.APIPublicLog{
		Status:   status,
		Endpoint: "/auth/login",
		Method:   "POST",
		Message:  message,
	}
	logger.PublicAPIMessage(data)
}

// LoginHandler is the handler function of the login endpoint
func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Can use RegisterRequest here, RealName will be an empty string
	// but this won't matter here
	var userInformation authTypes.RegisterRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Standard.Error("Error happened during parsing JSON in /auth/login")
		res.FinalizeError(w, err, http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &userInformation)
	if err != nil {
		logger.Standard.Error("Error happened during parsing JSON in /auth/login")
		res.FinalizeError(w, err, http.StatusBadRequest)
		return
	}
	err = requestValidate.UsernamePassword(userInformation)
	if err != nil {
		logFailLogin(http.StatusBadRequest, "User did not provide correct req.body.")
		res.FinalizeError(w, err, http.StatusBadRequest)
		return
	}
	var user dbModels.User
	err = db.Db.Where("user_name = ?", userInformation.UserName).First(&user).Error
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "record not found" {
			code = http.StatusNotFound
			err = errors.New("Incorrect userName or password.")
			logFailLogin(code, err.Error())
		}
		res.FinalizeError(w, err, code)
		return
	}

	isPasswordValid := security.ValidatePassword(
		user.PasswordHash,
		user.Salt,
		userInformation.Password,
	)
	if !isPasswordValid {
		err = errors.New("Incorrect userName or password.")
		logFailLogin(http.StatusUnauthorized, err.Error())
		res.FinalizeError(w, err, http.StatusUnauthorized)
		return
	}
	tokenString, jwtErr := jwtUtils.CreateToken(
		user.ID,
		user.UserName,
	)
	if jwtErr != nil {
		logger.Standard.Error("Error happened during JWT token creation in /auth/login")
		res.FinalizeError(w, jwtErr, http.StatusInternalServerError)
		return
	}
	w.Header().Set(configuration.Conf.JwtHeader, tokenString)
	logSuccessLogin(user)
	res.Finalize(w, nil)
}
