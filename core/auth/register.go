package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-auth/core/auth/types"
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/logger"
	"go-auth/utilities/response"
	"go-auth/utilities/security"
	"go-auth/utilities/validate"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func logSuccessRegister(user dbModels.User) {
	data := logger.APIPrivateLog{
		Status:   http.StatusOK,
		Endpoint: "/auth/register",
		Method:   "POST",
		UserName: user.UserName,
		ID:       user.ID,
	}
	logger.PrivateAPIMessage(data)
}

func logFailRegister(status int, message string) {
	data := logger.APIPublicLog{
		Status:   status,
		Endpoint: "/auth/register",
		Method:   "POST",
		Message:  message,
	}
	logger.PublicAPIMessage(data)
}

// RegisterHandler is the handler function of the register endpoint
func RegisterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userInformation authTypes.RegisterRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Standard.Error("Error happened during parsing JSON in /auth/register")
		res.FinalizeError(w, err, http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &userInformation)
	if err != nil {
		logger.Standard.Error("Error happened during parsing JSON in /auth/register")
		res.FinalizeError(w, err, http.StatusBadRequest)
		return
	}
	err = requestValidate.UsernamePassword(userInformation)
	if err != nil {
		logFailRegister(http.StatusBadRequest, "User did not provide correct req.body.")
		res.FinalizeError(w, err, http.StatusBadRequest)
		return
	}
	passwordHash, pwErr := security.GeneratePassword(userInformation.Password)
	if pwErr != nil {
		logger.Standard.Error("Error happened during generating passwordHash in /auth/register")
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
			logFailRegister(http.StatusBadRequest, fmt.Sprintf("%s user is already registered.", user.UserName))
		}
		res.FinalizeError(w, dbErr, http.StatusBadRequest)
		return
	}
	logSuccessRegister(user)
	res.Finalize(w, nil)
}
