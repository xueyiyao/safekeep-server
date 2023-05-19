package main

import (
	"fmt"
	"os"

	cors "github.com/gin-contrib/cors"
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
	initializers.SetupOAuthLogin()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers, user-id"},
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/admin/login", controllers.AdminLogin)
	routes.LoginRoutes(r)

	r.Use(middleware.RequireAuth)
	routes.AddRoutes(r)

	fmt.Println("Listening on Port", os.Getenv("PORT"))

	r.Run()
}
