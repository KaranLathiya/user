package response

// type GoogleAccessTokenResponse struct {
// 	AccessToken string `json:"access_token" validate:"required" `
// }

type GoogleUserInfo struct {
	Email     string `json:"email" validate:"required" `
	FirstName string `json:"firstname" validate:"required" `
	LastName  string `json:"lastname" validate:"required" `
	Message   string `json:"message"`
}

type GoogleAuthURL struct {
	AuthURL string `json:"authUrl" validate:"required" `
}
