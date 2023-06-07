package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/domain"
)

func (s *Server) registerItemRoutes(router *gin.Engine) {
	itemRouter := router.Group("/items")
	{
		itemRouter.POST("", s.CreateItem)
		itemRouter.GET("", s.ReadItems)
		itemRouter.GET("/:id", s.ReadItem)
		itemRouter.PUT("/:id", s.UpdateItem)
		itemRouter.PUT("/store/:id", s.StoreItem)
	}
}

func (s *Server) CreateItem(c *gin.Context) {
	var body struct {
		Name   string
		UserID uint
	}

	c.Bind(&body)

	item := domain.Item{Name: body.Name, UserID: body.UserID}
	result := s.ItemService.CreateItem(&item)

	if result != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

func (s *Server) ReadItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(("id")))

	if err != nil {
		c.Status(400)
		return
	}

	item, err := s.ItemService.FindItemByID(id)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

func (s *Server) ReadItems(c *gin.Context) {
	userIDStr := c.GetHeader("user-id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.Status(403)
		return
	}

	containerIDStr := c.Query("container")
	if err != nil {
		c.Status(403)
		return
	}

	var items []*domain.Item
	if len(containerIDStr) != 0 {
		containerID, err := strconv.Atoi(containerIDStr)
		if err != nil {
			c.Status(403)
			return
		}

		items, err = s.ItemService.FindItems(domain.ItemFilter{User_ID: &userID, Container_ID: &containerID})
	} else {
		items, err = s.ItemService.FindItems(domain.ItemFilter{User_ID: &userID})
	}

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}

func (s *Server) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(("id")))

	if err != nil {
		c.Status(400)
		return
	}

	var body struct {
		Name string
	}

	c.Bind(&body)
	item := domain.Item{Name: body.Name, ID: uint(id)}
	s.ItemService.UpdateItem(&item)
}

func (s *Server) StoreItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(("id")))

	if err != nil {
		c.Status(400)
		return
	}

	var body struct {
		ContainerID uint
	}
	c.Bind(&body)

	container, err := s.ContainerService.FindContainerByID(int(body.ContainerID))

	if err != nil {
		c.Status(400)
		return
	}

	item := domain.Item{ID: uint(id), ContainerID: body.ContainerID}

	if container.UserID != item.UserID {
		c.Status(403)
		return
	}

	s.ItemService.UpdateItem(&item)
}
