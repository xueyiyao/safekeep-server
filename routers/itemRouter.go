package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/controllers"
)

func ItemRoutes(router *gin.Engine) {
	itemRouter := router.Group("/items")
	{
		itemRouter.POST("/", controllers.CreateItem)
		itemRouter.GET("/", controllers.ReadItems)
		itemRouter.GET("/:id", controllers.ReadItem)
		itemRouter.PUT("/store/:id", controllers.StoreItem)
	}
}
