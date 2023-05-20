package google

type GoogleUserEmailResponse struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Picture  string `json:"picture"`
	Verified bool   `json:"verified_email"`
}
