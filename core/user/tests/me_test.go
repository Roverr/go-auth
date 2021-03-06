package user_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Roverr/go-auth/config"
	"github.com/Roverr/go-auth/core/user/types"
	"github.com/Roverr/go-auth/database"
	"github.com/Roverr/go-auth/utilities/test"
)

// This endpoint should be unreachable without login
func TestMeWithoutToken(t *testing.T) {
	request, reqError := http.NewRequest("GET", meURL, nil)

	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)

	if resError != nil {
		t.Error(resError)
	}
	// Me endpoint without token should be forbidden
	if res.StatusCode != 403 {
		t.Errorf("Response should have been 403.")
	}
}

// Logged in users should be able to get information from this endpoint
func TestMeWithValidToken(t *testing.T) {
	user := testUtils.CreateLoggedInUser()
	request, reqError := http.NewRequest("GET", meURL, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(configuration.Conf.JwtHeader, user.Token)
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
	body, ioErr := ioutil.ReadAll(res.Body)
	if ioErr != nil {
		t.Error(ioErr)
	}

	type data struct {
		Item userTypes.ClientObject `json:"item"`
	}
	type responseBody struct {
		Data data `json:"data"`
	}
	var resBody responseBody
	jsErr := json.Unmarshal(body, &resBody)
	if jsErr != nil {
		t.Error(jsErr)
	}
	if resBody.Data.Item.ID != user.User.ID {
		t.Errorf("User ID is not matching.")
	}
	if resBody.Data.Item.RealName != user.User.RealName {
		t.Errorf("User RealName is not matching.")
	}
	if resBody.Data.Item.UserName != user.User.UserName {
		t.Errorf("User UserName is not matching.")
	}
}

// Deleted users should not be able to use their valid token
func TestMeWithValidTokenAndDeletedUser(t *testing.T) {
	user := testUtils.CreateLoggedInUser()
	db.Db.Delete(&user.User)
	request, reqError := http.NewRequest("GET", meURL, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(configuration.Conf.JwtHeader, user.Token)
	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)
	if resError != nil {
		t.Error(resError)
	}
	if res.StatusCode != 403 {
		t.Errorf("Response should have been 403.")
	}
}
