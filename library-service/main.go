package main

import (
	"github.com/gin-gonic/gin"
	"library-service/configs"
	"library-service/routes"
	"log"
)

func main() {
	router := SetupAppRouter()
	log.Fatal(router.Run(":80"))
}

func SetupAppRouter() *gin.Engine {

	service := configs.NewDBService()
	db := service.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	api := router.Group("api/v1")

	routes.InitLibraryRoutes(db, api)
	routes.InitUserRoutes(db, api)

	return router
}
