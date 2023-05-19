package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/controllers"
)

func LoginRoutes(router *gin.Engine) {
	loginRouter := router.Group("/login")
	{
		loginRouter.POST("/google", controllers.HandleGoogleLogin)
	}
}
