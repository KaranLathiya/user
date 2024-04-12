package request

type Signup struct {
	Firstname   string `json:"firstname" validate:"required|min_len:2|max_len:50" `
	Lastname    string `json:"lastname" validate:"required|min_len:2|max_len:50" `
	Email       *string `json:"email" validate:"required_if:LoginType,email,google_login|email|required_without:PhoneNumber|max_len:320" `
	PhoneNumber *string `json:"phoneNumber" validate:"required_if:LoginType,phone|required_without:Email|min_len:4|max_len:15" `
	CountryCode *string `json:"countryCode" validate:"required_if:LoginType,phone|min_len:2|max_len:5" `
	LoginType   string `json:"loginType" validate:"required|in:phone,email"`
}
