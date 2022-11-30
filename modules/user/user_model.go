package user

import "gorm.io/gorm"

//User Model
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
