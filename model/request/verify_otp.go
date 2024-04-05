package request

type VerifyOTP struct {
	Firstname   string `json:"firstname"  `
	Lastname    string `json:"lastname" `
	Email       string `json:"email" validate:"email|required_without:PhoneNumber|max_len:320" `
	PhoneNumber string `json:"phoneNumber" validate:"required_without:Email|min_len:4|max_len:15"   `
	CountryCode string `json:"countryCode" `
	EventType   string `json:"eventType" validate:"required|in:signup,login,google_login,organization_delete"`
	SignupMode  string `json:"signupMode" validate:"required|in:phone,email,google_login"`
	OTP         string `json:"otp" validate:"required|len:6"`
}
