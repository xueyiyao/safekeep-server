package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/deprecated/initializers"
	"github.com/xueyiyao/safekeep/deprecated/models"
)

func CreateContainer(c *gin.Context) {
	var body struct {
		Name   string
		UserID uint
	}

	c.Bind(&body)

	container := models.Container{Name: body.Name, UserID: body.UserID}
	result := initializers.DB.Create(&container)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"container": container,
	})
}

func ReadContainers(c *gin.Context) {
	var body struct {
		UserID uint
	}

	c.Bind(&body)

	var containers []models.Container
	initializers.DB.Where("user_id <> ?", body.UserID).Find(&containers)

	c.JSON(200, gin.H{
		"containers": containers,
	})
}

func UpdateContainer(c *gin.Context) {
	containerId := c.Param(("id"))
	var container models.Container
	initializers.DB.First(&container, containerId)

	var body struct {
		Name string
	}

	c.Bind(&body)
	container.Name = body.Name
	initializers.DB.Save(&container)
}
