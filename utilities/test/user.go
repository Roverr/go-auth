package testUtils

import (
	"github.com/Roverr/go-auth/database"
	"github.com/Roverr/go-auth/database/user"
	"github.com/Roverr/go-auth/utilities/jwt"
	"github.com/Roverr/go-auth/utilities/security"
)

// CreateUser is a function for creating a random
// user to the database
func CreateUser() CreatedUser {
	var err error
	var user dbModels.User
	var password string
	var hash security.PasswordHash
	user.RealName, err = security.GenerateRandomString(5)
	CallFatalIfError(err)
	user.UserName, err = security.GenerateRandomString(5)
	CallFatalIfError(err)
	password, err = security.GenerateRandomString(5)
	CallFatalIfError(err)
	hash, err = security.GeneratePassword(password)
	CallFatalIfError(err)
	user.PasswordHash = hash.Hash
	user.Salt = hash.Salt
	dbErr := db.Db.Create(&user).Error
	CallFatalIfError(dbErr)
	return CreatedUser{
		User:     user,
		Password: password,
	}
}

// CreateLoggedInUser is the same as CreateUser
// except that it also adds a valid token to the
// return value which can be used to make API calls
func CreateLoggedInUser() CreatedUser {
	var err error
	user := CreateUser()
	user.Token, err = jwtUtils.CreateToken(user.User.ID, user.User.UserName)
	CallFatalIfError(err)
	return user
}
