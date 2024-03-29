package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xueyiyao/safekeep/domain"
	"golang.org/x/oauth2"
)

var OAuthStateString = "pseudo-random"
var testEmail = ""

func (s *Server) RegisterAuthRoutes(router *gin.Engine) {
	oauthRouter := router.Group("/oauth")
	{
		oauthRouter.POST("/google", s.handleOauthGoogleCallback)
	}
}

func (s *Server) handleOauthGoogleCallback(c *gin.Context) {
	var body struct {
		State *string `json:"state"`
		Code  *string `json:"code"`
	}

	c.Bind(&body)

	content, err := s.getUserInfo(*body.State, *body.Code)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	var user domain.User
	if content.Email != testEmail {
		// create new user?
		// initializers.DB.Where(domain.User{Email: content.Email}).FirstOrCreate(&user)
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

func (s *Server) getUserInfo(state string, code string) (*GoogleUserEmailResponse, error) {
	// TODO: Remember to randomize state string
	if state != OAuthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := s.OAuth2Config().Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	var contents *GoogleUserEmailResponse

	err = json.NewDecoder(response.Body).Decode(&contents)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}

type GoogleUserEmailResponse struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Picture  string `json:"picture"`
	Verified bool   `json:"verified_email"`
}
