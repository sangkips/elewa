package routes

import (
	"elewa/pkg/controllers"
	"elewa/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	router.Use(middleware.Authentication)
	router.POST("/book", controllers.CreateBook)
	router.GET("/books", controllers.GetAllBooks)
	router.GET("/book/:id", controllers.GetBookById)
	router.PATCH("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)
}
