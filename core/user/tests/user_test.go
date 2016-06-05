package user_test

import (
	"fmt"
	"go-auth/utilities/test"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server    *httptest.Server
	meURL     string
	deleteURL string
)

func TestMain(m *testing.M) {
	testRestServer := testUtils.StartServer()
	server = testRestServer.Server
	meURL = fmt.Sprintf("%s/user/me", server.URL)
	deleteURL = fmt.Sprintf("%s/user/delete", server.URL)
	ret := m.Run()
	server.Close()
	os.Exit(ret)
}
