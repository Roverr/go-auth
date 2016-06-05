package user_test

import (
	"go-auth/config"
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/test"
	"net/http"
	"testing"
)

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
