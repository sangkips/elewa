package routes

import (
	"elewa/pkg/controllers"
	"elewa/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	router.Use(middleware.Authentication)
	router.GET("/category/:category_id", controllers.GetCategoryByID)
	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/category", controllers.CreateCategory)
	router.PATCH("/category/:category_id", controllers.UpdateCategory)
	router.DELETE("/category/:category_id", controllers.DeleteCategory)
}
