package testUtils

import (
	"go-auth/database"
	"go-auth/database/user"
)

// DropDb is dropping all data from the database.
// It is used to clean database before tests.
func DropDb() {
	var err error
	err = db.Db.DropTableIfExists(&dbModels.User{}).Error
	CallFatalIfError(err)
	if !db.Db.HasTable(&dbModels.User{}) {
		err = db.InitalizeModels(db.Db)
		CallFatalIfError(err)
	}
}
