package routes

import (
	"elewa/pkg/controllers"
	"elewa/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.Use(middleware.Authentication)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:order_id", controllers.GetOrderByID)
	router.POST("/orders", controllers.CreateOrder)
	router.PATCH("/orders/:order_id", controllers.UpdateOrder)
}
