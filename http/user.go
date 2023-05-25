package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/domain"
	"github.com/xueyiyao/safekeep/logic"
)

func (s *Server) registerUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("", s.CreateUser)
		userRouter.GET("/:id", s.ReadUser)
		userRouter.PUT("/:id", s.UpdateUser)
	}
}

func (s *Server) CreateUser(c *gin.Context) {
	var body struct {
		Name  *string
		Email *string
	}

	c.Bind(&body)

	// check if required data is present
	if body.Name == nil || body.Email == nil {
		c.Status(400)
		return
	}

	user := domain.User{Name: *(body.Name), Email: *(body.Email)}

	// call logic layer
	err := s.UserService.CreateUser(&user)

	if err != nil {
		c.Status(400)
		return
	}

	c.Status(200)
}

func (s *Server) ReadUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(("id")))

	if err != nil {
		c.Status(400)
		return
	}

	user, err := logic.ReadUser(id)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func (s *Server) UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param(("id")))

	if err != nil {
		c.Status(400)
		return
	}

	user := domain.User{ID: uint(userId)}

	var body struct {
		Name  *string
		Email *string
	}

	c.Bind(&body)

	if body.Name != nil {
		user.Name = *(body.Name)
	}
	if body.Email != nil {
		user.Email = *(body.Email)
	}

	result, err := s.UserService.UpdateUser(&user)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}
