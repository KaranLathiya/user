package request

type GoogleAccessTokenRequest struct {
	Code         string `json:"code" validate:"required|min_len:10"  `
	ClientID     string `json:"client_id" validate:"required" `
	CLientSecret string `json:"client_secret" validate:"required" `
	RedirectURI  string `json:"redirect_uri" validate:"required" `
	GrantType    string `json:"grant_type" validate:"required" `
}
