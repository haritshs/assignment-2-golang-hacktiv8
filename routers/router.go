package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetAllOrder)
	router.POST("/orders", controllers.CreateOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)
	router.PUT("/orders/:orderID", controllers.UpdateOrderByID)

	return router
}
