package testUtils

import (
	"go-auth/config"
	"go-auth/core"
	"go-auth/database"
	"net/http/httptest"

	"github.com/jinzhu/gorm"
)

// ServerTest struct is used to describe
// the structure of the created test server properties
type ServerTest struct {
	Config configuration.Config
	Server *httptest.Server
	Db     *gorm.DB
}

// StartServer is a test utility function
// used to start the REST API with
// database connection
func StartServer() ServerTest {
	config := configuration.InitConfig()
	dbConn := db.CreateDbConnection()
	db.InitalizeModels(dbConn)
	router := routing.Init()
	server := httptest.NewServer(router)
	return ServerTest{
		Config: config,
		Server: server,
		Db:     dbConn,
	}
}
