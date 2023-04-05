package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/controllers"
	initializers "github.com/xueyiyao/safekeep/initializers"
	middleware "github.com/xueyiyao/safekeep/middleware"
	routes "github.com/xueyiyao/safekeep/routers"
)

// Runs before main
func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/admin/login", controllers.AdminLogin)

	r.Use(middleware.RequireAuth)
	routes.AddRoutes(r)

	fmt.Println("Listening on Port", os.Getenv("PORT"))

	r.Run()
}
