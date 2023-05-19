package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xueyiyao/safekeep/initializers"
	"golang.org/x/oauth2"
)

func HandleGoogleLogin(c *gin.Context) {
	var body struct {
		State *string `json:"state"`
		Code  *string `json:"code"`
	}

	c.Bind(&body)

	fmt.Println("HERE", *body.State, *body.Code)

	content, err := getUserInfo(*body.State, *body.Code)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	// fmt.Println("Content", string(content))
	c.JSON(200, content)
}

func getUserInfo(state string, code string) ([]byte, error) {
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
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Println(contents)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
