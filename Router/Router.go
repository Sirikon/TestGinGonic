package Router

import (
	"github.com/gin-gonic/gin"
	"../Views"
	"../Controllers"
)

func Init(){
	// Init a new Router
	router := gin.Default()
	// Enable template loading
	router.LoadHTMLGlob("Templates/*")

	router.GET("/",          Views.HomeView)
	router.POST("/sayhello", Controllers.SayHelloController)

	// Run in the port 8080
    router.Run(":8080")
}