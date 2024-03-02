package routes

import (
	"elewa/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	// Authentication
	router.POST("/login", controllers.LoginUser)
	router.POST("/register", controllers.RegisterUser)
}
