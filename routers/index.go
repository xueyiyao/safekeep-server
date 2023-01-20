package routers

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.Engine) {
	UserRoutes(router)
	ContainerRoutes(router)
	ItemRoutes(router)
}
