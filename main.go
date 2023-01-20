package main

import (
	"github.com/gin-gonic/gin"
	initializers "github.com/xueyiyao/safekeep/initializers"
	routes "github.com/xueyiyao/safekeep/routers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.AddRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
