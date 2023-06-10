package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/domain"
)

type Server struct {
	Router           *gin.Engine
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
	s.registerUserRoutes(r)
	s.RegisterContainerRoutes(r)
	s.registerItemRoutes(r)

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
