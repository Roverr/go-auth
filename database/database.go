package db

import (
	"fmt"
	"go-auth/config"
	"go-auth/database/user"
	"log"

	// Lightweight MySQL driver import
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// CreateDbConnection will be able to create connection
// to the SQLite database
func CreateDbConnection() *gorm.DB {
	config := configuration.Conf
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPass,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		fmt.Println("Error happened during opening connection with database")
		log.Fatal(err)
	}

	return db
}

// InitalizeModels is synchronizing the models
// into the database
func InitalizeModels(db *gorm.DB) {
	db.AutoMigrate(&dbModels.User{})
}
