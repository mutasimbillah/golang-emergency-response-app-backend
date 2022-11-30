package util

import (
	"github.com/mutasimbillah/go-er/modules/user"
	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
