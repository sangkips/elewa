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
	router.GET("/book/:book_id", controllers.GetBookById)
	router.PATCH("/book/:book_id", controllers.UpdateBook)
	router.DELETE("/book/:book_id", controllers.DeleteBook)
}
