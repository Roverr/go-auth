package dbModels

import "github.com/jinzhu/gorm"

// User model which describes the
// structure of the user in the database
type User struct {
	gorm.Model
	RealName     string
	Salt         string
	PasswordHash string
	UserName     string `gorm:"unique_index"`
}
