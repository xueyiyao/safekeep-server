package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/domain"
	"github.com/xueyiyao/safekeep/initializers"
	"github.com/xueyiyao/safekeep/repository"
)

type Server struct {
	Router      *gin.Engine
	UserService domain.UserService
}

func NewServer() *Server {
	s := &Server{
		Router:      gin.Default(),
		UserService: repository.NewUserService(initializers.DB),
	}

	return s
}

func (s *Server) Run() error {
	// register routes
	r := s.Router
	s.registerUserRoutes(r)

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
