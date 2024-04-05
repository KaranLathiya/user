package request

type Signup struct {
	FirstName   string `json:"firstname" validate:"required|min_len:2|max_len:50" `
	LastName    string `json:"lastname" validate:"required|min_len:2|max_len:50" `
	Email       string `json:"email" validate:"email|required_without:PhoneNumber|max_len:320" `
	PhoneNumber string `json:"phoneNumber" validate:"required_without:Email|min_len:4|max_len:15"   `
	CountryCode string `json:"countryCode" `
	LoginType   string `json:"loginType" validate:"required|in:phone,email,google_login"`
}
