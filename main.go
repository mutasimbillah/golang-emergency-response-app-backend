package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mutasimbillah/go-er/util"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func main() {
	log.Print("Welcome...")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	//connect to database
	dsn := config.DBSource
	util.ConnectDatabase(dsn)
	var db *gorm.DB = util.DB
	//migrate tables
	util.MigrateTables(db)
	//create gin server

	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to this project ...")
	})
	server.Run()
}
