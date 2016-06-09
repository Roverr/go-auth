package testUtils

import (
	"net/http/httptest"

	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/core"
	"github.com/roverr/go-auth/database"
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
		db.InitalizeModels()
	}
	router := routing.Init()
	server := httptest.NewServer(router)
	return ServerTest{
		Config: config,
		Server: server,
	}
}
