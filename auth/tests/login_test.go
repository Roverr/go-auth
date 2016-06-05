package auth_test

import (
	"bytes"
	"encoding/json"
	"go-auth/auth/types"
	"go-auth/config"
	"go-auth/utilities/security"
	"go-auth/utilities/test"
	"net/http"
	"testing"
)

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
