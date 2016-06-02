package auth_test

import (
	"fmt"
	"go-auth/core"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server      *httptest.Server
	loginURL    string
	registerURL string
	logoutURL   string
)

func TestMain(m *testing.M) {
	router := routing.Init()
	server = httptest.NewServer(router)
	loginURL = fmt.Sprintf("%s/auth/login", server.URL)
	registerURL = fmt.Sprintf("%s/auth/register", server.URL)
	logoutURL = fmt.Sprintf("%s/auth/logout", server.URL)
	ret := m.Run()
	server.Close()
	os.Exit(ret)
}
