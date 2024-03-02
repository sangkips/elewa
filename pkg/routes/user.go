package routes

import (
	"elewa/pkg/controllers"
	"elewa/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.Authentication)
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:user_id", controllers.GetUser)
	router.PATCH("/user/:user_id", controllers.UpdateUser)
	router.DELETE("/user/:user_id", controllers.DeleteUser)
}
