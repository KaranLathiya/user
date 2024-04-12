package request

type VerifyOTP struct {
	Firstname   string `json:"firstname"  validate:"required_if:EventType,signup,google_login|min_len:2|max_len:50"  `
	Lastname    string `json:"lastname"  validate:"required_if:EventType,signup,google_login|min_len:2|max_len:50"`
	Email       *string `json:"email" validate:"required_if:SignupMode,email,google_login|email|required_without:PhoneNumber|max_len:320" `
	PhoneNumber *string `json:"phoneNumber" validate:"required_if:SignupMode,phone|required_without:Email|min_len:4|max_len:15" `
	CountryCode *string `json:"countryCode" validate:"required_if:SignupMode,phone|min_len:2|max_len:5" `
	EventType   string `json:"eventType" validate:"required|in:signup,login,google_login,organization_delete"`
	SignupMode  string `json:"signupMode" validate:"required|in:phone,email,google_login"`
	OTP         string `json:"otp" validate:"required|len:6"`
}

type StoreOTP struct {
	Email       *string `json:"email" validate:"required_if:SignupMode,email,google_login|email|required_without:PhoneNumber|max_len:320" `
	PhoneNumber *string `json:"phoneNumber" validate:"required_if:SignupMode,phone|required_without:Email|min_len:4|max_len:15" `
	CountryCode *string `json:"countryCode" validate:"required_if:SignupMode,phone|min_len:2|max_len:5" `
	EventType   string `json:"eventType" validate:"required|in:signup,login,google_login,organization_delete"`
	LoginType   string `json:"loginType" validate:"required|in:phone,email,google_login"`
	HashedOTP   string `json:"HashedOTP" validate:"required"`
}
