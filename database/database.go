package db

import (
	"fmt"
	"log"

	// Lightweight MySQL driver import
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/database/user"
	"github.com/roverr/go-auth/utilities/logger"
)

// Conn is the exported singleton database connection
var (
	Db          *gorm.DB
	IsConnected = false
)

// CreateDbConnection will be able to create connection
// to the SQLite database
func CreateDbConnection() *gorm.DB {
	var err error
	config := configuration.Conf
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPass,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	Db, err = gorm.Open("mysql", connString)
	if err != nil {
		logger.Standard.Critical("Error happened during opening connection with database")
		log.Fatal(err)
	}
	IsConnected = true
	return Db
}

// InitalizeModels is synchronizing the models
// into the database
func InitalizeModels() error {
	err := Db.AutoMigrate(&dbModels.User{}).Error
	return err
}
