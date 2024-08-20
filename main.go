// main.go

package main

import (
	"baby-varkin-health/controllers"
	"baby-varkin-health/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/test", controllers.Test)
	router.POST("/line", controllers.WebHook)

	router.Run()
}
