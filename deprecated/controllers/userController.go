package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/deprecated/logic"
	"github.com/xueyiyao/safekeep/models"
)

func CreateUser(c *gin.Context) {
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

	user := models.User{Name: *(body.Name), Email: *(body.Email)}

	// call logic layer
	result, err := logic.CreateUser(&user)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": result,
	})
}

func ReadUser(c *gin.Context) {
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

func UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param(("id")))

	if err != nil {
		c.Status(400)
		return
	}

	user := models.User{ID: uint(userId)}

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

	err = logic.UpdateUser(&user)

	if err != nil {
		c.Status(400)
		return
	}

	c.Status(200)
}
