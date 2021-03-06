package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olegsobchuk/go-gin-app/controllers"
)

// Attach runs all existing routes
func Attach(router *gin.Engine) {
	router.GET("/", controllers.Root)
	router.GET("/login", controllers.LogIn)
}
