package testUtils

import "go-auth/database/user"

// CreatedUser is describing the structure
// of the newly created random user which is
// used in the tests
type CreatedUser struct {
	User     dbModels.User
	Password string
	Token    string
}
