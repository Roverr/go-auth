package user_test

import (
	"net/http"
	"testing"

	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/database"
	"github.com/roverr/go-auth/database/user"
	"github.com/roverr/go-auth/utilities/test"
)

// Should not be able to call this endpoint without login
func TestDeleteWithoutToken(t *testing.T) {
	request, reqError := http.NewRequest("DELETE", deleteURL, nil)

	if reqError != nil {
		t.Error(reqError)
	}

	res, resError := http.DefaultClient.Do(request)

	if resError != nil {
		t.Error(resError)
	}
	// Delete endpoint without token should be forbidden
	if res.StatusCode != 403 {
		t.Errorf("Response should have been 403.")
	}
}

// Should be able to delete own user after login
func TestDeleteWithValidToken(t *testing.T) {
	user := testUtils.CreateLoggedInUser()
	request, reqError := http.NewRequest("DELETE", deleteURL, nil)
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
	var deletedUser dbModels.User
	dbErr := db.Db.Unscoped().Where("ID = ?", user.User.ID).Find(&deletedUser).Error
	if dbErr != nil {
		t.Error(dbErr)
	}
	if deletedUser.DeletedAt == nil {
		t.Errorf("Deleted user should have a deletedAt field, with defined time")
	}
}
