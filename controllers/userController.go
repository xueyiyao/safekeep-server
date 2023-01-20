package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/initializers"
	"github.com/xueyiyao/safekeep/models"
)

func CreateUser(c *gin.Context) {
	var body struct {
		Name  string
		Email string
	}

	c.Bind(&body)

	user := models.User{Name: body.Name, Email: body.Email}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func ReadUser(c *gin.Context) {
	id := c.Param(("id"))

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}
