package auth_test

import (
	"bytes"
	"encoding/json"
	"go-auth/core/auth/types"
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/security"
	"net/http"
	"testing"
)

func TestRegisterWithoutBody(t *testing.T) {
	request, reqError := http.NewRequest("POST", registerURL, nil)

	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)

	if resError != nil {
		t.Error(resError)
	}
	// Register without body should be a json parse error
	if res.StatusCode != 400 {
		t.Errorf("Response should have been 400.")
	}
}

func TestRegisterWithInvalidBody(t *testing.T) {
	realName, rErr := security.GenerateRandomString(5)
	if rErr != nil {
		t.Error(rErr)
	}
	password, pErr := security.GenerateRandomString(5)
	if pErr != nil {
		t.Error(pErr)
	}
	body := authTypes.RegisterRequest{
		RealName: realName,
		Password: password,
	}
	js, jErr := json.Marshal(body)
	if jErr != nil {
		t.Error(jErr)
	}
	request, reqError := http.NewRequest("POST", registerURL, bytes.NewBuffer(js))
	request.Header.Set("Content-Type", "application/json")
	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)

	if resError != nil {
		t.Error(resError)
	}
	// Register with invalid body should be a bad request
	if res.StatusCode != 400 {
		t.Errorf("Response should have been 400.")
	}
}

func TestRegister(t *testing.T) {
	realName, rErr := security.GenerateRandomString(5)
	if rErr != nil {
		t.Error(rErr)
	}
	password, pErr := security.GenerateRandomString(5)
	if pErr != nil {
		t.Error(pErr)
	}
	userName, uErr := security.GenerateRandomString(5)
	if uErr != nil {
		t.Error(uErr)
	}
	body := authTypes.RegisterRequest{
		RealName: realName,
		Password: password,
		UserName: userName,
	}
	js, jErr := json.Marshal(body)
	if jErr != nil {
		t.Error(jErr)
	}
	request, reqError := http.NewRequest("POST", registerURL, bytes.NewBuffer(js))
	request.Header.Set("Content-Type", "application/json")
	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)

	if resError != nil {
		t.Error(resError)
	}
	// Register with invalid body should be a bad request
	if res.StatusCode != 200 {
		t.Errorf("Response should have been 200.")
	}
	var user dbModels.User
	dbErr := db.Db.Where("user_name = ?", body.UserName).Find(&user).Error
	if dbErr != nil {
		t.Error(dbErr)
	}
	realNameMatch := user.RealName == body.RealName
	userNameMatch := user.UserName == body.UserName
	if !realNameMatch || !userNameMatch {
		t.Errorf("Real name or user name is not matching.")
	}
}
