package http

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Server struct {
	Router *gin.Engine

	GoogleClientID     string
	GoogleClientSecret string

	UserService      domain.UserService
	ContainerService domain.ContainerService
	ItemService      domain.ItemService
}

func NewServer() *Server {
	s := &Server{
		Router: gin.Default(),
	}

	return s
}

func (s *Server) Run() error {
	// register routes
	r := s.Router

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers, user-id"},
	}))

	// Base route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "shhhh!",
		})
	})

	// Register unauthenticated routes
	{
		s.RegisterAuthRoutes(r)
	}

	// Register authenticated routes
	{
		s.registerUserRoutes(r)
		s.RegisterContainerRoutes(r)
		s.registerItemRoutes(r)
	}

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}

func (s *Server) OAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "https://35ff-135-180-67-224.ngrok-free.app/callback",
		ClientID:     s.GoogleClientID,
		ClientSecret: s.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
