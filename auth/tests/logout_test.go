package auth_test

import (
	"net/http"
	"testing"
)

func TestLogout(t *testing.T) {
	request, reqError := http.NewRequest("POST", logoutURL, nil)

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
}
