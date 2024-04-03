package response

type GoogleAccessTokenResponse struct {
	AccessToken string `json:"access_token" validate:"required" `
}

type GoogleUserInfo struct {
	Email string `json:"email" validate:"required" `
	GivenName string `json:"given_name" validate:"required" `
	FamilyName string `json:"family_name" validate:"required" `
}

type GoogleAuthURL struct {
	AuthURL string `json:"auth_url" validate:"required" `
}