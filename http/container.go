package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/domain"
)

func (s *Server) RegisterContainerRoutes(router *gin.Engine) {
	containerRouter := router.Group("/containers")
	{
		containerRouter.POST("", s.CreateContainer)
		containerRouter.GET("", s.ReadContainers)
		containerRouter.PUT("/:id", s.UpdateContainer)
	}
}

func (s *Server) CreateContainer(c *gin.Context) {
	var body struct {
		Name   string
		UserID uint
	}

	c.Bind(&body)

	container := domain.Container{Name: body.Name, UserID: body.UserID}

	err := s.ContainerService.CreateContainer(&container)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"container": container,
	})
}

func (s *Server) ReadContainers(c *gin.Context) {
	userIDStr := c.GetHeader("user-id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		c.Status(403)
		return
	}

	containers, err := s.ContainerService.FindContainers(userID)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"containers": containers,
	})
}

func (s *Server) UpdateContainer(c *gin.Context) {
	containerId, err := strconv.Atoi(c.Param(("id")))

	if err != nil {
		c.Status(400)
		return
	}

	container := domain.Container{ID: uint(containerId)}

	var body struct {
		Name   *string
		UserID *string
	}

	c.Bind(&body)

	if body.Name != nil {
		container.Name = *(body.Name)
	}
	if body.UserID != nil {
		userId, err := strconv.Atoi(c.Param(("id")))
		if err != nil {
			c.Status(400)
			return
		}
		container.UserID = uint(userId)
	}

	result, err := s.ContainerService.UpdateContainer(&container)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}
