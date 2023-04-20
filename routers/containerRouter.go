package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/controllers"
)

func ContainerRoutes(router *gin.Engine) {
	containerRouter := router.Group("/containers")
	{
		containerRouter.POST("/", controllers.CreateContainer)
		containerRouter.GET("/", controllers.ReadContainers)
		containerRouter.PUT("/:id", controllers.UpdateContainer)
	}
}
