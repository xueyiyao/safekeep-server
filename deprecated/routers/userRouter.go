package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/deprecated/controllers"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("", controllers.CreateUser)
		userRouter.GET("/:id", controllers.ReadUser)
		userRouter.PUT("/:id", controllers.UpdateUser)
	}
}
