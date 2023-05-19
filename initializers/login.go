package initializers

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	OAuthStateString  = "pseudo-random"
	GoogleOauthConfig *oauth2.Config
)

func SetupOAuthLogin() {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://35ff-135-180-67-224.ngrok-free.app/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
