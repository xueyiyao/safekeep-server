package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xueyiyao/safekeep/deprecated/initializers"
	"github.com/xueyiyao/safekeep/deprecated/models"
	"github.com/xueyiyao/safekeep/deprecated/models/google"
	"golang.org/x/oauth2"
)

func HandleGoogleLogin(c *gin.Context) {
	var body struct {
		State *string `json:"state"`
		Code  *string `json:"code"`
	}

	c.Bind(&body)

	content, err := getUserInfo(*body.State, *body.Code)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	var user models.User
	if content.Email != "alleny111@gmail.com" {
		// initializers.DB.Where(models.User{Email: content.Email}).FirstOrCreate(&user)
	}

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
}

func getUserInfo(state string, code string) (*google.GoogleUserEmailResponse, error) {
	if state != initializers.OAuthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := initializers.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	var contents *google.GoogleUserEmailResponse

	err = json.NewDecoder(response.Body).Decode(&contents)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
