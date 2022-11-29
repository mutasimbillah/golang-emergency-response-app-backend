package util

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(url string) {
	database, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
		//panic("Failed to connect to database!")
		return
	}

	DB = database
}
