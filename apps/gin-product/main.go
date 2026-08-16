package main

import (
	"github.com/gin-gonic/gin"
)

func applicationHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	registerSourceRoutes(router)
	registerExplicitSemanticRoutes(router)
	registerControllerRoutes(router)
	return router
}

func main() {
	_ = applicationHandler().Run(":8080")
}
