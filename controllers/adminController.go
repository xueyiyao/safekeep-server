package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xueyiyao/safekeep/initializers"
	"github.com/xueyiyao/safekeep/models"
)

func AdminLogin(c *gin.Context) {
	id := c.Request.Header.Get("id")
	auth := c.Request.Header.Get("Authorization")

	if id == os.Getenv("ADMIN") {
		var user models.User
		initializers.DB.First(&user, id)

		hash := md5.Sum([]byte(user.Email + os.Getenv("TIMESTAMP") + os.Getenv("SALT")))

		if strings.ToLower(hex.EncodeToString(hash[:])) == strings.ToLower(auth) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				// subject and expiration
				"sub": user.ID,
				"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
			})

			tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to create token",
				})
			}

			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", os.Getenv("ENV") != "LOCAL", true)

			c.JSON(http.StatusOK, gin.H{})
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
