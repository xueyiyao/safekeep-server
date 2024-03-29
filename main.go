package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xueyiyao/safekeep/deprecated/controllers"
	routes "github.com/xueyiyao/safekeep/deprecated/routers"
	HTTP "github.com/xueyiyao/safekeep/http"
	"github.com/xueyiyao/safekeep/middleware"
	"github.com/xueyiyao/safekeep/postgres"
)

// Runs before main
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// initializers.ConnectToDB()
	// initializers.SetupOAuthLogin()
}

func main() {
	m := NewMain()
	if err := m.Run(); err != nil {
		// Handle Gracefully
		fmt.Println(err)
	}
}

type Main struct {
	DB         *postgres.DB
	HTTPServer *HTTP.Server
}

func NewMain() *Main {
	return &Main{DB: postgres.NewDB(""), HTTPServer: HTTP.NewServer()}
}

func (m *Main) Run() (err error) {
	m.DB.DSN = os.Getenv("DB_URL")
	if err = m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	userService := postgres.NewUserService(m.DB.DB)
	containerService := postgres.NewContainerService(m.DB.DB)
	itemServce := postgres.NewItemService(m.DB.DB)

	m.HTTPServer.GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	m.HTTPServer.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	m.HTTPServer.UserService = userService
	m.HTTPServer.ContainerService = containerService
	m.HTTPServer.ItemService = itemServce

	if err = m.HTTPServer.Run(); err != nil {
		return err
	}
	return nil
}

func otherMain() {
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
