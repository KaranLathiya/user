package request

type Signup struct {
	FirstName   string `json:"firstName" validate:"required|min_len:2|max_len:50" `
	LastName    string `json:"lastName" validate:"required|min_len:2|max_len:50" `
	Email       string `json:"email" validate:"email|required_without:PhoneNumber|max_len:320" `
	PhoneNumber string `json:"phoneNumber" validate:"required_without:Email|min_len:4|max_len:15"   `
	CountryCode string `json:"countryCode" `
	EventType   string `json:"eventType" validate:"required|in:phone,email,google_login"`
}
