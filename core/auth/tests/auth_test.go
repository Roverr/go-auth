package auth_test

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/roverr/go-auth/utilities/test"
)

var (
	server      *httptest.Server
	loginURL    string
	registerURL string
	logoutURL   string
)

func TestMain(m *testing.M) {
	testRestServer := testUtils.StartServer()
	server = testRestServer.Server
	loginURL = fmt.Sprintf("%s/auth/login", server.URL)
	registerURL = fmt.Sprintf("%s/auth/register", server.URL)
	logoutURL = fmt.Sprintf("%s/auth/logout", server.URL)
	ret := m.Run()
	server.Close()
	os.Exit(ret)
}
