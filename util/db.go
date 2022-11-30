package util

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
		//panic("Failed to connect to database!")
		return
	}
	if err == nil {
		log.Info().Msg("Database Connected")
	}

	DB = db
}
