package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mereb/v4/database"
	"github.com/mereb/v4/routes"
)

func main() {

	database.ConnectDB()
	database.RunMigrations()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	routes.PersonRoutes(router)

	router.Run(":3000")
}
