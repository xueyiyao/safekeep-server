package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/deprecated/controllers"
)

func ItemRoutes(router *gin.Engine) {
	itemRouter := router.Group("/items")
	{
		itemRouter.POST("", controllers.CreateItem)
		itemRouter.GET("", controllers.ReadItems)
		itemRouter.GET("/:id", controllers.ReadItem)
		itemRouter.PUT("/:id", controllers.UpdateItem)
		itemRouter.PUT("/store/:id", controllers.StoreItem)
	}
}
