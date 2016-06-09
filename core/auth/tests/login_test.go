package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/core/auth/types"
	"github.com/roverr/go-auth/database"
	"github.com/roverr/go-auth/utilities/security"
	"github.com/roverr/go-auth/utilities/test"
)

// Login request without body should be invalid
func TestLoginWihoutBody(t *testing.T) {
	request, reqError := http.NewRequest("POST", loginURL, nil)

	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)

	if resError != nil {
		t.Error(resError)
	}
	if res.StatusCode != 400 {
		t.Errorf("Response should have been 400.")
	}
}

// Not registered users should not be able to login
func TestLoginWithoutRegisteredUser(t *testing.T) {
	password, pErr := security.GenerateRandomString(5)
	if pErr != nil {
		t.Error(pErr)
	}
	userName, uErr := security.GenerateRandomString(5)
	if uErr != nil {
		t.Error(uErr)
	}
	body := authTypes.RegisterRequest{
		Password: password,
		UserName: userName,
	}
	js, jErr := json.Marshal(body)
	if jErr != nil {
		t.Error(jErr)
	}
	request, reqError := http.NewRequest("POST", loginURL, bytes.NewBuffer(js))
	request.Header.Set("Content-Type", "application/json")
	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)
	if resError != nil {
		t.Error(resError)
	}

	if res.StatusCode != 404 {
		t.Errorf("Response should have been 404.")
	}
}

// Registered users should be able to login
func TestLoginWithValidRegisteredUser(t *testing.T) {
	createdUser := testUtils.CreateUser()
	body := authTypes.RegisterRequest{
		UserName: createdUser.User.UserName,
		Password: createdUser.Password,
	}
	js, jErr := json.Marshal(body)
	if jErr != nil {
		t.Error(jErr)
	}
	request, reqError := http.NewRequest("POST", loginURL, bytes.NewBuffer(js))
	request.Header.Set("Content-Type", "application/json")
	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)
	if resError != nil {
		t.Error(resError)
	}
	if res.StatusCode != 200 {
		t.Errorf("Response should have been 200.")
	}
	jwtToken := res.Header.Get(configuration.Conf.JwtHeader)
	if jwtToken == "" {
		t.Errorf("JWT Token should be valid after login.")
	}
}

// Deleted users should not be able to login
func TestLoginWithDeletedUser(t *testing.T) {
	createdUser := testUtils.CreateUser()
	db.Db.Delete(&createdUser.User)
	body := authTypes.RegisterRequest{
		UserName: createdUser.User.UserName,
		Password: createdUser.Password,
	}
	js, jErr := json.Marshal(body)
	if jErr != nil {
		t.Error(jErr)
	}
	request, reqError := http.NewRequest("POST", loginURL, bytes.NewBuffer(js))
	request.Header.Set("Content-Type", "application/json")
	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)
	if resError != nil {
		t.Error(resError)
	}
	if res.StatusCode != 404 {
		t.Errorf("Response should have been 404.")
	}
}
