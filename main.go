package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mutasimbillah/go-er/util"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Print("Welcome...")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	//connect to database
	dbUrl := config.DBSource
	util.ConnectDatabase(dbUrl)
	//create gin server
	server := gin.Default()
	server.Run()
}
