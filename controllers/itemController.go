package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/initializers"
	"github.com/xueyiyao/safekeep/models"
	"gorm.io/gorm"
)

func CreateItem(c *gin.Context) {
	var body struct {
		Name   string
		UserID uint
	}

	c.Bind(&body)

	item := models.Item{Name: body.Name, UserID: body.UserID}
	result := initializers.DB.Create(&item)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

func ReadItem(c *gin.Context) {
	itemId := c.Param(("id"))
	var item models.Item
	result := initializers.DB.First(&item, itemId)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

func ReadItems(c *gin.Context) {
	var items []models.Item

	userIDStr := c.GetHeader("user-id")
	userID, err := strconv.ParseUint(userIDStr, 0, 64)
	if err != nil {
		c.Status(403)
		return
	}

	containerID := c.Query("container")
	var container models.Container
	var result *gorm.DB
	if len(containerID) != 0 {
		initializers.DB.First(&container, containerID)

		if container.UserID != uint(userID) {
			c.Status(403)
			return
		}

		result = initializers.DB.Where("user_id = ? AND container_id = ?", userID, containerID).Find(&items)
	} else {
		result = initializers.DB.Where("user_id = ?", userID).Find(&items)
	}

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"items": items,
	})
}

func StoreItem(c *gin.Context) {
	itemId := c.Param(("id"))
	var item models.Item
	initializers.DB.First(&item, itemId)

	var body struct {
		ContainerID uint
	}
	c.Bind(&body)
	var container models.Container
	initializers.DB.First(&container, body.ContainerID)

	if container.UserID != item.UserID {
		c.Status(403)
		return
	}

	item.ContainerID = body.ContainerID
	initializers.DB.Model(&item).Update("container_id", body.ContainerID)
}
