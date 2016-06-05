package testUtils

import (
	"go-auth/config"
	"go-auth/core"
	"go-auth/database"
	"net/http/httptest"
)

// ServerTest struct is used to describe
// the structure of the created test server properties
type ServerTest struct {
	Config configuration.Config
	Server *httptest.Server
}

// Set some variables in config if
// the tests are running locally
func setLocalTestEnviroment() {
	if configuration.Conf.IsCodeShip {
		return
	}
	configuration.Conf.DbName = "go-auth-test"
}

// StartServer is a test utility function
// used to start the REST API with
// database connection
func StartServer() ServerTest {
	config := configuration.InitConfig()
	setLocalTestEnviroment()
	if !db.IsConnected {
		db.CreateDbConnection()
	}
	router := routing.Init()
	server := httptest.NewServer(router)
	return ServerTest{
		Config: config,
		Server: server,
	}
}
