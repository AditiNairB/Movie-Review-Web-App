package main

import (
	// homeRoutes "SPL/routes/homeRoutes"
	"GoLang-WebServer/models"
	"GoLang-WebServer/router/apiRouter"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// initialize mongodb server
	models.ConnectAndInitialize()

	router := gin.Default()

	// using default CORS until deployed
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(CORS())

	api := router.Group("/api")
	apiRouter.RouteApi(api)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
