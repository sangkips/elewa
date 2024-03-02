package routes

import (
	"elewa/pkg/controllers"
	"elewa/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(router *gin.Engine) {
	router.Use(middleware.Authentication)
	router.GET("/invoice/:invoice_id", controllers.GetInvoiceByID)
	router.POST("/invoice", controllers.CreateInvoice)
	router.PATCH("/invoice/:invoice_id", controllers.UpdateInvoice)
}
