package testUtils

import (
	"go-auth/database"
	"go-auth/database/user"
	"go-auth/utilities/security"
	"log"
)

func callFatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUser is a function for creating a random
// user to the database
func CreateUser() CreatedUser {
	var err error
	var user dbModels.User
	var password string
	var hash security.PasswordHash
	user.RealName, err = security.GenerateRandomString(5)
	callFatalIfError(err)
	user.UserName, err = security.GenerateRandomString(5)
	callFatalIfError(err)
	password, err = security.GenerateRandomString(5)
	callFatalIfError(err)
	hash, err = security.GeneratePassword(password)
	callFatalIfError(err)
	user.PasswordHash = hash.Hash
	user.Salt = hash.Salt
	dbErr := db.Db.Create(&user).Error
	callFatalIfError(dbErr)
	return CreatedUser{
		User:     user,
		Password: password,
	}
}
